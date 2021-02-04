package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//read
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("message sent: ")
		//reading the string and holding it in text variable
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message received: " + message)
		//closing the connection
		if strings.TrimSpace(string(text)) == "close" {
			fmt.Println("closed")
			return
		}
	}
}
