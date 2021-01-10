package easyaccess

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPrintStringToConsole(t *testing.T) {
	expectedString := "test"
	rescueStdout, r, w := openStdoutPipe()
	Ln(expectedString)
	out := getConsolePipeOutput(r, rescueStdout, w)
	actualString := strings.Trim(string(out), "\n")
	assert.True(t, string(actualString) == expectedString)
}

func TestPrintIntToConsole(t *testing.T) {
	expectedInt := 120
	rescueStdout, r, w := openStdoutPipe()
	Ln(expectedInt)
	out := getConsolePipeOutput(r, rescueStdout, w)
	actualString := strings.Trim(string(out), "\n")
	actualInt, _ := strconv.Atoi(actualString)
	assert.True(t, actualInt == expectedInt)
}

func TestPrintArrayToConsole(t *testing.T) {
	input := []int{1, 2, 3, 100, 50, 98464}
	expectedArrayString := "[1 2 3 100 50 98464]"
	rescueStdout, r, w := openStdoutPipe()
	Ln(input)
	out := getConsolePipeOutput(r, rescueStdout, w)
	actualString := strings.Trim(string(out), "\n")
	assert.True(t, actualString == expectedArrayString)
}
func getConsolePipeOutput(r *os.File, rescueStdout *os.File, w *os.File) []byte {
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	return out
}

func openStdoutPipe() (*os.File, *os.File, *os.File) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return rescueStdout, r, w
}
