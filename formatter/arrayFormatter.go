package formatter

import (
	"sort"
	"strings"
)

type Formatter interface {
	String() string
}

type ABCOrdered []string

func (a ABCOrdered) Len() int {
	return len(a)
}

func (a ABCOrdered) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ABCOrdered) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a* ABCOrdered) Sort() {
  sort.Sort(a)
}

func (a ABCOrdered) String() string {
	sort.Sort(a)

	return strings.Join(a, ", ")
}
