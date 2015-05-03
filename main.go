package tabulate

import "math"

func Tabulate(screenWidth int, items []string) string {

	// find the largest column in the data
	var colWidth = 0

	for _, item := range items {
		l := len(item)
		if l > colWidth {
			colWidth = l + 2
		}
	}

	var maxBase = float64(screenWidth) / float64(colWidth)
	var maxColumns = int(math.Floor(maxBase))

	var output = ""
	var length = len(items)

	if length == 0 {
		return ""
	}

	var minBase = float64(length) / float64(maxColumns)
	var minRows = int(math.Ceil(minBase))

	for row := 0; row < minRows; row++ {
		for col := 0; col < maxColumns; col++ {

			var idx = row*maxColumns + col

			if idx >= length {
				break
			}

			var item = items[idx]
			output += item

			// add the padding to the end of the item
			if col < maxColumns-1 {
				itemLen := len(item)
				for s := 0; s < colWidth-itemLen; s++ {
					output += " "
				}
			}
		}

		output += "\r\n"
	}

	return output
}
