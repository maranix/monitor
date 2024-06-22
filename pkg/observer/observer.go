package observer

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/maranix/monitor/pkg/config"
)

type Observer struct {
	watcher *fsnotify.Watcher
	config  config.Config
}

func NewObserver(cfg config.Config) (*Observer, error) {
	obs, err := new(cfg)
	if err != nil {
		msg := fmt.Sprintf("**Observer Failure:**\nTried to create an Observer but met an error.\n\n%s", err.Error())
		return nil, errors.New(msg)
	}

	return obs, nil
}

func (obs *Observer) Observe() {
	errChan := make(chan error)
	notifChan := make(chan bool)

	go obs.eventListener(notifChan, errChan)

	for {
		select {
		case shouldRestart, ok := <-notifChan:
			if !ok {
				// Channel was closed
				return
			}

			if !shouldRestart {
				continue
			}

			fmt.Println("File changed")
		case msg, ok := <-errChan:
			if !ok {
				// Channel was closed
				return
			}

			fmt.Println(msg.Error())
		}
	}

}

func (obs *Observer) Close() error {
	return obs.close()
}

func new(cfg config.Config) (*Observer, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	obs := Observer{
		watcher: watcher,
		config:  cfg,
	}

	return &obs, nil
}

func (obs *Observer) eventListener(notifChan chan bool, errChan chan error) {
	for {
		select {
		case event, ok := <-obs.watcher.Events:
			if !ok {
				// Channel was closed (Watcher.Close() was called).
				//
				// Close the listener channels
				close(notifChan)
				close(errChan)
				return
			}

			if event.Op == fsnotify.Chmod {
				notifChan <- false
			} else {
				notifChan <- true
			}

		case err, ok := <-obs.watcher.Errors:
			fmt.Println(err.Error())
			if !ok {
				// Channel was closed (Watcher.Close() was called).
				//
				// Close the listener channels
				close(notifChan)
				close(errChan)
				return
			}

			errChan <- err
		}
	}
}

func (obs *Observer) close() error {
	err := obs.watcher.Close()
	if err != nil {
		return err
	}

	return nil
}
