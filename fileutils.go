package fileutils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(filename string) []byte {
	res, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return res
}

func WriteFile(filename string, data []byte) {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}

func ReadContents(filename string) string {
	res := ReadFile(filename)
	return string(res)
}

func WriteContents(filename string, data string) {
	WriteFile(filename, []byte(data))
}

func RmRf(path string) error {
	return os.RemoveAll(path)
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func Mkdir(path string) error {
	return os.MkdirAll(path, 0755)
}

func Chdir(path string) error {
	return os.Chdir(path)
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// CopyDir 递归复制目录 src 到 dst
func CopyDir(src string, dst string) error {
	// 获取源目录信息
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// 确保源路径是目录
	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	// 创建目标目录
	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	// 遍历源目录
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算目标路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// 创建子目录
			return os.MkdirAll(destPath, info.Mode())
		} else {
			// 复制文件
			return CopyFile(path, destPath)
		}
	})

	return err
}

// copyFile 复制单个文件
func CopyFile(srcFile string, destFile string) error {
	// 打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 复制内容
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	// 获取文件权限并应用到目标文件
	srcInfo, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	return os.Chmod(destFile, srcInfo.Mode())
}

// 创建软链接
func Ln(src string, dst string) bool {
	err := os.Symlink(src, dst)
	return err != nil
}

func IsLink(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

func Basename(path string) string {
	return filepath.Base(path)
}

func Dirname(path string) string {
	return filepath.Dir(path)
}

func Extname(path string) string {
	return filepath.Ext(path)
}

func Absname(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	return abs
}

func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func Join(elem ...string) string {
	return filepath.Join(elem...)
}
