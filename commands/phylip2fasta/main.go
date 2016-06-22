package main

import (
	"os"
	"log"
	"fmt"

	"github.com/zxjsdp/Bioinfo-Go/converters"
	"github.com/zxjsdp/Bioinfo-Go/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Usage: /path/to/phylip2fasta input.phy output.fasta")
	}
	utils.CheckFileExists(args[0])
	fmt.Printf("Convert Phylip to Fasta:\n  %s => %s\n", args[0], args[1])
	converters.GenerateFasta(converters.ExtractSpeciesFromPhylipFile(args[0]), args[1])
}
