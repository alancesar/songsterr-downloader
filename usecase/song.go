package usecase

import (
	"fmt"
	"io"
	"log"
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
	song, err := uc.s.GetSongByID(id)
	if err != nil {
		return err
	}

	path := song.Path()
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

	if err := uc.d.Download(song.Source, writer); err != nil {
		return err
	}

	log.Printf("downloaded %s successfully", song.Fullname())
	return nil
}
