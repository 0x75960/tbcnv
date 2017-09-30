package main

import (
	"io"
	"strings"

	"github.com/0x75960/liner"
)

func Load(f io.Reader) (d []map[string]string) {

	d = []map[string]string{}

	if *filterTitle != "" {
		for _, c := range strings.Split(*filterTitle, ",") {
			toBeDisplayed[c] = true
		}
	}

	firstLine := true

	for line := range liner.LinesIn(f) {

		items := strings.Split(line, *delim)
		da := map[string]string{}

		if firstLine {
			title = items
			firstLine = false

			if *filterTitle == "" {
				for _, column := range title {
					toBeDisplayed[column] = true
				}
			}

			continue
		}

		for idx, item := range items {

			columnName := title[idx]
			if toBeDisplayed[columnName] == false {
				continue
			}

			da[columnName] = item

		}

		d = append(d, da)

	}

	return
}
