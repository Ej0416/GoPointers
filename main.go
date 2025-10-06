package main

import (
	"errors"
	"fmt"
)

//==================================================Chapter 10: Pointers

// -------------------------------------------------lesson 1
// type Message struct {
// 	Recipient string
// 	Text      string
// }

// func getMessageText(m Message) string {
// 	return fmt.Sprintf(`
// To: %v
// Message: %v
// `, m.Recipient, m.Text)
// }

// -------------------------------------------------lesson 2

// func removeProfanity(message *string) {
// 		*message = strings.ReplaceAll(*message, "fubb", "****")
// 		*message = strings.ReplaceAll(*message, "shiz", "****")
// 		*message = strings.ReplaceAll(*message, "witch", "*****")
// }

// ------------------------------------------------- lesson 3

// type Analytics struct {
// 	MessagesTotal     int
// 	MessagesFailed    int
// 	MessagesSucceeded int
// }

// type Message struct {
// 	Recipient string
// 	Success   bool
// }

// func analyzeMessage(a *Analytics, m Message) {
// 	if m.Success {
// 		a.MessagesSucceeded++
// 	} else {
// 		a.MessagesFailed++
// 	}

// 	a.MessagesTotal++
// }

// ------------------------------------------- lesson 6

// func removeProfanity(message *string) {
// 	if message != nil {
// 		messageVal := *message
// 		messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
// 		messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
// 		messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
// 		*message = messageVal
// 	}
// }

// ------------------------------------------- lesson 8
// func (e *email) setMessage(newMessage string) {
// 	e.message = newMessage
// }

// type email struct {
// 	message     string
// 	fromAddress string
// 	toAddress   string
// }

// ------------------------------------------ lesson 11

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customer *customer, transaction transaction) error {
	switch transaction.transactionType {
		case transactionDeposit:
			customer.balance += transaction.amount
		case transactionWithdrawal:
			if customer.balance < transaction.amount{
				return errors.New("insufficient funds")
			}
			customer.balance -= transaction.amount
		default:
			return errors.New("unknown transaction type")
	}
	return nil
}

func main() {
	fmt.Println("app started")
	var x int = 50
	var y *int = &x
	*y = 100
	fmt.Println(*y, x)
}
