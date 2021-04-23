package utils

import (
	"bytes"
)

func SliceString(cipher string, n int) []string {
	strArray := make([]string, n)
	for j := 0; j < n; j++ {
		var buffer bytes.Buffer
		for i := 0; i*n+j < len(cipher); i++ {
			pos := i*n + j
			buffer.WriteString(string(cipher[pos]))
		}
		strArray[j] = buffer.String()
	}
	return strArray
}

func SliceStringByInt(cipher string, keySize int) []string {
	var strArray []string

	for i := 1; i <= len(cipher); i++ {
		cropped := i * keySize
		initial := cropped - keySize
		if cropped < len(cipher) {
			strArray = append(strArray, cipher[initial:cropped])
		}
	}

	return strArray
}

func CalcDistance2Chars(start string, finish string) int {
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
	countAlphabet["w"] = 22
	countAlphabet["x"] = 23
	countAlphabet["y"] = 24
	countAlphabet["z"] = 25

	sum := 26 + countAlphabet[finish] - countAlphabet[start]
	return sum % 26
}

func ModStringWithDistance(start string, distance int) string {
	var cLetter = make(map[string]int)
	cLetter["a"] = 0
	cLetter["b"] = 1
	cLetter["c"] = 2
	cLetter["d"] = 3
	cLetter["e"] = 4
	cLetter["f"] = 5
	cLetter["g"] = 6
	cLetter["h"] = 7
	cLetter["i"] = 8
	cLetter["j"] = 9
	cLetter["k"] = 10
	cLetter["l"] = 11
	cLetter["m"] = 12
	cLetter["n"] = 13
	cLetter["o"] = 14
	cLetter["p"] = 15
	cLetter["q"] = 16
	cLetter["r"] = 17
	cLetter["s"] = 18
	cLetter["t"] = 19
	cLetter["u"] = 20
	cLetter["v"] = 21
	cLetter["w"] = 22
	cLetter["x"] = 23
	cLetter["y"] = 24
	cLetter["z"] = 25

	var countAlphabet = make(map[int]string)
	countAlphabet[0] = "a"
	countAlphabet[1] = "b"
	countAlphabet[2] = "c"
	countAlphabet[3] = "d"
	countAlphabet[4] = "e"
	countAlphabet[5] = "f"
	countAlphabet[6] = "g"
	countAlphabet[7] = "h"
	countAlphabet[8] = "i"
	countAlphabet[9] = "j"
	countAlphabet[10] = "k"
	countAlphabet[11] = "l"
	countAlphabet[12] = "m"
	countAlphabet[13] = "n"
	countAlphabet[14] = "o"
	countAlphabet[15] = "p"
	countAlphabet[16] = "q"
	countAlphabet[17] = "r"
	countAlphabet[18] = "s"
	countAlphabet[19] = "t"
	countAlphabet[20] = "u"
	countAlphabet[21] = "v"
	countAlphabet[22] = "w"
	countAlphabet[23] = "x"
	countAlphabet[24] = "y"
	countAlphabet[25] = "z"

	sum := (cLetter[start] + distance) % 26

	return countAlphabet[sum]
}

func ModStringWithDistanceInvert(start string, distance int) string {
	var cLetter = make(map[string]int)
	cLetter["a"] = 0
	cLetter["b"] = 1
	cLetter["c"] = 2
	cLetter["d"] = 3
	cLetter["e"] = 4
	cLetter["f"] = 5
	cLetter["g"] = 6
	cLetter["h"] = 7
	cLetter["i"] = 8
	cLetter["j"] = 9
	cLetter["k"] = 10
	cLetter["l"] = 11
	cLetter["m"] = 12
	cLetter["n"] = 13
	cLetter["o"] = 14
	cLetter["p"] = 15
	cLetter["q"] = 16
	cLetter["r"] = 17
	cLetter["s"] = 18
	cLetter["t"] = 19
	cLetter["u"] = 20
	cLetter["v"] = 21
	cLetter["w"] = 22
	cLetter["x"] = 23
	cLetter["y"] = 24
	cLetter["z"] = 25

	var countAlphabet = make(map[int]string)
	countAlphabet[0] = "a"
	countAlphabet[1] = "b"
	countAlphabet[2] = "c"
	countAlphabet[3] = "d"
	countAlphabet[4] = "e"
	countAlphabet[5] = "f"
	countAlphabet[6] = "g"
	countAlphabet[7] = "h"
	countAlphabet[8] = "i"
	countAlphabet[9] = "j"
	countAlphabet[10] = "k"
	countAlphabet[11] = "l"
	countAlphabet[12] = "m"
	countAlphabet[13] = "n"
	countAlphabet[14] = "o"
	countAlphabet[15] = "p"
	countAlphabet[16] = "q"
	countAlphabet[17] = "r"
	countAlphabet[18] = "s"
	countAlphabet[19] = "t"
	countAlphabet[20] = "u"
	countAlphabet[21] = "v"
	countAlphabet[22] = "w"
	countAlphabet[23] = "x"
	countAlphabet[24] = "y"
	countAlphabet[25] = "z"

	sum := ((26 + cLetter[start]) - distance) % 26

	return countAlphabet[sum]
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
