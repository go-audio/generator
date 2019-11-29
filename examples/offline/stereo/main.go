// sterep generator example
package main

import (
	"errors"
	"flag"
	"os"

	"fmt"
	"io"
	"strings"

	"github.com/go-audio/aiff"
	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/wav"
)

var (
	freqFlag     = flag.Float64("freq", 440, "frequency to generate")
	bitDepthFlag = flag.Int("bitDepth", 16, "bit size to use when generating the auid file")
	durationFlag = flag.Int("duration", 4, "duration of the generated file")
	formatFlag   = flag.String("format", "wav", "the audio format of the output file")
)

func main() {
	flag.Parse()
	var err error

	freq := *freqFlag
	fs := 44100
	biteDepth := *bitDepthFlag

	osc := generator.NewOsc(generator.WaveSine, float64(freq), fs)
	// our osc generates values from -1 to 1, we need to go back to PCM scale
	factor := float64(audio.IntMaxSignedValue(biteDepth))
	osc.Amplitude = factor
	data := make([]float64, fs**durationFlag*2)
	buf := &audio.FloatBuffer{Data: data, Format: audio.FormatStereo44100}
	osc.Fill(buf)

	// generate the sound file
	var outName string
	var format string
	switch strings.ToLower(*formatFlag) {
	case "aif", "aiff":
		format = "aif"
		outName = "stereoGenerated.aiff"
	default:
		format = "wav"
		outName = "stereoGenerated.wav"
	}

	o, err := os.Create(outName)
	if err != nil {
		panic(err)
	}
	defer o.Close()
	if err := encode(format, buf, o); err != nil {
		panic(err)
	}
	fmt.Println(outName, "generated")
}

func encode(format string, buf audio.Buffer, w io.WriteSeeker) error {
	switch format {
	case "wav":
		e := wav.NewEncoder(w,
			buf.PCMFormat().SampleRate,
			16,
			buf.PCMFormat().NumChannels, 1)
		if err := e.Write(buf.AsIntBuffer()); err != nil {
			return err
		}
		return e.Close()
	case "aiff":
		e := aiff.NewEncoder(w,
			buf.PCMFormat().SampleRate,
			16,
			buf.PCMFormat().NumChannels)
		if err := e.Write(buf.AsIntBuffer()); err != nil {
			return err
		}
		return e.Close()
	default:
		return errors.New("unknown format")
	}
}
