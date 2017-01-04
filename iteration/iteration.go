package iteration

import "github.com/Tolk-Haggard/go-bowling-kata/scorer"

type IterativeScorer struct {
}

func (I *IterativeScorer) RollBall(pins int) {
}

func (I *IterativeScorer) CalculateScore() int {
	return -1
}

func NewIterativeScorer() scorer.Scorer {
	return &IterativeScorer{}
}
