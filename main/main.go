package main

import (
	"log"
	"net"
)


//we create a new server
func main() {

	s := newServer()

	//go routine to interpret commands from the client
	go s.run()

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on port :9000")

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		//conn.Write([]byte(myTime))
		//create new clients when they join
		go s.newClient(conn)

	}
}
