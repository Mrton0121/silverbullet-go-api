package sbapi

import (
	"sort"
	"strings"
	"time"
)

func (pages *Pages) ReadFolder(folder string, oldestFirst ...bool) (PagesInFolder []Page) {
	PagesInFolder = []Page{}

	pageList := *pages.PageSlice

	for _, page := range pageList {
		if strings.HasPrefix(page.Name, folder) {
			PagesInFolder = append(PagesInFolder, page)
		}
	}

	sortAsc := false
	if len(oldestFirst) > 0 {
		sortAsc = oldestFirst[0]
	}

	// Sorting based on "Created" timestamp
	sort.Slice(PagesInFolder, func(i, j int) bool {
		timeI, errI := time.Parse("2006-01-02 15:04:05", PagesInFolder[i].Created)
		timeJ, errJ := time.Parse("2006-01-02 15:04:05", PagesInFolder[j].Created)

		// If parsing fails, keep original order
		if errI != nil || errJ != nil {
			return false
		}

		if sortAsc {
			return timeI.Before(timeJ) // Oldest first
		}
		return timeJ.Before(timeI) // Newest first
	})

	return PagesInFolder
}
