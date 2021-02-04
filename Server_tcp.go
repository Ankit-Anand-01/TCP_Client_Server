package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//start the server
	ln, err := net.Listen("tcp", ":8080")
	//error handling
	if err != nil {
		fmt.Println(err)
		return
	}
	//closing the listener
	defer ln.Close()
	//Accept waits and returns the next listener
	c, err := ln.Accept()
	//error handling
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//reading the string
		netData, err := bufio.NewReader(c).ReadString('\n')
		//error handling
		if err != nil {
			fmt.Println(err)
			return
		}
		//closing the connection
		if strings.TrimSpace(string(netData)) == "close" {
			fmt.Println("closed")
			return
		}

		fmt.Print("message received: ", string(netData))
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("message sent: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
	}
}
