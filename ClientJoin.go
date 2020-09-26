package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	conn, err := net.Dial("tcp", "localhost:"+CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		
		//connect client to the currently existing server
		go server.newClient(conn)

	}
}
