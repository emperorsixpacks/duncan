package routers

import "strings"

func del[T interface{}](i int, K []T) []T {
	newArray := make([]T, len(K))
	copy(newArray, K)
	return append(newArray[:i], newArray[i+1:]...)
}

// NOTE this returns all a list of int that identifes all the path params in the path
func returnDestinationPath(regPath string) (string, []Param) {
	var params []Param
	var newPathItems []string
	pathSlice := strings.Split(strings.Trim(regPath, "/"), "/")
	for i, r := range pathSlice {
		if PathParamRegex.Match([]byte(r)) {
			if AssignParamRegex.Match([]byte(pathSlice[i])) {
				params = append(params, Param{
					key:   strings.Split(pathSlice[i], "=")[1],
					index: i,
				})
				continue
			}
			params = append(params, Param{
				key:   r,
				index: i,
			})
			continue
		}
		newPathItems = append(newPathItems, r)
	}
	return strings.Join(newPathItems, "/"), params
}

/*
	func cleanPath(rp string, pathParam []Param) (map[string]interface{}, error) {
		returnMap := make(map[string]interface{})
		if len(pathParam) == 0 {
			returnMap["path"] = rp
			returnMap["params"] = nil
			return returnMap, nil
		}
		//	reqPSplit := strings.Split(strings.Trim("/", reqP), "/")
		for _, x := range pathParam {
			// NOTE this could return an error is that index does not exist
			//		x.value = reqPSplit[x.index]
			fmt.Println(x)
		}
		return //reqP, nil
	}
*/
func commonPrefix(str1 string, str2 string) (string, bool) {
	var i int
	max_ := min(len(str1), len(str2))
	var common string
	for i = 0; i < max_ && str1[i] == str2[i]; i++ {
		continue
	}
	if common = str1[:i]; common == "" {
		return common, false
	}
	return common, true
}
