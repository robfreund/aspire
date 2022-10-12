package aspire

import (
	"fmt"
	"strings"
	. "unicode"
)

// CapitalizeEveryThirdAlphanumericChar will capitalize every 3 alpha char and return string
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	count := 0
	var sb strings.Builder
	for _, v := range s {

		// is alpha numeric inc count
		if IsDigit(v) || IsLetter(v) {
			count++
		}
		// if count is reached upper
		if count%3 == 0 {
			v = ToUpper(v)
			count = 0
		} else {
			v = ToLower(v)
		}
		sb.WriteRune(v)
	}
	return sb.String()
}

// Interface for transforming rune slice
type Interface interface {
	// TransformRune will transform at pos
	TransformRune(pos int)
	// GetValueAsRuneSlice returns the rune slice
	GetValueAsRuneSlice() []rune
}

// MapString will perform transform on rune slice
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

// Skipper implements Interface and will perform skip logic on the provided value
type Skipper struct {
	skip  int
	value []rune
	ct    *int
}

// NewSkipString is a skipper that will transform by capitalizing skip [alpha , numeric] characters in value
func NewSkipString(skip int, value string) Skipper {
	return Skipper{
		skip:  skip,
		value: []rune(value),
		ct:    new(int),
	}
}

// isAlphaNumeric checks if r is alpha numeric
func isAlphaNumeric(r rune) bool {
	return IsDigit(r) || IsLetter(r)
}

// TransformRune will perform transform at pos
func (s Skipper) TransformRune(pos int) {
	r := s.value[pos]
	r = ToLower(r)
	if isAlphaNumeric(r) {
		*s.ct++
	}
	// if we reached skip count upper
	if *s.ct%s.skip == 0 && pos > 0 {
		r = ToUpper(r)
		*s.ct = 0
	}
	s.value[pos] = r
}

// GetValueAsRuneSlice will return the rune slice of skipper value
func (s Skipper) GetValueAsRuneSlice() []rune {
	return s.value
}

// String is the stringer implementation
func (s Skipper) String() string {
	return fmt.Sprintf("%s", string(s.value))
}
