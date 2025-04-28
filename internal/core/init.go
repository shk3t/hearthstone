package core

import (
	"hearthstone/internal/config"
	"hearthstone/internal/logging"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"os"
	"os/signal"
	"sync"
)

func InitAll() {
	initMutex.Lock()
	defer initMutex.Unlock()
	if initialized {
		return
	}
	initialized = true

	initAll()

	interruptChan = make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	go func() {
		_, ok := <-interruptChan
		if ok {
			DeinitAll()
			os.Exit(0)
		}
	}()
}

func DeinitAll() {
	initMutex.Lock()
	defer initMutex.Unlock()
	if !initialized {
		return
	}
	initialized = false

	deinitAll()

	signal.Stop(interruptChan)
	close(interruptChan)
}

func initAll() {
	config.Init()
	logging.Init()
	DisplayFrame = sugar.If(config.Config.Debug, ui.PrintFrame, ui.UpdateFrame)
}

func deinitAll() {
	logging.Deinit()
}

var initialized = false

var initMutex sync.Mutex

var interruptChan chan os.Signal