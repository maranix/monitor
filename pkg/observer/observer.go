package observer

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/maranix/monitor/pkg/config"
	"github.com/maranix/monitor/pkg/debouncer"
	"github.com/maranix/monitor/pkg/runner"
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
	watcher := obs.watcher
	target := obs.config.GetTarget()

	notifChan := make(chan bool)
	errChan := make(chan error)

	debouncer, err := debouncer.NewDebouncer(obs.config.GetDebounce())
	if err != nil {
		errChan <- err
	}

	go eventListener(watcher, debouncer, notifChan, errChan)

	if err := watchTarget(watcher, target); err != nil {
		fmt.Println(err.Error())
	}

	for {
		select {
		case shouldRestart, ok := <-notifChan:
			if !ok {
				// Channel was closed
				return
			}

			if shouldRestart {
				runner.Run(obs.config.GetRunner())
			}
		case msg, ok := <-errChan:
			fmt.Println(msg.Error())
			if !ok {
				// Channel was closed
				return
			}
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

func (obs *Observer) close() error {
	err := obs.watcher.Close()
	if err != nil {
		return err
	}

	return nil
}

func eventListener(watcher *fsnotify.Watcher, debouncer *debouncer.Debouncer, notifChan chan bool, errChan chan error) {
	if watcher == nil {
		errChan <- errors.New("Expected a reference to watcher but received nil.\nThis scenario should never happen, please file a bug report or open an issue in the source repostiory for further debugging and support.")
		return
	}

	if debouncer == nil {
		errChan <- errors.New("Expected a reference to debouncer but received nil.\nThis scenario should never happen, please file a bug report or open an issue in the source repostiory for further debugging and support.")
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				// Channel was closed (Watcher.Close() was called).
				//
				// Close the listener channels
				close(notifChan)
				close(errChan)
				return
			}

			debouncer.Call(func() {
				if event.Op == fsnotify.Chmod {
					notifChan <- false
				} else {
					notifChan <- true
				}
			})
		case err, ok := <-watcher.Errors:
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

func watchTarget(watcher *fsnotify.Watcher, target string) error {
	err := watcher.Add(target)
	if err != nil {
		return err
	}

	return nil
}
