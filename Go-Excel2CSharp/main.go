package main

import (
	"excel2csharp/excel2json"
	export_config_info "excel2csharp/export-config-info"
	"excel2csharp/tocsharp"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var waitGroup sync.WaitGroup

//  excel目录路径
var excelDirPath = flag.String("excelPath", "", "input excel path")

// 输出json目录路径
var outJsonDirPath = flag.String("outJsonPath", "", "out json path")

//  输出C#目录路径
var outCSharpDirPath = flag.String("outCSharpPath", "", "out csharp path")

// 输出json反序列化C#文件
var outCSharpConfigPath = flag.String("outCSharpConfigPath", "", "out csharp config path")

func toJsonInfo(file fs.FileInfo, excelDirPath, outJsonDirPath, outCSharpDirPath string) {
	defer waitGroup.Done()
	if file.IsDir() {
		return
	}
	if strings.HasPrefix(file.Name(), "~$") {
		return
	}
	excelPath := filepath.Join(excelDirPath, file.Name())

	fileSuffix := path.Ext(excelPath)
	filenameWithSuffix := filepath.Base(excelPath)
	fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	outCSharpPath := filepath.Join(outCSharpDirPath, fmt.Sprintf("T%s.cs", fileName))

	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("打开excel失败", err)
		return
	}
	rows := f.GetRows(fileName)

	if strings.HasSuffix(excelPath, "CommonConfig.xlsx") {

		return
	}

	outJsonPath := filepath.Join(outJsonDirPath, fileName+".json")
	fmt.Println("toPath ", outJsonPath)
	excel2json.ExportJsonConfig(rows, outJsonPath)

	tocsharp.ExportCSharpConfig(rows, outCSharpPath)

	export_config_info.SaveConfigInfo(fileName)
}

func main() {
	flag.Parse()
	excelDirPath := *excelDirPath
	outJsonDirPath := *outJsonDirPath
	outCSharpDirPath := *outCSharpDirPath
	outCSharpConfigPath := *outCSharpConfigPath

	//excelDirPath = "E:\\project\\excel2csharp\\UnityExcel2CsharpDemo\\DataTable"
	//outJsonDirPath = "E:\\project\\excel2csharp\\UnityExcel2CsharpDemo\\Assets\\Resources\\Json"
	//outCSharpDirPath = "E:\\project\\excel2csharp\\UnityExcel2CsharpDemo\\Assets\\Scripts\\DataTable"
	//outCSharpConfigPath = "E:\\project\\excel2csharp\\UnityExcel2CsharpDemo\\Assets\\Scripts\\DataTable\\DataTableManager.Init.cs"

	fmt.Println("excelPath", excelDirPath)

	files, _ := ioutil.ReadDir(excelDirPath)
	for _, file := range files {
		waitGroup.Add(1)
		go toJsonInfo(file, excelDirPath, outJsonDirPath, outCSharpDirPath)
	}
	waitGroup.Wait()
	export_config_info.ExportCSharpInit(outCSharpConfigPath)
}
