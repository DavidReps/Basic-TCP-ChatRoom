package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)


//we create a new server to listen and create the client structure when they join
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide server port number: \n ")
	}
	temp := bufio.NewReader(os.Stdin)
	port, _ := temp.ReadString('\n')

	//create new server
	s := newServer()
	//go routine to interpret commands from the client
	go s.run()

	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on port :" + port)

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		//create new clients as they join the server
		go s.newClient(conn)

	}
}
