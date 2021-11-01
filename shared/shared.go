package shared

// AnalyzeData function
// separation of data unique and duplicates
func AnalyzeData(array ...[]string) ([]string, []string) {
	temp := make(map[string]int)

	for _, arrs := range array {
		for _, val := range arrs {
			if _, ok := temp[val]; !ok {
				temp[val] = 1
			} else {
				temp[val] += 1
			}
		}
	}

	unique := make([]string, 0)
	duplicate := make([]string, 0)

	for name, isDuplicate := range temp {
		if isDuplicate > 1 {
			duplicate = append(duplicate, name)
		} else {
			unique = append(unique, name)
		}
	}

	return unique, duplicate
}
