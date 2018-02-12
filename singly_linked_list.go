//Needs constructor method, search method and delete method

package main

import (
	"fmt"
)

type Link struct {
	val int
	next *Link
}

type List struct {
	first *Link
}

func (l List) appendLink(num int) {
	newLink :=  new(Link)
	newLink.val, newLink.next = num, nil
	
	var iter *Link = l.first
	if iter == nil {
		l.first = newLink
	}
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = newLink
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

func main() {
	listLink := Link{10, nil}
	linkList := List{&listLink}
	linkList.appendLink(5)
	linkList.appendLink(7)
	linkList.appendLink(9)
	linkList.printList()
}
