package synth

import "math"

// Mixer combines streams.
type Mixer []StaticStreamer

func (self Mixer) Stream(sampling int, sink chan<- int16) error {
	if len(self) == 0 {
		return nil
	}

	// Create a channel per sound.
	channels := make([]chan int16, len(self))
	errCh := make(chan error, len(self))

	// Start streaming all sounds concurrently.
	for i, sound := range self {
		ch := make(chan int16, 1024)
		channels[i] = ch

		go func(i int, s Streamer, c chan int16) {
			err := s.Stream(sampling, c)
			close(c)
			errCh <- err
		}(i, sound, ch)
	}

	// Mix samples until all channels are closed.
	for {
		activeChannels := 0
		var sum int32 = 0

		for _, ch := range channels {
			sample, ok := <-ch
			if ok {
				sum += int32(sample)
				activeChannels++
			}
		}

		if activeChannels == 0 {
			break
		}

		// Normalize by number of active channels to avoid clipping.
		if activeChannels > 1 {
			sum /= int32(activeChannels)
		}

		// Clamp to int16 range.
		if sum > math.MaxInt16 {
			sum = math.MaxInt16
		} else if sum < math.MinInt16 {
			sum = math.MinInt16
		}

		sink <- int16(sum)
	}

	// Collect errors from all goroutines
	for range self {
		if err := <-errCh; err != nil {
			return err
		}
	}

	return nil
}

func (self Mixer) Time() Seconds {
	maxDur := 0.0
	for _, streamer := range self {
		if dur := streamer.Time(); dur > maxDur {
			maxDur = dur
		}
	}
	return maxDur
}
