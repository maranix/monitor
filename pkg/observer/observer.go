package observer

import (
	"github.com/fsnotify/fsnotify"
)

type Observer struct {
	watcher *fsnotify.Watcher
}

func Create(p string, c string) (*Observer, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	obs := &Observer{
		watcher: watcher,
	}

	return obs, nil
}
