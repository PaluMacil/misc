package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	plaintext, err := ioutil.ReadFile("raw.txt")
	if err != nil {
		panic(err.Error())
	}

	// this is a key
	key := []byte("example key 1234")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// create a new file for saving the encrypted data.
	f, err := os.Create("encrypted.aes.txt")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}
