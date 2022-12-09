/*
D. Родословная: подсчет уровней

В генеалогическом древе у каждого человека, кроме родоначальника, есть ровно один родитель.
Каждому элементу дерева сопоставляется целое неотрицательное число, называемое высотой.
У родоначальника высота равна 0, у любого другого элемента высота на 1 больше, чем у его родителя.
Вам дано генеалогическое древо, определите высоту всех его элементов.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	familyTree := make(map[string]string)
	counter := make(map[string]int)

	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	for i := 0; i < n-1; i++ {
		var child, parent string
		_, err2 := fmt.Scan(&child, &parent)
		if err2 != nil {
			return
		}
		familyTree[child] = parent
		counter[child] = 0
		counter[parent] = 0
	}

	for key, _ := range familyTree {
		tmp := key
		_, ok := familyTree[tmp]
		for ok {
			tmp = familyTree[tmp]
			counter[key] += 1
			_, ok = familyTree[tmp]
		}
	}

	names := make([]string, 0, len(counter))
	for k := range counter {
		names = append(names, k)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name, counter[name])
	}

}
