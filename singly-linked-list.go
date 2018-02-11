//This is pretty broken right now

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

func appendLink(num int) {
	newLink := *Link
	newLink.val, newLink.next = num, nil
	
	var iter *Link = List.first
	if iter == nil {
		List.first = x
	}
	for x.next != nil {
		x = x.next
	}
	var appended = Link{val, nil}
	x.next = &appended
}

//return int slice containing values of list
func (l List) printList() {
	x := l.first
	s := []int{}
	for x.next != nil {
		s = append(s, x.val)
		x = x.next
	}
	fmt.Println(s)
}

func main() {
	listLink := Link{10, nil}
	linkList := List{&listLink}
	linkList.appendLink(5)
	linkList.printList()
}
