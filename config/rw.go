package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/kisinga/mzurihealth/models"
	"go.mongodb.org/mongo-driver/bson"
)

const pass = "zero-q/mzurihealth"

// Chasis is the package instance, it contains config settings,
type Chasis struct {
	debug bool
}

//CreateHosp will encrypt the hospital struct and save it to a file
func CreateHosp(chasis *Chasis, hosp models.Hospital) (err error) {
	b, _ := bson.Marshal(hosp)
	ciphertext := encrypt(b, pass)
	if chasis.debug {
		fmt.Printf("Encrypted: %x\n", ciphertext)
	}
	if _, _ = os.Stat("config.txt"); os.IsNotExist(err) {
		err = errors.New("File Alreasy exists")
	}
	err = writeToFile("config.txt", ciphertext)
	plaintext := decrypt(ciphertext, pass)
	if chasis.debug {
		fmt.Printf("Decrypted: %s\n", plaintext)
	}

	return
}

//New is the Constructor that defines packange-wide Chasis config
func New(debugstate bool) *Chasis {
	debug := &Chasis{
		debug: debugstate,
	}
	return debug
}

//ReadFile reads the config file and returns the decripted data or (and) errors
func ReadFile(chasis *Chasis) (config models.Hospital, err error) {
	data, returnerr := decryptFile("config.txt", pass)
	if returnerr != nil {
		empty := models.Hospital{}
		return empty, returnerr
	}
	err = bson.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error Unmarshaing Config ", err)
	}
	if chasis.debug {
		fmt.Printf("Decrypted: %#v\n", config)
	}
	return
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
	_, _ = file.Write(data)
	return file.Sync()
}

func decryptFile(filename string, passphrase string) (data []byte, err error) {
	encrypteddata, readerr := ioutil.ReadFile(filename)
	if readerr != nil {
		return []byte{}, readerr
	}
	data = decrypt(encrypteddata, passphrase)
	return
}
