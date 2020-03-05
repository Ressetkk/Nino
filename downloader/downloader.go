package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type DownloadEntry struct {
	streamUrl,
	coverUrl,
	title string
}

func DownloadFile(entry *DownloadEntry) error {
	resp, err := http.Get(entry.streamUrl)
	if err != nil {
		return fmt.Errorf("could not get download body %w", err)
	}
	defer resp.Body.Close()
	out, err := os.Create(fmt.Sprintf("%x", &entry.title))
	if err != nil {
		return fmt.Errorf("could not create temporary file %w", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func AddToQueue(w http.ResponseWriter, r *http.Request) {
	b := []byte{}
	r.Body.Read(b)
	w.Write(b)
}
