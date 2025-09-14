noisewars - just enough synthesizer to piss off your neighbors!

# Usage

```bash
go run ./cmd/noisewars | aplay -f S16_LE -r 44100 -c 1
```

Check [presets.go](./internal/synth/presets.go) for built-in sequences. To play
them, use flag `-preset`:

```bash
go run ./cmd/noisewars -preset RandomPunchFast -loop | aplay -f S16_LE -r 44100 -c 1 -V mono
```
