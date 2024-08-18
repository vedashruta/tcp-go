package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// if len(os.Args) < 3 {
	// 	fmt.Println("Error : Please specify address with port & concurrent connections")
	// }
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:8989")
	if err != nil {
		log.Fatalln(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	// concurrency, err := strconv.ParseInt(os.Args[2], 10, 0)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Printf("TCP Server Running on Port : %d \n", addr.Port)
	fmt.Println("Ready to accept connections... ")
	ch := make(chan struct{}, 1)
	defer close(ch)
	for {
		connection, err := listener.Accept()
		if err != nil {
			return
		}
		go handler(connection, ch)
	}
}

func handler(conn net.Conn, ch chan struct{}) {
	defer func() {
		<-ch
	}()
	buffer := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {
		message, err := buffer.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading from connection:", err)
			}
			break
		}
		fmt.Printf("%[1]s-%[2]s /> %[3]s", "client", time.Now(), message)
		writer.WriteString(message)
		err = writer.Flush()
		if err != nil {
			fmt.Println("write error :", err)
		}
	}
}
