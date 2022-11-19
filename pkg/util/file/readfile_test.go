package file_test

import (
	"testing"

	fileutil "github.com/ralucas/advent-of-code/pkg/util/file"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	rf := fileutil.ReadFile("../../../test/testdata/2020/3/test-input1.txt")
	assert.NotNil(t, rf)
}

func TestReadFileToArray(t *testing.T) {
	rfArray := fileutil.ReadFileToArray("../../../test/testdata/2020/3/test-input1.txt", "\n")
	assert.NotNil(t, rfArray)
	assert.IsType(t, rfArray, []string{})
	assert.True(t, len(rfArray) > 1)
}
