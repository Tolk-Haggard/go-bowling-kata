package correctness

import "github.com/Tolk-Haggard/go-bowling-kata/scorer"

type CorrectScorer struct{}

func (c *CorrectScorer) RollBall(pins int) {
}

func (c *CorrectScorer) CalculateScore() (score int) {
	return 0
}

func NewCorrectScorer() scorer.Scorer {
	return &CorrectScorer{}
}
