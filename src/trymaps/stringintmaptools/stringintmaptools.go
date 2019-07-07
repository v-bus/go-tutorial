package stringintmaptools

//CopySIMap copies map[string]int to map[string]int
func CopySIMap(src map[string]int) map[string]int {
	dst := make(map[string]int)
	if src != nil {
		for key, value := range src {
			dst[key] = value
		}
	}
	return dst
}

//EqualSIMap returns true if firstMap and secondMap are equal to each other
func EqualSIMap(firstMap, secondMap map[string]int) bool {
	equal := false

	switch {
	case ((len(firstMap) == len(secondMap)) && (len(firstMap) > 0 && len(secondMap) > 0)):
		for key, valueFirstMap := range firstMap {
			valueSecondMap, hasKey := secondMap[key]
			if hasKey {
				if valueFirstMap == valueSecondMap {
					equal = true
				} else {
					equal = false
					break
				}
			} else {
				equal = false
				break
			}
		}
	case firstMap == nil && secondMap == nil:
		equal = true
	}

	return equal
}
