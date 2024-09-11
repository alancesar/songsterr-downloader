package infra

import (
	"bytes"
	"net/http"
	"songsterr-downloader/pkg/util"
	"testing"
)

func TestDownloader_Download(t *testing.T) {
	t.Run("copy downloaded content to the writer", func(t *testing.T) {
		expectedContent := "some content"
		client := util.NewFakeHTTPClient([]byte(expectedContent), http.StatusOK)
		downloader := NewDownloader(client)
		writer := new(bytes.Buffer)
		if err := downloader.Download("https://some.url", writer); err != nil {
			t.Fatal(err)
		}

		if writer.String() != expectedContent {
			t.Errorf("Downloader.Download returned unexpected body: got %v want %v",
				writer.String(), expectedContent)
		}
	})
}
