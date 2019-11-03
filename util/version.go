package util

import (
	"fmt"
	"time"
)

const na string = "N/A"

// Version data
type Version struct {
	Semantic   string     `json:"semantic"`
	BuildDate  string     `json:"build-date"`
	BuildTime  *time.Time `json:"build-time"`
	GitCommit  string     `json:"git-commit"`
	GitBranch  string     `json:"git-branch"`
	GitSummary string     `json:"git-summary"`
}

// BuildString ...
func BuildString(version *Version) string {
	if version.GitCommit == "" {
		return na
	}
	return fmt.Sprintf("%s@%s, %s", version.GitBranch, version.GitCommit, version.BuildDate)
}

// GetBuildTime ...
func GetBuildTime(version *Version) *time.Time {
	return version.BuildTime
}

// VersionString ...
func VersionString(version *Version) string {
	if version.Semantic == "" {
		return na
	}
	return version.Semantic
}

// VersionedBuildString ...
func VersionedBuildString(version *Version) string {
	v := version.Semantic
	gc := version.GitCommit
	if v == "" {
		v = na
	}
	if gc == "" {
		gc = na
	}
	return fmt.Sprintf("%s, %s@%s, %s", v, version.GitBranch, gc, version.BuildDate)
}
