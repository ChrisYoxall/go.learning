// Decorator is a structural design pattern that lets you attach new behaviors to objects.
// Really not sure if this is a good example...
package main

import "log"

type Notifier interface {
	Send(msg *string)
}

type LogNotifier struct{}

func (l *LogNotifier) Send(msg *string) {
	log.Println("Sending notification message:", *msg)
}

type SmsNotifier struct {
	Notifier
	PhoneNumber string
}

func (s *SmsNotifier) Send(msg *string) {
	s.Notifier.Send(msg)
	log.Println("Sending sms notification to:", s.PhoneNumber)
}

type EmailNotifier struct {
	Notifier
	Email string
}

func (e *EmailNotifier) Send(msg *string) {
	e.Notifier.Send(msg)
	log.Print("Sending email notification to: ", e.Email)
}

type SlackNotifier struct {
	Notifier
	SlackID string
}

func (s *SlackNotifier) Send(msg *string) {
	s.Notifier.Send(msg)
	log.Println("Sending slack notification to ID:", s.SlackID)
}

func main() {

	// Create a notifier. This is a little contrived as it just logs a message.
	n := Notifier(&LogNotifier{})

	// At runtime could add different combinations of notifiers as needed. Casting to Notifier interface.
	n = Notifier(&SmsNotifier{Notifier: n, PhoneNumber: "0123456789"})
	n = Notifier(&EmailNotifier{Notifier: n, Email: "first.last@fake.com"})
	n = Notifier(&SlackNotifier{Notifier: n, SlackID: "fake.slack.id"})

	// Send the message.
	message := "Hello World!"
	n.Send(&message)

}
