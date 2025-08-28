package synth

// LinearAD is a simple envelope that controls fade-in and fade-out.
type LinearAD struct {
	Sound  FiniteSound
	Attack float64
	Decay  float64
}

func (self *LinearAD) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Sound.Time() * float64(sampling))
	attackSamples := int(self.Attack * float64(sampling))
	decaySamples := int(self.Decay * float64(sampling))
	internalSink := make(chan int16, 1024)

	go self.Sound.Stream(sampling, internalSink)
	var inputSample int16
	var envelopeFactor float64
	for i := range totalSamples {
		inputSample = <-internalSink
		if i < attackSamples {
			envelopeFactor = float64(i) / float64(attackSamples)
		}
		if i >= totalSamples-attackSamples {
			envelopeFactor = 1.0 - float64(i-(totalSamples-attackSamples))/float64(decaySamples)
		}
		sink <- int16(float64(inputSample) * envelopeFactor)
	}
	return nil
}
