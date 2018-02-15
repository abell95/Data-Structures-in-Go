//Needs search method and delete method

package main

import (
	"fmt"
)

// Link is the main component of list
type Link struct {
	val   int
	index int
	next  *Link
}

// List points to head of list
type List struct {
	first *Link
	size  int
}

func createEmptyList() *List {
	newList := new(List)
	newList.first = nil
	return newList
}

func createListContaining(args ...int) *List {
	newList := createEmptyList()
	for _, n := range args {
		newList.appendLink(n)
	}
	newList.size = len(args)
	return newList
}

func (l *List) appendLink(num int) {
	newLink := new(Link)
	newLink.val, newLink.next = num, nil

	iter := l.first
	if iter == nil {
		l.first = newLink
		newLink.index = 0
		return
	}
	indexCtr := 1
	for iter.next != nil {
		iter = iter.next
		indexCtr++
	}
	iter.next = newLink
	newLink.index = indexCtr
	l.size = indexCtr + 1
}

func (l *List) appendMultipleLinks(args ...int) {
	for _, x := range args {
		l.appendLink(x)
	}
}

//add link to start of list, push rest down
func (l *List) prependLink(num int) { // broken
	newLink := new(Link)
	if l.first == nil {
		l.first = newLink
	}
	newLink.val = num
	temp := l.first
	l.first = newLink
	newLink.next = temp
	l.size++
}

func (l List) prependMultipleLinks(args ...int) {
	for _, n := range args {
		l.prependLink(n)
	}
}

// func (l List) insertAtPosition(index, num int) {

// }

//this breaks everything :(
// func (l *List) removeFromList(remove int) {
// 	iter := l.first
// 	if iter == nil {
// 		fmt.Println("List is empty")
// 		return
// 	} else if iter.val == remove {
// 		l.first = nil
// 		if l.first.next != nil {
// 			l.first = l.first.next
// 		}
// 	}
// 	for iter.next.val != remove {
// 		iter = iter.next
// 	}
// 	iter.next = iter.next.next
// }

// func (l *List) removeAllInstancesOf(num int) {

// }

// func (l *List) removeFirstXInstancesOf(num, qty int) {

// }

func (l List) getListSize() int {
	return l.size
}

//fill and print slice containing values of list
func (l List) printList() {
	printer := l.first
	if printer == nil {
		fmt.Println("List empty")
		return
	}
	s := []int{}
	for printer.next != nil {
		s = append(s, printer.val)
		printer = printer.next
	}
	s = append(s, printer.val) //catches last element
	fmt.Println(s)
}

//print elems from zero-based position x to y in the list
func (l List) printFromXToY(x, y int) {
	if y < x {
		fmt.Println("I can't print backwards, dummy")
		return
	}
	printer := l.first
	if printer == nil {
		fmt.Println("List empty")
		return
	}
	ctr := 0
	for ctr < x {
		printer = printer.next
		ctr++
		if printer == nil {
			fmt.Println("Inner bound does not exist")
			return
		}
	}
	s := []int{}
	s = append(s, printer.val)
	for ctr < y {
		printer = printer.next
		ctr++
		if printer == nil {
			fmt.Println("Outer bound does not exist")
			fmt.Println(s)
			return
		}
		s = append(s, printer.val)
	}
	fmt.Println(s)
}

func (l List) numIsInList(num int) bool {
	iter := l.first
	if iter == nil {
		fmt.Println("Empty list")
		return false
	}
	for iter != nil {
		if iter.val == num {
			return true
		}
		iter = iter.next
	}
	return false
}

func main() {
	linkList := createEmptyList()
	linkList.appendLink(5)
	linkList.appendLink(7)
	linkList.appendLink(9)
	linkList.prependLink(12)
	linkList.printList()
	//linkList.removeFromList(5)
	linkList.appendMultipleLinks(1, 2, 3, 4, 5)
	linkList.printFromXToY(2, 5)
	linkList.printList()
	fmt.Println(linkList.getListSize())

	anotherList := createEmptyList()
	anotherList.printList()
	anotherList.prependMultipleLinks(1, 5, 9, 13)
	anotherList.printList()

	moreList := createListContaining(4, 5, 6, 7, 8)
	moreList.printList()
	fmt.Println(moreList.getListSize())
	fmt.Println(moreList.numIsInList(7))  //outputs true
	fmt.Println(moreList.numIsInList(8))  //outputs true
	fmt.Println(moreList.numIsInList(12)) //outputs false
}
