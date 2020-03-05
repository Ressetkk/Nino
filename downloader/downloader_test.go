package downloader

import "testing"

func TestDownloadFile(t *testing.T) {
	t.Run("DownloadFile should download desired file without errors", func(t *testing.T) {
		entry := DownloadEntry{streamUrl:"https://vignette.wikia.nocookie.net/meme/images/1/12/Ricardo.jpg/revision/latest?cb=20190616150107", title:"ricardo milos"}
		err := DownloadFile(&entry)
		if err != nil {
			t.Fail()
		}
	})
}