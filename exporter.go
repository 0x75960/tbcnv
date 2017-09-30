package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func PrintTable(d []map[string]string, sep string) {

	outTitle := []string{}

	for _, column := range title {
		if toBeDisplayed[column] == false {
			continue
		}

		outTitle = append(outTitle, column)
	}

	fmt.Println(strings.Join(outTitle, sep))

	for _, data := range d {
		line := []string{}

		for _, c := range outTitle {
			if val, ok := data[c]; ok {
				line = append(line, val)
			}
		}

		fmt.Println(strings.Join(line, sep))
	}
}

func PrintCSV(d []map[string]string) {
	PrintTable(d, ",")
}

func PrintTSV(d []map[string]string) {
	PrintTable(d, "\t")
}

func PrintJson(d []map[string]string) {
	json.NewEncoder(os.Stdout).Encode(d)
}

func PrintYaml(d []map[string]string) {
	buf, err := yaml.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(os.Stdout, bytes.NewBuffer(buf))
}
