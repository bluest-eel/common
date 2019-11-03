package util_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/bluest-eel/common/components"
	"github.com/bluest-eel/common/util"
	"github.com/stretchr/testify/suite"
)

type utilTestSuite struct {
	components.TestBase
	version *util.Version
}

func (suite *utilTestSuite) SetupSuite() {
	buildTime, err := time.Parse(time.RFC3339, "1924-01-08T01:08:00+05:30")
	if err != nil {
		panic(err)
	}
	suite.version = &util.Version{
		Semantic:   "1.2.3-alpha4",
		BuildDate:  "Today",
		BuildTime:  &buildTime,
		GitCommit:  "abc1234",
		GitBranch:  "master",
		GitSummary: "super-great",
	}
}

func TestUtilTestSuite(t *testing.T) {
	suite.Run(t, &utilTestSuite{})
}

func (suite *utilTestSuite) TestVersion() {
	suite.Equal("1.2.3-alpha4", suite.version.Semantic)
	suite.Equal("abc1234", suite.version.GitCommit)
}

func (suite *utilTestSuite) TestJSONVersion() {
	result, err := json.Marshal(suite.version)
	if err != nil {
		panic(err)
	}
	expected := `{"semantic":"1.2.3-alpha4","build-date":"Today",` +
		`"build-time":"1924-01-08T01:08:00+05:30","git-commit":"abc1234",` +
		`"git-branch":"master","git-summary":"super-great"}`
	suite.Equal(expected, string(result))
}
