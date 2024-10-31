package nx

import (
	"os"
	"path"
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

func TestLn(t *testing.T) {
	fileutils.Ln("../.tmp", "../.tmp3")
	if fileutils.IsLink("../.tmp3") == false {
		t.Error("not a link")
	}
}

func TestBasename(t *testing.T) {
	res := fileutils.Basename("../README.md")
	if res != "README.md" {
		t.Error("not equal")
	}
}

func TestDirname(t *testing.T) {
	res := fileutils.Dirname("../README.md")
	t.Log(res)
	if res != ".." {
		t.Error("not equal")
	}
}
func TestAbsname(t *testing.T) {
	res := fileutils.Absname("../README.md")
	cwd, _ := os.Getwd()
	t.Log(res)
	if res != path.Join(cwd, "../README.md") {
		t.Error("not equal")
	}
}
