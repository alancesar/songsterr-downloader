package infra

import (
	"fmt"
	"io"
	"net/http"
	pkg "songsterr-downloader/pkg/util"
)

type (
	Downloader struct {
		client *http.Client
	}
)

func NewDownloader(client *http.Client) *Downloader {
	return &Downloader{
		client: client,
	}
}

func (d Downloader) Download(url string, writer io.Writer) error {
	resp, err := d.client.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d from %s", pkg.ErrBadStatus, resp.StatusCode, url)
	}

	if _, err := io.Copy(writer, resp.Body); err != nil {
		return err
	}

	return nil
}
