package infra

import "testing"

func TestFileHandler_Exist(t *testing.T) {
	t.Run("return false when try to open a file that not exists", func(t *testing.T) {
		fh := NewFileHandler("")
		exists, err := fh.Exist("do_not_exist.txt")
		if err != nil {
			t.Fatal(err)
		}

		if exists {
			t.Fatal("expected file to not exist")
		}
	})

	t.Run("return true when try to open a file that exists", func(t *testing.T) {
		fh := NewFileHandler("testdata")
		exists, err := fh.Exist("file.txt")
		if err != nil {
			t.Fatal(err)
		}

		if !exists {
			t.Fatal("expected file to exist")
		}
	})
}
