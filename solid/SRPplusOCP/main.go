/*
🎯 Scenario:
You’re building a notification system for an enterprise product.

Your system should:

Accept and process user alerts (e.g., password reset, low balance, new login).

Send these alerts via different channels: Email, SMS, Push notification, WhatsApp, etc.

More channels (like Slack, Telegram) may be added in the future.

Logs should be handled by a separate logger (console or file—simulate via fmt.Println).

No part of your existing alert/notification code should need to change when a new channel or new alert type is added.
*/

package main

import "fmt"

type Alert interface {
	Message() string
}

type PasswordResetAlert struct{}
type LowBalanceAlert struct{}
type NewLoginAlert struct{}

func (pra PasswordResetAlert) Message() string {
	return "Your password has been reset"
}

func (lba LowBalanceAlert) Message() string {
	return "Your Balance is low"
}

func (nla NewLoginAlert) Message() string {
	return "New login has been attempt"
}

type NotificationChannel interface {
	Notify(string) string
}

type EmailChannel struct{}
type SMSChannel struct{}
type WhatsAppChannel struct{}

func (ec EmailChannel) Notify(message string) string {
	return fmt.Sprintf("Email Notification '%s'", message)
}

func (sc SMSChannel) Notify(message string) string {
	return fmt.Sprintf("SMS Notification '%s'", message)
}

func (wc WhatsAppChannel) Notify(message string) string {
	return fmt.Sprintf("WhatsApp Notification '%s'", message)
}

type Logger interface {
	Log(string)
}

type ConsoleLogger struct{}
type FileLogger struct{}
type RemoteServerLogger struct{}
type DBLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("Logging to Console: ", message)
}

func (fl FileLogger) Log(message string) {
	fmt.Println("Logging to File: ", message)
}

func (rsl RemoteServerLogger) Log(message string) {
	fmt.Println("Logging to Remote Server: ", message)
}

func NotifyAndLog(a Alert, c NotificationChannel, l Logger) {
	message := a.Message()
	notification := c.Notify(message)
	l.Log(notification)
}

func main() {
	alerts := []Alert{PasswordResetAlert{}, NewLoginAlert{}, LowBalanceAlert{}}
	channels := []NotificationChannel{EmailChannel{}, WhatsAppChannel{}}
	logger := ConsoleLogger{}

	for _, alert := range alerts {
		for _, channel := range channels {
			NotifyAndLog(alert, channel, logger)
		}
	}

}
