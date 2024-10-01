package arrays

import "testing"

func TestSumArr(t *testing.T) {
	// var arr = []int{1, 2, 3, 4, 5}
	// sum := SumArr(arr)
	// expected := 15

	// if sum != expected {
	// 	t.Errorf("Expected %d, got %d", expected, sum)
	// }

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numb := []int{1, 2, 3, 4, 5}
		got := SumArr(numb)
		want := 15
		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numb := []int{1, 2, 3}
		got := SumArr(numb)
		want := 6
		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})
}

// func BenchmarkSumArr(b *testing.B) {
// 	arr := []int{1, 2, 3, 3, 4, 4, 5, 4, 3, 4, 4, 4, 5, 4, 3, 4, 5, 6, 4, 3, 3, 5, 6, 6, 6, 3, 3, 4, 5, 6, 44, 3, 3, 4, 5, 5, 4, 3, 2, 3, 54, 6, 6, 4, 56, 6, 6, 6, 6, 3, 2, 33, 3}
// 	for i := 0; i < b.N; i++ {
// 		SumArr(arr)
// 	}
// }
