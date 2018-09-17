package tenpin

import "github.com/Tolk-Haggard/go-bowling-kata/scorer"

type TenPinScorer struct {
}

func (i *TenPinScorer) RollBall(pins int) {
}

func (i *TenPinScorer) CalculateScore() (score int) {
	return -1
}

func NewTenPinScorer() scorer.Scorer {
	return &TenPinScorer{}
}
