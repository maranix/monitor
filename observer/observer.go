package observer

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/maranix/monitor/fsutil"
	"github.com/maranix/monitor/runner"
)

type Observer struct {
	watcher    *fsnotify.Watcher
	observable string
	command    string
}

func Create(p string, c string) (*Observer, error) {
	path, err := fsutil.AbsPath(p)
	if err != nil {
		return nil, err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	obs := &Observer{
		watcher:    watcher,
		observable: path,
		command:    c,
	}

	return obs, nil
}

func (obs *Observer) Observe() {
	slog.Info("Starting observer")
	dir, err := fsutil.IsDirectory(obs.observable)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if dir {
		obs.watchDirEvents()
	} else {
		files := []string{obs.observable}
		obs.watchFileEvents(files)
	}
}

func (obs *Observer) add(p string) {
	err := obs.watcher.Add(p)
	if err != nil {
		slog.Error(fmt.Sprintf("%q: %s", p, err))
		os.Exit(1)
	}
}

func (obs *Observer) watchDirEvents() {
	subdirs, err := fsutil.GetSubDirPaths(obs.observable)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	for _, dir := range subdirs {
		if !strings.Contains(dir, "/.") {
			obs.add(dir)
		}
	}

	slog.Info(fmt.Sprintf("Watching %d dirs", len(obs.watcher.WatchList())))
	slog.Info(fmt.Sprintf("WatchList %v", obs.watcher.WatchList()))

	defer obs.watcher.Close()
	// Start listening for events.
	go func() {
		var timer *time.Timer

		for {
			select {
			case event, ok := <-obs.watcher.Events:
				if timer != nil {
					timer.Stop()
				}

				if !ok || event.Op == fsnotify.Chmod {
					continue
				}

				timer = time.AfterFunc(time.Second*1, func() {
					slog.Info(fmt.Sprintf("event: %q, modified :%q", event, event.Name))
					slog.Info(fmt.Sprintf("executing command %s", obs.command))

					runner.Run(obs.command)
				})
			case err, ok := <-obs.watcher.Errors:
				if !ok {
					slog.Error("error:", err)
					return
				}
			}
		}
	}()

	<-make(chan struct{})
}

func (obs *Observer) watchFileEvents(files []string) {
	filemap := make(map[string]string)
	dirs := make([]string, 0)

	for _, f := range files {
		name := fsutil.GetFileFromPath(f)
		dir := fsutil.GetParentDir(f)
		filemap[name] = f

		if !slices.Contains(dirs, dir) {
			dirs = append(dirs, dir)
			obs.add(dir)
		}
	}

	slog.Info(fmt.Sprintf("Watching %d files", len(obs.watcher.WatchList())))
	slog.Info(fmt.Sprintf("WatchList %v", obs.watcher.WatchList()))

	defer obs.watcher.Close()
	// Start listening for events.
	go func() {
		var timer *time.Timer

		for {
			select {
			case event, ok := <-obs.watcher.Events:

				if timer != nil {
					timer.Stop()
				}

				if !ok || event.Op == fsnotify.Chmod {
					continue
				}

				timer = time.AfterFunc(time.Second*1, func() {
					slog.Info(fmt.Sprintf("event: %q, modified :%q", event, event.Name))
					slog.Info(fmt.Sprintf("executing command %s", obs.command))

					runner.Run(obs.command)
				})
			case err, ok := <-obs.watcher.Errors:
				if !ok {
					slog.Error("error:", err)
					return
				}
			}
		}
	}()

	<-make(chan struct{})
}
