package ic

import (
	"sort"
)

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
