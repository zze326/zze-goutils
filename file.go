package zze_goutils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//
//  CopyFile
//  @Description: 拷贝单个文件
//  @param src 源路径
//  @param dst 目标路径
//  @return error 错误信息
//
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

//
//  CopyDir
//  @Description: 拷贝目录
//  @param src 源路径
//  @param dst 目标路径
//  @return error 错误信息
//
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

//
//
//  ReplaceStrInAllFile
//  @Description: 在指定目录下递归所有目录将文件名符合 patterns 规则的文件替换文件内容中的指定字符串为新字符串
//  @param dir 递归的根目录
//  @param old 要替换的字符串
//  @param new 替换后的字符串
//  @param patterns 文件名规则，如 *.yaml、*.json
//  @return error 错误信息
//
func ReplaceStrInAllFile(dir, old, new string, patterns ...string) error {
	return filepath.Walk(dir, refactorFunc(old, new, patterns))
}

func refactorFunc(old, new string, filePatterns []string) filepath.WalkFunc {
	return filepath.WalkFunc(func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !!fi.IsDir() {
			return nil
		}

		var matched bool
		for _, pattern := range filePatterns {
			var err error
			matched, err = filepath.Match(pattern, fi.Name())
			if err != nil {
				return err
			}

			if matched {
				read, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				log.Printf("Refactoring: %s\n", path)

				newContents := strings.Replace(string(read), old, new, -1)

				err = ioutil.WriteFile(path, []byte(newContents), 0)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

//
//  IsExist
//  @Description: 判断指定文件是否存在
//  @param path 文件路径
//  @return bool true 为存在
//
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
