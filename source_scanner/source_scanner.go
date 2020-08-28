package source_scanner

import (
	"bufio"
	"github.com/ypetya/web-components-maturity-checker/model"
	"os"
	"strings"
)

type folderWithSource struct {
	folder string
	source *model.SourceFolder
}

type fileWithSource struct {
	file   string
	source *model.SourceFolder
}

func Scan(sources *[]model.SourceFolder, components *[]string) Stats {
	stats := Stats{}

	for _, c := range *components {
		stats.Result = make(map[string]ComponentStat)
		stats.Result[c] = ComponentStat{}
	}

	done := make(chan bool)
	folders, files, matches := make(chan folderWithSource), make(chan fileWithSource), make(chan Match)

	go func(matches <-chan Match) {
		for match := range matches {
			x, found := stats.Result[*match.component]

			if !found {
				x = ComponentStat{}
				x.matches = []Match{}
				x.Summary = 0
			}
			x.Summary += match.count
			x.matches = append(x.matches, match)
			stats.Result[*match.component] = x
		}

		done <- true
	}(matches)

	go func(files <-chan fileWithSource, matches chan<- Match) {
		for fx := range files {
			scanFile(fx.file, components, fx.source, matches)
		}
		close(matches)
	}(files, matches)

	go func(folders <-chan folderWithSource, files chan<- fileWithSource) {
		for fx := range folders {
			filterFiles(fx.folder, fx.source, files)
		}
		close(files)
	}(folders, files)

	for i, s := range *sources {
		walkFolder(s.Folder, &((*sources)[i]), folders)
	}
	close(folders)

	<-done

	return stats
}

func scanFile(file string, components *[]string, source *model.SourceFolder, matches chan<- Match) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lN := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range *components {
			count := strings.Count(line, c)
			if count > 0 {
				m := Match{
					component: &((*components)[i]),
					file:      &file,
					source:    source,
					line:      lN,
					count:     count}
				matches <- m
			}
		}
		lN = lN + 1
	}
}

func matchFileExtension(fileName string, extensions *[]string) bool {
	for _, ext := range *extensions {
		if strings.HasSuffix(fileName, "."+ext) {
			return true
		}
	}
	return false
}

func filterFiles(folder string, source *model.SourceFolder, interested chan<- fileWithSource) {
	f, err := os.Open(folder)
	if err != nil {
		panic(err)
	}
	list, err := f.Readdir(-1)
	f.Close()
	for _, fileInfo := range list {
		fn := fileInfo.Name()
		if matchFileExtension(fn, &(*source).Extensions) {
			filePath := folder + "/" + fn
			fx := fileWithSource{file: filePath, source: source}
			interested <- fx
		}
	}
}

func isExcluded(folder string, sourceInfo *model.SourceFolder) bool {
	for _, x := range (*sourceInfo).ExcludeFolders {
		if strings.Compare(folder, x) == 0 {
			return true
		}
	}
	return false
}

func walkFolder(folder string, sourceInfo *model.SourceFolder, out chan<- folderWithSource) error {
	match := folderWithSource{folder: folder, source: sourceInfo}
	out <- match

	f, err := os.Open(folder)
	if err != nil {
		return err
	}
	list, err := f.Readdir(-1)
	f.Close()
	for _, fileInfo := range list {
		if fileInfo.IsDir() {
			folderName := fileInfo.Name()
			if !isExcluded(folderName, sourceInfo) && folderName[0] != '.' {
				err := walkFolder(folder+"/"+folderName, sourceInfo, out)
				if err != nil {
					return err
				}
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}
