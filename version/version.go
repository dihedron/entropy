package version

import (
	"log/slog"
	"os"
	"path"
	"runtime/debug"
)

// NOTE: some of these variables are populated at compile time by using the -ldflags
// linker flag:
//
//	$> go build -ldflags "-X github.com/dihedron/entropy/version.VersionMajor=$(major_ver)"
//
// in order to get the package path to the GitHash variable to use in the
// linker flag, use the nm utility and look for the variable in the built
// application symbols, then use its path in the linker flag:
//
//	$> nm ./entropy | grep VersionMajor
//	0000000000677fe0 B github.com/dihedron/entropy/version.VersionMajor
var (
	// BuildTime is the time at which the application was built.
	BuildTime string
	// VersionMajor is the major version of the application.
	VersionMajor = "0"
	// VersionMinor is the minor version of the application.
	VersionMinor = "0"
	// VersionPatch is the patch or revision level of the application.
	VersionPatch = "0"
	// Name is the name of the application or plugin.
	Name string = "application name placeholder"
	// Description is a one-liner description of the application or plugin.
	Description string = "application name placeholder"
	// Copyright is the copyright clause of the application or plugin.
	Copyright string = "copyright placeholder"
	// License is the license under which the code is released.
	License string = "license placeholder"
	// LicenseURL is the URL at which the license is available.
	LicenseURL string = "license URL placeholder"
	// GitCommit is the commit of this version of the application.
	GitCommit string
	// GitTime is the modification time associated with the Git commit.
	GitTime string
	// GitModified reports whether the repository had outstanding local changes at time of build.
	GitModified string
	// GoVersion is the version of the Go compiler used in the build process.
	GoVersion string
	// GoOS is the target operating system of the application.
	GoOS string
	// GoOS is the target architecture of the application.
	GoArch string
)

func init() {

	if Name == "" {
		Name = path.Base(os.Args[0])
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		slog.Error("no build info available")
		return
	}

	GoVersion = bi.GoVersion

	for _, setting := range bi.Settings {
		switch setting.Key {
		case "GOOS":
			GoOS = setting.Value
		case "GOARCH":
			GoArch = setting.Value
		case "vcs.revision":
			GitCommit = setting.Value
		case "vcs.time":
			GitTime = setting.Value
		case "vcs.modified":
			GitModified = setting.Value
		}
	}
}
