package combination

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	return NewCombination(candidates, target).GetResults()
}

func NewCombination(candidates []int, target int) *Combination {
	c := new(Combination)
	c.target = target
	c.candidates = candidates
	c.results = [][]int{}
	c.sortCandidates()
	c.fillResults(target, []int{})
	return c
}

type Combination struct {
	target     int
	candidates []int
	results    [][]int
}

func (c Combination) GetResults() [][]int {
	return c.results
}

func (c *Combination) sortCandidates() {
	sort.Ints(c.candidates)
}

func (c *Combination) fillResults(remainTarget int, candidateResult []int) {
	switch {
	case remainTarget < 0:
		return
	case remainTarget == 0:
		result := make([]int, len(candidateResult))
		copy(result, candidateResult)
		c.results = append(c.results, result)
	default:
		for i := 0; i < len(c.candidates); i++ {
			if remainTarget < c.candidates[i] {
				return
			}
			if len(candidateResult) != 0 && candidateResult[len(candidateResult)-1] > c.candidates[i] {
				continue
			}
			c.fillResults(remainTarget-c.candidates[i], append(candidateResult, c.candidates[i]))
		}
	}
}
