package main

import (
	"fmt"
	"strings"
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

// fubb	****
// shiz	****
// witch   *****

func removeProfanity(message *string) {
		*message = strings.ReplaceAll(*message, "fubb", "****")
		*message = strings.ReplaceAll(*message, "shiz", "****")
		*message = strings.ReplaceAll(*message, "witch", "*****")
}

func main() {
	fmt.Println("app started")
	x := "witch is so freakingly stupid what the fubb this shiz bro"
	removeProfanity(&x)
	fmt.Println(x)
}
