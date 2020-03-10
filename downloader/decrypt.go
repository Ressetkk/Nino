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
	for i, v := range nonce {
		iv[i] = v
	}
	c, _ := aes.NewCipher(key)
	d := cipher.NewCTR(c, iv)
	return &TidalDecipher{
		reader:   r,
		decipher: d,
	}
}
func (d *TidalDecipher) Read(p []byte) (n int, err error) {
	n, err = d.reader.Read(p)
	d.decipher.XORKeyStream(p, p)
	return
}

func DecryptSecurityKey(securityToken string) (key, nonce []byte, e error) {
	decodedMasterKey, err := base64.StdEncoding.DecodeString(MasterKey)
	if err != nil {
		e = fmt.Errorf("base64 decoding error %w", err)
		return
	}
	log.Debugf("Decoded MasterKey: %x\n", decodedMasterKey)
	decodedSecurityToken, err := base64.StdEncoding.DecodeString(securityToken)
	if err != nil {
		e = fmt.Errorf("base64 decoding error %w", err)
		return
	}
	log.Debugf("Decoded SecurityKey: %x\n", decodedSecurityToken)

	iv := decodedSecurityToken[:16]
	encryptedSt := decodedSecurityToken[16:]
	log.Debugf("%x | %x", iv, encryptedSt)

	var decrypter cipher.BlockMode
	if d, err := aes.NewCipher(decodedMasterKey); err == nil {
		decrypter = cipher.NewCBCDecrypter(d, iv)
	}
	decryptedData := make([]byte, 32)
	decrypter.CryptBlocks(decryptedData, encryptedSt)
	log.Debugf("%x", decryptedData)
	key = decryptedData[:16]
	nonce = decryptedData[16:24]

	return
}
