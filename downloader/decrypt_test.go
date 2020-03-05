package downloader

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	securityKey = "mcBE4Cx0EOhByT1UhFwNooGY1DTbNMzCihb03i93yE+XgIuL3PK1AlPct3kVYaGT"
	wantedKey   = []byte{0x21, 0xf, 0xc7, 0xbb, 0x81, 0x86, 0x39, 0xac, 0x48, 0xa4, 0xc6, 0xaf, 0xa2, 0xf1, 0x58, 0x1a}
	wantedNonce = []byte{0x58, 0xb1, 0xed, 0xd7, 0x4d, 0x7c, 0xa9, 0x2}
)

func TestDecryptSecurityKey(t *testing.T) {
	t.Run("DecryptSecurityKey returns proper key, nonce pair", func(t *testing.T) {

		key, nonce, err := DecryptSecurityKey(securityKey)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		fmt.Printf("%x %x", key, nonce)
		if !bytes.Equal(key, wantedKey) {
			t.Errorf("key/wantedKey mismatch!\nGot: %x\nWanted: %x\n", key, wantedKey)
		}
		if !bytes.Equal(nonce, wantedNonce) {
			t.Errorf("nonce/wantedNonce mismatch!\nGot: %x\nWanted: %x\n", nonce, wantedNonce)
		}
	})
}
