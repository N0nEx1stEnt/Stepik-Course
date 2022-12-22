package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

/*type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser struct{}

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}
*/
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}

func main() {
	t := rover{}
	fmt.Println(t.talk())
	shout(t)
}
