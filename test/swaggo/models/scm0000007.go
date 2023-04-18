package models

type CM000007I struct {
	FileId string `json:"fileId" description:"File id"`
}

type CM000007O struct {
	UrlStr string `json:"urlStr" description:"Url of file"`
	Buf    string `json:"buf" description:"Buffer of file"`
}

func (*CM000007I) GetServiceKey() string {
	return "cm000007"
}
