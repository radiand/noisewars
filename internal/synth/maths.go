package synth

import "math"

// Integrator integrates underlying sound.
type Integrator struct {
	Sound StaticStreamer
}

func (self *Integrator) Stream(sampling int, sink chan<- int16) error {
	internalSink := make(chan int16, 1024)
	errCh := make(chan error, 1)

	go func() {
		errCh <- self.Sound.Stream(sampling, internalSink)
		close(internalSink)
	}()

	const alpha = 0.01 // smoothing factor, tweak for desired effect
	var integrated float64 = 0

	for sample := range internalSink {
		normSample := float64(sample) / math.MaxInt16

		// Leaky integration (lowpass smoothing)
		integrated = min((1-alpha)*integrated + alpha*normSample, 1)

		out := int16(integrated * math.MaxInt16)
		sink <- out
	}

	return <-errCh
}

func (self *Integrator) Time() Seconds {
	return self.Sound.Time()
}
