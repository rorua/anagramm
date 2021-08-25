package anagramm_service

import "testing"

func Test_find(t *testing.T) {

	dict := []string{"foobar", "aabb", "baba", "boofar", "test"}

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "test 1", input: "foobar", expected: []string{"foobar", "boofar"}},
		{name: "test 2", input: "raboof", expected: []string{"foobar", "boofar"}},
		{name: "test 3", input: "abba", expected: []string{"aabb", "baba"}},
		{name: "test 4", input: "test", expected: []string{"test"}},
		{name: "test 5", input: "qwerty", expected: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := findAll(dict, tt.input)
			if len(output) != len(tt.expected) {
				t.Fatalf("expected len %d, got %d", len(dict), len(tt.expected))
			}

			for i, v := range tt.expected {
				if v != output[i] {
					t.Fatalf("expected len %d, got %d", len(dict), len(tt.expected))
				}
			}
		})
	}
}
