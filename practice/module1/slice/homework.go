package slice

import "fmt"

// ReplaceStringInSlice 替换切片中指定文字
func ReplaceStringInSlice(param []string) []string {
	if param == nil {
		fmt.Println("param is nil")
		return nil
	}

	for index, val := range param {
		if val == "stupid" {
			param[index] = "smart"
		} else if val == "weak" {
			param[index] = "strong"
		}
	}

	return param
}
