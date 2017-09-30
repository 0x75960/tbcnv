package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var targetTypes = []string{
	"tsv",
	"csv",
	"json",
	"yaml",
}

var title = []string{}

var targetFlags = map[string]bool{}
var toBeDisplayed = map[string]bool{}

var delim = flag.String("d", "\t", "デリミタを指定します。")
var outputType = flag.String("t", "tsv", "出力形式を指定します。指定可能な形式: "+strings.Join(targetTypes, ","))
var filterTitle = flag.String("f", "", "出力する列名を指定します。指定されない場合、すべての列を出力します。")

func init() {
	for _, item := range targetTypes {
		targetFlags[item] = true
	}
}

func main() {
	flag.Parse()

	var data []map[string]string

	if flag.NArg() == 0 {
		data = Load(os.Stdin)
	} else {

		fIdx := flag.NArg() - 1
		filename := flag.Arg(fIdx)

		f, err := os.Open(filename)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		data = Load(f)
	}

	switch *outputType {
	case "tsv":
		PrintTSV(data)
	case "csv":
		PrintCSV(data)
	case "json":
		PrintJson(data)
	case "yaml":
		PrintYaml(data)
	default:
		log.Fatalln("not supported output type.")
	}
}
