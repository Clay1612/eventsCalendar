package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Clay1612/eventsCalendar/calendar"
	"github.com/Clay1612/eventsCalendar/events"
	"github.com/Clay1612/eventsCalendar/helpers"
	"github.com/Clay1612/eventsCalendar/logger"
	"github.com/Clay1612/eventsCalendar/storage"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
)

type Cmd struct {
	Calendar        *calendar.Calendar
	TerminalHistory [][]string
	Storage         storage.Store
}

var mutex sync.Mutex

func NewCmd(c *calendar.Calendar, logStorage storage.Store) *Cmd {
	return &Cmd{
		Calendar:        c,
		TerminalHistory: [][]string{},
		Storage:         logStorage,
	}
}

func (c *Cmd) SaveTerminalHistory() error {
	data, err := json.Marshal(c.TerminalHistory)
	if err != nil {
		return fmt.Errorf("save terminal history func error: %w", err)
	}

	err = c.Storage.Save(data)
	if err != nil {
		return fmt.Errorf("save terminal history func error: %w", err)
	}

	return nil
}

func (c *Cmd) LoadTerminalHistory() error {
	data, err := c.Storage.Load()
	if err != nil {
		return fmt.Errorf("load terminal history func error: %w", err)
	}

	err = json.Unmarshal(data, &c.TerminalHistory)
	if err != nil {
		return fmt.Errorf("load terminal history func error: %w", err)
	}

	return nil
}

