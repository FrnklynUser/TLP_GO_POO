package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) error
	GetFee() float64
}

type CreditCardProcessor struct {
	CardNumber string
	FeeRate    float64
}

func (cc CreditCardProcessor) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("monto inválido")
	}
	fmt.Printf("💳 Procesando $%.2f con tarjeta ****%s\n", amount, cc.CardNumber[len(cc.CardNumber)-4:])
	return nil
}

func (cc CreditCardProcessor) GetFee() float64 {
	return cc.FeeRate
}

type PayPalProcessor struct {
	Email string
}

func (pp PayPalProcessor) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("monto inválido")
	}
	fmt.Printf("🅿 Procesando $%.2f via PayPal (%s)\n", amount, pp.Email)
	return nil
}

func (pp PayPalProcessor) GetFee() float64 {
	return 0.029 // 2.9%
}

type CryptoProcessor struct {
	WalletAddress string
	Currency      string
}

func (cp CryptoProcessor) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("monto inválido")
	}
	fmt.Printf("🪙 Procesando $%.2f en %s (wallet: %s...)\n", amount, cp.Currency, cp.WalletAddress[:10])
	return nil
}

func (cp CryptoProcessor) GetFee() float64 {
	return 0.01 // 1%
}

func ProcessOrder(processor PaymentProcessor, amount float64) error {
	fee := processor.GetFee()
	total := amount + (amount * fee)

	fmt.Printf("🧮 Total con comisión: $%.2f (comisión: %.1f%%)\n", total, fee*100)
	return processor.Process(amount)
}

func SelectBestProcessor(amount float64, processors ...PaymentProcessor) PaymentProcessor {
	if len(processors) == 0 {
		return nil
	}

	bestProcessor := processors[0]
	lowestFee := bestProcessor.GetFee()

	for _, p := range processors[1:] {
		if p.GetFee() < lowestFee {
			bestProcessor = p
			lowestFee = p.GetFee()
		}
	}

	fmt.Printf("Seleccionado procesador con %.1f%% de comisión\n", lowestFee*100)
	return bestProcessor
}

func main() {
	creditCard := CreditCardProcessor{
		CardNumber: "1234567890123456",
		FeeRate:    0.035,
	}

	paypal := PayPalProcessor{
		Email: "franklinUss@example.com",
	}

	crypto := CryptoProcessor{
		WalletAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		Currency:      "BTC",
	}

	best := SelectBestProcessor(100, creditCard, paypal, crypto)
	ProcessOrder(best, 100)
}
