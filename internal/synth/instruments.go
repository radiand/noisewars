package synth

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

func Punch(amplitude Amplitude, duration Seconds, frequency Frequency) Sound {
	return &LinearADSR{
		Sound:   &Sine{Amplitude: amplitude, Frequency: frequency, Duration: duration},
		Attack:  0.05,
		Decay:   0.04,
		Sustain: 0.4,
		Release: 0.2,
	}
}
