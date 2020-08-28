package component_detector

import (
	"io/ioutil"
	"strings"
)

// Tries to find files with componentExtension extension in the given directory.
// It cuts the extension off, the remaining prefix represents a webcomponent.
// Returns component names on the channel.
func Detect(componentFolder string,
	componentExtension string) (<-chan string, error) {

	components := make(chan string)

	files, err := ioutil.ReadDir(componentFolder)

	if err == nil {
		go func() {
			for _, file := range files {
				fN := file.Name()
				ix := len(fN) - len(componentExtension)
				if ix > 0 && strings.LastIndex(fN, componentExtension) == ix {
					componentName := string(fN[0:ix])
					components <- componentName
				}
			}
			close(components)
		}()
	} else {
		close(components)
	}
	return components, err
}
