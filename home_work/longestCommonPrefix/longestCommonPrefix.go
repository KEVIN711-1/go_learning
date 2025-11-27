package main

import (
	"fmt"
)

func CommonPrefix(strA string, strB string) string {
	var min_len int

	lenA := len(strA)
	lenB := len(strB)

	if lenA < lenB {
		min_len = lenA
	} else {
		min_len = lenB
	}

	for i := 0; i < min_len; i++ {
		if strA[i] != strB[i] {
			return strA[:i]
		}
	}

	if lenA < lenB {
		return strA[:min_len]
	} else {
		return strB[:min_len]
	}
}

func longestCommonPrefix(strs []string) string {
	CommonStr := strs[0]

	for i := 0; i < len(strs); i++ {

		CommonStr = CommonPrefix(CommonStr, strs[i])
		if CommonStr == "" {
			break
		}
	}
	return CommonStr
}

func main() {
	// 测试代码
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("最长公共前缀:", longestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Println("2最长公共前缀:", longestCommonPrefix(strs))
}
