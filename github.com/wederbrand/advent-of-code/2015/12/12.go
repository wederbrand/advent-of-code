package main

import (
	"encoding/json"
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"time"
)

func main() {
	start := time.Now()
	inFile := util.GetFileContents("2015/12/input.txt", "\n")

	var in interface{}
	err := json.Unmarshal([]byte(inFile[0]), &in)
	if err != nil {
		panic(err)
	}

	part1 := getCount(in, false)
	fmt.Println("part1", part1, "in", time.Since(start))
	part2 := getCount(in, true)
	fmt.Println("part2", part2, "in", time.Since(start))
}

func getCount(in interface{}, ignoreRed bool) int {
	switch lv := in.(type) {
	case float64:
		return int(lv)
	case int:
		return lv
	case string:
		return 0
	case []interface{}:
		cnt := 0
		for _, entry := range lv {
			cnt += getCount(entry, ignoreRed)
		}
		return cnt
	case map[string]interface{}:
		cnt := 0
		for _, entry := range lv {
			if ignoreRed && entry == "red" {
				return 0
			}
		}

		for _, entry := range lv {
			cnt += getCount(entry, ignoreRed)
		}
		return cnt
	default:
		fmt.Printf("unhandled type %t", in)
	}
	return 0
}
