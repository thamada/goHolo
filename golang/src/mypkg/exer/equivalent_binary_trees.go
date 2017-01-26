//Time-stamp: <2017-01-27 02:02:09 hamada>
package exer

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	var dwarf func(*tree.Tree)

	dwarf = func(t *tree.Tree) {
		if t.Right != nil {
			dwarf(t.Right)
		}
		ch <- t.Value
		if t.Left != nil {
			dwarf(t.Left)
		}
	}

	fmt.Printf("top.Value = %v\n", t.Value)
	dwarf(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func Equivalent_binary_trees() {

	if true {
		t := tree.New(2)
		fmt.Printf("%v, %T\n", t, t)
		c := make(chan int)
		go Walk(t, c)
		for {
			n, ok := <-c
			fmt.Printf("%v\n", n)
			if !ok {
				break
			}
		}
	}

	if false {
		for i := 0; i < 10; i++ {
			ta := tree.New(i)
			tb := tree.New(i)
			fmt.Printf("%v, %T\n", ta, ta)
			fmt.Printf("%v, %T\n", tb, tb)
			is_same := Same(ta, tb)
			fmt.Printf("%v\n\n", is_same)
		}
	}
	fmt.Println("end.")
}
