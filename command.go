package main

type commandID int

const (
	CMD_Name commandID = iota
	CMD_Join
	CMD_Rooms
	CMD_Msg
	CMD_Private_Message
	CMD_Exit
	CMD_ListMembers
)

type command struct {
	id     commandID
	client *client
	args   []string
}