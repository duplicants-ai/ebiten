// Copyright 2022 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/duplicants-ai/ebiten"
	"github.com/duplicants-ai/ebiten/audio"
	"github.com/duplicants-ai/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	sampleRate   = 48000
)

type Game struct {
	audioContext *audio.Context
	player       *audio.Player
	sineWave     *SineWave
}

type SineWave struct {
	frequency    int
	minFrequency int
	maxFrequency int

	// pos is the position in the wave length in the range of [0, 1).
	pos float64

	m sync.Mutex
}

func NewSineWave() *SineWave {
	return &SineWave{
		frequency:    440,
		minFrequency: 440,
		maxFrequency: 880,
	}
}

func (s *SineWave) Update(raisePitch bool) {
	s.m.Lock()
	defer s.m.Unlock()

	if raisePitch {
		if s.frequency < s.maxFrequency {
			s.frequency += 10
		}
	} else {
		if s.frequency > s.minFrequency {
			s.frequency -= 10
		}
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()

	const bytesPerSample = 8

	n := len(buf) / bytesPerSample * bytesPerSample
	buf = buf[:n]

	length := sampleRate / float64(s.frequency)
	p := float64(length * s.pos)
	for i := 0; i < n/bytesPerSample; i++ {
		v := math.Float32bits(float32(math.Sin(2 * math.Pi * p / length)))
		buf[8*i] = byte(v)
		buf[8*i+1] = byte(v >> 8)
		buf[8*i+2] = byte(v >> 16)
		buf[8*i+3] = byte(v >> 24)
		buf[8*i+4] = byte(v)
		buf[8*i+5] = byte(v >> 8)
		buf[8*i+6] = byte(v >> 16)
		buf[8*i+7] = byte(v >> 24)
		p++
	}

	_, s.pos = math.Modf(p / length)

	return n, nil
}

func NewGame() *Game {
	return &Game{
		audioContext: audio.NewContext(sampleRate),
	}
}

func (g *Game) Update() error {
	if g.audioContext == nil {
		g.audioContext = audio.NewContext(sampleRate)
	}
	if g.player == nil {
		g.sineWave = NewSineWave()
		p, err := g.audioContext.NewPlayerF32(g.sineWave)
		if err != nil {
			return err
		}
		g.player = p
		g.player.Play()

		// Adjust the buffer size to reflect the audio source changes in real time.
		// Note that Ebitengine doesn't guarantee the audio quality when the buffer size is modified.
		// 1/20[s] should work in most cases, but this might cause glitches in some environments.
		g.player.SetBufferSize(time.Second / 20)
	}
	g.sineWave.Update(ebiten.IsKeyPressed(ebiten.KeyA))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "This is an example of a real time PCM.\nPress and hold the A key.")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Real Time PCM (Ebitengine Demo)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
