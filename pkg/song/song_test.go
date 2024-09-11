package song

import (
	"testing"
)

func TestSong_Path(t *testing.T) {
	t.Run("obtain the path like artist/song.extension", func(t *testing.T) {
		expectedFilename := "Some Artist/Some Song.gp"
		s := Song{
			Artist: "Some Artist",
			Title:  "Some Song",
			Source: "https://songs.com/tab.gp",
		}

		if filename := s.Path(); filename != expectedFilename {
			t.Errorf("expected filename %s, got %s", expectedFilename, filename)
		}
	})
}
