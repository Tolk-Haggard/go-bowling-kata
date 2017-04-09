package statemachine

type Daisy struct {
	state string
}

func NewDaisy() *Daisy {
	return &Daisy{state: "Sleeping"}
}

func (d *Daisy) HandleEvent(event string) {
	d.state = "Barking"
}

func (d *Daisy) GetState() string {
	return d.state
}
