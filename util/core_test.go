package util_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/bluest-eel/common/components"
	"github.com/bluest-eel/common/util"
	"github.com/stretchr/testify/suite"
)

type utilTestSuite struct {
	components.TestBase
}

func TestUtilTestSuite(t *testing.T) {
	suite.Run(t, &utilTestSuite{})
}

func (suite *utilTestSuite) TestCallerPaths() {
	paths := util.CallerPaths(runtime.Caller(0))
	suite.True(strings.HasSuffix(paths.DotPath, "/util"))
}
