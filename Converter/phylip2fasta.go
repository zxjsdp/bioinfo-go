package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
	"io/ioutil"
)

var (
	print = fmt.Println
)

type Species struct {
	name, sequence string
}

func extractSpeciesFromPhylipFile(phylipFile string) []Species {
	var species []Species
	var line string
	file, err := os.Open(phylipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			elements := strings.Fields(line)
			if len(elements) != 2 {
				log.Fatal("Invalid Phylip file!")
			}
			if !isStringIntType(elements[1]) {
				species = append(species, Species{elements[0], elements[1]})
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return species;
}

func generateFasta(species []Species, outputFile string) {
	if species == nil {
		log.Panic("No species!")
		return
	}
	var fastaLines []string

	for _, each := range species {
		fastaLines = append(fastaLines, fmt.Sprintf(">%s", each.name))
		fastaLines = append(fastaLines, each.sequence + "\n")
	}
	fastaContent := strings.Join(fastaLines, "\n")
	err := ioutil.WriteFile(outputFile, []byte(fastaContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func isStringIntType(stringToCheck string) bool {
	if _, err := strconv.Atoi(stringToCheck); err == nil {
		return true;
	}
	return false;
}

func checkPhylipFileExists(fileName string) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: /path/to/phylip2fasta input.phy output.fasta")
	}
	checkPhylipFileExists(args[0])
	fmt.Printf("Convert Phylip to Fasta:\n  %s => %s\n", args[0], args[1])
	generateFasta(extractSpeciesFromPhylipFile(args[0]), args[1])
}
