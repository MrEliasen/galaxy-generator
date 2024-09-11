package utils

import (
	"fmt"
	"strings"
)

func ToTable(headings []string, data [][]string) []string {
	longest := []int{}
	padding := []string{}
	spacer := []string{}

	for _, h := range headings {
		longest = append(longest, len(h))
		spacer = append(spacer, "")
	}

	tableLines := [][]string{
		spacer,
		headings,
		spacer,
	}

	for _, d := range data {
		row := []string{}
		for i, c := range d {
			l := len(c)

			if l > longest[i] {
				longest[i] = l
			}

			row = append(row, c)
		}

		tableLines = append(tableLines, row)
	}

	for _, d := range longest {
		padding = append(padding, "%-"+fmt.Sprintf("%d", d+1)+"s")
	}

	lineLng := 0
	result := []string{}

	for _, line := range tableLines {
		for i, cell := range line {
			pd := padding[i]
			line[i] = fmt.Sprintf(pd, cell)
		}

		str := "| " + strings.Join(line, "| ") + " |"

		if len(str) > lineLng {
			lineLng = len(str)
		}

		result = append(result, str)
	}

	separator := strings.Repeat("-", lineLng)
	result[0] = separator
	result[2] = separator
	result = append(result, separator)
	return result
}
