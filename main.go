package main

import (
	"bearded-ironman/rom"
	"fmt"
	"strings"
)

func main() {
	var lines int
	_, err := fmt.Scanf("%d", &lines)
	if err != nil {
		panic("could not read number of lines")
	}
	for ii := 0; ii < lines; ii++ {
		var text string
		_, err := fmt.Scanf("%s", &text)
		if err != nil {
			panic("could not read equation")
		}
		tmp := strings.Split(text[:len(text)-1], "+")
		ans := rom.Add(tmp[0], tmp[1])
		fmt.Printf("%v%v\n\n", text, ans)
	}
}
