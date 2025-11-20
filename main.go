package main

import (
	"fmt"

	"github.com/Clay1612/eventsCalendar/calendar"
	"github.com/Clay1612/eventsCalendar/cmd"
	"github.com/Clay1612/eventsCalendar/logger"
	"github.com/Clay1612/eventsCalendar/storage"
)

func main() {
	var err error

	var (
		calendarStorage        = storage.NewZipStorage("save/calendar.zip")
		terminalHistoryStorage = storage.NewJsonStorage("save/terminalHistoryLog.json")
	)

	eventsCalendar := calendar.NewCalendar(calendarStorage)
	err = eventsCalendar.Load()
	if err != nil {
		fmt.Println(err)
	}

	cli := cmd.NewCmd(eventsCalendar, terminalHistoryStorage)
	err = cli.LoadTerminalHistory()
	if err != nil {
		fmt.Println(err)
	}

	err = logger.Init()
	if err != nil {
		fmt.Println(err)
	}

	cli.Run()
}
