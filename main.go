package main

//LinkedList Intersection Found
import "fmt"

type Node struct {
	Next *Node
	Val  int
}
type LinkedList struct {
	Head Node
}

func main() {
	//4->3->2->0
	n1 := &Node{Val: 0, Next: &Node{}}
	n2 := &Node{Val: 2, Next: n1}
	n3 := &Node{Val: 3, Next: n2}
	n4 := &Node{Val: 4, Next: n3}
	l1 := &LinkedList{Head: *n4}

	//1->2->0
	//nn1 := &Node{Val: 0, Next: &Node{}}
	//nn2 := &Node{Val: 2, Next: nn1}
	nn3 := &Node{Val: 1, Next: n2}
	l2 := &LinkedList{Head: *nn3}

	intersectionPoint, _ := FindIntersection(*l1, *l2)
	fmt.Println("IntersectionPoint", intersectionPoint)

	//Node{Node.Val: 1, No}
}

func FindIntersection(L1 LinkedList, L2 LinkedList) (int, error) {

	LinkedMap := make(map[int]int)

	for L1.Head.Next != nil {
		LinkedMap[L1.Head.Val] = -1
		L1.Head = *L1.Head.Next
	}
	for L2.Head.Next != nil {
		if _, ok := LinkedMap[L2.Head.Val]; ok {
			return L2.Head.Val, nil
		}
		L2.Head = *L2.Head.Next
	}

	// for L1.Head.Next != nil {
	// 	tempL2 := L2
	// 	for L2.Head.Next != nil {
	// 		if L1.Head.Val == L2.Head.Val {
	// 			return L1.Head.Val, nil
	// 		}
	// 		//L1.Head = *L1.Head.Next
	// 		L2.Head = *L2.Head.Next
	// 	}
	// 	L1.Head = *L1.Head.Next
	// 	L2 = tempL2
	// }

	return -1, nil
}
