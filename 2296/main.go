package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Building struct {
	x, y, cost int
}

// 백준 플랫폼 이용 시 표준 입출력 처리 시간으로 인해 버퍼 도입 (bufio)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func input() (n int, buildings []Building) {
	fmt.Fscanln(reader, &n)
	buildings = make([]Building, 0, n)

	for i := 0; i < n; i++ {
		var x, y, c int
		fmt.Fscanln(reader, &x, &y, &c)
		buildings = append(buildings, Building{x: x, y: y, cost: c})
	}

	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x < buildings[j].x
	})

	return
}

func main() {
	defer writer.Flush()

	var answer int
	n, buildings := input()
	// dp[i][0] 은 i번째 빌딩까지 있을 때의 오름차순의 최대 이익
	// dp[i][1] 은 i번째 빌딩까지 있을 때의 내림차순의 최대 이익
	dp := make([][2]int, n*2)

	for i, building := range buildings {
		dp[i][0] = building.cost
		dp[i][1] = building.cost
		answer = max(answer, dp[i][0], dp[i][1])
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if buildings[i].y > buildings[j].y {
				dp[i][0] = max(dp[i][0], dp[j][0]+buildings[i].cost)
				answer = max(answer, dp[i][0])
			} else {
				dp[i][1] = max(dp[i][1], dp[j][1]+buildings[i].cost)
				answer = max(answer, dp[i][1])
			}
		}
	}

	fmt.Fprint(writer, answer)
}
