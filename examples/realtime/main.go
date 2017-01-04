// demo package simulating a realtime generation and processing.
// Start the example from your terminal and type a letter + enter.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"

	"github.com/go-audio/audio"
	"github.com/go-audio/generator"
	"github.com/go-audio/transforms"
	"github.com/gordonklaus/portaudio"
)

func main() {
	bufferSize := 512
	buf := &audio.FloatBuffer{
		Data:   make([]float64, bufferSize),
		Format: audio.FormatMono44100,
	}
	currentNote := 440.0
	osc := generator.NewOsc(generator.WaveSine, currentNote, buf.Format.SampleRate)
	osc.Amplitude = 1

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	gainControl := 0.0
	currentVol := osc.Amplitude

	fmt.Println(`This is a demo, press a key followed by enter, the played note should change.
Use the - and + keys follow by enter to decrease or increase the volume\nPress q or ctrl-c to exit.
Note that the sound will come out of your default sound card.`)

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			if len(scanner.Text()) > 0 {
				k := scanner.Text()[0]
				switch k {
				case 'q':
					sig <- os.Interrupt
				case '+':
					gainControl += 0.10
				case '-':
					gainControl -= 0.10
				default:
					v := float64(math.Abs(float64(int(k - 100))))
					currentNote = 440.0 * math.Pow(2, (v)/12.0)
					fmt.Printf("switching oscillator to %.2f Hz\n", currentNote)
					if currentNote > 22000 {
						currentNote = 440.0
					}
					osc.SetFreq(currentNote)
				}
			}
		}
	}()

	// Audio output
	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]float32, bufferSize)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	if err := stream.Start(); err != nil {
		log.Fatal(err)
	}
	defer stream.Stop()
	for {

		// populate the out buffer
		if err := osc.Fill(buf); err != nil {
			log.Printf("error filling up the buffer")
		}
		// apply vol control if needed (applied as a transform instead of a control
		// on the osc)
		if gainControl != 0 {
			currentVol += gainControl
			if currentVol < 0.1 {
				currentVol = 0
			}
			if currentVol > 6 {
				currentVol = 6
			}
			fmt.Printf("new vol %f.2", currentVol)
			gainControl = 0
		}
		transforms.Gain(buf, currentVol)

		f64ToF32Copy(out, buf.Data)

		// write to the stream
		if err := stream.Write(); err != nil {
			log.Printf("error writing to stream : %v\n", err)
		}
		select {
		case <-sig:
			fmt.Println("\tCiao!")
			return
		default:
		}
	}
}

// portaudio doesn't support float64 so we need to copy our data over to the
// destination buffer.
func f64ToF32Copy(dst []float32, src []float64) {
	for i := range src {
		dst[i] = float32(src[i])
	}
}
