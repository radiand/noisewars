package synth

type Amplitude = float64
type Frequency = float64
type Seconds = float64

// Sound generates 16bit samples at given sampling rate.
type Sound interface {
	Stream(sampling int, sink chan<- int16) error
}

// Finite defines sounds that have known duration.
type Finite interface {
	Time() Seconds
}

// FiniteSound is streamable and finite.
type FiniteSound interface {
	Sound
	Finite
}

// Sequence organizes Sounds in order.
type Sequence []Sound

func (self Sequence) Stream(sampling int, sink chan<- int16) error {
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
	Sound Sound
}

func (self *Infinite) Stream(sampling int, sink chan<- int16) error {
	for {
		err := self.Sound.Stream(sampling, sink)
		if err != nil {
			return err
		}
	}
}
