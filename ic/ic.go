// Author: Marcelo Bernardy de Azevedo
// Email: marcelo.bernardy@edu.pucrs.br

package ic

import (
	"VigenereCipher/utils"
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
)

type sizeIC struct {
	size int
	ic   float64
}

func GetSizeKey(encryptedMessage string, attempts int) int {
	var lstKeys []sizeIC

	for m := 1; m <= attempts; m++ {
		arrayCiphers := utils.SliceString(encryptedMessage, m)

		for _, cipherSliced := range arrayCiphers {
			letters := CalcFrequencyLetters(cipherSliced)

			_, icCalc := CalcIC(letters, len(cipherSliced))

			if icCalc > 0.070 && icCalc < 0.075 {
				lstKeys = append(lstKeys, sizeIC{m, icCalc})
			}
		}
	}

	sort.Slice(lstKeys, func(i, j int) bool {
		return math.Abs(lstKeys[i].ic-0.070) < math.Abs(lstKeys[j].ic-0.073)
	})
	return lstKeys[0].size
}

func GetKeyOfMessage(strParsed []string) string {
	var firstKey string
	for _, newTexts := range strParsed {
		firstLetter := FirstLetterFrequency(string(newTexts))
		letterPosition := reflect.ValueOf(utils.CalcDistance2Chars("a", firstLetter))
		firstKey += reflect.Indirect(letterPosition).FieldByName("letter").String()
	}

	var secondKey string
	for _, newTexts := range strParsed {
		firstLetter := FirstLetterFrequency(string(newTexts))
		letterPosition := reflect.ValueOf(utils.CalcDistance2Chars("e", firstLetter))
		secondKey += reflect.Indirect(letterPosition).FieldByName("letter").String()
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite a possível chave escolhendo uma das duas letras apresentadas em cada palavra: ex: avelino\n")
	fmt.Println("-> " + firstKey)
	fmt.Println("-> " + secondKey + "\n")
	choosedKey, _ := reader.ReadString('\n')
	fmt.Println("\nChave escolhida >>> " + choosedKey)

	return choosedKey
}

func Decrypt(choosedKey string, strParsed []string, sizeMessage int, keySize int) string {
	var strFinal []string
	for l, newTexts := range strParsed {
		firstLetter := FirstLetterFrequency(string(newTexts))
		getLetterChoosed := reflect.ValueOf(utils.CalcDistance2Chars(string(choosedKey[l]), firstLetter))
		letterChoosed := reflect.Indirect(getLetterChoosed).FieldByName("letter").String()
		letterPosition := reflect.ValueOf(utils.CalcDistance2Chars(letterChoosed, firstLetter))
		position := reflect.Indirect(letterPosition).FieldByName("position").Int()
		var buffer bytes.Buffer

		for _, letter := range string(newTexts) {
			newString := utils.ModStringWithDistanceInvert(string(letter), int(position))
			buffer.WriteString(newString)
		}

		strFinal = append(strFinal, buffer.String())
	}

	var strResponse string

	for p := 0; p < sizeMessage/keySize; p++ {
		for j := 0; j < keySize; j++ {
			strResponse = strResponse + string((strFinal[j][p]))
		}
	}

	// fmt.Println(strResponse)

	return strResponse
}

func CalcFrequencyLetters(cipher string) map[string]int {
	var countAlphabet = make(map[string]int)
	countAlphabet["a"] = 0
	countAlphabet["b"] = 0
	countAlphabet["c"] = 0
	countAlphabet["d"] = 0
	countAlphabet["e"] = 0
	countAlphabet["f"] = 0
	countAlphabet["g"] = 0
	countAlphabet["h"] = 0
	countAlphabet["i"] = 0
	countAlphabet["j"] = 0
	countAlphabet["k"] = 0
	countAlphabet["l"] = 0
	countAlphabet["m"] = 0
	countAlphabet["n"] = 0
	countAlphabet["o"] = 0
	countAlphabet["p"] = 0
	countAlphabet["q"] = 0
	countAlphabet["r"] = 0
	countAlphabet["s"] = 0
	countAlphabet["t"] = 0
	countAlphabet["u"] = 0
	countAlphabet["v"] = 0
	countAlphabet["x"] = 0
	countAlphabet["w"] = 0
	countAlphabet["y"] = 0
	countAlphabet["z"] = 0

	for _, element := range cipher {
		countAlphabet[string(element)] = countAlphabet[string(element)] + 1
	}

	// fmt.Println(countAlphabet)

	return countAlphabet
}

func CalcIC(freqLetters map[string]int, size int) (map[string]float64, float64) {
	var calcIC = make(map[string]float64)
	var cipherSum = float64(size * (size - 1))
	var acc float64
	var ic float64

	for k, v := range freqLetters {
		var sumLetter = float64(v * (v - 1))

		acc += sumLetter
		calcIC[k] = sumLetter
	}

	ic += acc / cipherSum

	return calcIC, ic
}

func FirstLetterFrequency(firstString string) string {
	freqLetters := CalcFrequencyLetters(firstString)

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for k, v := range freqLetters {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss[0].Key
}
