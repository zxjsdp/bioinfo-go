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
		log.Fatal("Usage: /path/to/fasta2phylip input.fasta output.phy")
	}
	utils.CheckFileExists(args[0])
	fmt.Printf("Convert FASTA to Phylip:\n  %s => %s\n", args[0], args[1])
	converters.GeneratePhylip(converters.ExtractSpeciesFromFastaFile(args[0]), args[1])
}
