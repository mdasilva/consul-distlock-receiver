package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {

	fmt.Println("Got new connection!")
	defer func() {
		fmt.Println("Closing connection")
		conn.Close()
	}()

	bufReader := bufio.NewReader(conn)

	for {
		message, err := bufReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s\n", strings.TrimSuffix(message, "\n"))
	}
}

func main() {
	bindPtr := flag.String("bind", ":9000", "")
	flag.Parse()

	fmt.Printf("Listening on %s ...\n", *bindPtr)
	ln, err := net.Listen("tcp", *bindPtr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		ln.Close()
		fmt.Println("Listerner closed!")
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}
