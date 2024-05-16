package main

import (
	"regexp"
)

func generateWikiTest(g int) func() []segment {

	words := regexp.MustCompile("\\s+").Split(extractText(), -1)

	return func() []segment {
		segments := make([]segment, g)
		for i := 0; i < g; i++ {
			segments[i] = segment{cleanWiki(words), ""}
		}
		return segments
	}

}
