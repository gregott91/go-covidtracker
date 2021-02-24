package covidtracker

// FilterCsv filters a csv array
func FilterCsv(data [][]string, filter string, column int) (ret [][]string) {
	for _, s := range data {
		if s[column] == filter {
			ret = append(ret, s)
		}
	}
	return
}

// ReduceCsvColumns reduces a csv array
func ReduceCsvColumns(data [][]string, columns []int) (ret [][]string) {
	for _, s := range data {
		reduced := []string{}

		for _, column := range columns {
			reduced = append(reduced, s[column])
		}

		ret = append(ret, reduced)
	}
	return
}
