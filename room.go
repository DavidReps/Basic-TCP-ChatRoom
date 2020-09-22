package main

import (
	"net"
)
//rooms are maps filled with address as keys and pointer to client as value
type room struct {
	name    string
	members map[net.Addr]*client
}

//allow messages to be broadcasted to the entire room
func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}