package downloader

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

const MasterKey = "UIlTTEMmmLfGowo/UC60x2H45W6MdGgTRfo/umg4754="

type TidalDecipher struct {
	decipher cipher.Stream
	reader   io.Reader
}

func NewReader(r io.Reader, key, nonce []byte) io.Reader {
	iv := make([]byte, aes.BlockSize)
	copy(iv, nonce)
	c, _ := aes.NewCipher(key)
	d := cipher.NewCTR(c, iv)
	return &TidalDecipher{
		reader:   r,
		decipher: d,
	}
}

// Read reads and decrypts incoming Reader.
// It's used for decrypting audio stream
func (d *TidalDecipher) Read(p []byte) (n int, err error) {
	n, err = d.reader.Read(p)

	// if finished reading reader
	if err == io.EOF {
		return
	}
	// if something went wrong
	if err != nil {
		return 0, err
	}
	d.decipher.XORKeyStream(p, p)
	return
}

func DecryptSecurityKey(securityToken string) (key, nonce []byte, e error) {
	decodedMasterKey, err := base64.StdEncoding.DecodeString(MasterKey)
	if err != nil {
		return key, nonce, fmt.Errorf("base64 decoding error %w", err)
	}
	log.Debugf("Decoded MasterKey: %x\n", decodedMasterKey)
	decodedSecurityToken, err := base64.StdEncoding.DecodeString(securityToken)
	if err != nil {
		return key, nonce, fmt.Errorf("base64 decoding error %w", err)
	}
	log.Debugf("Decoded SecurityKey: %x\n", decodedSecurityToken)

	iv := decodedSecurityToken[:aes.BlockSize]
	encryptedSt := decodedSecurityToken[aes.BlockSize:]
	log.Debugf("%x | %x", iv, encryptedSt)

	var decipher cipher.BlockMode
	if d, err := aes.NewCipher(decodedMasterKey); err != nil {
		return key, nonce, err
	} else {
		decipher = cipher.NewCBCDecrypter(d, iv)
	}
	decryptedData := make([]byte, 32)
	decipher.CryptBlocks(decryptedData, encryptedSt)
	log.Debugf("%x", decryptedData)

	key = decryptedData[:aes.BlockSize]
	nonce = decryptedData[aes.BlockSize:24]
	return
}
