package main

import (
	"os"
)

var gridCellDuration float64 = .08125

func freedomMotif(f *os.File) {
	generateTone(f, NoteB, 3, gridCellDuration*14, 1)
	generateTone(f, NoteB, 3, gridCellDuration*2, 0)
	generateTone(f, NoteB, 3, gridCellDuration*2, 1)
	generateTone(f, NoteASh, 3, gridCellDuration*2, 1)
	generateTone(f, NoteB, 3, gridCellDuration*2, 1)
	generateTone(f, NoteFSh, 4, gridCellDuration*4, 1)
	generateTone(f, NoteASh, 3, gridCellDuration*2, 1)
	generateTone(f, NoteB, 3, gridCellDuration*16, 1)
	generateTone(f, NoteB, 3, gridCellDuration*2, 0)
	generateTone(f, NoteB, 3, gridCellDuration*2, 1)
	generateTone(f, NoteASh, 3, gridCellDuration*2, 1)
	generateTone(f, NoteGSh, 3, gridCellDuration*2, 1)
	generateTone(f, NoteFSh, 3, gridCellDuration*4, 1)
	generateTone(f, NoteASh, 3, gridCellDuration*4, 1)
	generateTone(f, NoteB, 3, gridCellDuration*2, 1)
	generateTone(f, NoteASh, 3, gridCellDuration*2, 1)
	generateTone(f, NoteB, 3, gridCellDuration*2, 1)
	generateTone(f, NoteFSh, 4, gridCellDuration*6, 1)
	generateTone(f, NoteFSh, 4, gridCellDuration*2, 1) // ?
	generateTone(f, NoteGSh, 4, gridCellDuration*2, 1)
	generateTone(f, NoteFSh, 4, gridCellDuration*4, 1)
	generateTone(f, NoteF, 4, gridCellDuration*2, 1)
	generateTone(f, NoteDSh, 4, gridCellDuration*6, 1)
	generateTone(f, NoteB, 3, gridCellDuration*4, 1)
	generateTone(f, NoteCSh, 4, gridCellDuration*6, 1)
	generateTone(f, NoteGSh, 3, gridCellDuration*6, 1)
	generateTone(f, NoteCSh, 4, gridCellDuration*4, 1)
	generateTone(f, NoteDSh, 4, gridCellDuration*6, 1)
	generateTone(f, NoteF, 4, gridCellDuration, 1)
	generateTone(f, NoteDSh, 4, gridCellDuration, 1)
	generateTone(f, NoteGSh, 4, gridCellDuration*8, 1)
}
