package tenpin_test

import (
	"testing"

	"github.com/Tolk-Haggard/go-bowling-kata/tenpin"
	"github.com/stretchr/testify/assert"
)

func Test_ReturnsScoreGutterGame(t *testing.T) {
	testObject := tenpin.NewTenPinScorer()

	testObject.RollBall(0)

	assert.Equal(t, 0, testObject.CalculateScore())
}
