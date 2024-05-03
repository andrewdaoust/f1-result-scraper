package result

import (
	"strconv"
	"strings"
)

type Results map[string]int

func ParseResult(rows []string) (Results, error) {
	results := Results{}
	for _, row := range rows {
		splitRow := strings.Split(row, ",")
		
		var position int
		var err error
		if splitRow[0] == "NC" {
			position = 0
		} else {
			position, err = strconv.Atoi(splitRow[0])
			if err != nil {
				return nil, err
			}
		}
		name := splitRow[2]
		shortName := strings.Split(name, " ")[2]

		results[shortName] = position
	}

	return results, nil
}
