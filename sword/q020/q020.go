package q020

func Answer(matrix [][]int) (res []int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	total := n * m
	var layer, startN, endN, startM, endM int
	for total > 0 {
		startN, startM, endN, endM = layer, layer, n-layer-1, m-layer-1
		for i := startM; i <= endM && total > 0; i++ {
			res = append(res, matrix[startN][i])
			total--
		}
		for i := startN + 1; i <= endN && total > 0; i++ {
			res = append(res, matrix[i][endM])
			total--
		}
		for i := endM - 1; i >= startM && total > 0; i-- {
			res = append(res, matrix[endN][i])
			total--
		}
		for i := endN - 1; i >= startN+1 && total > 0; i-- {
			res = append(res, matrix[i][startM])
			total--
		}
		layer++
	}
	return
}
