package config

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kisinga/mzurihealth/models"
)

type hospconfig struct {
	name string
	id   string
}

func ReadFile() (config models.Hospital, err error) {
	file, _ := ioutil.ReadFile("config.txt")
	err = json.Unmarshal([]byte(file), &config)
	return
}

const pass = "zero-q/mzurihealth"

//Create will encrypt the hospital struct and save it to a file
func Create(hosp models.Hospital) {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(hosp); err != nil {
		panic(err)
	}
	ciphertext := encrypt([]byte(b.Bytes()), pass)

	fmt.Printf("Encrypted: %x\n", ciphertext)

	writeToFile("config.txt", ciphertext)

	plaintext := decrypt(ciphertext, pass)

	fmt.Printf("Decrypted: %s\n", plaintext)

	fmt.Println(string(decryptFile("config.txt", "sdkmk")))
}
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

// writeToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func writeToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return file.Sync()
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}
