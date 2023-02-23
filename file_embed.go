package main

import (
	"embed"
	"io"
	"os"
	"path/filepath"
)

// 通过go:embed 嵌入的资源
func CopyDirectoryEmbed(scrDir, dest string, local embed.FS) error {
	// 从local获取目录
	entries, err := local.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			CopyDirectoryEmbed(sourcePath, destPath, local)
		} else {
			CopyEmbed(sourcePath, destPath, local)
		}
	}
	return nil
}

// 拷贝嵌入的资源到本地
func CopyEmbed(srcFile, dstFile string, local embed.FS) error {

	println("craete file ", dstFile)
	CreateDir(dstFile)

	out, err := os.Create(dstFile)

	if err != nil {
		println(err.Error())
		return err
	}
	defer out.Close()
	in, err := local.Open(srcFile)

	defer func() {
		if in != nil {
			in.Close()
		}
	}()

	if err != nil {
		return err
	}
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return nil
}

func CreateDir(path string) {
	dir, _ := filepath.Split(path)
	os.MkdirAll(dir, os.ModePerm)
}
