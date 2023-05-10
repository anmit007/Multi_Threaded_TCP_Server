package main

import (
	"log"
	"net"
	"time"
)

func worker(id int, jobs chan net.Conn) {
	for {
		conn := <-jobs

		// Set a deadline for the incoming connection
		if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
			log.Printf("worker %d: failed to set connection deadline: %s", id, err)
			conn.Close()
			continue
		}

		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("worker %d: failed to read from connection: %s", id, err)
			conn.Close()
			continue
		}

		log.Printf("worker %d: processing the request", id)
		time.Sleep(8 * time.Second)

		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
		conn.Close()
	}
}

func main() {
	jobs := make(chan net.Conn)

	for i := 1; i <= 10; i++ {
		go worker(i, jobs)
	}

	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("waiting for a client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client connected")

		jobs <- conn
	}
}
