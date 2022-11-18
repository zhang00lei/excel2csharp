package export_config_info

import (
	"excel2csharp/export_type"
	"excel2csharp/util"
	"fmt"
	"sync"
)

var allConfigName []string

var mutex sync.Mutex

func SaveConfigInfo(configName string) {
	mutex.Lock()
	allConfigName = append(allConfigName, configName)
	mutex.Unlock()
}

func ExportCSharpInit(outPath string) {
	write, file := util.GetFileWriter(outPath, export_type.CSharp)
	defer file.Close()
	write.WriteString(`using UnityEngine;

public partial class DataTableManager
{
    private TextAsset textAsset;

    private void InitData()
    {`)
	write.WriteString("\n")
	for _, configName := range allConfigName {
		infoTemp := fmt.Sprintf("        textAsset = Resources.Load<TextAsset>(\"Json/%s\");\n", configName)
		write.WriteString(infoTemp)
		infoTemp = fmt.Sprintf("        T%sHelper.InitData(textAsset.text);\n", configName)
		write.WriteString(infoTemp)
	}
	write.WriteString(`    }
}`)
	write.Flush()
}
