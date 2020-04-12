package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO edit this struct to match the final resource
type DownloadEntry struct {
	streamUrl,
	coverUrl,
	securityKey,
	title string
}

// TODO write proper downloading functions based on a resource
func DownloadFile(entry *DownloadEntry) error {
	resp, err := http.Get(entry.streamUrl)
	if err != nil {
		return fmt.Errorf("could not get download body %w", err)
	}
	defer resp.Body.Close()
	out, err := os.Create(fmt.Sprintf("%v-e", entry.title))
	if err != nil {
		return fmt.Errorf("could not create temporary file %w", err)
	}
	defer out.Close()
	//key, nonce, err := DecryptSecurityKey(entry.securityKey)
	//_, err = io.Copy(&TidalDecipher{key, nonce, out}, resp.Body)
	_, err = io.Copy(out, resp.Body)
	return err
}
