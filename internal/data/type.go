package data

import "time"

type Project struct {
	Name        string
	Description string

	RootPath string
	VCS      string
	Repo     string

	Versions []string

	LatestCommitTime time.Time
	LatestCommitHash string
}
