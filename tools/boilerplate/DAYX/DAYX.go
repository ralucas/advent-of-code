package DAYX

import (
	"log"

	"github.com/ralucas/advent-of-code/pkg/utils"
)

// TODO: Alter this for actual implementation
func PrepareData(filepath string) []string {
	if filepath == "" {
		log.Fatalf("Missing input file")
	}
	data := utils.ReadFileToArray(filepath, "\n")

	return data
}
