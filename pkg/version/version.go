package version

import (
	"fmt"
)

type VersionInfo struct {
	version string
	commit  string
	date    string
	builtBy string
}

var versionInfo VersionInfo

func Init(version, commit, date, builtBy string) {
	versionInfo = VersionInfo{
		version: version,
		commit:  commit,
		date:    date,
		builtBy: builtBy,
	}
}

// Print version info
func Print() {
	fmt.Printf("Version information\nVersion: %s, Commit: %s\n", Version(), Commit())
	fmt.Printf("Build date: %s, Build by: %s\n", BuildDate(), BuiltBy())
}

// Commit returns git commit
func Commit() string {
	return versionInfo.commit
}

// Version returns application version
func Version() string {
	return versionInfo.version
}

func BuildDate() string {
	return versionInfo.date
}

func BuiltBy() string {
	return versionInfo.builtBy
}
