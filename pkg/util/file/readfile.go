package file

import (
	"io/ioutil"
	"log"
	"strings"

	arrayutil "github.com/ralucas/advent-of-code/pkg/util/array"
)

func ReadFile(filepath string) string {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading in file %v", err)
	}

	return string(f)
}

func ReadFileToArray(filepath string, sep string) []string {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading in file %v", err)
	}

	return arrayutil.Filter(strings.Split(string(f), sep), func(v string) bool {
		return v != ""
	})
}
