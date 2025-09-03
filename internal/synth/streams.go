package synth

// Streamer generates 16bit samples at given sampling rate.
type Streamer interface {
	Stream(sampling int, sink chan<- int16) error
}

// Timer defines sounds that have known duration.
type Timer interface {
	Time() Seconds
}

// StaticStreamer is streamable and has known, deterministic duration.
type StaticStreamer interface {
	Streamer
	Timer
}

// StaticSequence organizes streams in order.
type StaticSequence []StaticStreamer

func (self StaticSequence) Stream(sampling int, sink chan<- int16) error {
	for _, event := range self {
		err := event.Stream(sampling, sink)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self StaticSequence) Time() Seconds {
	overall := 0.0
	for _, streamer := range self {
		overall += streamer.Time()
	}
	return overall
}

// DynamicSequence organizes streams in order. It is more relaxed than
// StaticSequence because it does not require streams to have known,
// deterministic duration.
type DynamicSequence []Streamer

func (self DynamicSequence) Stream(sampling int, sink chan<- int16) error {
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
