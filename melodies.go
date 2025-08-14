package main

import (
	"os"
)

func melodyOne(f *os.File) {
	generateTone(f, getNoteFreq(NoteC), .5, .1)
	generateTone(f, getNoteFreq(NoteD), .5, .1)
	generateTone(f, getNoteFreq(NoteE), 1, .1)
	generateTone(f, getNoteFreq(NoteG), 1, .1)
	generateTone(f, getNoteFreq(NoteE), 1, .1)
	generateTone(f, getNoteFreq(NoteC), 1, .1)
	generateTone(f, getNoteFreq(NoteD), 2, .1)
	generateTone(f, getNoteFreq(NoteE), 2, .1)
	generateTone(f, getNoteFreq(NoteC), 1, .1)
}

func melodyTwo(f *os.File) {
	generateTone(f, getNoteFreq(NoteC), .3, 10)
	generateTone(f, getNoteFreq(NoteD), .2, 100)
	generateTone(f, getNoteFreq(NoteE), .2, 10)
	generateTone(f, getNoteFreq(NoteF), .3, 5)
	generateTone(f, getNoteFreq(NoteG), .1, 1)
	generateTone(f, getNoteFreq(NoteA), .1, 1)
	generateTone(f, getNoteFreq(NoteB), .4, 20)
	generateTone(f, getNoteFreq(NoteC), .3, 10)
	generateTone(f, getNoteFreq(NoteA), .2, 5)
	generateTone(f, getNoteFreq(NoteG), .2, 3)
	generateTone(f, getNoteFreq(NoteF), .3, 3)
	generateTone(f, getNoteFreq(NoteE), .1, 1)
	generateTone(f, getNoteFreq(NoteD), .1, 5)
	generateTone(f, getNoteFreq(NoteC), .5, 10)
}
