//Needs search method and delete method

package main

import (
	"fmt"
)

// Link is the main component of list
type Link struct {
	val  int
	next *Link
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
		return
	}
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = newLink
	l.size++
}

func (l *List) appendMultipleLinks(args ...int) {
	for _, x := range args {
		l.appendLink(x)
		l.size++
	}
}

//add link to start of list, push rest down
func (l List) prependLink(num int) { // broken
	newLink := new(Link)
	newLink.val = num
	temp := l.first
	l.first = newLink
	newLink.next = temp
}

// func (l List) prependMultipleLinks(args ...int) {
// 	for _, x := range args {
// 		//prependLink(x)
// 		x++
// 		l.size++
// 	}
// }

// func (l *List) removeFromList(remove int) {
// 	iter := l.first
// 	if iter == nil {
// 		fmt.Println("List is empty")
// 		return
// 	} else if iter.val == remove {
// 		l.first = nil
// 	}
// 	for iter.next.val != remove {
// 		iter = iter.next
// 	}
// 	iter.next = iter.next.next
// }

func (l List) getListSize() int { //this needs work
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

//print elems from location x to y in the list
func (l List) printFromXToY(x, y int) {
	printer := l.first
	for i := 0; i < x; i++ {
		if printer.next == nil {
			fmt.Println("Outside of list range")
			return
		}
		printer = printer.next
	}
	//s := []int{}

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
	// linkList.removeFromList(5)
	linkList.appendMultipleLinks(1, 2, 3, 4, 5)
	linkList.printList()
	fmt.Println(linkList.getListSize())

	anotherList := createEmptyList()
	anotherList.printList()

	moreList := createListContaining(4, 5, 6, 7, 8)
	moreList.printList()
	fmt.Println(moreList.getListSize())
	fmt.Println(moreList.numIsInList(7))
	fmt.Println(moreList.numIsInList(12))
}
