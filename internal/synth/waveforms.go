package synth

import "math"

type Sine struct {
	Amplitude Amplitude
	Duration  Seconds
	Frequency Frequency
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

// SweepSine linearly changes frequency over duration.
type SweepSine struct {
	Amplitude Amplitude
	Duration  Seconds
	StartFreq Frequency
	EndFreq   Frequency
}

func (self *SweepSine) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration * float64(sampling))
	for i := range totalSamples {
		t := float64(i) / float64(sampling)
		freq := self.StartFreq + (self.EndFreq-self.StartFreq)*(t/self.Duration)
		phase := 2 * math.Pi * freq * t
		sink <- int16(self.Amplitude * math.Sin(phase) * math.MaxInt16)
	}
	return nil
}

func (self *SweepSine) Time() float64 {
	return self.Duration
}
