package correctness_test

import (
	"testing"

	"github.com/Tolk-Haggard/go-bowling-kata/correctness"
	"github.com/stretchr/testify/assert"
)

func Test_ReturnsCorrectScoreGutterGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	for i := 0; i < 20; i++ {
		testObject.RollBall(0)
	}

	assert.Equal(t, 0, testObject.CalculateScore())
}

func Test_ReturnsCorrectScoreComplexGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	testObject.RollBall(10) //strike
	testObject.RollBall(10) //strike
	testObject.RollBall(5)  //spare 1
	testObject.RollBall(5)  //spare 2
	testObject.RollBall(1)
	testObject.RollBall(1)
	testObject.RollBall(9)
	testObject.RollBall(0)

	for i := 0; i < 13; i++ {
		testObject.RollBall(0)
	}

	assert.Equal(t, 25+20+11+2+9, testObject.CalculateScore())
}

func Test_ReturnsCorrectScorePerfectGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	for i := 0; i < 12; i++ {
		testObject.RollBall(10) //strike
	}

	assert.Equal(t, 300, testObject.CalculateScore())
}

func Test_ReturnsCorrectScoreHeartBreakerGame(t *testing.T) {
	testObject := correctness.NewCorrectScorer()

	for i := 0; i < 10; i++ {
		testObject.RollBall(10) //strike
	}
	testObject.RollBall(9)
	testObject.RollBall(1)

	assert.Equal(t, 289, testObject.CalculateScore())
}
