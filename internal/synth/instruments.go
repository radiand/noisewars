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

func (self *Pause) Time() Seconds {
	return self.Duration
}

func Punch(amplitude Amplitude, duration Seconds, frequency Frequency) FiniteStreamer {
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
	Sound    Streamer
	MinPause Seconds
	MaxPause Seconds
}

func (self *Chaotic) Stream(sampling int, sink chan<- int16) error {
	self.Sound.Stream(sampling, sink)
	minPauseMs := int(self.MinPause * 1000)
	maxPauseMs := int(self.MaxPause * 1000)
	pause := &Pause{Duration: float64(rand.Intn(maxPauseMs-minPauseMs)+minPauseMs) / 1000}
	pause.Stream(sampling, sink)
	return nil
}
