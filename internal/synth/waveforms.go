package synth

import "math"

type Sine struct {
	Amplitude float64
	Frequency float64
	Duration  float64
}

func (self *Sine) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration * float64(sampling))
	angularFreq := 2 * math.Pi * self.Frequency / float64(sampling)

	for i := range totalSamples {
		sink <- int16(self.Amplitude * math.Sin(angularFreq*float64(i)) * math.MaxInt16)
	}
	return nil
}

func (self *Sine) Time() float64 {
	return self.Duration
}
