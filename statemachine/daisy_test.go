package statemachine_test

import (
	"testing"

	"github.com/Tolk-Haggard/go-bowling-kata/statemachine"
	"github.com/stretchr/testify/assert"
)

func Test_GetStateReturnsInitialState(t *testing.T) {
	testObject := statemachine.NewDaisy()

	assert.Equal(t, "Sleeping", testObject.GetState())
}

func Test_MakingNoiseWakesASleepingDaisy(t *testing.T) {
	testObject := statemachine.NewDaisy()

	testObject.HandleEvent("MakeNoise")

	assert.Equal(t, "Barking", testObject.GetState())
}

func Test_PettingABarkingDaisyStartsChewing_1(t *testing.T) {
	testObject := statemachine.NewDaisy()

	testObject.HandleEvent("MakeNoise")
	testObject.HandleEvent("Pet")

	assert.Equal(t, "Chewing", testObject.GetState())
}

func Test_PettingABarkingDaisyStartsChewing_2(t *testing.T) {
	testObject := statemachine.NewDaisy()

	testObject.SetState("Barking") //Method introduced for testing
	testObject.HandleEvent("Pet")

	assert.Equal(t, "Chewing", testObject.GetState())
}
