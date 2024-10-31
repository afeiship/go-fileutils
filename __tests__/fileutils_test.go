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

func TestReadContents(t *testing.T) {
	res := fileutils.ReadContents("../README.md")
	if strings.Contains(res, ("# go-fileutils")) == false {
		t.Error("not contains")
	}
}

func TestWriteFile(t *testing.T) {
	// mkdir
	fileutils.Mkdir("../.tmp")
	fileutils.WriteFile("../.tmp/test.txt", []byte("hello world"))
}

func TestCopyDir(t *testing.T) {
	fileutils.CopyDir("../.tmp", "../.tmp2")
	if fileutils.IsDir("../.tmp2") == false {
		t.Error("not a dir")
	}
}
