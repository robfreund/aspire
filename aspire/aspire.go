package aspire

import (
	"fmt"
	"strings"
	. "unicode"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	ct := 0
	var sb strings.Builder
	for i, v := range s {
		if IsDigit(v) || IsLetter(v) {
			ct++
		}
		if ct%3 == 0 && i > 0 {
			v = ToUpper(v)
			ct = 0
		} else {
			v = ToLower(v)
		}
		sb.WriteRune(v)
	}
	return sb.String()
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

type Skipper struct {
	skip  int
	value []rune
	ct    *int
}

func NewSkipString(skip int, value string) Skipper {
	return Skipper{
		skip:  skip,
		value: []rune(value),
		ct:    new(int),
	}
}

func (s Skipper) TransformRune(pos int) {
	r := s.value[pos]
	r = ToLower(r)
	if IsDigit(r) || IsLetter(r) {
		*s.ct++
	}
	if *s.ct%s.skip == 0 && pos > 0 {
		r = ToUpper(r)
		*s.ct = 0
	}
	s.value[pos] = r
}

func (s Skipper) GetValueAsRuneSlice() []rune {
	return s.value
}

func (s Skipper) String() string {
	return fmt.Sprintf("%s", string(s.value))
}
