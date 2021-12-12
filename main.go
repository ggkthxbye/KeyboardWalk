/*
	KeyboardWalk
	 - Small Go program to generate keyboard walks to quickly get any low hanging fruit while password cracking
	 - Not an all inclusive key walk list.
	 - Designed to be used with `john` by passing the wordlist via stdin
	 - Example: ./KeyboardWalk -op=4x4 | john --stdin hashes.txt

	 Author: Jason Glenn
	 https://github.com/cyk1k
	 https://www.linkedin.com/in/ggkthxbye/
*/

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed walks/four
var fourWalk string

//go:embed walks/three
var threeWalk string

type Pair struct {
	walk  *string
	count int
}

var opOrder = []string{"4x4", "3x4", "4x5", "3x5", "4x3"}

func main() {
	var opMap map[string]Pair
	var pool []string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s is a small Go program to generate keyboard walks to quickly get any low hanging fruit while password cracking.\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Example: \n\t%s -op=4x3 | john --stdin hashes.txt", os.Args[0])
	}

	op := flag.String("op", "all", "Optionally, explicitly specify keywalk operation to perform. [4x4, 3x4, 4x5, 3x5, 4x3]")
	flag.Parse()
	if !strings.EqualFold(*op, "all") {
		opOrder = []string{*op}
	}

	opMap = map[string]Pair{
		"4x4": Pair{&fourWalk, 4},
		"3x4": Pair{&threeWalk, 4},
		"4x5": Pair{&fourWalk, 5},
		"3x5": Pair{&threeWalk, 5},
		"4x3": Pair{&fourWalk, 3},
	}
	for _, each := range opOrder {
		pool = nil
		for i := 0; i < opMap[each].count; i++ {
			pool = append(pool, strings.Split(strings.Trim(*opMap[each].walk, "\n"), "\n")...)
			pool = append(pool, strings.Split(strings.Trim(ReverseString(*opMap[each].walk), "\n"), "\n")...)
		}
		GetCombos(pool, opMap[each].count, 0, make([]string, len(pool[1])*opMap[each].count))
	}
}

func GetCombos(arr []string, length int, start int, results []string) {
	if length == 0 {
		var out string
		for _, r := range results {
			out += r
		}
		fmt.Println(out)
		return
	}
	for i := start; i <= len(arr)-length; i++ {
		results[len(results)-length] = arr[i]
		GetCombos(arr, length-1, i+1, results)
	}
}

func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
