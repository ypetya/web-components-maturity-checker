package Config

import (
	"encoding/json"
	"errors"
	"github.com/ypetya/web-components-maturity-checker/model"
	"io/ioutil"
)

type Config struct {
	// This extension identifies components in dist folder
	ComponentExtension string `json:"componentExtension"`
	// Distribution folder of webcomponents
	ComponentFolder string               `json:"componentFolder"`
  // Name of component-lib
  ComponentLibraryName string
	SourceFolders   []model.SourceFolder `json:"sources"`
}

func Load(fileName string) (Config, error) {
	c := Config{}
	c.load(fileName)

	if c.ComponentFolder == "" {
		return c, errors.New(fileName + ": componentFolder is empty")
	}

	if len(c.SourceFolders) == 0 {
		return c, errors.New(fileName + ": sourceFolders is empty")
	}

	var err error
	return c, err
}

func (c *Config) load(fileName string) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &c)

	if err != nil {
		panic(err)
	}
}
