// Author: Marcelo Bernardy de Azevedo

package main

import (
	"VigenereCipher/ic"
	"VigenereCipher/utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

const ICPortuguese = 0.072723

func main() {
	var countAlphabet = make(map[string]int)
	countAlphabet["a"] = 0
	countAlphabet["b"] = 1
	countAlphabet["c"] = 2
	countAlphabet["d"] = 3
	countAlphabet["e"] = 4
	countAlphabet["f"] = 5
	countAlphabet["g"] = 6
	countAlphabet["h"] = 7
	countAlphabet["i"] = 8
	countAlphabet["j"] = 9
	countAlphabet["k"] = 10
	countAlphabet["l"] = 11
	countAlphabet["m"] = 12
	countAlphabet["n"] = 13
	countAlphabet["o"] = 14
	countAlphabet["p"] = 15
	countAlphabet["q"] = 16
	countAlphabet["r"] = 17
	countAlphabet["s"] = 18
	countAlphabet["t"] = 19
	countAlphabet["u"] = 20
	countAlphabet["v"] = 21
	countAlphabet["x"] = 22
	countAlphabet["w"] = 23
	countAlphabet["y"] = 24
	countAlphabet["z"] = 25

	filePath := os.Args[1]

	//read file
	encyptedMessage, err := ioutil.ReadFile(filePath)
	utils.Check(err)

	//get lenght of chars
	sizeMessage := utf8.RuneCountInString(string(encyptedMessage))

	key := 0

	// run check IC
	for m := 1; m <= 25; m++ {
		arrayCiphers := utils.SliceString(string(encyptedMessage), m)
		// fmt.Println(arrayCiphers)

		for _, cipherSliced := range arrayCiphers {
			letters := ic.CalcFrequencyLetters(cipherSliced)

			// fmt.Println(cipherSliced)

			_, icCalc := ic.CalcIC(letters, len(cipherSliced))

			// fmt.Println(cipherSliced, icCalc, m)

			// if m == 4 {
			// 	fmt.Println(cipherSliced)
			// }

			if icCalc > 0.062 && icCalc < 0.068 {
				//refactor after if

				fmt.Println("entrou", m, icCalc)

				key = m

				// firstLetter := ic.FirstLetterFrequency(arrayCiphers, m)
				// distance := utils.CalcDistance2Chars("e", firstLetter)

				// var buffer bytes.Buffer

				// for k := 0; k < len(encyptedMessage); k++ {
				// 	buffer.WriteString(utils.ModStringWithDistance(string(encyptedMessage[k]), distance))
				// }
				// result := buffer.String()

				// fmt.Println(result, firstLetter, distance)
				// fmt.Println(i, icCalc)
				break
			}
		}

		// TEST
		if key != 0 {
			break
		}
	}

	if key != 0 {
		slicedCipher := utils.SliceStringByInt(string(encyptedMessage), key, sizeMessage)
		strArray := make([]string, key)

		for l := 0; l < key; l++ {
			var buffer bytes.Buffer
			for _, text := range slicedCipher {
				if text != "" {
					buffer.WriteString(string(text[l]))
				}
			}
			strArray[l] = buffer.String()
		}
		fmt.Println(strArray, key)
	}
}
