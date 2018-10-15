package medium

const (
	MEDIUM_TAG = "https://medium.com/tag/"
)

type MediumRecord struct {
	Url string `json:"url"`
}

type MediumArchive struct {
	records map[string][]MediumRecord `json:"records"`
}

type MediumTag struct {
	Tag string `json:"tag"`
}
