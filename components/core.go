package components

import (
	"github.com/geomyidia/reverb"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// BaseApp ...
type BaseApp struct {
	AppName     string
	AppAbbv     string
	ProjectPath string
	ConfigFile  string
}

// BaseCLI ...
type BaseCLI struct {
	Args    []string
	Command string
	RawArgs []string
}

// BaseGRPC ...
type BaseGRPC struct {
	GRPCD *reverb.Reverb
}

// BaseLogger component collection
type BaseLogger struct {
	Logger *logger.Logger
}

// TestBase component that keeps stdout clean
type TestBase struct {
	suite.Suite
}

// TestGRPC ...
type TestBaseGRPC struct {
	BaseGRPC
	TestBase
}
