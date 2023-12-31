package main

import (
	"log"
	"testing"
)

/*
Usefull notes(page 100):
	Go provides a built-in testing framework that simplifies the process of testing Go code. Here is
	an overview of the file structure, naming conventions, and how testing works in Go:
	File Structure and Naming Conventions:
	1. Test files: Test files live in the same package as the code they test. They are named with
	the following convention: `filename_test.go` where filename is the name of the file that
	contains the code you want to test.
	2. Test functions: Test functions must start with the word `Test` followed by a word that starts
	with a capital letter. The signature of a test function is `func TestXxx(*testing.T)`, where Xxx
	does not start with a lowercase letter.
Unit tests:
	A unit test is a type of software testing where individual components or units of a
	software are tested. The purpose is to validate that each unit of the software performs as
	designed. A unit is the smallest testable part of any software. It usually has one or a few
	inputs and usually a single output.
	In Go programming language, a unit test usually tests a single function, method, or struct. The
	goal is to confirm that the behavior is correct.
	Unit tests in Go are typically written using the built-in testing package, `testing`. This
	package does not require any third-party libraries, but it's somewhat limited compared to
	some other languages testing tools. Despite its simplicity, it has enough features to construct
	effective unit tests.
	When you ask if this differs from a "test" in Go, it's worth clarifying that a "unit test" is a subset
	of "test". Tests in software can take many forms such as unit tests, integration tests,
	functional tests, system tests, etc.
	An "Integration Test", for example, in contrast to a "Unit Test", would test the interaction
	between multiple components of the system, to ensure they work together correctly.
	So, to summarize, in Go, a unit test is just one kind of test you can conduct, focused on
	verifying the correct behavior of a small, isolated piece of functionality. Other types of tests
	have different focuses and may require different tools or techniques.

*/

func TestAdd(t *testing.T) {
	expected := 22
	total := Add(10, 12)
	if total != expected {
		t.Errorf("Function 'Add' was problematic. It gives %d instead of %d.", total, expected)
	}
}

func TestSubstract(t *testing.T) {
	expected := 5
	result := Substract(10, 5)
	if result != expected {
		t.Errorf("Function 'Substract' was problematic. It gives %d instead of %d.", result, expected)
	}
}

func TestDoMath(t *testing.T) {
	expected := 12
	result := doMath(10, 2, Add)
	if result != expected {
		t.Errorf("Function 'doMath' was problematic. It gives %d instead of %d.", result, expected)
	}
}

func TestDisplayMyName(t *testing.T) {
	expected := "Panos"
	result := displayMyName("Efraim")
	if result != expected {
		log.Fatalf("Error: Function 'TestDisplayMyName' was problematic. It gives %s instead of %s.", result, expected)
	}

}
