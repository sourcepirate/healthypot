package medium

type MediumRecord struct {
	Url string `json:"url"`
}

type MediumArchive struct {
	records map[string][]MediumRecords `json:"records"`
}

type MediumTag struct {
	Tag string `json:"tag"`
}
