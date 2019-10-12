package chapter03

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
}

func Example_report() {
	report("scruple")
	// Output:
	// U+2108	℈	SCRUPLE
	// 1 character found
}

func Test_CharName_String(t *testing.T) {
	want := "U+0041\tA\tLATIN CAPITAL LETTER A"
	cn := CharName{'A', "LATIN CAPITAL LETTER A"}
	got := fmt.Sprint(cn)
	assert.Equal(t, want, got)
}

func Test_scan(t *testing.T) {
	testCases := []struct {
		label string
		start rune
		end   rune
		want  []CharName
	}{
		{"A", 'A', 'B', []CharName{{'A', "LATIN CAPITAL LETTER A"}}},
		{"ABC", 'A', 'D', []CharName{
			{'A', "LATIN CAPITAL LETTER A"},
			{'B', "LATIN CAPITAL LETTER B"},
			{'C', "LATIN CAPITAL LETTER C"},
		}},
		{"unassigned", '\u0377', '\u037B', []CharName{
			{'\u0377', "GREEK SMALL LETTER PAMPHYLIAN DIGAMMA"},
			{'\u037A', "GREEK YPOGEGRAMMENI"},
		}},
		{"unnamed", '\x1E', '\x22', []CharName{
			{' ', "SPACE"},
			{'!', "EXCLAMATION MARK"},
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.label, func(t *testing.T) {
			got := scan(tc.start, tc.end)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_filter(t *testing.T) {
	testCases := []struct {
		start rune
		end   rune
		query []string
		want  []CharName
	}{
		{' ', unicode.MaxRune, []string{"MADEUPWORD"}, []CharName{}},
		{'\u2108', unicode.MaxRune, []string{"SCRUPLE"}, []CharName{
			{'\u2108', "SCRUPLE"}}},
		{'A', 'C', []string{"A"}, []CharName{
			{'A', "LATIN CAPITAL LETTER A"}}},
		{'?', 'C', []string{"LETTER", "CAPITAL"}, []CharName{
			{'A', "LATIN CAPITAL LETTER A"},
			{'B', "LATIN CAPITAL LETTER B"},
		}},
		{',', '/', []string{"MINUS"}, []CharName{
			{'-', "HYPHEN-MINUS"},
		}},
		{'?', 'C', []string{"LeTter", "capital"}, []CharName{
			{'A', "LATIN CAPITAL LETTER A"},
			{'B', "LATIN CAPITAL LETTER B"},
		}},
	}
	for _, tc := range testCases {
		t.Run(strings.Join(tc.query, "+"), func(t *testing.T) {
			sample := scan(tc.start, tc.end)
			got := filter(sample, tc.query)
			assert.Equal(t, tc.want, got)
		})
	}
}
