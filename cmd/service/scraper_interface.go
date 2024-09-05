package service

func NewScraperInterface() InterfaceScraper {
	return &interfaceScraper{}
}

type InterfaceScraper interface {
	ScrapJobs(selectors *Selectors, searchUrl string) []Job
}

type interfaceScraper struct {
}

type Selectors struct {
	TitleSelector    string
	SrcSelector      string
	CompanySelector  string
	LocationSelector string
	PostedAtSelector string
	JobSelector      string
}

type Job struct {
	Title      string `json:"title"`
	Src        string `json:"src"`
	Company    string `json:"company"`
	CompanySrc string `json:"companySrc"`
	Location   string `json:"location"`
	PostedAt   string `json:"postedAt"`
}
