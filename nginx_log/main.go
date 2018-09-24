package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	reg, err := regexp.Compile(`^([0-9]{1,3}\.){3}[0-9]{1,3}`)
	if err != nil {
		log.Fatal(err)
	}
	input.Split(bufio.ScanWords)
	for input.Scan() {
		if reg.MatchString(input.Text()) {
			counts[input.Text()]++
		}

	}
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%d\t%s\n",  kv.Value, kv.Key)
	}

}
