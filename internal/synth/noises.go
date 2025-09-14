package synth

import "math/rand"
import "math"
import "time"

// WhiteNoise generates random samples. Use with caution as it may damage your
// hearing.
type WhiteNoise struct {
	Amplitude Amplitude
	Duration  Seconds
}

func (self *WhiteNoise) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Duration * float64(sampling))
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for range totalSamples {
		randNumber := 2*rng.Float64() - 1
		sample := int16(self.Amplitude * randNumber * math.MaxInt16)
		sink <- sample
	}
	return nil
}

func (self *WhiteNoise) Time() Seconds {
	return self.Duration
}
