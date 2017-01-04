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
	roll := 0
	for j := 0; j < 10; j++ {
		if i.rolls[roll] == 10 {
			score += 10 + i.rolls[j+1] + i.rolls[j+2]
			roll++
		} else if i.rolls[roll]+i.rolls[roll+1] == 10 {
			score += 10 + i.rolls[roll+2]
			roll += 2
		} else {
			score += i.rolls[roll] + i.rolls[roll+1]
			roll += 2
		}
	}
	return
}

func NewIterativeScorer() scorer.Scorer {
	return &IterativeScorer{}
}
