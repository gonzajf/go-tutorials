package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectRepetition := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeating 5 times character 'a'", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectRepetition(t, got, want)
	})

	t.Run("repeating 2 times character 'a'", func(t *testing.T) {
		got := Repeat("a", 2)
		want := "aa"

		assertCorrectRepetition(t, got, want)
	})

	t.Run("repeating 0 times character 'a'", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		assertCorrectRepetition(t, got, want)
	})
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}