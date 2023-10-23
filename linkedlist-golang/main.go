package main

import "fmt"


type Node struct{
	data int
	nextNode *Node
}


type LinkedList struct{
	head *Node
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
	
	ll.insertAtHead(46)
	ll.insertAtHead(40)
	ll.insertAtHead(99)
	ll.insertFromTail(69420)
	fmt.Println("")
	ll.print()
}