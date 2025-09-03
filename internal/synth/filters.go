package synth

import "fmt"

// Fade is a simple envelope that controls fade-in and fade-out.
type Fade struct {
	Sound StaticStreamer
	In    Seconds
	Out   Seconds
}

func (self *Fade) Stream(sampling int, sink chan<- int16) error {
	totalSamples := int(self.Sound.Time() * float64(sampling))
	attackSamples := int(self.In * float64(sampling))
	releaseSamples := int(self.Out * float64(sampling))
	internalSink := make(chan int16, 1024)

	go self.Sound.Stream(sampling, internalSink)
	var inputSample int16
	var envelopeFactor float64
	for i := range totalSamples {
		inputSample = <-internalSink
		if i < attackSamples {
			envelopeFactor = float64(i) / float64(attackSamples)
		}
		if i >= totalSamples-attackSamples {
			envelopeFactor = 1.0 - float64(i-(totalSamples-attackSamples))/float64(releaseSamples)
		}
		sink <- int16(float64(inputSample) * envelopeFactor)
	}
	return nil
}

func (self *Fade) Time() Seconds {
	return self.Sound.Time()
}

// Envelope implements linear ADSR envelope.
type Envelope struct {
	Sound   StaticStreamer
	Attack  Seconds
	Decay   Seconds
	Sustain Amplitude
	Release Seconds
}

func (self *Envelope) Stream(sampling int, sink chan<- int16) error {
	if err := self.verify(); err != nil {
		return err
	}

	totalSamples := self.Sound.Time() * float64(sampling)
	attackSamples := self.Attack * float64(sampling)
	decaySamples := self.Decay * float64(sampling)
	releaseSamples := self.Release * float64(sampling)
	sustainSamples := totalSamples - attackSamples - decaySamples - releaseSamples
	internalSink := make(chan int16, 1024)

	go self.Sound.Stream(sampling, internalSink)
	var inputSample int16
	var envelopeFactor float64 = 0.0

	attackPoint := int(attackSamples)
	decayPoint := attackPoint + int(decaySamples)
	sustainPoint := decayPoint + int(sustainSamples)

	for i := range int(totalSamples) {
		inputSample = <-internalSink
		if i < attackPoint {
			envelopeFactor = float64(i) / attackSamples
		}
		if i >= attackPoint && i < decayPoint {
			offset := i - attackPoint
			envelopeFactor = 1.0 - (1.0-self.Sustain)*float64(offset)/decaySamples
		}
		if i >= decayPoint && i < sustainPoint {
			envelopeFactor = self.Sustain
		}
		if i >= sustainPoint {
			offset := i - sustainPoint
			envelopeFactor = self.Sustain * (1.0 - float64(offset)/releaseSamples)
		}
		sink <- int16(float64(inputSample) * envelopeFactor)
	}
	return nil
}

func (self *Envelope) Time() Seconds {
	return self.Sound.Time()
}

func (self *Envelope) verify() error {
	fxDuration := self.Attack + self.Decay + self.Release
	if fxDuration > self.Sound.Time() {
		return fmt.Errorf(
			"Envelope parameters are longer than sound (%.2f + %.2f + %.2f > %.2f)",
			self.Attack, self.Decay, self.Release, self.Sound.Time(),
		)
	}
	return nil
}
