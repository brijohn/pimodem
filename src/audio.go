package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gobuffalo/packr"
	"time"
)

func playAudio(file string, cb func()) error {
	box := packr.NewBox("./assets")
	f, err := box.Open(file)
	if err != nil {
		return err
	}
	s, format, err := wav.Decode(f)
	if err != nil {
		return err
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	playing := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
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
