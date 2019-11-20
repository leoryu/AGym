package combination

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	return NewCombination(candidates, target).GetResults()
}

func NewCombination(candidates []int, target int) *Combination {
	c := new(Combination)
	c.target = target
	c.candidates = candidates
	c.results = [][]int{}
	c.sortCandidates()
	c.fillResults(0, target, []int{})
	return c
}

type Combination struct {
	target     int
	mem        int
	candidates []int
	results    [][]int
}

func (c Combination) GetResults() [][]int {
	return c.results
}

func (c *Combination) sortCandidates() {
	sort.Ints(c.candidates)
}

func (c *Combination) fillResults(index, remainTarget int, candidateResult []int) {
	switch {
	case remainTarget < 0:
		return
	case remainTarget == 0:
		result := make([]int, len(candidateResult))
		copy(result, candidateResult)
		c.results = append(c.results, result)
		fallthrough
	default:
		var mem = -1
		for i := index; i < len(c.candidates); i++ {
			if c.candidates[i] == mem {
				continue
			}
			mem = c.candidates[i]
			if c.candidates[i] > remainTarget {
				return
			}
			c.fillResults(i+1, remainTarget-c.candidates[i], append(candidateResult, c.candidates[i]))
		}
	}
}
