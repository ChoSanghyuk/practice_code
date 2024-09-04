package functionaloption

type Scraper struct {
	ScrapeOption func() (string, error)
}

func NewScraper(option func() (string, error)) *Scraper {

	return &Scraper{
		ScrapeOption: option,
	}
}

func (s *Scraper) Scrape() (string, error) {
	return s.ScrapeOption()
}
