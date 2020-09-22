package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {

	rooms    map[string]*room
	commands chan command

}

func newServer() *server {

	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),

	}
}

//how the server will interpret and execute the commands form the client
func (s *server) run() {

	for cmd := range s.commands {

		switch cmd.id {

		case CMD_Name:
			s.name(cmd.client, cmd.args[1])

		case CMD_Join:
			s.join(cmd.client, cmd.args[1])

		case CMD_Msg:
			s.msg(cmd.client, cmd.args)

		case CMD_Private_Message:
			s.FindClient(cmd.client, cmd.args[1])

			s.PrivateMessage(cmd.client, cmd.args[2:])

		case CMD_Rooms:
			s.listRooms(cmd.client)

//		case CMD_ListMembers:
//			s.listMembers(cmd.client)

		case CMD_Exit:
			s.quit(cmd.client)
		}
	}
}

//identify the client from command argument
func (s *server) FindClient(c *client, name string) string{

//find remote address of desired client
	if val, ok := c.room.members[name.net.Addr]; ok {
		return val.conn.RemoteAddr().String()
	}

	//c.DisplayMsg("Desired client is not in the room")
	invalid := "Desired client is not in the room"
	return 	invalid

}


//print on server side the remote address of new clients who connect
func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		name:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

//modify client structure's name field
func (s *server) name(c *client, name string) {
	c.name = name
	c.msg(fmt.Sprintf("Hello %s", name))
}

//join a room is equal to adding client to the map of that room
func (s *server) join(c *client, roomName string) {

	//check if current room already exists
	r, ok := s.rooms[roomName]
	if !ok {
		//if current room doesn't exist we create it
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	//notify all members in the chat of a new arrival
	r.broadcast(c, fmt.Sprintf("%s has joined the chat", c.name))

	//welcome the client to the room
	c.DisplayMsg(fmt.Sprintf("-----Welcome to room: %s-----", roomName))
}

/*
func (s *server) listMembers(c *client){
	var CurrentMembers []string
	ppl := c.room.members
	for members := range ppl{
		CurrentMembers = append(CurrentMembers, members)
	}

	c.msg(fmt.Sprintf("available rooms: %s", strings.Join(CurrentMembers, ", ")))

}
*/

//list available rooms
func (s *server) listRooms(c *client) {
	var rooms []string

	//print all maps that exist within server
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.DisplayMsg(fmt.Sprintf("----------Available rooms:---------- \n%s ", strings.Join(rooms, "\n")))
}

//send a private message to a client within the room
func (s *server) PrivateMessage(c *client, args []string){
	message := strings.Join(args[1:len(args)], " ")
	c.msg(message)
}

//standard room message broadcast
func (s *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:len(args)], " ")
	c.room.broadcast(c, c.name + ": "+msg)
}

//leaving the room protocol
func (s *server) quit(c *client) {
	log.Printf("----------Client has left the chat:---------- \n%s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("later fader")
	c.conn.Close()
}

//exit from current room
func (s *server) quitCurrentRoom(c *client) {

	if c.room != nil {

		oldRoom := s.rooms[c.room.name]

		//use built in delete function for maps
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("----------%s has left the room----------\n", c.name))
	}
}