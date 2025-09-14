package synth

import "math"

// LeakyIntegrator smoothens signal and applies kind of low pass filter.
type LeakyIntegrator struct {
	Sound StaticStreamer
}

func (self *LeakyIntegrator) Stream(sampling int, sink chan<- int16) error {
	internalSink := make(chan int16, 1024)
	errCh := make(chan error, 1)

	go func() {
		errCh <- self.Sound.Stream(sampling, internalSink)
		close(internalSink)
	}()

	const alpha = 0.01 // smoothing factor
	var integrated float64 = 0

	for inSample := range internalSink {
		normalizedInSample := float64(inSample) / math.MaxInt16
		integrated = min((1-alpha)*integrated + alpha*normalizedInSample, 1)
		outSample := int16(integrated * math.MaxInt16)
		sink <- outSample
	}

	return <-errCh
}

func (self *LeakyIntegrator) Time() Seconds {
	return self.Sound.Time()
}
