package song

import (
	"fmt"
	"path/filepath"
	"time"
)

type (
	Author struct {
		PersonID    int    `json:"personId"`
		Name        string `json:"name"`
		IsModerator bool   `json:"isModerator"`
	}

	Track struct {
		InstrumentID int    `json:"instrumentId"`
		Instrument   string `json:"instrument"`
		Views        int    `json:"views"`
		Name         string `json:"name"`
		Tuning       []int  `json:"tuning,omitempty"`
		Hash         string `json:"hash"`
		Difficulty   int    `json:"difficulty,omitempty"`
	}

	Video struct {
		ID      int    `json:"id"`
		Status  string `json:"status"`
		Feature string `json:"feature,omitempty"`
		VideoID string `json:"videoId"`
	}

	Report struct {
		Kind        string `json:"kind"`
		Summary     string `json:"summary"`
		ByModerator bool   `json:"byModerator"`
	}

	Song struct {
		CreatedAt          time.Time `json:"createdAt"`
		RevisionID         int       `json:"revisionId"`
		SongID             int       `json:"songId"`
		Artist             string    `json:"artist"`
		ArtistID           int       `json:"artistId"`
		Title              string    `json:"title"`
		Author             Author    `json:"author"`
		Description        string    `json:"description"`
		Source             string    `json:"source"`
		Restriction        string    `json:"restriction"`
		HasPlayer          bool      `json:"hasPlayer"`
		HasTracks          bool      `json:"hasTracks"`
		HasChords          bool      `json:"hasChords"`
		Tracks             []Track   `json:"tracks"`
		DefaultTrack       int       `json:"defaultTrack"`
		PopularTrack       int       `json:"popularTrack"`
		IsBlank            bool      `json:"isBlank"`
		Videos             []Video   `json:"videos"`
		PopularTrackGuitar int       `json:"popularTrackGuitar"`
		PopularTrackBass   int       `json:"popularTrackBass"`
		PopularTrackDrum   int       `json:"popularTrackDrum"`
		PrevRevisionID     int       `json:"prevRevisionId"`
		Tags               []string  `json:"tags"`
		Views              int       `json:"views"`
		Image              string    `json:"image"`
		Lyrics             bool      `json:"lyrics"`
		Audio              string    `json:"audio"`
		AudioV2            string    `json:"audioV2"`
		AudioV2Generated   string    `json:"audioV2Generated"`
		AudioV2MIDI        string    `json:"audioV2Midi"`
		AudioV4            string    `json:"audioV4"`
		AudioV4Generated   string    `json:"audioV4Generated"`
		AudioV4MIDI        string    `json:"audioV4Midi"`
		IsBlocked          bool      `json:"isBlocked"`
		IsOnModeration     bool      `json:"isOnModeration"`
		IsOnReview         bool      `json:"isOnReview"`
		IsCollaborative    bool      `json:"isCollaborative"`
	}

	Revision struct {
		SongID         int       `json:"songId"`
		RevisionID     int       `json:"revisionId"`
		CreatedAt      time.Time `json:"createdAt"`
		Artist         string    `json:"artist"`
		Title          string    `json:"title"`
		Author         Author    `json:"author"`
		Description    string    `json:"description"`
		Source         string    `json:"source"`
		TracksCount    int       `json:"tracksCount"`
		CommentsCount  int       `json:"commentsCount"`
		IsDeleted      bool      `json:"isDeleted"`
		IsBlocked      bool      `json:"isBlocked"`
		IsOnModeration bool      `json:"isOnModeration"`
		IsOnReview     bool      `json:"isOnReview"`
		Reports        []Report  `json:"reports"`
		GPImport       bool      `json:"gpImport"`
		AIGenerated    bool      `json:"aiGenerated"`
		Person         string    `json:"person"`
		PersonID       int       `json:"personId"`
	}

	SearchSongResult struct {
		SongID             int     `json:"songId"`
		ArtistID           int     `json:"artistId"`
		Artist             string  `json:"artist"`
		Title              string  `json:"title"`
		HasChords          bool    `json:"hasChords"`
		HasPlayer          bool    `json:"hasPlayer"`
		Tracks             []Track `json:"tracks"`
		DefaultTrack       int     `json:"defaultTrack"`
		PopularTrack       int     `json:"popularTrack"`
		PopularTrackGuitar int     `json:"popularTrackGuitar"`
		PopularTrackBass   int     `json:"popularTrackBass"`
		PopularTrackDrum   int     `json:"popularTrackDrum"`
	}
)

func (s Song) Path() string {
	ext := filepath.Ext(s.Source)
	filename := fmt.Sprintf("%s%s", s.Title, ext)
	return filepath.Join(s.Artist, filename)
}

func (s Song) Fullname() string {
	return fmt.Sprintf("%s - %s", s.Artist, s.Title)
}
