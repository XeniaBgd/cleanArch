package application

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Application struct {
	MainFunc func(ctx context.Context, halt <-chan struct{}) error

	appState int32
	mux      sync.Mutex
	err      error
	halt     chan struct{}
	done     chan struct{}
}

const (
	appStateInit int32 = iota
	appStateRunning
	appStateHalt
	appStateShutdown
)

var (
	errAppWrongState = errors.New("Wrong app state")
	errShutdown      = errors.New("Shutdown")
	errTermTimeout   = errors.New("Term timeout")
)

func (a *Application) checkState(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&a.appState, old, new)
}

func (a *Application) Run() error {
	if !a.checkState(a.appState, appStateRunning) {
		return errAppWrongState
	}

	if err := a.init(); err != nil {
		a.err = err
		a.appState = appStateShutdown
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	a.setError(a.run(sig))

	return a.getError()
}

type AppContext struct{}

func (a *Application) Value(key interface{}) interface{} {
	var appContext = AppContext{}
	if key == appContext {
		return a
	}
	return nil
}

func (a *Application) Err() error {
	if err := a.getError(); err != nil {
		return err
	}

	if atomic.LoadInt32(&a.appState) == appStateShutdown {
		return errShutdown
	}
	return nil
}

func (a *Application) Done() <-chan struct{} {
	return a.done
}

func (a *Application) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (a *Application) Halt() {
	if a.checkState(appStateRunning, appStateHalt) {
		close(a.halt)
	}
}

func (a *Application) Shutdown() {
	a.Halt()
	if a.checkState(appStateHalt, appStateShutdown) {
		close(a.done)
	}
}

func (a *Application) run(sig <-chan os.Signal) error {
	defer a.Halt()

	errRun := make(chan error, 1)

	// запуск сервера через объект приложения и ожидание корректного завершенеия, либо ошибки
	go func() {
		defer close(errRun)
		if err := a.MainFunc(a, a.halt); err != nil {
			errRun <- err
		}
	}()

	errHalt := make(chan error, 1)

	go func() {
		defer close(errHalt)

		select {
		case <-sig:
			log.Println("Get OS signal. Work interrupted")
			a.Halt()
			select {
			case <-time.After(time.Second * 15):
				// завершение по таймауту
				log.Println("Termination Timeout. Stop working...")
				errHalt <- errTermTimeout
			case <-a.done:
				// корректное завершение
			}
		case <-a.done:
			// завершение без участия операционной системы
		}
	}()

	select {
	case err, ok := <-errRun:
		if ok && err != nil {
			return err
		}
	case err, ok := <-errHalt:
		if ok && err != nil {
			return err
		}
	case <-a.done:
		log.Println("Shutdown")
	}

	return nil
}

func (a *Application) init() error {
	a.halt = make(chan struct{})
	a.done = make(chan struct{})

	return nil
}

func (a *Application) setError(err error) {
	if err == nil {
		return
	}

	a.mux.Lock()
	if a.err == nil {
		a.err = err
	}
	a.mux.Unlock()
	a.Shutdown()
}

func (a *Application) getError() error {
	var err error
	a.mux.Lock()
	err = a.err
	a.mux.Unlock()
	return err
}
