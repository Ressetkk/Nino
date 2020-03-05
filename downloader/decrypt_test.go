package downloader

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDecryptSecurityKey(t *testing.T) {
	t.Run("DecryptSecurityKey returns proper key, nonce pair", func(t *testing.T) {
		rt := make([]byte, 32)
		_, _ = rand.Read(rt)
		fakeSecurityKey := base64.StdEncoding.EncodeToString(rt)
		key, nonce, err := DecryptSecurityKey(fakeSecurityKey)
		if err != nil{
			fmt.Println(err)
			t.Fail()
		}
		fmt.Printf("%x %x\n", key, nonce)
	})
}