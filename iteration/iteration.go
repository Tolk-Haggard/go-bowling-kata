package iteration

import "github.com/Tolk-Haggard/go-bowling-kata/scorer"

type IterativeScorer struct {
	score int
}

func (i *IterativeScorer) RollBall(pins int) {
	i.score += pins
}

func (i *IterativeScorer) CalculateScore() int {
	return i.score
}

func NewIterativeScorer() scorer.Scorer {
	return &IterativeScorer{}
}
