// Author: Marcelo Bernardy de Azevedo

package main

import (
	"VigenereCipher/ic"
	"VigenereCipher/utils"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

func main() {
	filePath := os.Args[1]

	encyptedMessage, err := ioutil.ReadFile(filePath)
	fmt.Println(string(encyptedMessage))
	utils.Check(err)

	// arrayCiphers := utils.SliceString(string(encyptedMessage), 1)

	letters := ic.CalcFrequencyLetters(string(encyptedMessage))

	ic.CalcIC(letters, utf8.RuneCountInString(string(encyptedMessage)))
}
