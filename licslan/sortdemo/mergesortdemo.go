package main

func main() {

	//归并排序  将数据分为左右2半  分别归并排序  再把2个有序数据归并  2路归并
	//[1,3,4,7],[2,4,1,5,6]  -->1
	//[3,4,7],[2,4,1,5,6]  -->1
	//[3,4,7],[2,4,5,6]  -->2
	//[3,4,7],[4,5,6]  -->3
	//[4,7],[4,5,6]  -->4
	//......
	//-------------------> 1 1 2 3 4 4 5 6 7

	//外部排序  k路归并  多个节点排序后进行归并  利用堆进行排序

	//pipeline
	//source--->节点--->节点--->节点---->节点---> sink (只出不进)
	//传统语言  对象 方法调用 对象   go语言  goroutine channel goroutine



	
}
