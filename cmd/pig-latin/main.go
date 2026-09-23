package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/computerscienceforhumans/pig-latin-go/internal/translator"
)

var defaultInFile = "small-sample-text.txt"
var defaultOutFile = "result.txt"

func main() {
	var readFile = flag.String("readFile", defaultInFile, "The file you want to translate.")
	var writeFile = flag.String("writeFile", defaultOutFile, "The file to write the result to.")
	var language = flag.String("language", "pigLatin", "Options: pigLatin, oldEnglish.")
	flag.Parse()
	defaultInFile = *readFile
	defaultOutFile = *writeFile
	var lang = *language

	inFile, err := os.Open(defaultInFile)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(defaultOutFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	var i = 0
	for scanner.Scan() {
		var translated = ""
		switch lang {
		case "pigLatin":
			translated = translator.TranslatePig(scanner.Text()) + "\n"
		case "oldEnglish":
			translated = translator.TranslateOldEnglish(scanner.Text()) + "\n"
		default:
			fmt.Println("???")

		}

		_, err = outFile.WriteString(translated)
		if err != nil {
			log.Fatal(err)
		}

		printProgress(i)
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("\n\nFinished.")
}

func printProgress(i int) {
	if i%900 == 0 {
		fmt.Print("\n")
	}
	if i%300 == 0 {
		fmt.Print(" ")
	}
	if i%100 == 0 {
		fmt.Print(".")
	}
}
