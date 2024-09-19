package songsterr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"songsterr-downloader/pkg/song"
	"songsterr-downloader/pkg/util"
	"strconv"
	"strings"
)

const (
	apiBasePath = "https://www.songsterr.com/api"
)

type (
	Service struct {
		client *http.Client
	}

	Pagination struct {
		Limit  int
		Offset int
	}
)

func NewService(client *http.Client) *Service {
	return &Service{
		client: client,
	}
}

func GetIDFromURL(rawURL string) (int, error) {
	base := filepath.Base(rawURL)
	elem := strings.Split(base, "-")
	rawID := strings.TrimPrefix(elem[len(elem)-1], "s")
	return strconv.Atoi(rawID)
}

func (p Pagination) Next() Pagination {
	return Pagination{
		Limit:  p.Limit,
		Offset: p.Offset + p.Limit,
	}
}

func (s Service) GetSongByID(songID int) (song.Song, error) {
	parsedURL := fmt.Sprintf("%s/meta/%d", apiBasePath, songID)

	var output song.Song
	if err := s.doGet(parsedURL, &output); err != nil {
		return song.Song{}, err
	}

	return output, nil
}

func (s Service) GetRevisions(songID string) ([]song.Song, error) {
	parsedURL := fmt.Sprintf("%s/meta/%s/revisions", apiBasePath, songID)

	var songs []song.Song
	if err := s.doGet(parsedURL, &songs); err != nil {
		return nil, err
	}

	return songs, nil
}

func (s Service) SearchSongsByArtistID(artistID int, p Pagination) ([]song.SearchSongResult, error) {
	parsedURL, err := url.Parse(fmt.Sprintf("%s/artist/%d/songs", apiBasePath, artistID))
	if err != nil {
		return nil, err
	}

	query := parsedURL.Query()
	query.Set("size", strconv.Itoa(p.Limit))
	query.Set("from", strconv.Itoa(p.Offset))
	parsedURL.RawQuery = query.Encode()

	var results []song.SearchSongResult
	if err := s.doGet(parsedURL.String(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (s Service) doGet(parsedURL string, output any) error {
	resp, err := s.client.Get(parsedURL)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d from %s", util.ErrBadStatus, resp.StatusCode, parsedURL)
	}

	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return err
	}

	return nil
}
