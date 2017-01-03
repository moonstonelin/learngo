package main

import (
	"fmt"
)

func main() {
    var grid [][]byte =
    {
      {1,1,1,1,0},
      {1,1,0,1,0},
      {1,1,0,0,0},
      {0,0,0,0,0}
    }
    
    ans := numIslands(grid)
}

func numIslands(grid [][]byte) int { 
	R := cap(grid)
	C := cap(grid[0])
	
	color := 2
	
	for i:=0; i<R; i++ {
		for j:=0; j<C; j++ {
			if grid[i][j] == 1 {
				Helper(&grid, i, j, color)
				color += 1
			}
		}
	}
	
	return color - 2
}

func Helper(grid *[][]byte, i int, j int, color int) {
	R := cap(*grid)
	C := cap((*grid)[0])
	
	(*grid)[i][j] = byte(color)
	
	var dirx []int = []int{0, -1, 0, 1}
	var diry []int = []int{-1, 0, 1, 0}
	
	for i:=0; i<4; i++ {
		r := i + dirx[i]
		c := j + diry[i]
		
		if r>=0 && r<R && c>=0 && c<C && (*grid)[r][c] == 1 {
			Helper(grid, r, c, color)
		}
	}
}
