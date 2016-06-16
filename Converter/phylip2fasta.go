package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
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

func isStringIntType(stringToCheck string) bool {
	if _, err := strconv.Atoi(stringToCheck); err == nil {
		return true;
	}
	return false;
}

func main() {
	print(extractSpeciesFromPhylipFile("data/test.phy"))
}
