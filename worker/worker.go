package worker

import (
	"github.com/Ryan-Gosusluging/conveyor/types"
	"fmt"
	"time"
	"sync"
)

type Worker struct{
	ID int
	TrashBox *types.TrashBox
	Delay time.Duration
}

func NewWorker(id int, tb *types.TrashBox, delay time.Duration) *Worker{
	return &Worker{
		ID: id,
		TrashBox: tb,
		Delay: delay,
	}
}

func (w *Worker) StartWork(wg *sync.WaitGroup, stopChan <-chan struct{}){
	defer wg.Done()
	for{
		select{
		case <-stopChan:
			fmt.Printf("Работник %d завершает работу\n", w.ID)
			return
		default:
			time.Sleep(w.Delay)
			w.TrashBox.Mutex.Lock()
			if len(w.TrashBox.Channel) < w.TrashBox.Capacity{
				w.TrashBox.Channel <- 1
				fmt.Printf("Работник %d добавил мусор. Текущий уровень: %d/%d\n", w.ID, len(w.TrashBox.Channel), w.TrashBox.Capacity)
			}
			w.TrashBox.Mutex.Unlock()
		}
	}
}