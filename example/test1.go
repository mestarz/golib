/* 假设我有 n 个线程，m 个任务，任务的执行时间是 (t1,t2,t3,t4.....tm)。一个线程在执行任务期间不能再执行其他任务。
   怎么投放任务，可以使得 m 个任务整体最快被执行完 */

package main

import "fmt"

func getR(n, m int, tasks []int) [][]int {
    result := make([][]int, m)
	return result
}

func main() {
	var m int
    var tmp int
	var tasks []int

	fmt.Scanln(&m)
    for i := 0; i < m; i++ {
        fmt.Scan(&tmp)
        tasks = append(tasks, tmp)
    }
}