func (c *Cmd) LogTerminalMsg(msg string) {
	mutex.Lock()
	defer mutex.Unlock()

	terminalMsg, err := shlex.Split(msg)
	if err != nil {
		return
	}

	c.TerminalHistory = append(c.TerminalHistory, terminalMsg)
}

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)
	if err != nil {
		return
	}

	c.LogTerminalMsg(input)
	logger.Input(input)

	cmd := strings.ToLower(parts[0])

	switch cmd {
	case addCommand:
		if len(parts) != 4 {
			fmt.Println(addCommandUsage)
			c.LogTerminalMsg(addCommandUsage)
			logger.Output(addCommandUsage)
			return
		}

		_, err = c.Calendar.AddEvent(
			parts[1],
			parts[2],
			events.Priority(parts[3]),
		)

		if err != nil {
			helpers.ErrorHandler(err)
			c.LogTerminalMsg(err.Error())
			logger.Error(err.Error())
		} else {
			fmt.Println(successfulAdded)
			c.LogTerminalMsg(successfulAdded)
			logger.Output(successfulAdded)
		}

	case updateCommand:
		if len(parts) != 4 {
			fmt.Println(updateCommandUsage)
			c.LogTerminalMsg(updateCommandUsage)
			logger.Output(updateCommandUsage)
			return
		}

		err = c.Calendar.EditEvent(
			parts[1],
			parts[2],
			events.Priority(parts[3]))

		if err != nil {
			helpers.ErrorHandler(err)
			c.LogTerminalMsg(err.Error())
			logger.Error(err.Error())
		} else {
			fmt.Println(successfulUpdated)
			c.LogTerminalMsg(successfulUpdated)
			logger.Output(successfulUpdated)
		}

	case removeCommand:
		if len(parts) != 2 {
			fmt.Println(removeCommandUsage)
			c.LogTerminalMsg(removeCommandUsage)
			logger.Output(removeCommandUsage)
			return
		}

		err = c.Calendar.DeleteEvent(parts[1])

		if err != nil {
			helpers.ErrorHandler(err)
			c.LogTerminalMsg(err.Error())
			logger.Error(err.Error())
		} else {
			fmt.Println(successfulRemoved)
			c.LogTerminalMsg(successfulRemoved)
			logger.Output(successfulRemoved)
		}

	case listCommand:
		if len(parts) > 1 {
			fmt.Println(listCommandUsage)
			c.LogTerminalMsg(listCommandUsage)
			logger.Output(listCommandUsage)
			return
		}

		var eventsSlice []string
		for _, event := range c.Calendar.GetEvents() {
			eventString := fmt.Sprintf("%s (%s) - %s", event.Title, event.Priority, event.StartAt.Format(TimeFormat))
			fmt.Println(eventString)

			eventsSlice = append(eventsSlice, eventString)
		}
		c.TerminalHistory = append(c.TerminalHistory, eventsSlice)

	case helpCommand:
		if len(parts) != 1 {
			fmt.Println(helpCommandUsage)
			c.LogTerminalMsg(helpCommandUsage)
			logger.Output(helpCommandUsage)
			return
		}

		var helpSlice = []string{
			addCommandUsage,
			updateCommandUsage,
			removeCommandUsage,
			addReminderUsage,
			removeReminderUsage,
			listCommandUsage,
			historyCommandUsage,
			exitCommandUsage,
		}

		for _, command := range helpSlice {
			fmt.Println(command)
		}
		c.TerminalHistory = append(c.TerminalHistory, helpSlice)

	case reminderCommand:
		if len(parts) == 1 {
			fmt.Println(reminderCommandUsage)
			c.LogTerminalMsg(reminderCommandUsage)
			logger.Output(reminderCommandUsage)
			return
		}

		if len(parts) == 2 && parts[1] == addCommand {
			fmt.Println(addReminderUsage)
			c.LogTerminalMsg(addReminderUsage)
			logger.Output(addReminderUsage)
			return
		}

		if len(parts) == 2 && parts[1] == removeCommand {
			fmt.Println(removeReminderUsage)
			c.LogTerminalMsg(removeReminderUsage)
			logger.Output(removeReminderUsage)
			return
		}

		if parts[1] == addCommand {
			if len(parts) != 5 {
				fmt.Println(addReminderUsage)
				c.LogTerminalMsg(addReminderUsage)
				logger.Output(addReminderUsage)
				return
			}

			err = c.Calendar.SetEventReminder(parts[2], parts[3], parts[4])

			if err != nil {
				helpers.ErrorHandler(err)
				c.LogTerminalMsg(err.Error())
				logger.Error(err.Error())
			} else {
				fmt.Println(successfulReminderAdded)
				c.LogTerminalMsg(successfulReminderAdded)
				logger.Output(successfulReminderAdded)
			}
		}

		if parts[1] == removeCommand {
			if len(parts) != 3 {
				fmt.Println(removeReminderUsage)
				c.LogTerminalMsg(removeReminderUsage)
				logger.Output(removeReminderUsage)
				return
			}

			err = c.Calendar.CancelEventReminder(parts[2])

			if err != nil {
				helpers.ErrorHandler(err)
				c.LogTerminalMsg(err.Error())
				logger.Error(err.Error())
			} else {
				fmt.Println(successfulReminderRemoved)
				c.LogTerminalMsg(successfulReminderRemoved)
				logger.Output(successfulReminderRemoved)
			}
		}

	case historyCommand:
		if len(parts) != 1 {
			fmt.Println(historyCommandUsage)
			c.LogTerminalMsg(historyCommandUsage)
			logger.Output(historyCommandUsage)
		}

		for _, logMsg := range c.TerminalHistory {
			fmt.Println(logMsg)
		}

	case exitCommand:
		close(c.Calendar.Notification)

		err = c.Calendar.Save()
		if err != nil {
			fmt.Println(fmt.Errorf("save calendar error on exit: %w", err))
			c.LogTerminalMsg(err.Error())
			logger.Output(err.Error())
		}

		err = c.SaveTerminalHistory()
		if err != nil {
			fmt.Println(fmt.Errorf("save terminal history error on exit: %w", err))
		}

		err = logger.File.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("close logger file error on exit: %w", err))
		}

		os.Exit(0)

	default:
		fmt.Println(unknownCommand)
		c.LogTerminalMsg(unknownCommand)
		logger.Output(unknownCommand)
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: addCommand, Description: "Add a new event"},
		{Text: updateCommand, Description: "Change event"},
		{Text: removeCommand, Description: "Delete event"},
		{Text: reminderCommand + " " + "add", Description: "Event reminder installation"},
		{Text: reminderCommand + " " + "remove", Description: "Event reminder deleting"},
		{Text: listCommand, Description: "Show all events"},
		{Text: helpCommand, Description: "show command help"},
		{Text: historyCommand, Description: "Show terminal history"},
		{Text: exitCommand, Description: "quit program"},
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix(">"),
	)

	go func() {
		for msg := range c.Calendar.Notification {
			fmt.Println(msg)
			c.LogTerminalMsg(msg)
			logger.Output(msg)
		}
	}()

	p.Run()
}
