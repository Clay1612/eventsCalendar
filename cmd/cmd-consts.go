package cmd

const TimeFormat = "02.01.2006-15:04"

const (
	addCommand      = "add"
	updateCommand   = "update"
	removeCommand   = "remove"
	listCommand     = "list"
	helpCommand     = "help"
	historyCommand  = "hist"
	exitCommand     = "exit"
	reminderCommand = "reminder"
)

const (
	addCommandUsage      = "Usage: add <event name> <date time> <priority>"
	updateCommandUsage   = "Usage: update <event name> <date time> <priority>"
	removeCommandUsage   = "Usage: remove <event name>"
	listCommandUsage     = "Usage: list"
	helpCommandUsage     = "Usage: help to show all commands"
	reminderCommandUsage = "Usage: reminder add/remove"
	historyCommandUsage  = "Usage: hist"
	addReminderUsage     = "Usage: reminder add <event name> <delay before event in minutes> <reminder message>"
	removeReminderUsage  = "Usage: reminder remove <event name>"
	exitCommandUsage     = "Usage: exit"
	unknownCommand       = "Unknown command"
)

const (
	successfulAdded           = "Event added"
	successfulUpdated         = "Event updated"
	successfulRemoved         = "Event removed"
	successfulReminderAdded   = "Reminder added"
	successfulReminderRemoved = "Reminder removed"
)
