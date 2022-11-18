package util

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetFileWriter(outPath string) (*bufio.Writer, *os.File) {
	if _, err := os.Stat(outPath); err == nil {
		os.Remove(outPath)
	}
	file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return nil, nil
	}
	write := bufio.NewWriter(file)
	write.WriteString("// this file is generate by tools,do not modify it.\n")
	return write, file
}

//
// GetFileNameNoExtensionByPath
//  @Description: 获取文件名，后缀
//  @param filePath 文件路径
//  @return string,string 文件名 后缀
//
func GetFileNameNoExtensionByPath(filePath string) (string, string) {
	fileSuffix := path.Ext(filePath)
	filenameWithSuffix := filepath.Base(filePath)
	fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return fileName, fileSuffix
}
