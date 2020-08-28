package model

import (
  "io/ioutil"
  "encoding/json"
  "log"
  )

type SourceFolder struct {
	Name       string
  Version string
	Folder     string
	Extensions []string
  ExcludeFolders []string
}

type packageJson struct {
  Name string
  Dependencies map[string]interface{}
}

// Loads Name and Version from package.json
func (f* SourceFolder) ResolvePackageJson(dependencyName string) {
  fileName := f.Folder + "/package.json"
	buf, err := ioutil.ReadFile(fileName)
	if err == nil {
    p := packageJson{}

    err = json.Unmarshal(buf, &p)

    if err != nil {
      panic(err)
    }

    f.Name = p.Name
    if ver, found := p.Dependencies[dependencyName] ; found {
      f.Version = ver.(string)
    } else {
      log.Println("No dependency for", dependencyName, "in project", f.Name)
    }
  } else {
    log.Fatal("Could not resolve:", fileName)
  }
}
