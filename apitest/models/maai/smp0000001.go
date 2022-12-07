package models

// swagger:parameters SMP0000001Request
type SMP0000001Request struct {
	// in: body
	// swagger:allOf
	SMP0000001I SMP0000001I
}

// swagger:parameters SMP0000001
// in: body
type SMP0000001I struct {
	LatestVersion int      `json:"latestVersion" description:"latest version" validate:"max=10"`
	DocGroup      []string `json:"docGroup" description:"doc group" validate:"max=100"`
}

// swagger:response SMP0000001O
type SMP0000001O struct {
	// in: body
	Documents []DLCM0ALL004OList `json:"documents" description:"document list"`
}
type DLCM0ALL004OList struct {
	DocGroup     string          `json:"docGroup" description:"document group"`
	DocId        string          `json:"docId" description:"document id"`
	DocVersion   int             `json:"docVersion" description:"document version"`
	LanguageList []LanguageList2 `json:"languageList" description:"language document content list"`
	Remark       string          `json:"remark" description:"document remark"`
}
type LanguageList2 struct {
	Language   string `json:"language" description:"document language"`
	DocContent string `json:"docContent" description:"document content"`
}

func (i SMP0000001I) GetServiceKey() string {
	return "MP000001"
}
