// Author: Marcelo Bernardy de Azevedo
// Email: marcelo.bernardy@edu.pucrs.br

package main

import (
	"VigenereCipher/ic"
	"VigenereCipher/utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//read path file
	filePath := os.Args[1]

	//read file
	encryptedMessage, err := ioutil.ReadFile(filePath)
	//check file is valid
	utils.Check(err)

	//get size of key
	keySize := ic.GetSizeKey(string(encryptedMessage), 10)

	if keySize != 0 {
		//generate subtring spliting by keySize of message
		slicedCipher := utils.SliceStringByInt(string(encryptedMessage), keySize)

		var strParsed []string
		//generate array of strings by each position until size of key
		for l := 0; l < keySize; l++ {
			var buffer bytes.Buffer
			for _, text := range slicedCipher {
				if text != "" && l < len(text) {
					buffer.WriteString(string(text[l]))
				}
			}
			strParsed = append(strParsed, buffer.String())
		}

		//chosed key with 2 options of letters "a" and "e"
		choosedKey := ic.GetKeyOfMessage(strParsed)

		//decrypted message with a choosed key and strings parsed by position
		decryptedMessage := ic.Decrypt(choosedKey, strParsed, len(string(encryptedMessage)), keySize)

		//write finally message in new file inside folder output
		f, _ := os.Create("./output/output.txt")

		defer f.Close()

		_, err := f.WriteString(decryptedMessage)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(">>> Salvo com sucesso saída gerada em: [PATH]/output/output.txt")

	}
}
