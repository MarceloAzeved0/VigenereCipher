// Author: Marcelo Bernardy de Azevedo

package main

import (
	"VigenereCipher/ic"
	"VigenereCipher/utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"unicode/utf8"
)

// const ICPortuguese = 0.072723
const FirstLetterING = "e"

func main() {
	filePath := os.Args[1]

	//read file
	encyptedMessage, err := ioutil.ReadFile(filePath)
	utils.Check(err)

	//get lenght of chars
	sizeMessage := utf8.RuneCountInString(string(encyptedMessage))

	type sizeIC struct {
		size int
		ic   float64
	}

	var lstKeys []sizeIC

	// run check IC
	for m := 1; m <= 10; m++ {
		arrayCiphers := utils.SliceString(string(encyptedMessage), m)

		for _, cipherSliced := range arrayCiphers {
			letters := ic.CalcFrequencyLetters(cipherSliced)

			_, icCalc := ic.CalcIC(letters, len(cipherSliced))

			if icCalc > 0.064 && icCalc < 0.069 {
				lstKeys = append(lstKeys, sizeIC{m, icCalc})
			}
		}

		// TEST
		// if key != 0 {
		// 	fmt.Println(key)
		// 	break
		// }
	}

	if len(lstKeys) != 0 {
		fmt.Println(lstKeys)
		sort.Slice(lstKeys, func(i, j int) bool {
			return math.Abs(lstKeys[i].ic-0.0667) < math.Abs(lstKeys[j].ic-0.0667)
		})

		key := lstKeys[0].size

		slicedCipher := utils.SliceStringByInt(string(encyptedMessage), key)
		var strArray []string
		var strFinal []string

		for l := 0; l < key; l++ {
			var buffer bytes.Buffer
			for _, text := range slicedCipher {
				if text != "" {
					buffer.WriteString(string(text[l]))
				}
			}
			strArray = append(strArray, buffer.String())
		}

		// fmt.Println(strArray)

		for z, newTexts := range strArray {
			firstLetter := ic.FirstLetterFrequency(newTexts)
			if z == 1 {

			}
			distance := utils.CalcDistance2Chars(FirstLetterING, firstLetter)

			var buffer bytes.Buffer

			for _, letter := range newTexts {
				newString := ""
				if z == 1 {
					newString = utils.ModStringWithDistanceInvert(string(letter), 7)
				} else {
					newString = utils.ModStringWithDistanceInvert(string(letter), distance)
				}
				buffer.WriteString(newString)
			}

			strFinal = append(strFinal, buffer.String())
		}

		var strResponse string

		for p := 0; p < sizeMessage/key; p++ {
			for j := 0; j < key; j++ {
				strResponse = strResponse + string((strFinal[j][p]))
			}
		}

		fmt.Println(strResponse)
	}
}
