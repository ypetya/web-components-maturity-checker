package source_scanner

import (
	"fmt"
	"github.com/ypetya/web-components-maturity-checker/formatter"
	"github.com/ypetya/web-components-maturity-checker/model"
	"strconv"
	"time"
)

type (
	Stats struct {
		Result map[string]ComponentStat
	}
	ComponentStat struct {
		matches []Match
		Summary int
	}
	Match struct {
		component *string
		source    *model.SourceFolder
		file      *string
		line      int
		count     int
	}
)

func (cs ComponentStat) Debug() {
	fmt.Println(strconv.Itoa(cs.Summary))
	for _, m := range cs.matches {
		fmt.Println(*m.file + ":" + strconv.Itoa(m.line) + " " + strconv.Itoa(m.count))
	}
}

func (s Stats) Debug() {
	for k, v := range s.Result {
		fmt.Println(k)
		v.Debug()
	}
}

func includes(p *model.SourceFolder, ps *[]*model.SourceFolder) bool {
	for _, v := range *ps {
		if v == p {
			return true
		}
	}
	return false
}

func (st ComponentStat) count(p *model.SourceFolder) int {
	c := 0
	for _, m := range st.matches {
		if m.source == p {
			c += m.count
		}
	}
	return c
}

func (s *Stats) PrintMarkDown() {

	currentTime := time.Now()
	fmt.Println("# Components")

	var components []string
	var projects []*model.SourceFolder

	//TODO one iteration should be enough to construct the table
	// collect components and versions
	for k, v := range (*s).Result {
		components = append(components, k)

		for _, m := range v.matches {
			if !includes(m.source, &projects) {
				projects = append(projects, m.source)
			}
		}
	}

	// prepare header
	var prjCols []string //ABCOrdered

	for _, v := range projects {
		prjCols = append(prjCols, (*v).Name+" "+(*v).Version)
	}

	// print header
	fmt.Printf("| Component ")

	for _, v := range prjCols {
		fmt.Printf("| %s ", v)
	}
	fmt.Printf("|\n")

	// print divider
	fmt.Printf("| --- ")

	for i := 0; i < len(prjCols); i++ {
		fmt.Printf("| --- ")
	}
	fmt.Printf("|\n")

	// print content
	var componentsOrdered formatter.ABCOrdered
	componentsOrdered = components
	componentsOrdered.Sort()

	for _, c := range componentsOrdered {
		fmt.Printf("| %s ", c)

		for _, p := range projects {
			fmt.Printf("| %d ", s.Result[c].count(p))
		}

		fmt.Printf("|\n")
	}

	fmt.Println("* Report generated:", currentTime.String())
}
