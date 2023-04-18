package models

type CM000006I struct {
	FileId        string          `json:"fileId" description:"File id"`
	DosFileBase64 []DosFileBase64 `json:"dosFileBase64" description:"Dos file base64"`
	Flag          string          `json:"flag" description:"Flag(2-Multiple files)"`
	File          []byte          `json:"file" description:"File"`
}

type CM000006O struct {
	Key    []File `json:"key" description:"Key"`
	Status string `json:"status" description:"Status(Ok)"`
}

type DosFileBase64 struct {
	Key string `json:"key" description:"Key"`
	Buf string `json:"buf" description:"Buf"`
}

type File struct {
	FileId   string `json:"fileId" description:"File id"`
	FileName string `json:"fileName" description:"File name"`
}

func (*CM000006I) GetServiceKey() string {
	return "cm000006"
}
