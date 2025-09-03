package synth

// Pause generates silence for given duration.
type Pause struct {
	Duration Seconds
}

func (self *Pause) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration * float64(sampling))
	for range totalSamples {
		sink <- 0
	}
	return nil
}

func (self *Pause) Time() Seconds {
	return self.Duration
}

// VaryingPause generates silence for given duration. Duration may be
// nondeterministic, thus Time() is not implemented.
type VaryingPause struct {
	Duration Provider[Seconds]
}

func (self *VaryingPause) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration() * float64(sampling))
	for range totalSamples {
		sink <- 0
	}
	return nil
}

func Punch(amplitude Amplitude, duration Seconds, frequency Frequency) StaticStreamer {
	return &Envelope{
		Sound:   &Sine{Amplitude: amplitude, Duration: duration, Frequency: frequency},
		Attack:  0.05,
		Decay:   0.04,
		Sustain: 0.4,
		Release: 0.2,
	}
}
