package main

type commandID int

const (
	CMD_Name commandID = iota
	CMD_Join
	CMD_Rooms
	CMD_Msg
	CMD_Private_Message
	CMD_Exit
	CMD_Back
)

//each command needs to know the id, name of client calling it, and the actual string which may be required
type command struct {
	id     commandID
	client *client
	args   []string
}
