package utils

import (
	"bytes"
)

func SliceString(cipher string, n int) []string {
	strArray := make([]string, n)
	for j := 0; j < n; j++ {
		var buffer bytes.Buffer
		for i := j; i*n < len(cipher); i++ {
			buffer.WriteString(string(cipher[i*n]))
		}
		strArray[j] = buffer.String()
	}
	return strArray
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
