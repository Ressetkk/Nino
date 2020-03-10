package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/Ressetkk/nino/downloader"
	"io/ioutil"
	"log"
	"os"
)

const masterKey = "UIlTTEMmmLfGowo/UC60x2H45W6MdGgTRfo/umg4754="

var (
	file = flag.String("file", "", "Path to file. [Required]")
)

func main() {
	flag.Parse()
	if *file == "" {
		log.Fatal("filepath can't be empty")
	} else if stat, err := os.Stat(*file); os.IsNotExist(err) {
		log.Fatal("file does not exist")
	} else if stat.IsDir() {
		log.Fatal("path is a directory")
	}

	filepath, key, _ := encryptFlac(*file)
	fmt.Printf("%s %x\n", filepath, key)
	encryptedKey := encryptKey(key)
	fmt.Printf("Encrypted base64 Key: %s\n", encryptedKey)
	key, nonce, _ := downloader.DecryptSecurityKey(encryptedKey)
	fmt.Printf("%x | %x\n", key, nonce)

}

func decryptFlac(file string, key []byte) string {
	c, _ := aes.NewCipher(key[:16])
	cip := cipher.NewCTR(c, key[16:])

	f, _ := ioutil.ReadFile(file)
	out := make([]byte, len(f))
	cip.XORKeyStream(out, f)
	of, _ := os.Create("decrypted-" + file)
	_, _ = of.Write(out)
	defer of.Close()
	return of.Name()
}

func encryptFlac(file string) (string, []byte, error) {
	secKey := make([]byte, 32)
	if _, err := rand.Read(secKey[:24]); err != nil {
		return "", nil, err
	}
	fmt.Printf("Got a key: %x\n", secKey)
	c, _ := aes.NewCipher(secKey[:16])
	cip := cipher.NewCTR(c, secKey[16:])

	// open and encrypt File
	f, _ := ioutil.ReadFile(file)
	out := make([]byte, len(f))
	cip.XORKeyStream(out, f)
	of, _ := os.Create("encrypted-" + file)
	_, _ = of.Write(out)
	defer of.Close()

	return of.Name(), secKey, nil
}

func encryptKey(k []byte) string {
	decodedMasterKey, _ := base64.StdEncoding.DecodeString(masterKey)
	iv := make([]byte, 16)
	_, _ = rand.Read(iv)
	cip, _ := aes.NewCipher(decodedMasterKey)
	enc := cipher.NewCBCEncrypter(cip, iv)
	encryptedData := make([]byte, 64)
	dataToEncrypt := make([]byte, 64)
	for i, v := range k {
		dataToEncrypt[i] = v
	}
	for i, v := range iv {
		dataToEncrypt[i+16] = v
	}
	enc.CryptBlocks(encryptedData, dataToEncrypt)

	securityKey := make([]byte, 32)
	copy(securityKey, encryptedData)
	fmt.Printf("%x\n", securityKey)
	return base64.StdEncoding.EncodeToString(securityKey)
}
