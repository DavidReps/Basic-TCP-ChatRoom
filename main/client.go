package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// define what is the client
type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

//read input and execute desired commands
func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		//parse out the message segments
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

	//identify the actual writing of each command as well as reading the parameter if needed
		switch cmd {

		//change name command
		case "/name":
			c.commands <- command{
				id:     CMD_Name,
				client: c,
				args:   args,
			}

		//join room command
		case "/join":
			c.commands <- command{
				id:     CMD_Join,
				client: c,
				args:   args,
			}

		//list available rooms command
		case "/rooms":
			c.commands <- command{
				id:     CMD_Rooms,
				client: c,
			}
/*
		//list members within a current room
		case "/listM":
			c.commands <- command{
			id:			CMD_ListMembers,
			client: 	c,
			}
*/
		//send message to the room
		case "/msg":
			c.commands <- command{
				id:     CMD_Msg,
				client: c,
				args:   args,
			}

		//send a private message
		case "/pmsg":
			c.commands <- command{
			id: 		CMD_Private_Message,
			client: 	c,
			args: 		args,
			}


		//exit the room
		case "EXIT":
			c.commands <- command{
				id:     CMD_Exit,
				client: c,
			}

		//protocol if unknown command is enters
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

//how errors are displayed
func (c *client) err(err error) {
	c.conn.Write([]byte("error: " + err.Error() + "\n"))
}

//how the messages are displayed in the room
func (c *client) msg(msg string) {
	c.conn.Write([]byte("-> " + msg + "\n"))
}

//used for presenting messages not from members
func (c *client) DisplayMsg(msg string) {
	c.conn.Write([]byte(msg + "\n"))
