package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
//	"strings"
)

func main(){

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host port number")
	}

	temp := bufio.NewReader(os.Stdin)
	port, _ := temp.ReadString('\n')

	conn, err := net.Dial("tcp", "localhost:" + port)
	if err != nil {
		fmt.Println(err)
		//return nil
	}
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")


		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text)

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
		//if strings.TrimSpace(string(text)) == "STOP" {
		//	fmt.Println("TCP client exiting...")
		}
	}
