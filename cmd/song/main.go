package main

import (
	"flag"
	"log"
	"net/http"
	"songsterr-downloader/internal/infra"
	"songsterr-downloader/pkg/songsterr"
	"songsterr-downloader/usecase"
)

var (
	url string
)

func init() {
	flag.StringVar(&url, "url", "", "song url")
	flag.Parse()

	if url == "" {
		log.Fatal("url param is required")
	}
}

func main() {
	root, err := infra.DefaultRootPath()
	if err != nil {
		log.Fatalln(err)
	}

	defaultClient := http.DefaultClient
	s := songsterr.NewService(defaultClient)
	fh := infra.NewFileHandler(root)
	d := infra.NewDownloader(defaultClient)

	uc := usecase.NewSongUseCase(s, fh, d)
	if err := uc.DownloadTabByURL(url); err != nil {
		log.Fatalln(err)
	}
}
