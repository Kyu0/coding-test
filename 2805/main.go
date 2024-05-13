package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// 백준 플랫폼 이용 시 표준 입출력 처리 시간으로 인해 버퍼 도입 (bufio)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

// trees: 입력받은 나무들의 길이가 저장된 슬라이스, mid: 자를 높이
func getTotalLength(trees []int64, mid int64) int64 {
	result := int64(0)

	for _, value := range trees {
		if value > mid {
			result += value - mid
		}
	}

	return result
}

// n: 입력받을 나무의 개수, m: 가져갈 나무의 총 길이, trees: 입력받은 나무들의 길이가 저장된 슬라이스
func input() (n int, m int64, trees []int64) {
	fmt.Fscanln(reader, &n, &m)
	trees = make([]int64, 0, n)

	for i := 0; i < n; i++ {
		var temp int64
		fmt.Fscan(reader, &temp)
		trees = append(trees, temp)
	}

	slices.Sort(trees)

	return n, m, trees
}

func main() {
	defer writer.Flush()

	var answer int64
	_, m, trees := input()

	left, right := int64(0), trees[len(trees)-1]

	// 이분 탐색으로 적절한 길이의 최댓값을 찾는다.
	for left <= right {
		mid := (left + right) / 2

		if m <= getTotalLength(trees, mid) {
			answer = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Fprint(writer, answer)
}
