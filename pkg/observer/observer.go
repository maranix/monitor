package observer

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/maranix/monitor/pkg/config"
)

type Observer struct {
	watcher *fsnotify.Watcher
	cfg     config.Config
}

func NewObserver(cfg config.Config) (*Observer, error) {
	obs, err := new(cfg)
	if err != nil {
		msg := fmt.Sprintf("**Observer Failure:**\nTried to create an Observer but met an error.\n\n%s", err.Error())
		return nil, errors.New(msg)
	}

	return obs, nil
}

func (obs *Observer) Observe() error {
	return nil
}

func new(cfg config.Config) (*Observer, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	obs := Observer{
		watcher: watcher,
		cfg:     cfg,
	}

	return &obs, nil
}
