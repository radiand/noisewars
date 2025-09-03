package synth

type Amplitude = float64
type Frequency = float64
type Seconds = float64

// Streamer generates 16bit samples at given sampling rate.
type Streamer interface {
	Stream(sampling int, sink chan<- int16) error
}

// Finite defines sounds that have known duration.
type Finite interface {
	Time() Seconds
}

// FiniteStreamer is streamable and finite.
type FiniteStreamer interface {
	Streamer
	Finite
}

// Sequence organizes Streamers in order.
type Sequence []FiniteStreamer

func (self Sequence) Stream(sampling int, sink chan<- int16) error {
	for _, event := range self {
		err := event.Stream(sampling, sink)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self Sequence) Time() Seconds {
	overall := 0.0
	for _, streamer := range self {
		overall += streamer.Time()
	}
	return overall
}

// Stream sets Streamers in order and does not require them to have known duration.
type Stream []Streamer

func (self Stream) Stream(sampling int, sink chan<- int16) error {
	for _, event := range self {
		err := event.Stream(sampling, sink)
		if err != nil {
			return err
		}
	}
	return nil
}

// Infinite plays same sound infinitely.
type Infinite struct {
	Sound Streamer
}

func (self *Infinite) Stream(sampling int, sink chan<- int16) error {
	for {
		err := self.Sound.Stream(sampling, sink)
		if err != nil {
			return err
		}
	}
}
