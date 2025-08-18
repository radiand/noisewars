/*
noisewars - just enough synthesizer to piss off your neighbors!
*/
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"sync"
)
import . "github.com/radiand/noisewars/internal/synth"

// Player manages playback, looping, and graceful shutdown.
type Player struct {
	sampleRate int
	out        chan int16
	done       chan struct{}
	wg         sync.WaitGroup
}

func NewPlayer(sampleRate int, out chan int16) *Player {
	return &Player{
		sampleRate: sampleRate,
		out:        out,
		done:       make(chan struct{}),
	}
}

// Play runs the given Sound until done or error.
func (self *Player) Play(s Sound) error {
	self.wg.Add(1)
	defer self.wg.Done()
	defer close(self.out)

	playErrCh := make(chan error, 1)

	go func() {
		playErrCh <- s.Stream(self.sampleRate, self.out)
	}()

	select {
	case <-self.done:
		return nil
	case err := <-playErrCh:
		return err
	}
}

// Stop signals the player to stop playback.
func (self *Player) Stop() {
	close(self.done)
	self.wg.Wait()
}

var presets = map[string]Sequence{
	"TestPreset": {
		&Infinite{
			Sound: &Sequence{
				&Punch{Attack: 0.01, Decay: 0.1, Sustain: 0.3, Release: 0.2, Duration: 0.4, Frequency: 80, Amplitude: 1.0},
				&Pause{Duration: 0.4},
				&Punch{Attack: 0.01, Decay: 0.1, Sustain: 0.3, Release: 0.2, Duration: 0.4, Frequency: 60, Amplitude: 1.0},
				&Pause{Duration: 0.4},
			},
		},
	},
}

func main() {
	presetFlag := flag.String("preset", "TestPreset", "Name of preset to play")
	rateFlag := flag.Int("rate", 44100, "Sample rate in Hz")
	loopFlag := flag.Bool("loop", false, "Play sequence repeatedly until interrupted")
	flag.Parse()

	sequence, ok := presets[*presetFlag]
	if !ok {
		fmt.Fprintf(os.Stderr, "Preset %q not found\n", *presetFlag)
		os.Exit(1)
	}

	var wrappedSequence Sound = sequence
	if *loopFlag {
		wrappedSequence = &Infinite{Sound: sequence}
	}

	sampling := *rateFlag
	sink := make(chan int16, 1024)

	go func() {
		for sample := range sink {
			err := binary.Write(os.Stdout, binary.LittleEndian, sample)
			if err != nil {
				break
			}
		}
	}()

	errorChannel := make(chan error, 1)
	go func() {
		errorChannel <- wrappedSequence.Stream(sampling, sink)
	}()

	select {
	case err := <-errorChannel:
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error playing sound: %v\n", err)
		}
	}
}
