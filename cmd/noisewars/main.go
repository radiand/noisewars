/*
noisewars - just enough synthesizer to piss off your neighbors!
*/
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)
import . "github.com/radiand/noisewars/internal/synth"

func main() {
	presetFlag := flag.String("preset", "TestPreset", "Name of preset to play")
	rateFlag := flag.Int("rate", 44100, "Sample rate in Hz")
	loopFlag := flag.Bool("loop", false, "Play sequence repeatedly until interrupted")
	flag.Parse()

	sequence, ok := Presets[*presetFlag]
	if !ok {
		fmt.Fprintf(os.Stderr, "Preset %q not found\n", *presetFlag)
		os.Exit(1)
	}

	var wrappedSequence Streamer = sequence
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
