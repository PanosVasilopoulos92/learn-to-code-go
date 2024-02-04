package dog

import "testing"

func TestYears(t *testing.T) {
	expected := 70
	actual := Years(10)
	if actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func TestYearsTwo(t *testing.T) {
	expected := 70
	actual := YearsTwo(10)
	if actual != expected {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

/*
Remember to BET
-Benchmark --> BenchmarkCat(b *testing.B)
-Example   --> ExampleCat()
-Test	   --> TestCat(t *testing.T)

Commands
godoc -http=:8080

go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
*/
