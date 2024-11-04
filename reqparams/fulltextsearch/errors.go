package fulltextsearch

import "fmt"

type ErrorQuery struct {
	Index int
	At    string
}

func (e ErrorQuery) Error() string {
	return fmt.Sprintf("query incorrect format {index:'%v',value:'%v'}", e.Index, e.At)
}
