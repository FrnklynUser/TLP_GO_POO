package main

import "fmt"

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println("LOG:", message)
}

type Notifier interface {
	Send(message string)
}

// EmailNotifier compone Logger
type EmailNotifier struct {
	Logger
	EmailAddress string
}

func (e EmailNotifier) Send(message string) {
	e.Log("Sending email...")
	fmt.Printf("📧 Enviado a %s: %s\n", e.EmailAddress, message)
}

// SMSNotifier compone Logger
type SMSNotifier struct {
	Logger
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) {
	s.Log("Sending SMS...")
	fmt.Printf("📱 Enviado a %s: %s\n", s.PhoneNumber, message)
}

func SendNotification(n Notifier, message string) {
	n.Send(message)
}

func main() {
	email := EmailNotifier{EmailAddress: "admin@example.com"}
	sms := SMSNotifier{PhoneNumber: "+123456789"}

	SendNotification(email, "Bienvenido al sistema")
	SendNotification(sms, "Tu pedido está listo")
}
