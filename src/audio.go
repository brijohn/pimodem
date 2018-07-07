package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gobuffalo/packr"
	"time"
)

func playAudio(file string, volume int, cb func()) error {
	box := packr.NewBox("./assets")
	if volume < 0 || volume > 3 {
		return fmt.Errorf("Volume must be between 0 and 3")
	}
	f, err := box.Open(file)
	if err != nil {
		return err
	}
	s, format, err := wav.Decode(f)
	if err != nil {
		return err
	}
	v := &effects.Volume{
		Streamer: s,
		Base:     2,
		Volume:   0,
		Silent:   volume == 0,
	}
	if volume == 1 {
		v.Volume -= 2
	} else if volume == 3 {
		v.Volume += 3
	}
	fmt.Printf("volume struct %v\n", v)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	playing := make(chan struct{})
	speaker.Play(beep.Seq(v, beep.Callback(func() {
		close(playing)
	})))
	ticker := time.NewTicker(time.Millisecond * 500)
done:
	for {
		select {
		case <-ticker.C:
			if cb != nil {
				cb()
			}
		case <-playing:
			ticker.Stop()
			playing = nil
			break done
		}
	}
	return nil
}
