package loader

import (
	"fmt"
	"time"
	"sync"
	"github.com/Ryan-Gosusluging/conveyor/types"
)

type Loader struct{
	TrashBox *types.TrashBox
	MinDelay time.Duration
	CheckPeriod time.Duration
}

func NewLoader(tb *types.TrashBox, minDelaySec int, period time.Duration) *Loader{
	return &Loader{
		TrashBox: tb,
		MinDelay: time.Duration(minDelaySec)*time.Second,
		CheckPeriod: period,
	}
}

func (l *Loader) StartWork(wg *sync.WaitGroup, stopChan <-chan struct{}){
	defer wg.Done()
	lastLoad := time.Now()
	for {
		select {
		case <-stopChan:
			fmt.Println("Погрузчик завершает работу")
			return
		default:
			time.Sleep(l.CheckPeriod)
			
			l.TrashBox.Mutex.Lock()
			current := len(l.TrashBox.Channel)
			full := current == l.TrashBox.Capacity
			timePassed := time.Since(lastLoad) >= l.MinDelay
			l.TrashBox.Mutex.Unlock()

			if full || timePassed {
				l.TrashBox.Mutex.Lock()
				loadAmount := len(l.TrashBox.Channel)
				for i := 0; i < loadAmount; i++ {
					<-l.TrashBox.Channel
				}
				lastLoad = time.Now()
				l.TrashBox.Mutex.Unlock()

				if loadAmount > 0 {
					fmt.Printf("Погрузчик забрал %d единиц мусора. Следующая загрузка через %v\n", loadAmount, l.MinDelay)
				}
			}
		}
	}
}