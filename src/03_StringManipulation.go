package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

/*
separate arguments by space, indent them, and print each index
then urlencode each word an join them
*/

func main() {
	var accum = ""
	for i, arg := range os.Args[1:] {
		for j, word := range strings.Fields(arg) {
			fmt.Printf(strings.Repeat(" ", j)+"[%d][%d] → %s\n", i, j, word)
			accum += strings.ToLower(url.PathEscape(word)) + "%20"
		}
	}
	fmt.Println(accum)
}
