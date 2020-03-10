package downloader

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

func TestTidalDecrypter_Write(t *testing.T) {
	// TODO write test for fake encrypted file and key pair
	f, _ := os.Open("Zettai Zetsumei-e")
	defer f.Close()
	out, _ := os.Create("ZT-dec.flac")
	defer out.Close()
	key, nonce, _ := DecryptSecurityKey("0c1BiPpbtfNvUX+28tFZnng2dHoBIGcx/u31jUnchAypcnnbdZ6ab9/qGBV46tb1")
	d := NewReader(f, key, nonce)
	if _, err := io.Copy(out, d); err != nil {
		t.Errorf("Caught an error %v\n", err)
	}
}
