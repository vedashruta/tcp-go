package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Error : Please specify address with port of server")
	// }
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:8989")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Connected to : %s\n", conn.RemoteAddr().String())
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	consoleReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to be transmitted : ")
		message, err := consoleReader.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
		}
		writer.WriteString(message + "\n")
		writer.Flush()
		payload, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
		}
		fmt.Printf("%[1]s-%[2]v /> %[3]s\n", "server", time.Now(), payload)
	}
}
