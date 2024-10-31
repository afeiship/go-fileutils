package nx

import (
	"testing"

	"github.com/afeiship/go-fileutils"
)

func TestReadFile(f *testing.T) {
	res := fileutils.ReadFile("../README.md")
	f.Log(string(res))
}
