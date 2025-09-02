package synth

import "math/rand"

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

func Punch(amplitude Amplitude, duration Seconds, frequency Frequency) Streamer {
	return &Envelope{
		Sound:   &Sine{Amplitude: amplitude, Duration: duration, Frequency: frequency},
		Attack:  0.05,
		Decay:   0.04,
		Sustain: 0.4,
		Release: 0.2,
	}
}

// Chaotic adds random pause after the sound.
type Chaotic struct {
	Sound Streamer
	Pause Bound[Milliseconds, Milliseconds]
}

func (self *Chaotic) Stream(sampling int, sink chan<- int16) error {
	self.Sound.Stream(sampling, sink)
	pause := &Pause{Duration: float64(rand.Intn(self.Pause.Max-self.Pause.Min)+self.Pause.Min) / 1000}
	pause.Stream(sampling, sink)
	return nil
}
