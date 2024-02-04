package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	expected := "Welcome dear friend!"
	result := Greet("friend")
	if expected != result {
		t.Error("Expected outcome does not much with the actual outcome.")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("friend"))
	// Output:
	// Welcome dear friend!
}

// Benchmarking the function ('go test -bench <name of function>' or 'go test -bench .' for bench every function.)
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("friend")
	}
}
