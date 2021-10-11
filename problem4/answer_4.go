package problem4

import (
	"sort"
	"strings"
)

func GroupAnagram(arrStr []string) map[string][]string {
	groupAnagram := make(map[string][]string)

	for _, str := range arrStr {
		s := strings.Split(str, "")
		sort.Strings(s)
		strSort := strings.Join(s, "")

		if _, exist := groupAnagram[strSort]; exist {
			groupAnagram[strSort] = append(groupAnagram[strSort], str)
		} else {
			groupAnagram[strSort] = []string{str}
		}
	}

	return groupAnagram
}
