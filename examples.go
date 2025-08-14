package main

import (
	"os"
)

func melodyOne(f *os.File) {
	generateTone(f, getNoteFreq(NoteC), .5)
	generateTone(f, getNoteFreq(NoteD), .5)
	generateTone(f, getNoteFreq(NoteE), 1)
	generateTone(f, getNoteFreq(NoteG), 1)
	generateTone(f, getNoteFreq(NoteE), 1)
	generateTone(f, getNoteFreq(NoteC), 1)
	generateTone(f, getNoteFreq(NoteD), 2)
	generateTone(f, getNoteFreq(NoteE), 2)
	generateTone(f, getNoteFreq(NoteC), 1)
}

func melodyTwo(f *os.File) {
	generateTone(f, getNoteFreq(NoteC), .3)
	generateTone(f, getNoteFreq(NoteD), .2)
	generateTone(f, getNoteFreq(NoteE), .2)
	generateTone(f, getNoteFreq(NoteF), .3)
	generateTone(f, getNoteFreq(NoteG), .1)
	generateTone(f, getNoteFreq(NoteA), .1)
	generateTone(f, getNoteFreq(NoteB), .4)
	generateTone(f, getNoteFreq(NoteC), .3) // high octave
	generateTone(f, getNoteFreq(NoteA), .2)
	generateTone(f, getNoteFreq(NoteG), .2)
	generateTone(f, getNoteFreq(NoteF), .3)
	generateTone(f, getNoteFreq(NoteE), .1)
	generateTone(f, getNoteFreq(NoteD), .1)
	generateTone(f, getNoteFreq(NoteC), .5)
}
