package downloader

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	securityKey = "0UI0Fys+AmRBk1qecAXO5FFhp3z2xVzsLJ2fgXDpF/ApKDyce7n6DNZCilCsBEks"
	wantedKey   = []byte{0x84, 0xbc, 0x2c, 0xda, 0xd6, 0x10, 0x50, 0x66, 0xfe, 0xed, 0xbe, 0xca, 0xc, 0xa9, 0xc5, 0x2, 0x4b, 0xd0, 0xbe, 0x13, 0x2b, 0xa5, 0x85, 0x3d, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
)

func TestDownloadFile(t *testing.T) {
	// TODO fix this test
	t.Run("DownloadFile should download desired file without errors", func(t *testing.T) {
		entry := DownloadEntry{
			streamUrl:   "https://github.com/Ressetkk/Nino/raw/master/downloader/testdata/encrypted-172960.flac",
			title:       "testfile",
			securityKey: securityKey,
		}
		err := DownloadFile(&entry)
		if err != nil {
			t.Fail()
		}
	})
}

func TestDecryptSecurityKey(t *testing.T) {
	t.Run("DecryptSecurityKey returns proper key, nonce pair", func(t *testing.T) {

		key, nonce, err := DecryptSecurityKey(securityKey)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		fmt.Printf("%x %x", key, nonce)
		if !bytes.Equal(key, wantedKey[:aes.BlockSize]) {
			t.Errorf("key/wantedKey mismatch!\nGot: %x\nWanted: %x\n", key, wantedKey[:aes.BlockSize])
		}
		if !bytes.Equal(nonce, wantedKey[aes.BlockSize:24]) {
			t.Errorf("nonce/wantedNonce mismatch!\nGot: %x\nWanted: %x\n", nonce, wantedKey[aes.BlockSize:24])
		}
	})
}

func TestTidalDecrypter_Write(t *testing.T) {
	wanted, _ := ioutil.ReadFile(filepath.FromSlash("testdata/172960.flac"))
	encryptedFile, _ := os.Open(filepath.FromSlash("testdata/encrypted-172960.flac"))

	// anonymous func to silent warnings
	defer func() {
		if err := encryptedFile.Close(); err != nil {
			t.Errorf("Error during close %v\n", err)
		}
	}()

	key, nonce := wantedKey[:aes.BlockSize], wantedKey[aes.BlockSize:24]
	d := NewReader(encryptedFile, key, nonce)

	out, _ := ioutil.ReadAll(d)

	if !bytes.Equal(wanted, out) {
		t.Errorf("Wanted file and Decrypted file do not match!")
	}
}
