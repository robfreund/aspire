package main

import (
	"flag"
	"fmt"
	. "github.com/robfreund/aspire/aspire"
)

var (
	w string
	s int
)

func init() {
	flag.StringVar(&w, "w", "Aspiration.com", "word to perform logic on")
	flag.IntVar(&s, "s", 3, "skip length")
}
func main() {
	flag.Parse()
	s := NewSkipString(s, w)
	MapString(&s)
	fmt.Println(s)
}
