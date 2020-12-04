/** audio.go
 * Copyright (C) 2018-2020  Brian Johnson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 **/

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
