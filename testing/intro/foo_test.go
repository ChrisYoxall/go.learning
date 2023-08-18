// Example taken from https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
// More information in that blg post on skipping slow running integration tests, the Cleanup and Helper methods and
// code coverage reports.

package main

import "testing"

// TestFoo tests the Foo function. This is the recommended naming convention for tests.
func TestFoo(t *testing.T) {
	result := Foo(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

// TestFooTableDriven tests the Foo function using a table-driven approach.
func TestFooTableDriven(t *testing.T) {

	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo(tt.input); got != tt.want {
				t.Errorf("Foo() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFooParallel runs two tests in parallel. The GOMAXPROCS environment variable environment defines how many
// tests can run in parallel at one time, and by default this number is equal to the number of CPUs. To avoid
// duplication use the table-driven approach.
func TestFooParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Foo(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Foo(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}

// FuzzFoo is a fuzz test. It must start with 'Fuzz' and take a testing.F object. To run from the command line
// do 'go test -fuzz FuzzFoo -run ^$' which will run tests matching the regex 'FuzzFoo' and only run those tests.
// When running in a CI pipeline use the '-fuzztime' flag to specify how long to run the fuzz tests as by default
// they just keep running. For example 'go test -fuzz FuzzFoo -run ^$ -fuzztime 10s'.
func FuzzFoo(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {

		// Could just call Foo(a) here to keep running to try to find panic situations.
		result := Foo(a)

		if a%3 == 0 && result != "Foo" {
			t.Errorf("Result was incorrect. With: %d got: %s, want: %s.", a, result, "Foo")
		}
	})
}
