package leetcode

// Задача №1436: Destination City
// https://leetcode.com/problems/destination-city/

func DestCity(paths [][]string) string {
	cityMap := make(map[string]bool)

	// add departure cities only
	for i := range paths {
		cityMap[paths[i][0]] = true
	}

	// check for only city that is destination but not departure
	for i := range paths {
		if _, ok := cityMap[paths[i][1]]; !ok {
			return paths[i][1]
		}

	}

	return ""
}
