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
	id int
)

func init() {
	flag.IntVar(&id, "id", 0, "artist id")
	flag.Parse()

	if id == 0 {
		log.Fatal("id param is required")
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
	if err := uc.DownloadTabsByArtistID(id); err != nil {
		log.Fatalln(err)
	}
}
