package dna

import (
	"reflect"
	"testing"
)

func TestCounts(t *testing.T) {
	for _, tc := range testCases {
		dna := DNA(tc.strand)
		s, err := dna.Counts()
		switch {
		case tc.errorExpected:
			if err == nil {
				t.Fatalf("FAIL: %s\nCounts(%q)\nExpected error\nActual: %#v",
					tc.description, tc.strand, s)
			}
		case err != nil:
			t.Fatalf("FAIL: %s\nCounts(%q)\nExpected: %#v\nGot error: %q",
				tc.description, tc.strand, tc.expected, err)
		case !reflect.DeepEqual(s, tc.expected):
			t.Fatalf("FAIL: %s\nCounts(%q)\nExpected: %#v\nActual: %#v",
				tc.description, tc.strand, tc.expected, s)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func benchmarkCounts(strand string, t *testing.B) {
	for n := 0; n < t.N; n++ {
		dna := DNA(strand)
		dna.Counts()
	}
}

func BenchmarkCounts1(b *testing.B) { benchmarkCounts("ACGT", b) }
func BenchmarkCounts2(b *testing.B) { benchmarkCounts("ACGTACGT", b) }
func BenchmarkCounts3(b *testing.B) { benchmarkCounts("ACGTACGTACGT", b) }
func BenchmarkCounts4(b *testing.B) { benchmarkCounts("ACGTACGTACGTACGT", b) }
func BenchmarkCounts5(b *testing.B) { benchmarkCounts("ACGTACGTACGTACGTACGT", b) }
