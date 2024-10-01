package iteration

import "testing"

func TestIteration(t *testing.T) {
	char := Iteration("a")
	wanted := "aaaaa"
	if char != wanted {
		t.Errorf("Expected %s, Got %s", wanted, char)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a")
	}
}
