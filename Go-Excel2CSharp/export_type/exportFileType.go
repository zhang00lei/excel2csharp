package export_type

type FileType int

const (
	// 导出json
	Json FileType = iota
	CSharp
)
