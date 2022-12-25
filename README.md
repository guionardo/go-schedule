# Go Simple Schedule Package

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guionardo/go-schedule)
![GitHub last commit](https://img.shields.io/github/last-commit/guionardo/go-schedule)
[![Go](https://github.com/guionardo/go-schedule/actions/workflows/go.yml/badge.svg)](https://github.com/guionardo/go-schedule/actions/workflows/go.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=guionardo_go-schedule&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=guionardo_go-schedule)


## Setup

```bash
❯ go get github.com/guionardo/go-schedule
go: downloading github.com/guionardo/go-schedule v0.0.3
go: added github.com/guionardo/go-schedule v0.0.3
```

On your project, you have to setup the schedules.

```golang
package main

import (
	"context"
	"fmt"
    "log"
	"os"
	"os/signal"
	"time"

	"github.com/guionardo/go-schedule/schedule"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	scheduler := schedule.NewScheduler().SetLogger(logger)
	// Add a task that runs every 15 minutes
	scheduler.AddSchedule(schedule.NewSchedule("First Task").Every(15 * time.Minute))
	// Add a task that runs every 20 minutes between 06:00 and 08:00
	scheduler.AddSchedule(schedule.NewSchedule("Second Task").Every(20 * time.Minute).
		DontRunBefore(6 * time.Hour).
		DontRunAfter(8 * time.Hour))

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	scheduler.Run(ctx, func(sch *schedule.Schedule) {
		fmt.Printf("Running %v", sch)
	})

}
```

The first example shows the event loop blocked, but you can start as a goroutine and control the exit with your context.

You, also, can use the ```scheduler.RunWithChannel(ctx, channel)``` to populate your channel and control the loop event in your application.