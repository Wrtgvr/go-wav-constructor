# WAV file constructor! Make a simple melody with nothing but a code!
## Requirements:
- Golang 1.24.0+
## Installing
```bash
git clone https://github.com/Wrtgvr/go-wav-constructor
cd go-wav-constructor
```
## Open melodies.go and start creating!!
Example:
```go
func melodyOne(f *os.File) {
	// f - .wav file, NoteC - music note, 4 - octave, .5 - note duration in seconds, 4 - volume
	generateTone(f, NoteC, 4, .5, 4) // you CAN set volume to 1+
	generateTone(f, NoteC, 4, .3, 10)
	generateTone(f, NoteD, 4, .2, 100)
	generateTone(f, NoteF, 4, .3, .5)
	generateTone(f, NoteG, 4, .1, 1)
}
```
## "Why do I need it?"
- I don't know but its cool!!
