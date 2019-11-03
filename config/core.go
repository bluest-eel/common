package config

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	cfg "github.com/spf13/viper"
)

// Configuration related constants
const (
	ConfigFileEnvKey  string = "BEAP_CONFIG"
	DefaultConfigDir  string = "configs"
	DefaultConfigFile string = "app"
	DefaultConfigType string = "yml"
	ExtensionSep      string = "."
	ConfigReadError   string = "Couldn't read config file"
)

// Options for use in configuration functions
type Options struct {
	AbsPath     string // The absolute path to the config file
	AppName     string
	File        string // The name of the config file, minus any path or extension
	ProjectPath string // The absolute path to the project being configured
	RelPath     string // The relative path to the config file
	Type        string // The config file type, usually extension
}

// Bootstrap ...
func Bootstrap(appName string, projectPath string) {
	opts := &Options{
		AbsPath:     AbsConfigPath("", projectPath, DefaultConfigDir),
		AppName:     appName,
		File:        DefaultConfigFile,
		ProjectPath: projectPath,
		RelPath:     DefaultConfigFile,
		Type:        DefaultConfigType,
	}
	setConfiguration(opts)
}

// Setup ...
func Setup(appName string, projectPath string, filepath string) {
	opts := ParseConfigFilename(filepath, projectPath)
	opts.AppName = appName
	opts.ProjectPath = projectPath
	setConfiguration(opts)
}

func setConfiguration(opts *Options) {
	cfg.AddConfigPath(DefaultConfigDir)
	cfg.AddConfigPath(opts.RelPath)
	cfg.SetConfigName(opts.File)
	cfg.SetConfigType(opts.Type)
	cfg.SetEnvPrefix(opts.AppName)
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	cfg.Set("Verbose", true)
	cfg.AutomaticEnv()
	cfg.AddConfigPath(opts.AbsPath) // Computed path
	cfg.AddConfigPath("/")          // support for Docker

	err := cfg.ReadInConfig()
	if err != nil {
		log.Panicf("%s: %s", ConfigReadError, err)
	}
	log.Tracef("Env vars: %v", os.Environ())
}

// ParseConfigFilename takes a given string and parses it as a filename,
// separating prefix directories, name of file, and file extension.
func ParseConfigFilename(filename string, projectPath string) *Options {
	relPath, parsedFilename := filepath.Split(filepath.Clean(filename))
	extension := filepath.Ext(parsedFilename)
	return &Options{
		AbsPath:     AbsConfigPath(filename, projectPath, relPath),
		File:        strings.TrimSuffix(parsedFilename, extension),
		ProjectPath: projectPath,
		RelPath:     relPath,
		Type:        Extension2Type(extension),
	}
}

// AbsConfigPath ...
func AbsConfigPath(filename string, projectPath string, relPath string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(projectPath, relPath)
}

// Extension2Type ...
func Extension2Type(ext string) string {
	parts := strings.Split(ext, ExtensionSep)
	return parts[len(parts)-1]
}

// EnvConfigFile extracts the value associated with the environment variable
// that specifies which config file to use when starting up a system.
func EnvConfigFile() string {
	return os.Getenv(ConfigFileEnvKey)
}
