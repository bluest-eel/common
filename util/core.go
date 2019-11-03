package util

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// CallerTree returns data that precomputes paths for convenience
type CallerTree struct {
	DotPath          string
	DotDotPath       string
	DotDotDotPath    string
	DotDotDotDotPath string
}

// CallerPaths takes as input the results from calling runtime.Caller(0) and
// returns a data structure of current caller's path and a set of parent paths.
// In a Golang-standard project directory for this project, the following
// would be true:
//
// paths := CallerPaths(pc, file, line, ok)
// strings.HasSuffix(paths.DotPath, "/github.com/bluest-eel/common/util")
// strings.HasSuffix(paths.DotDotPath, "/github.com/bluest-eel/common")
// strings.HasSuffix(paths.DotDotDotPath, "/github.com/bluest-eel")
// strings.HasSuffix(paths.DotDotDotDotPath, "/github.com")
func CallerPaths(pc uintptr, file string, line int, ok bool) *CallerTree {
	if !ok {
		log.Warn("Couldn't find caller")
		return &CallerTree{}
	}
	callerPath := filepath.Dir(file)
	dotCallerPath := filepath.Dir(callerPath)
	dotDotCallerPath := filepath.Dir(dotCallerPath)
	dotDotDotCallerPath := filepath.Dir(dotDotCallerPath)

	return &CallerTree{
		DotPath:          callerPath,
		DotDotPath:       dotCallerPath,
		DotDotDotPath:    dotDotCallerPath,
		DotDotDotDotPath: dotDotDotCallerPath,
	}
}
