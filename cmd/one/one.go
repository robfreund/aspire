package main

import (
	"flag"
	"fmt"
	. "github.com/robfreund/aspire/aspire"
)

var (
	w string
)

func init() {
	flag.StringVar(&w, "w", "Aspiration.com", "word to perform logic on")
}
func main() {
	flag.Parse()
	fmt.Println(CapitalizeEveryThirdAlphanumericChar(w))
}
