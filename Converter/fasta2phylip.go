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

const freeSpaceNum int = 4

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

func generatePhylip(species []Species, outputFile string) {
	if species == nil {
		log.Panic("No species!")
		return
	}
	var phylipLines []string
	longestNameLength := getLongestNameLength(species)
	speciesNum := len(species)
	charNum := len(species[0].sequence)

	// Add (speciesNum  charNum) to top line of output Phylip file
	phylipLines = append(phylipLines, fmt.Sprintf("%d  %d", speciesNum, charNum))

	// Add (species  sequence) to each line
	for _, each := range species {
		spacesForCurrentSpecies := generateSpacesForAlignment(longestNameLength, len(each.name))
		phylipLines = append(phylipLines, each.name + spacesForCurrentSpecies + each.sequence)
	}
	phylipContent := strings.Join(phylipLines, "\n")

	err := ioutil.WriteFile(outputFile, []byte(phylipContent), 0644)
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
}

func getLongestNameLength(species []Species) int {
	longestNameLength := 0
	for _, each := range species {
		if len(each.name) > longestNameLength {
			longestNameLength = len(each.name)
		}
	}
	return longestNameLength
}

func generateSpacesForAlignment(longestNameLength, currentNameLen int) string {
	spaceNum := longestNameLength - currentNameLen + freeSpaceNum
	return strings.Repeat(" ", spaceNum)
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: ./FastaToPhylip-Go input.fasta output.phy")
	}
	checkFileExists(args[0])
	fmt.Printf("Convert FASTA to Phylip:\n  %s => %s\n", args[0], args[1])
	generatePhylip(extractSpecies(args[0]), args[1])
}
