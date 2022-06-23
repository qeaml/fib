package config

import (
	"bytes"
	"io"
	"os"

	"github.com/qeaml/naml"
)

type config struct {
	Port uint16
}

var GlobalConfig *config

func LoadConfig() error {
	f, err := os.Open("config.naml")
	if err != nil {
		return err
	}
	defer f.Close()
	src, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	GlobalConfig = &config{}
	d := naml.NewDecoder(bytes.NewReader(src))
	return d.Decode(GlobalConfig)
}
