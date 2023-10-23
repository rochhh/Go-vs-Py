package main

import (
	"fmt"
	"log"
)


type Node struct{
	data int
	nextNode *Node
}


type LinkedList struct{
	head *Node
}


func ( ll *LinkedList ) insertBetween ( data int , valueInLL int ) error {

	newNode := &Node{ data: data , nextNode: nil }
	node := ll.head

	if ll.head == nil {
		ll.insertAtHead( data )
		return nil
	}

	for node != nil {
		if node.data == valueInLL {
			newNode.nextNode = node.nextNode
			node.nextNode = newNode
			return nil
		}

		node = node.nextNode
	}

	return fmt.Errorf("value does not exist in  LinkedList")

}


func ( ll *LinkedList ) clear (){
	ll.head = nil
}


func ( ll *LinkedList ) insertFromTail (data int ) {

	if ll.head == nil {
		ll.head = &Node{ data: data , nextNode: nil }
		return
	}

	node := ll.head
	newNode := &Node{ data: data , nextNode: nil }

	for node.nextNode != nil {
		node = node.nextNode
	}

	node.nextNode  = newNode
	
}


func ( ll *LinkedList ) insertAtHead( data int ){

	if ll.head == nil {
		ll.head = &Node{ data: data , nextNode: nil }
		return
	}

	newNode := &Node{ data: data  , nextNode: ll.head }

	ll.head = newNode

}


func ( ll *LinkedList ) print (){
	node := ll.head

	for node != nil {
		fmt.Printf( "%v  ->" , node.data)
		node = node.nextNode
	}
 
}



func main() {
	
	node4 := Node{ 4 ,  nil }
	node3 := Node{ 3 , &node4  }
	node2 := Node{ 2 , &node3  }
	node1 := Node{ 1 , &node2  }

	ll := LinkedList{head: &node1}
	// ll := LinkedList{}
	
	ll.insertAtHead(46)
	ll.insertAtHead(6)
	ll.insertAtHead(40)
	ll.insertAtHead(99)
	ll.insertFromTail(69420)
	ll.insertBetween(34 , 8 )

	err := ll.insertBetween(34 , 46)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	ll.print()
	
	fmt.Println("")
	fmt.Println("below ")
	ll.clear()
	fmt.Println("printing below -> ")
	ll.print()
	fmt.Println("----")

}