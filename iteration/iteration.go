package iteration

import "github.com/Tolk-Haggard/go-bowling-kata/scorer"

type IterativeScorer struct {
	rolls       [21]int
	currentRoll int
}

func (i *IterativeScorer) RollBall(pins int) {
	i.rolls[i.currentRoll] = pins
	i.currentRoll++
}

func (i *IterativeScorer) CalculateScore() (score int) {
	for j := 0; j < 21; j++ {
		score += i.rolls[j]
	}
	return
}

func NewIterativeScorer() scorer.Scorer {
	return &IterativeScorer{}
}
