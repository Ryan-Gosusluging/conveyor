package settings

import (
	"time"
)

type Settings struct{
	Workers int
	Pallets int
	LoaderDelay int
	WorkerDelay time.Duration
	LoaderPeriod time.Duration
}

func NewSettings() *Settings{
	return &Settings{
		Workers: 5,
		Pallets: 10,
		LoaderDelay: 5,
		WorkerDelay: 2000 * time.Millisecond,
		LoaderPeriod: 100 * time.Millisecond,
	}
}