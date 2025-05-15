package model

type ICommon interface {
	SetID(id string)
}

type DocumentChange struct {
	Action string
	Sheet  Sheet
}

var ActionDict = map[int]string{
	0: "ADD",
	1: "DELETED",
	2: "MODIFIED",
}
