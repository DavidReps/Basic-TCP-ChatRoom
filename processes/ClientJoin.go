package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	// "strings"
)

func read(c net.Conn){

	for {
		msg := bufio.NewScanner(bufio.NewReader(c))
		for msg.Scan(){
			fmt.Println(msg.Text())
		}
	}
}


func write(c net.Conn){

	for{

		cmd,err := bufio.NewReader(os.Stdin).ReadString('\n')

		if cmd == "EXIT"{
			return
		}
		if err != nil {
			fmt.Println("ERROR!")
			return
		}
		c.Write([]byte(cmd))

	}
}
func main(){

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host port number:")
	}

	temp := bufio.NewReader(os.Stdin)
	port, _ := temp.ReadString('\n')

	conn, err := net.Dial("tcp", "localhost:" + port)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	go read(conn)
	write(conn)

}
