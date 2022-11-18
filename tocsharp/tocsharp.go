package tocsharp

import (
	"bufio"
	"excel2csharp/export_type"
	"excel2csharp/util"
	"fmt"
	"strings"
)

//
// ExportCSharpConfig
//  @Description: 生成C#相关文件
//  @param excelData 表信息
//  @param outCSharpPath C#文件路径
//
func ExportCSharpConfig(excelData [][]string, outCSharpPath string) {
	write, file := util.GetFileWriter(outCSharpPath, export_type.CSharp)
	if file == nil {
		return
	}
	defer file.Close()
	fileName, _ := util.GetFileNameNoExtensionByPath(outCSharpPath)
	write.WriteString(`using System;
using System.Collections.Generic;
using UnityEngine;
`)

	write.WriteString("\n")
	write.WriteString(fmt.Sprintf("public class %s\n", fileName))
	write.WriteString("{\n")

	dataLength := len(excelData[3])
	for j := 0; j < dataLength; j++ {
		fieldName := excelData[2][j]
		if strings.HasSuffix(fieldName, "#") || fieldName == "" {
			continue
		}
		fieldType := excelData[3][j]
		fieldNotes := excelData[1][j]
		infoTemp := getFieldInfo(fieldName, fieldType, fieldNotes)
		write.WriteString(infoTemp)
	}
	write.WriteString("}\n\n")
	exportCSharpHelper(fileName, write)
	write.Flush()
}

func exportCSharpHelper(fileName string, write *bufio.Writer) {
	write.WriteString(fmt.Sprintf("public static class %sHelper\n", fileName))
	write.WriteString("{\n")
	write.WriteString(fmt.Sprintf("    private static List<%s> DataList;\n", fileName))
	write.WriteString("\n")
	write.WriteString(`    public static void InitData(string jsonStr)
    {
`)
	write.WriteString(fmt.Sprintf("        DataList = LitJson.JsonMapper.ToObject<List<%s>>(jsonStr);\n", fileName))
	write.WriteString(`        if (DataList == null || DataList.Count == 0)
        {
            Debug.LogError("反序列化异常");
        }
    }

`)
	write.WriteString(fmt.Sprintf("    public static List<%s> GetAll()", fileName))
	write.WriteString("\n")
	write.WriteString(`    {
        return DataList;
    }

`)
	write.WriteString(fmt.Sprintf("    public static %s GetById(int id)\n", fileName))
	write.WriteString(`    {
        var info = GetByCondition(x => x.Id == id);
        if (info == null || info.Count == 0)
        {
            return null;
        }

        return info[0];
    }

`)
	write.WriteString(fmt.Sprintf("    public static List<%s> GetByCondition(Predicate<%s> predicate)\n", fileName, fileName))
	write.WriteString(`    {
        return DataList.FindAll(predicate);
    }

`)
	write.WriteString(fmt.Sprintf("    public static %s GetOneByCondition(Predicate<%s> predicate)\n", fileName, fileName))
	write.WriteString(`    {
        var temp = GetByCondition(predicate);
        if (temp == null || temp.Count == 0)
        {
            return null;
        }

        return temp[0];
    }
}`)
}

//
// getFieldInfo
//  @Description: 获取字段详情
//  @param fieldName 字段名称
//  @param fieldType 字段类型
//  @param fieldNotes 字段注释
//  @return string 字段详情
//
func getFieldInfo(fieldName, fieldType, fieldNotes string) string {
	notes := `
    /// <summary>
    /// %s
    /// </summary>
`
	if strings.Contains(fieldNotes, "\n") {
		fieldNotes = strings.Replace(fieldNotes, "\n", "", -1)
	}
	notes = fmt.Sprintf(notes, fieldNotes)
	info := "    public %s %s { get; set; }\n"
	typeTemp := ""
	if fieldType == "int" {
		typeTemp = "int"
	} else if fieldType == "float" {
		typeTemp = "float"
	} else if fieldType == "bool" {
		typeTemp = "bool"
	} else {
		typeTemp = "string"
	}
	if fieldName == "id" {
		fieldName = "Id"
	}
	info = fmt.Sprintf(info, typeTemp, fieldName)
	return notes + info
}
