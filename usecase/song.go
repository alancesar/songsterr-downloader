package usecase

import (
	"fmt"
	"io"
	"log"
	"songsterr-downloader/pkg/song"
	"songsterr-downloader/pkg/songsterr"
	"songsterr-downloader/pkg/util"
)

type (
	FileHandler interface {
		Create(filename string) (io.Writer, error)
		Exist(filename string) (bool, error)
	}

	Downloader interface {
		Download(url string, writer io.Writer) error
	}

	SongUseCase struct {
		s  *songsterr.Service
		fh FileHandler
		d  Downloader
	}
)

func NewSongUseCase(s *songsterr.Service, fh FileHandler, d Downloader) *SongUseCase {
	return &SongUseCase{
		s:  s,
		fh: fh,
		d:  d,
	}
}

func (uc SongUseCase) DownloadTabByURL(url string) error {
	id, err := songsterr.GetIDFromURL(url)
	if err != nil {
		return err
	}

	return uc.DownloadTabByID(id)
}

func (uc SongUseCase) DownloadTabByID(id int) error {
	s, err := uc.s.GetSongByID(id)
	if err != nil {
		return err
	}

	path := s.Path()
	exists, err := uc.fh.Exist(path)
	if err != nil {
		return err
	} else if exists {
		return fmt.Errorf("%w: %s", util.ErrAlreadyExists, path)
	}

	writer, err := uc.fh.Create(path)
	if err != nil {
		return err
	}

	if err := uc.d.Download(s.Source, writer); err != nil {
		return err
	}

	log.Printf("downloaded %s successfully", s.Path())
	return nil
}

func (uc SongUseCase) DownloadTabsByArtistID(artistID int) error {
	results := uc.streamArtist(artistID)

	for result := range results {
		if err := uc.DownloadTabByID(result.SongID); err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (uc SongUseCase) streamArtist(artistID int) chan song.SearchSongResult {
	ssr := make(chan song.SearchSongResult)

	go func() {
		if err := uc.fetchArtist(artistID, songsterr.Pagination{
			Limit:  50,
			Offset: 0,
		}, ssr); err != nil {
			log.Println(err)
		}

		close(ssr)
	}()

	return ssr
}

func (uc SongUseCase) fetchArtist(artistID int, p songsterr.Pagination, ssr chan<- song.SearchSongResult) error {
	results, err := uc.s.SearchSongsByArtistID(artistID, p)
	if err != nil {
		return err
	}

	for _, r := range results {
		ssr <- r
	}

	if len(results) == p.Limit {
		return uc.fetchArtist(artistID, p.Next(), ssr)
	}

	return nil
}
