package main

import "fmt"

func heapify(tree []int, n, i int) {
	if i > n {
		return
	}

	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}

	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i {
		tree[max], tree[i] = tree[i], tree[max]
		heapify(tree, n, max)
	}
}

func buildHeap(tree []int, n int) {
	lastNode := n - 1
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

/*
第一个节点最大，和最后一个节点做交换，
砍断它，进行heapify，然后继续这个流程
*/
func heapSort(tree []int, n int) {
	buildHeap(tree, n)
	for i := n - 1; i >= 0; i-- {
		tree[i], tree[0] = tree[0], tree[i]
		heapify(tree, i, 0)
	}
}

func main() {
	// 这里只做第一个节点的heapify
	tree := []int{4, 10, 3, 5, 1, 2}
	n := 6
	//heapify(tree, n, 0)
	//for i := 0; i < n; i++ {
	//	fmt.Println(tree[i])
	//}

	// 构建最后一个堆,只要知道最后一个父节点，往上 heapify 即可
	//buildHeap(tree, n)
	//for i := 0; i < n; i++ {
	//	fmt.Println(tree[i])
	//}

	heapSort(tree, n)
	for i := 0; i < n; i++ {
		fmt.Println(tree[i])
	}

}
