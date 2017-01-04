package correctness_test

import (
	"testing"

	"github.com/Tolk-Haggard/go-bowling-kata/correctness"
	"github.com/Tolk-Haggard/go-bowling-kata/scorer"
	"github.com/stretchr/testify/assert"
)

func Test_ReturnsCorrectScoreGutterGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	rollGutterBalls(testObject, 20)

	assert.Equal(t, 0, testObject.CalculateScore())
}

func Test_ReturnsCorrectScoreComplexGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	rollStrike(testObject)
	rollStrike(testObject)
	rollSpare(testObject)
	testObject.RollBall(1)
	testObject.RollBall(1)
	testObject.RollBall(9)
	testObject.RollBall(0)

	rollGutterBalls(testObject, 13)

	assert.Equal(t, 25+20+11+2+9, testObject.CalculateScore())
}

func Test_ReturnsCorrectScorePerfectGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	for i := 0; i < 12; i++ {
		rollStrike(testObject)
	}

	assert.Equal(t, 300, testObject.CalculateScore())
}

func Test_ReturnsCorrectScoreHeartBreakerGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	for i := 0; i < 10; i++ {
		rollStrike(testObject)
	}
	testObject.RollBall(9)
	testObject.RollBall(1)

	assert.Equal(t, 289, testObject.CalculateScore())
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
