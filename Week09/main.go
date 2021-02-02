package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Message msg for chan
type Message struct {
	MsgChan chan string
}

func sendMsg(conn net.Conn, ch <-chan string) {
	wr := bufio.NewWriter(conn)

	for msg := range ch {
		wr.WriteString("res: ")
		wr.WriteString(msg)
		wr.WriteString("\n")
		wr.Flush()
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	rd := bufio.NewReader(conn)

	msg := &Message{make(chan string, 8)}
	go sendMsg(conn, msg.MsgChan)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read line fail: %v\n", err)
			return
		}
		fmt.Printf("read line: %s\n", line)
		msg.MsgChan <- string(line)
	}
}

func main() {
	port := "3000"
	listen, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("tcp running, listen %s\n", port)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept fail: %v\n", err)
			continue
		}
		go handleConn(conn)
	}
}
