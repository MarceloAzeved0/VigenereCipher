// Author: Marcelo Bernardy de Azevedo

package main

import (
	"VigenereCipher/ic"
	"VigenereCipher/utils"
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"sort"
)

type sizeIC struct {
	size int
	ic   float64
}

func main() {
	//read path file
	filePath := os.Args[1]

	//read file
	encyptedMessage, err := ioutil.ReadFile(filePath)
	//check file is valid
	utils.Check(err)

	//get lenght of chars
	var lstKeys []sizeIC

	// run check IC
	for m := 1; m <= 10; m++ {
		arrayCiphers := utils.SliceString(string(encyptedMessage), m)

		for _, cipherSliced := range arrayCiphers {
			letters := ic.CalcFrequencyLetters(cipherSliced)

			_, icCalc := ic.CalcIC(letters, len(cipherSliced))

			if icCalc > 0.070 && icCalc < 0.075 {
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
			return math.Abs(lstKeys[i].ic-0.070) < math.Abs(lstKeys[j].ic-0.073)
		})

		key := lstKeys[0].size

		slicedCipher := utils.SliceStringByInt(string(encyptedMessage), key)

		var strParsed []string
		var strFinal []string

		for l := 0; l < key; l++ {
			var buffer bytes.Buffer
			for _, text := range slicedCipher {
				if text != "" && l < len(text) {
					buffer.WriteString(string(text[l]))
				}
			}
			strParsed = append(strParsed, buffer.String())
		}

		// fmt.Println(strParsed)

		//get key

		var firstKey string
		for _, newTexts := range strParsed {
			firstLetter := ic.FirstLetterFrequency(newTexts)
			letterPosition := reflect.ValueOf(utils.CalcDistance2Chars("a", firstLetter))
			firstKey += reflect.Indirect(letterPosition).FieldByName("letter").String()
		}

		var secondKey string
		for _, newTexts := range strParsed {
			firstLetter := ic.FirstLetterFrequency(newTexts)
			letterPosition := reflect.ValueOf(utils.CalcDistance2Chars("e", firstLetter))
			secondKey += reflect.Indirect(letterPosition).FieldByName("letter").String()
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Digite a possível chave escolhendo uma das duas letras apresentadas:")
		fmt.Println("-> " + firstKey)
		fmt.Println("-> " + secondKey + "\n")
		choosedKey, _ := reader.ReadString('\n')
		fmt.Println("\nChave escolhida >>> " + choosedKey)

		for l, newTexts := range strParsed {
			firstLetter := ic.FirstLetterFrequency(newTexts)
			getLetterChoosed := reflect.ValueOf(utils.CalcDistance2Chars(string(choosedKey[l]), firstLetter))
			letterChoosed := reflect.Indirect(getLetterChoosed).FieldByName("letter").String()
			letterPosition := reflect.ValueOf(utils.CalcDistance2Chars(letterChoosed, firstLetter))
			position := reflect.Indirect(letterPosition).FieldByName("position").Int()
			var buffer bytes.Buffer

			for _, letter := range newTexts {
				newString := utils.ModStringWithDistanceInvert(string(letter), int(position))
				buffer.WriteString(newString)
			}

			strFinal = append(strFinal, buffer.String())
		}

		var strResponse string

		for p := 0; p < len(string(encyptedMessage))/key; p++ {
			for j := 0; j < key; j++ {
				strResponse = strResponse + string((strFinal[j][p]))
			}
		}

		fmt.Println(strResponse)
	}
}
