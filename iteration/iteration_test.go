package iteration_test

import (
	"testing"

	"github.com/Tolk-Haggard/go-bowling-kata/iteration"
	"github.com/Tolk-Haggard/go-bowling-kata/scorer"
	"github.com/stretchr/testify/assert"
)

func Test_ReturnsIterativeScoreGutterGame(t *testing.T) {
	testObject := iteration.NewIterativeScorer()

	rollGutterBalls(testObject, 20)

	assert.Equal(t, 0, testObject.CalculateScore())
}

func Test_ReturnsIterativeScoreSinglePinGame(t *testing.T) {
	testObject := iteration.NewIterativeScorer()

	testObject.RollBall(1)
	rollGutterBalls(testObject, 19)

	assert.Equal(t, 1, testObject.CalculateScore())
}

func Test_ReturnsIterativeScoreSingleSpareGame(t *testing.T) {
	testObject := iteration.NewIterativeScorer()

	rollSpare(testObject)
	testObject.RollBall(3)
	rollGutterBalls(testObject, 17)

	assert.Equal(t, 16, testObject.CalculateScore())
}

func Test_ReturnsIterativeScoreSingleStrikeGame(t *testing.T) {
	testObject := iteration.NewIterativeScorer()

	rollStrike(testObject)
	testObject.RollBall(3)
	testObject.RollBall(3)
	rollGutterBalls(testObject, 17)

	assert.Equal(t, 22, testObject.CalculateScore())
}

func rollStrike(s scorer.Scorer) {
	s.RollBall(10)
}

func rollSpare(s scorer.Scorer) {
	s.RollBall(5)
	s.RollBall(5)
}

func rollGutterBalls(s scorer.Scorer, balls int) {
	for i := 0; i < balls; i++ {
		s.RollBall(0)
	}
}
