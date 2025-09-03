package synth

var Presets = map[string]Streamer{
	"Sweep": &Envelope{
		&StaticSequence{
			&SweepSine{1.0, 0.4, 50, 80},
			&Sine{1.0, 0.2, 80},
			&SweepSine{1.0, 0.4, 80, 50},
		},
		0.05,
		0.05,
		0.6,
		0.1,
	},
	"RandomPunchFast": &DynamicSequence{
		&Envelope{
			&VaryingSine{randomf64(0.6, 1.0), 0.2, randomf64(50.0, 80.0)},
			0.02,
			0.02,
			0.6,
			0.05,
		},
		&VaryingPause{randomf64(0.01, 0.1)},
	},
	"RandomPunchSlow": &DynamicSequence{
		&Envelope{
			&VaryingSine{randomf64(0.6, 1.0), 0.5, randomf64(40.0, 80.0)},
			0.15,
			0.10,
			0.45,
			0.08,
		},
		&VaryingPause{randomf64(0.5, 2.0)},
	},
}
