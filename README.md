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
	generateTone(f, getNoteFreq(NoteC), .5) // f - .wav file, getNoteFreq - frequency of given note, .5 - note duration in seconds
	generateTone(f, getNoteFreq(NoteD), .5)
	generateTone(f, getNoteFreq(NoteE), 1)
	generateTone(f, getNoteFreq(NoteG), 1)
	generateTone(f, getNoteFreq(NoteE), 1)
	generateTone(f, getNoteFreq(NoteC), 1)
	generateTone(f, getNoteFreq(NoteD), 2)
	generateTone(f, getNoteFreq(NoteE), 2)
	generateTone(f, getNoteFreq(NoteC), 1)
}
```
## "Why do I need it?"
- I don't know but its cool!!
