package synth

var Presets = map[string]Streamer{
	"P2": &Sequence{
		&Fade{
			Sound: &Sine{Amplitude: 1.0, Frequency: 120.0, Duration: 1},
			In:    0.4,
			Out:   0.4,
		},
		&Pause{Duration: 0.2},
	},
	"P3": &Sequence{
		Punch(1.0, 0.3, 60),
		&Pause{Duration: 0.2},
	},
	"P4": &Chaotic{Punch(1.0, 0.3, 60), 0.01, 0.4},
	"P5": &Envelope{
		&Sequence{
			&SweepSine{1.0, 0.4, 50, 80},
			&Sine{1.0, 0.2, 80},
			&SweepSine{1.0, 0.4, 80, 50},
		},
		0.05,
		0.05,
		0.6,
		0.1,
	},
	"P6": &Sequence{
		&Envelope{
			&VaryingSine{randomf64(0.5, 1.0), 0.3, randomf64(40.0, 60.0)},
			0.05,
			0.04,
			0.6,
			0.1,
		},
		&Pause{0.2},
	},
}
