package synth

import "math"

type Pause struct {
	Duration float64
}

func (self *Pause) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration * float64(sampling))
	for range totalSamples {
		sink <- 0
	}
	return nil
}

type Punch struct {
	Attack, Decay, Sustain, Release float64
	Duration, Frequency             float64
	Amplitude                       float64
}

func (self *Punch) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration * float64(sampling))
	attackSamples := int(self.Attack * float64(sampling))
	decaySamples := int(self.Decay * float64(sampling))
	releaseSamples := int(self.Release * float64(sampling))
	sustainSamples := max(totalSamples-attackSamples-decaySamples-releaseSamples, 0)

	angularFreq := 2 * math.Pi * self.Frequency / float64(sampling)

	for i := range totalSamples {
		var env float64
		switch {
		case i < attackSamples:
			env = float64(i) / float64(attackSamples)
		case i < attackSamples+decaySamples:
			pos := i - attackSamples
			env = 1.0 - (1.0-self.Sustain)*float64(pos)/float64(decaySamples)
		case i < attackSamples+decaySamples+sustainSamples:
			env = self.Sustain
		case i < totalSamples:
			pos := i - (attackSamples + decaySamples + sustainSamples)
			env = self.Sustain * (1.0 - float64(pos)/float64(releaseSamples))
		default:
			env = 0
		}

		sample := self.Amplitude * env * math.Sin(angularFreq*float64(i))
		intSample := int16(sample * math.MaxInt16)

		sink <- intSample
	}
	return nil
}
