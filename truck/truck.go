package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// 商品
type Item struct {
	Id     int
	Weight int
}

func NewItem(commandArg string) *Item {
	s := strings.Split(commandArg, ":")
	item := new(Item)
	var err error
	item.Id, err = strconv.Atoi(s[0])
	if err != nil {
		// TODO: エラーを適切に処理する
		fmt.Println("Can not parse item id.")
		//panic("Can not parse item id.")
	}
	item.Weight, err = strconv.Atoi(s[1])
	if err != nil {
		// TODO: エラーを適切に処理する
		fmt.Println("Can not parse weight.")
		//panic("Can not parse weight.")
	}
	return item
}

type Truck struct {
	Id     int
	Items  []Item
	Weight int
}

// トラックに商品を積載する
func (truck *Truck) Load(item Item) {
	truck.Items = append(truck.Items, item)
	truck.Weight += item.Weight
}

func (truck *Truck) Print() {
	fmt.Println(truck.Items)
}

func FindLightestTruck(trucks []Truck) int {
	id := -1
	// intの最大値など極めて大きい値にしておくと良い
	max := 10000
	for i, t := range trucks {
		if t.Weight < max {
			max = t.Weight
			id = i
		}
	}
	return id
}

func main() {
	// 3台のトラックを用意
	trucks := []Truck{
		{Id: 0, Weight: 0}, {Id: 1, Weight: 0}, {Id: 2, Weight: 0},
	}

	// コマンド引数の読み込み
	flag.Parse()
	for _, v := range flag.Args() {
		item := NewItem(v)
		// 毎回ソーティングする案
		//sort.Slice(trucks, func(i, j int) bool {
		//	return trucks[i].Weight < trucks[j].Weight
		//})
		id := FindLightestTruck(trucks)
		trucks[id].Load(*item)
	}

	// 結果の表示
	for _, t := range trucks {
		t.Print()
	}
}
