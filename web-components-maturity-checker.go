package main

import (
	"fmt"
	cd "github.com/ypetya/web-components-maturity-checker/component_detector"
	config "github.com/ypetya/web-components-maturity-checker/config"
	s "github.com/ypetya/web-components-maturity-checker/source_scanner"
	"os"
)

func debug() bool {
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		return true
	}
	return false
}

func main() {
	if c, err := config.Load("./config.json"); err != nil {
		panic(err)
	} else {
		// print config
		if debug() {
			fmt.Println("Component library name:", c.ComponentLibraryName)
			fmt.Println("Component folder:", c.ComponentFolder)
			fmt.Println("Source folders:")
		}
		for i, _ := range c.SourceFolders {
			c.SourceFolders[i].ResolvePackageJson(c.ComponentLibraryName)
			if debug() {
				v := c.SourceFolders[i]
				fmt.Println(i, v.Name, ":", v.Folder, "Version:", v.Version)
			}
		}

		// discover components
		components, _ := cd.Detect(c.ComponentFolder, c.ComponentExtension)

		var componentKeys []string
		for component := range components {
			if debug() {
				fmt.Println("Component:", component)
			}
			componentKeys = append(componentKeys, component)
		}

		// detect
		stats := s.Scan(&c.SourceFolders, &componentKeys)

		// print output
		if debug() {
			fmt.Println(" *** Stats ")
			stats.Debug()
		}
		stats.PrintMarkDown()
	}
}
