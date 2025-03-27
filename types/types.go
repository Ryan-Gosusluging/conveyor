package types

import (
	"sync"
)

type TrashBox struct{
	Capacity int
	Channel chan int
	Mutex sync.Mutex
}

func NewTrashBox (capacity int) *TrashBox{
	return &TrashBox{
		Capacity: capacity,
		Channel: make(chan int, capacity),
	}
}