package main

import (
	"github.com/Ryan-Gosusluging/conveyor/loader"
	"github.com/Ryan-Gosusluging/conveyor/worker"
	"github.com/Ryan-Gosusluging/conveyor/settings"
	"github.com/Ryan-Gosusluging/conveyor/types"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main(){
	stt := settings.NewSettings()
	tb := types.NewTrashBox(stt.Pallets)
	stopChan := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nПолучен сигнал завершения, останавливаю систему...")
		close(stopChan)
	}()

	var wg sync.WaitGroup

	fmt.Printf("Запускаю %d работников...\n", stt.Workers)
	for i := 0; i < stt.Workers; i++ {
		wg.Add(1)
		w := worker.NewWorker(i, tb, stt.WorkerDelay)
		go w.StartWork(&wg, stopChan)
	}

	fmt.Println("Запускаю погрузчик...")
	wg.Add(1)
	l := loader.NewLoader(tb, stt.LoaderDelay, stt.LoaderPeriod)
	go l.StartWork(&wg, stopChan)

	fmt.Println("Система работает...")
	time.Sleep(10 * time.Second)
	close(stopChan)

	wg.Wait()
	fmt.Println("Система остановлена")
}