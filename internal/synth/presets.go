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
	"Kick": &StaticSequence{
		&Envelope{
			Sound: &Mixer{
				&SweepSine{
					Amplitude: 1.0,
					Duration:  0.3,
					StartFreq: 120,
					EndFreq:   50,
				},
				&Sine{
					Amplitude: 1.0,
					Duration:  0.3,
					Frequency: 50,
				},
			},
			Attack:  0.005,
			Decay:   0.1,
			Sustain: 0.8,
			Release: 0.015,
		},
		&Pause{0.2},
	},
	"WhiteNoise": &WhiteNoise{
		Amplitude: 1.0,
		Duration:  1.0,
	},
	"BrownNoise": &LowPass{
		Sound: &LeakyIntegrator{
			Sound: &WhiteNoise{
				Amplitude: 1.0,
				Duration:  10.0,
			},
		},
		CutoffFrequency: 200.0,
	},
}
