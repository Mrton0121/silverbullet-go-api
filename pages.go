package sbapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type RawPage struct {
	ContentType  string `json:"contentType"`
	Created      int64  `json:"created"`
	LastModified int64  `json:"lastModified"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
}

// Page struct which uses
type Page struct {
	ContentType  string
	Created      string
	LastModified string
	Name         string
	Size         int
}

type Pages struct {
	PageSlice *[]Page
}

func (pages Pages) GetLatestCreated() (*Page, error) {
	var times []time.Time
	for _, page := range *pages.PageSlice {
		// Parse dates into time.Time
		t, err := time.Parse(time.DateTime, page.Created)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}
		times = append(times, t)
	}

	latestIndex := 0
	for i, t := range times {
		if t.After(times[latestIndex]) {
			latestIndex = i
		}
	}

	latestPage := *pages.PageSlice

	return &latestPage[latestIndex], nil
}

func (pages Pages) GetLatestModified() (*Page, error) {
	var times []time.Time
	for _, page := range *pages.PageSlice {
		// Parse dates into time.Time
		t, err := time.Parse(time.DateTime, page.LastModified)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return nil, err
		}
		times = append(times, t)
	}

	latestIndex := 0
	for i, t := range times {
		if t.After(times[latestIndex]) {
			latestIndex = i
		}
	}

	latestPage := *pages.PageSlice

	return &latestPage[latestIndex], nil
}

// GetPages gets and formats the /index.json response into a list of Page slice
// skipStrings are handled as prefixes
func (client *SBClient) GetPages(skip bool, skipStrings ...string) (pages *Pages, err error) {
	rawPages := []RawPage{}

	resp, err := client.Get("index.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into RawPage structs
	err = json.Unmarshal([]byte(resp), &rawPages)
	if err != nil {
		return nil, err
	}

	// Convert RawPage to Page with formatted timestamps
	pagesSlice := []Page{}
	for _, rawPage := range rawPages {
		page := Page{
			ContentType:  rawPage.ContentType,
			Created:      formatTimestamp(rawPage.Created),
			LastModified: formatTimestamp(rawPage.LastModified),
			Name:         rawPage.Name,
			Size:         rawPage.Size,
		}
		pagesSlice = append(pagesSlice, page)
	}

	// if skip is true and there are user listed strings to skip
	if skip && skipStrings != nil {
		filteredPages := []Page{}
		for _, page := range pagesSlice {
			// skip listed pages
			if containsPrefix(page.Name, skipStrings) {
				continue
			}
			filteredPages = append(filteredPages, page)
		}
		return &Pages{PageSlice: &filteredPages}, nil
	}

	return &Pages{PageSlice: &pagesSlice}, err
}
