package excel2json

import (
	"excel2csharp/export_type"
	"excel2csharp/util"
	"fmt"
	"strings"
)

func ExportJsonConfig(excelData [][]string, outJsonPath string) {
	write, file := util.GetFileWriter(outJsonPath, export_type.Json)
	defer file.Close()
	write.WriteString("[\n")
	for i, rowData := range excelData {
		if i < 4 {
			continue
		}
		write.WriteString("  {")
		dataLength := len(excelData[3])

		rowCount := 0
		for j := 0; j < dataLength; j++ {
			fieldName := excelData[2][j]
			if fieldName == "" {
				break
			}
			rowCount++
		}

		for j := 0; j < rowCount; j++ {
			cell := ""
			if j < len(rowData) {
				cell = rowData[j]
			}
			fieldName := excelData[2][j]
			if strings.HasSuffix(fieldName, "#") || fieldName == "" {
				continue
			}
			fieldType := excelData[3][j]
			if fieldType == "idx1" {
				write.WriteString(fmt.Sprintf("\"%s\":%s", "Id", cell))
			} else if fieldType == "int" || fieldType == "float" {
				if cell == "" {
					cell = "0"
				}
				write.WriteString(fmt.Sprintf("\"%s\":%s", fieldName, cell))
			} else if fieldType == "bool" {
				tempInfo := "True"
				if cell == "0" {
					tempInfo = "False"
				} else if cell == "1" {
					tempInfo = "True"
				} else {
					if cell == "" {
						tempInfo = ""
					} else if strings.ToUpper(cell) == "TRUE" {
						tempInfo = "True"
					} else {
						tempInfo = "False"
					}
				}
				write.WriteString(fmt.Sprintf("\"%s\":\"%s\"", fieldName, tempInfo))
			} else {
				if strings.Contains(cell, "\"") {
					cell = strings.Replace(cell, "\"", `\"`, -1)
				}
				if strings.Contains(cell, "\n") {
					cell = strings.Replace(cell, "\n", "\\n", -1)
				}
				write.WriteString(fmt.Sprintf("\"%s\":\"%s\"", fieldName, cell))
			}
			if j != rowCount-1 {
				write.WriteString(",")
			}
		}
		write.WriteString("}")
		if i != len(excelData)-1 {
			write.WriteString(",")
		}
		write.WriteString("\n")
	}
	write.WriteString("]")
	write.Flush()
}
