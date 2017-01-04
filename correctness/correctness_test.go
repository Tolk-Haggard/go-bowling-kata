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
