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
			score += i.handleStrike(roll)
			roll--
		} else if i.rolls[roll]+i.rolls[roll+1] == 10 {
			score += i.handleSpare(roll)
		} else {
			score += i.rolls[roll] + i.rolls[roll+1]
		}
		roll += 2
	}
	return
}

func (i *IterativeScorer) handleStrike(roll int) int {
	return 10 + i.rolls[roll+1] + i.rolls[roll+2]
}

func (i *IterativeScorer) handleSpare(roll int) int {
	return 10 + i.rolls[roll+2]
}

func NewIterativeScorer() scorer.Scorer {
	return &IterativeScorer{}
}
