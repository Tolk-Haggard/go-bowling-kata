package correctness

import "github.com/Tolk-Haggard/go-bowling-kata/scorer"

type CorrectScorer struct {
	firstRoll   bool
	doublers    [21]int
	currentRoll int
	rolls       [21]int
	frameCount  int
}

func (c *CorrectScorer) RollBall(pins int) {
	c.rolls[c.currentRoll] = pins
	if c.frameCount >= 10 {
		c.doublers[c.currentRoll] = c.doublers[c.currentRoll] - 1
	} else {
		if c.firstRoll {
			if pins == 10 {
				c.doublers[c.currentRoll+1] = c.doublers[c.currentRoll+1] + 1
				c.doublers[c.currentRoll+2] = c.doublers[c.currentRoll+2] + 1
				c.firstRoll = false
				c.frameCount++
			}
		} else {
			if pins+c.rolls[c.currentRoll-1] == 10 {
				c.doublers[c.currentRoll+1] = c.doublers[c.currentRoll+1] + 1
			}
			c.frameCount++
		}
	}

	c.firstRoll = !c.firstRoll
	c.currentRoll++
}

func (c *CorrectScorer) CalculateScore() (score int) {
	for i := 0; i < 21; i++ {
		score = score + c.rolls[i]*(c.doublers[i]+1)
	}
	return
}

func NewCorrectScorer() scorer.Scorer {
	return &CorrectScorer{firstRoll: true}
}
