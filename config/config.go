package config

import (
	"dotamaster/cmd"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	App App
	Log Log
}

type App struct {
	Host      string
	Port      int
	Debug     bool
	Whitelist []string
}

type Log struct {
	Prefix     string
	Dir        string
	LevelDebug bool
}

func init() {
	// Init CLI commands
	cmd.Root().Use = "bin/dotaapi --config <Config path>"
	cmd.Root().Short = "dotaapi - Provide API for dotastats"
	cmd.Root().Long = "dotaapi"

	cmd.SetRunFunc(load)
}

func load() {
	once.Do(func() {
		if err := cmd.GetViper().Unmarshal(&conf); err != nil {
			panic(err)
		}
	})
}

func Load() {
	load()
}

func Get() Config {
	load()
	return conf
}
