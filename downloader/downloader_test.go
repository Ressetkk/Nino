package downloader

import (
	"testing"
)

func TestDownloadFile(t *testing.T) {
	t.Run("DownloadFile should download desired file without errors", func(t *testing.T) {
		entry := DownloadEntry{
			streamUrl:   "http://en-pr-ak.audio.tidal.com/v2/0/927debf2fca45693d97b59063d67f034_26.flac?__token__=exp=1583709589~hmac=8a4e4120ad84e5916237cffacdc4f9bd2f7815bff310a07825855035e2e394a4",
			title:       "Zettai Zetsumei",
			securityKey: "0c1BiPpbtfNvUX+28tFZnng2dHoBIGcx/u31jUnchAypcnnbdZ6ab9/qGBV46tb1",
		}
		err := DownloadFile(&entry)
		if err != nil {
			t.Fail()
		}
	})
}
