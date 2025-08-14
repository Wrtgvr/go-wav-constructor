package main

import (
	"encoding/binary"
	"io"
	"log/slog"
	"math"
	"os"
)

const (
	SampleRate = 44100 // Hz
	Volume     = 0.8   // 0.0 - 1.0
)

func main() {
	newWav, err := os.Create("newWav.wav")
	if err != nil {
		slog.Error("failed to create file", "err", err)
		return
	}
	defer newWav.Close()

	dataStartPos, _ := newWav.Seek(0, io.SeekCurrent)

	melodyTwo(newWav)

	// write header after data
	dataEndPos, _ := newWav.Seek(0, io.SeekCurrent)
	dataSize := dataEndPos - dataStartPos
	newWav.Seek(0, io.SeekStart)
	writeWAVHeader(newWav, 1, 16, int(dataSize))
}

func getNoteFreq(note Note) float64 {
	return math.Pow(2, float64(note)/12) * 440
}

func generateTone(f *os.File, frequency float64, duration float64) {
	numSamples := int(SampleRate * duration)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / SampleRate
		v := math.Sin(2 * math.Pi * frequency * t)
		sample := int16(v * Volume * 32767)
		binary.Write(f, binary.LittleEndian, sample)
	}
}

func writeWAVHeader(f *os.File, numChannels, bitsPerSample uint16, dataSize int) {
	f.Write([]byte("RIFF"))
	binary.Write(f, binary.LittleEndian, uint32(36+dataSize))
	f.Write([]byte("WAVEfmt "))
	binary.Write(f, binary.LittleEndian, uint32(16)) // fmt
	binary.Write(f, binary.LittleEndian, uint16(1))  // PCM
	binary.Write(f, binary.LittleEndian, numChannels)
	binary.Write(f, binary.LittleEndian, uint32(SampleRate))
	binary.Write(f, binary.LittleEndian, uint32(SampleRate*uint32(numChannels)*uint32(bitsPerSample)/8))
	binary.Write(f, binary.LittleEndian, uint16(numChannels*uint16(bitsPerSample/8)))
	binary.Write(f, binary.LittleEndian, bitsPerSample)
	f.Write([]byte("data"))
	binary.Write(f, binary.LittleEndian, uint32(dataSize))
}
