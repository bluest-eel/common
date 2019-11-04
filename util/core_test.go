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

type TestData struct {
	Thing1 string
	Thing2 string
}

func (suite *utilTestSuite) TestGOB() {
	testData := &TestData{Thing1: "stuff", Thing2: "other stuff"}
	encdoed, _ := util.GOBEncode(testData)
	suite.Equal(69, len(encdoed))
	decoded := &TestData{}
	_ = util.GOBDecode(encdoed, decoded)
	suite.Equal("stuff", decoded.Thing1)
	suite.Equal("other stuff", decoded.Thing2)
}
