package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/hima398/recruitment-exam/monster/client"
)

type Monster struct {
	Name string
}

// if m2 wins return true
func Compare(m1, m2 string) bool {
	return m2 == client.Battle(m1, m2).Winner
}

func main() {
	flag.Parse()
	monsters := flag.Args()
	// sort.Sliceは内部でQuick Sortを使用
	sort.Slice(monsters, func(i, j int) bool {
		return monsters[j] == client.Battle(monsters[i], monsters[j]).Winner
	})

	fmt.Println("Loser <-- ---- --> Winner")
	fmt.Println(monsters)
}
