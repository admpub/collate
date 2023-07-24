package collate

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	if len(SupportedLangs()) == 0 {
		t.Fatal("expected something greater than zero")
	}
	less := IndexString("ENGLISH_CI")
	if !less("a", "b") {
		t.Fatal("expected true, got false")
	}
}

func ExampleIndexString() {
	var nameA = "Miller"
	var nameB = "anderson"
	less := IndexString("ENGLISH_CI")
	fmt.Printf("%t\n", less(nameA, nameB))
	fmt.Printf("%t\n", less(nameB, nameA))
	// Output:
	// false
	// true
}

func ExampleSpanish() {
	less := IndexString("SPANISH_CI")
	fmt.Printf("%t\n", less("Hola", "hola"))
	fmt.Printf("%t\n", less("hola", "Hola"))
	// Output:
	// false
	// false
}
