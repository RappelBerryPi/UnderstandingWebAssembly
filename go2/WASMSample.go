package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"syscall/js"
)

func encrypttest(this js.Value, args []js.Value) interface{} {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte(args[0].String())

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv, _ := hex.DecodeString("00000000000000000000000000000000")
	// iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	printString := hex.EncodeToString(ciphertext)
	fmt.Println(printString)
	return js.ValueOf(printString)

}

func registerCallbacks() {
	js.Global().Set("encrypt", js.FuncOf(encrypttest))
}

func main() {
	registerCallbacks()
	fmt.Printf("Hello World\n")
	c := make(chan struct{}, 0)
	<-c
}
