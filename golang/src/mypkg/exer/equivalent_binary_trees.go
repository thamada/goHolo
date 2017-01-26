//Time-stamp: <2017-01-27 03:01:16 hamada>
package exer

import (
	"fmt"
	"github.com/Sirupsen/logrus"
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
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		n1, ok1 := <-c1
		n2, ok2 := <-c2

		fmt.Printf("%v, %v\n", n1, n2)
		if n1 != n2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	fmt.Printf("%v, %T\n", logger, logger)
	return true
}

var logger *logrus.Logger = logrus.New()

func Equivalent_binary_trees() {
	
	//	logger.Formatter = new(logrus.JSONFormatter)

	logger.WithFields(logrus.Fields{
		"comment": "this is a test",
		"pi": 3.141592,
	}).Info("INFO")

	if false {
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

	if true {
		for i := 0; i < 1; i++ {
			ta := tree.New(i + 1)
			tb := tree.New(i + 1)
			fmt.Printf("%v, %T\n", ta, ta)
			fmt.Printf("%v, %T\n", tb, tb)
			is_same := Same(ta, tb)
			fmt.Printf("%v\n\n", is_same)
		}
	}
	fmt.Println("end.")
}
