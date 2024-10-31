package nx

import (
	"strings"
	"testing"

	"github.com/afeiship/go-fileutils"
)

func TestReadFile(f *testing.T) {
	res := fileutils.ReadFile("../README.md")
	if strings.Contains(string(res), ("# go-fileutils")) == false {
		f.Error("not contains")
	}
}

func TestRmRf(t *testing.T) {
	err := fileutils.RmRf("../.tmp")
	if err != nil {
		t.Error(err)
	}
}

func TestGetContents(t *testing.T) {
	res := fileutils.GetContents("../README.md")
	if strings.Contains(res, ("# go-fileutils")) == false {
		t.Error("not contains")
	}
}
