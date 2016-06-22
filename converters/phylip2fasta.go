package converters

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"io/ioutil"

	"github.com/zxjsdp/Bioinfo-Go/utils"
)

func ExtractSpeciesFromPhylipFile(phylipFile string) []Species {
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
			if !utils.IsStringIntType(elements[1]) {
				species = append(species, Species{elements[0], elements[1]})
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return species;
}

func GenerateFasta(species []Species, outputFile string) {
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



