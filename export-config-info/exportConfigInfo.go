package export_config_info

import (
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
	write, file := util.GetFileWriter(outPath)
	defer file.Close()
	write.WriteString(`using MxxM.GameClient;
using UnityEngine;

public partial class DataTableManager
{
    private TextAsset textAsset;

    private void InitData()
    {`)
	write.WriteString("\n")
	for _, configName := range allConfigName {
		infoTemp := fmt.Sprintf("        textAsset = Context.Game.Loader.LoadAsset<TextAsset>(\"Assets/DevHere/Datas/Json/%s.json\");\n", configName)
		write.WriteString(infoTemp)
		infoTemp = fmt.Sprintf("        T%sHelper.InitData(textAsset.text);\n", configName)
		write.WriteString(infoTemp)
	}
	write.WriteString(`    }
}`)
	write.Flush()
}
