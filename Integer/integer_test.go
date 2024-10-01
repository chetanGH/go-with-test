package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(002, 002)
	expected := 4

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}

func ExampleAdd() {
	sum := Add(001, 001)
	fmt.Println(sum)
}
