package observe

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/maranix/monitor/fsutil"
)

type Observer struct {
	watcher        *fsnotify.Watcher
	observablePath string
}

func Create(p string) (*Observer, error) {
	path, err := fsutil.AbsPath(p)
	if err != nil {
		return nil, err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	obs := &Observer{
		watcher:        watcher,
		observablePath: path,
	}

	return obs, nil
}

func (obs *Observer) Observe() {
	slog.Info("Starting observer")
	dir, err := fsutil.IsDirectory(obs.observablePath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if dir {
		subdirs, err := fsutil.GetSubDirPaths(obs.observablePath)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		for _, d := range subdirs {
			obs.add(d)
		}
	}

	slog.Info(fmt.Sprintf("Watching %d dirs", len(obs.watcher.WatchList())))

	defer obs.watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-obs.watcher.Events:

				if !ok {
					return
				}
				slog.Info(fmt.Sprintf("event:%q", event))
				if event.Has(fsnotify.Write) {
					slog.Info(fmt.Sprintf("modified file:%q", event.Name))
				}
			case err, ok := <-obs.watcher.Errors:
				if !ok {
					return
				}
				slog.Error("error:", err)
			}
		}
	}()

	<-make(chan struct{})
}

func (obs *Observer) add(p string) {
	err := obs.watcher.Add(p)
	if err != nil {
		slog.Error(fmt.Sprintf("%q: %s", p, err))
		os.Exit(1)
	}
}
