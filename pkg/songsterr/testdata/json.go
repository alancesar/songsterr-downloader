package testdata

const GetSongByArtistResponse = `
[
    {
        "songId": 1,
        "artistId": 123,
        "artist": "Some Artist",
        "title": "Some Song I",
        "hasChords": true,
        "hasPlayer": true,
        "tracks": [
            {
                "instrumentId": 1,
                "instrument": "Electric Guitar",
                "views": 100,
                "name": "Guitarist",
                "tuning": [
                    64,
                    59,
                    55,
                    50,
                    45,
                    40
                ],
                "difficulty": 1,
                "hash": "guitar_hash"
            },
            {
                "instrumentId": 2,
                "instrument": "Electric Bass",
                "views": 100,
                "name": "Bassist",
                "tuning": [
                    43,
                    38,
                    33,
                    28
                ],
				"difficulty": 2,
                "hash": "bass_hash"
            },
            {
                "instrumentId": 3,
                "instrument": "Drums",
                "views": 100,
                "name": "Drummer",
				"difficulty": 5,
                "hash": "drums_hash"
            }
        ],
		"defaultTrack": 1,
		"popularTrack": 2,
		"popularTrackGuitar": 1,
		"popularTrackBass": 2,
		"popularTrackDrum": 3
    },
    {
        "songId": 2,
        "artistId": 123,
        "artist": "Some Artist",
        "title": "Some Song II",
        "hasChords": true,
        "hasPlayer": true,
        "tracks": [
            {
                "instrumentId": 1,
                "instrument": "Electric Guitar",
                "views": 100,
                "name": "Guitarist",
                "tuning": [
                    64,
                    59,
                    55,
                    50,
                    45,
                    40
                ],
                "difficulty": 2,
                "hash": "guitar_hash"
            },
            {
                "instrumentId": 2,
                "instrument": "Electric Bass",
                "views": 100,
                "name": "Bassist",
                "tuning": [
                    43,
                    38,
                    33,
                    28
                ],
                "difficulty": 1,
                "hash": "bass_hash"
            },
            {
                "instrumentId": 3,
                "instrument": "Drums",
                "views": 100,
                "name": "Drummer",
                "difficulty": 3,
                "hash": "drums_hash"
            }
        ]
    }
]`
