package hash_helper

import "sort"

func RankMapStringFloat(values map[string]float64) []string {
	type keyValuePair struct {
		Key   string
		Value float64
	}

	var sortedSlice []keyValuePair

	for key, value := range values {
		sortedSlice = append(sortedSlice, keyValuePair{key, value})
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	rankedKeys := make([]string, len(values))

	for index, keyValuePair := range sortedSlice {
		rankedKeys[index] = keyValuePair.Key
	}

	return rankedKeys
}
