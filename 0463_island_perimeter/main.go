package _463_island_perimeter

func islandPerimeter(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++{
		for j:=0; j < len(grid[i]); j++{

			if grid[i][j] == 1 {

				if i+1 > len(grid)-1 || grid[i+1][j] == 0 { res++ }
				if i-1 < 0 || grid[i-1][j] == 0 { res++ }
				if j+1 > len(grid[i])-1 || grid[i][j+1] == 0 { res++ }
				if j-1 < 0 || grid[i][j-1] == 0 { res++ }
			}
		}
	}

	return res
}
