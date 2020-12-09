// 启动两个http server，监听linux sig关闭servers并退出进程
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return startServer(ctx, 3000)
	})
	g.Go(func() error {
		return startServer(ctx, 4000)
	})
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		// kill -2 or kill -15
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		fmt.Printf("receive sig: %s\n", sig)
		cancel()
		return nil
	})
	if err := g.Wait(); err != nil {
		log.Fatal("fatal ", err)
	}
	fmt.Print("done")
}

func startServer(pCtx context.Context, port int) error {
	ctx := context.WithValue(pCtx, "", "")
	portStr := strconv.Itoa(port)
	log.Print("startServer: ", portStr)

	var srv http.Server
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			if err := srv.Shutdown(ctx); err != nil {
				log.Printf("HTTP server %s Shutdown error: %v", srv.Addr, err)
			}
			log.Printf("HTTP server %s Shutdown successfully", srv.Addr)
			return
		}
	}(ctx)

	srv.Addr = ":" + portStr
	return srv.ListenAndServe()
}
