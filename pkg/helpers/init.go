package helpers

import (
	"os"
	"os/signal"
	"sync"
)

type Initializer struct {
	init          func()
	deinit        func()
	up            bool
	mutex         sync.Mutex
	interruptChan chan os.Signal
}

func NewInitializer(init func(), deinit func()) *Initializer {
	return &Initializer{
		init:          init,
		deinit:        deinit,
		up:            false,
		interruptChan: make(chan os.Signal, 1),
	}
}

func (i *Initializer) Init() {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.up {
		return
	}
	i.up = true

	i.init()

	signal.Notify(i.interruptChan, os.Interrupt)
	go func() {
		_, ok := <-i.interruptChan
		if ok {
			i.deinit()
			os.Exit(0)
		}
	}()
}

func (i *Initializer) Deinit() {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if !i.up {
		return
	}
	i.up = false

	i.deinit()

	signal.Stop(i.interruptChan)
	close(i.interruptChan)
}