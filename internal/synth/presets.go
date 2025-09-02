package synth

var Presets = map[string]Streamer{
	"P2": &Sequence{
		&LinearAD{
			Sound:  &Sine{Amplitude: 1.0, Frequency: 120.0, Duration: 1},
			Attack: 0.4,
			Decay:  0.4,
		},
		&Pause{Duration: 0.2},
	},
	"P3": &Sequence{
		Punch(1.0, 0.3, 60),
		&Pause{Duration: 0.2},
	},
	"P4": &Chaotic{Punch(1.0, 0.3, 60), Bound[Milliseconds, Milliseconds]{10, 400}},
}
