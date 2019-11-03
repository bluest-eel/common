package config_test

import (
	"testing"

	"github.com/bluest-eel/common/components"
	"github.com/bluest-eel/common/config"
	"github.com/stretchr/testify/suite"
)

type configTestSuite struct {
	components.TestBase
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, &configTestSuite{})
}

func (suite *configTestSuite) TestExtension2Type() {
	fileType := config.Extension2Type(".blurg")
	suite.Equal("blurg", fileType)
}

func (suite *configTestSuite) TestAbsConfigPath() {
	path1 := config.AbsConfigPath("/etc/project/configs/app.yml", "/my/project", "configs")
	suite.Equal("/etc/project/configs/app.yml", path1)
	path2 := config.AbsConfigPath("/my/project/configs/app.yml", "/my/project", "configs")
	suite.Equal("/my/project/configs/app.yml", path2)
	path3 := config.AbsConfigPath("dummy", "/my/project", "configs")
	suite.Equal("/my/project/configs", path3)
	path4 := config.AbsConfigPath("", "/my/project/configs", "")
	suite.Equal("/my/project/configs", path4)
	path5 := config.AbsConfigPath("", "/my/project", "configs")
	suite.Equal("/my/project/configs", path5)
	path6 := config.AbsConfigPath("", "/my/project", "")
	suite.Equal("/my/project", path6)
}

func (suite *configTestSuite) TestParseConfigFilename() {
	opts1 := config.ParseConfigFilename("/etc/project/app.yml", "/my/project")
	suite.Equal("/etc/project/app.yml", opts1.AbsPath)
	suite.Equal("app", opts1.File)
	suite.Equal("/my/project", opts1.ProjectPath)
	suite.Equal("/etc/project/", opts1.RelPath)
	suite.Equal("yml", opts1.Type)
	opts2 := config.ParseConfigFilename("configs/app.yml", "/my/project")
	suite.Equal("/my/project/configs", opts2.AbsPath)
	suite.Equal("app", opts2.File)
	suite.Equal("/my/project", opts2.ProjectPath)
	suite.Equal("configs/", opts2.RelPath)
	suite.Equal("yml", opts2.Type)
	opts3 := config.ParseConfigFilename("app.yml", "/my/project")
	suite.Equal("/my/project", opts3.AbsPath)
	suite.Equal("app", opts3.File)
	suite.Equal("/my/project", opts3.ProjectPath)
	suite.Equal("", opts3.RelPath)
	suite.Equal("yml", opts3.Type)
}
