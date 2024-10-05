package maps

import "testing"

func TestDictionary(t *testing.T) {
	dict := Dictionary{"test": "Hello!"}
	t.Run("Search key in the List", func(t *testing.T) {
		want := "Hello!"
		got, err := dict.Search("test")
		if err != nil {
			t.Errorf("%s", err)
		}
		assetStrings(t, got, want)
	})

	t.Run("Add new Key successfully to the List", func(t *testing.T) {
		want := "24"
		errors := dict.Add("24", "24")
		if errors != nil {
			t.Errorf("adding failed %s", errors.Error())
		}
		got, err := dict.Search("24")
		if err != nil {
			t.Error(err)
		}
		assetStrings(t, got, want)
	})

	t.Run("Test If key existed in list", func(t *testing.T) {
		want := "Key already exists in the list!."
		err := dict.Add("test", "24")
		if err == nil {
			t.Error(err)
		}
		assetStrings(t, err.Error(), want)
	})

	// t.Run("Remove Existed Key from List", func(t *testing.T) {
	// 	word := "dog"
	// 	err := dict.Add(word, "bow bow")

	// 	dict.Remove(word)
	// 	_, err := dict.Search(word)
	// 	assetStrings(t, err.Error(), errorNotExists)
	// })

}

func assetStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
