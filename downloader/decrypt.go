package downloader

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

const MasterKey = "UIlTTEMmmLfGowo/UC60x2H45W6MdGgTRfo/umg4754="

func DecryptSecurityKey(securityToken string) (key, nonce []byte, e error) {
	decodedMasterKey, err := base64.StdEncoding.DecodeString(MasterKey)
	if err != nil {
		e = fmt.Errorf("base64 decoding error %w", err)
		return
	}
	log.Printf("Decoded MasterKey: %x\n", decodedMasterKey)
	decodedSecurityToken, err := base64.StdEncoding.DecodeString(securityToken)
	if err != nil {
		e = fmt.Errorf("base64 decoding error %w", err)
		return
	}
	log.Printf("Decoded SecurityKey: %x\n", decodedSecurityToken)


	iv := decodedSecurityToken[:16]
	encryptedSt := decodedSecurityToken[16:]
	log.Printf("%x | %x", iv, encryptedSt)

	var decrypter cipher.BlockMode
	if d, err := aes.NewCipher(decodedMasterKey); err == nil {
		decrypter = cipher.NewCBCDecrypter(d, iv)
	}
	decryptedData := make([]byte, 32)
	decrypter.CryptBlocks(decryptedData, encryptedSt)
	log.Printf("%x", decryptedData)
	key = decryptedData[:16]
	nonce = decryptedData[16:24]
	return
}