// 启动两个http server，监听linux sig关闭servers并退出进程
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return startServer(ctx, "3000")
	})
	g.Go(func() error {
		return startServer(ctx, "4000")
	})
	g.Go(func() error {
		return listenSig(ctx)
	})
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("it's time to exit")
	})
	err := g.Wait()
	log.Printf("exit, err: %v\n", err)
}

func startServer(ctx context.Context, port string) error {
	log.Printf("startServer: %s\n", port)
	srv := http.Server{Addr: ":" + port}
	go func(ctx context.Context) {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server %s Shutdown error: %v", srv.Addr, err)
		}
		log.Printf("HTTP server %s Shutdown successfully", srv.Addr)
	}(ctx)

	return srv.ListenAndServe()
}

func listenSig(ctx context.Context) error {
	// kill -2 or kill -15
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fmt.Printf("no sig")
		return nil
	case s := <-sigs:
		fmt.Printf("receive sig: %s\n", s)
		return nil
	}
	return nil
}
