package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"io/ioutil"
	"unicode"
)

var (
	print = fmt.Println
	fileInfo *os.FileInfo
	err      error
)

type Species struct {
	name, sequence string
}

func extractSpecies(fastaFile string) []Species {
	var species []Species
	var title, sequence, line string
	file, err := os.Open(fastaFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line != "" {
			if strings.HasPrefix(line, ">") {
				if sequence != "" {
					species = append(species, Species{title, sequence})
					sequence = ""
				}
				title = strings.TrimLeft(line, ">")
				title = replaceBlankChars(title)
			} else {
				sequence = sequence + strings.TrimSpace(line)
			}
		}
	}
	if title != "" && sequence != "" {
		species = append(species, Species{title, sequence})
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return species
}

func generatePhylip(species []Species) {
	if species == nil {
		log.Panic("No species!")
		return
	}
	var phylipLines []string
	speciesNum := len(species)
	charNum := len(species[0].sequence)
	// Add (speciesNum  charNum) to top line of output Phylip file
	phylipLines = append(phylipLines, fmt.Sprintf("%d  %d", speciesNum, charNum))

	for _, each := range species {
		phylipLines = append(phylipLines, each.name + "\t\t" + each.sequence)
	}
	phylipContent := strings.Join(phylipLines, "\n")

	err := ioutil.WriteFile("out.phy", []byte(phylipContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func replaceBlankChars(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '_'
		}
		return r
	}, str)
}

func checkFileExists(fileName string) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	// log.Println(fileInfo)
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Usage: ./FastaToPhylip-Go file.fasta")
	}
	checkFileExists(args[0])
	generatePhylip(extractSpecies(args[0]))
}
