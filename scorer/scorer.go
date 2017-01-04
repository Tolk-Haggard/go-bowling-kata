package scorer

type Scorer interface {
	RollBall(pins int)
	CalculateScore() int
}
