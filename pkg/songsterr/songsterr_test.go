package songsterr

import (
	"net/http"
	"reflect"
	"songsterr-downloader/pkg/song"
	"songsterr-downloader/pkg/songsterr/testdata"
	"songsterr-downloader/pkg/util"
	"testing"
)

func TestGetIDFromURL(t *testing.T) {
	t.Run("get the ID from url", func(t *testing.T) {
		givenURL := "https://www.songsterr.com/a/wsa/some-song-tab-s123"
		expectedID := 123
		retrievedID, err := GetIDFromURL(givenURL)
		if err != nil {
			t.Fatal(err)
		}

		if expectedID != retrievedID {
			t.Fatalf("expected: %d, retrieved: %d", expectedID, retrievedID)
		}
	})
}

func TestService_SearchSongsByArtistID(t *testing.T) {
	t.Run("search some artist by id", func(t *testing.T) {
		client := util.NewFakeHTTPClient([]byte(testdata.GetSongByArtistResponse), http.StatusOK)
		service := NewService(client)
		results, err := service.SearchSongsByArtistID(123, Pagination{
			Limit:  10,
			Offset: 0,
		})
		if err != nil {
			t.Fatal(err)
		}

		expected := []song.SearchSongResult{
			{
				SongID:    1,
				ArtistID:  123,
				Artist:    "Some Artist",
				Title:     "Some Song I",
				HasChords: true,
				HasPlayer: true,
				Tracks: []song.Track{
					{
						InstrumentID: 1,
						Instrument:   "Electric Guitar",
						Views:        100,
						Name:         "Guitarist",
						Tuning:       []int{64, 59, 55, 50, 45, 40},
						Hash:         "guitar_hash",
						Difficulty:   1,
					},
					{
						InstrumentID: 2,
						Instrument:   "Electric Bass",
						Views:        100,
						Name:         "Bassist",
						Tuning:       []int{43, 38, 33, 28},
						Hash:         "bass_hash",
						Difficulty:   2,
					},
					{
						InstrumentID: 3,
						Instrument:   "Drums",
						Views:        100,
						Name:         "Drummer",
						Hash:         "drums_hash",
						Difficulty:   5,
					},
				},
				DefaultTrack:       1,
				PopularTrack:       2,
				PopularTrackGuitar: 1,
				PopularTrackBass:   2,
				PopularTrackDrum:   3,
			},
			{
				SongID:    2,
				ArtistID:  123,
				Artist:    "Some Artist",
				Title:     "Some Song II",
				HasChords: true,
				HasPlayer: true,
				Tracks: []song.Track{
					{
						InstrumentID: 1,
						Instrument:   "Electric Guitar",
						Views:        100,
						Name:         "Guitarist",
						Tuning:       []int{64, 59, 55, 50, 45, 40},
						Hash:         "guitar_hash",
						Difficulty:   2,
					},
					{
						InstrumentID: 2,
						Instrument:   "Electric Bass",
						Views:        100,
						Name:         "Bassist",
						Tuning:       []int{43, 38, 33, 28},
						Hash:         "bass_hash",
						Difficulty:   1,
					},
					{
						InstrumentID: 3,
						Instrument:   "Drums",
						Views:        100,
						Name:         "Drummer",
						Hash:         "drums_hash",
						Difficulty:   3,
					},
				},
			},
		}

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("\nexpected = %#v\ngot      = %#v", results, expected)
		}
	})
}
