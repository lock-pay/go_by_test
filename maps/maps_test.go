package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dict := Dictionnary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dict.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unkwon word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionnary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionnary{"idk": "idc"}
		err := dict.Add("idk", "man")

		assertErrors(t, err, ErrNotFound)
		assertDefinition(t, dict, "idk", "idc")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {

		dict := Dictionnary{"test": "old"}
		err := dict.Update("test", "new")

		assertErrors(t, err, nil)
		assertDefinition(t, dict, "test", "new")
	})

	t.Run("invalid", func(t *testing.T) {
		dict := Dictionnary{}
		err := dict.Update("test", "new")

		assertErrors(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		dict := Dictionnary{"test": "old"}
		err := dict.Delete("test")

		assertErrors(t, err, nil)
		_, err = dict.Search("test")
		assertErrors(t, err, ErrNotFound)
	})

	t.Run("invalid", func(t *testing.T) {
		dict := Dictionnary{}
		err := dict.Delete("test")

		assertErrors(t, err, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionnary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
