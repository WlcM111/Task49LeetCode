package groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	var result [][]string

	if len(strs) == 0 {
		return result
	}
	if len(strs) == 1 {
		result = append(result, []string{strs[0]})
		return result
	}

	anagrams := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		runes := []rune(strs[i])

		sort.Slice(runes, func(x, y int) bool {
			return runes[x] < runes[y]
		})

		s := string(runes)

		if _, ok := anagrams[s]; ok {
			anagrams[s] = append(anagrams[s], strs[i])
		} else {
			anagrams[s] = append(anagrams[s], strs[i])
		}
	}

	for _, item := range anagrams {
		result = append(result, item)
	}

	return result
}
