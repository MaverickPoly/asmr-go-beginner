package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	str := strings.Split(strings.ToLower(s), "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func isAnagram(s1, s2 string) bool {
	return sortString(s1) == sortString(s2)
}

func main() {
	fmt.Println(sortString("listeN"))
	fmt.Println(sortString("silenT"))
	fmt.Println(isAnagram("LiStEN", "SILent"))
}
