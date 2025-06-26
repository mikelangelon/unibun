package game

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/mikelangelon/unibun/assets"
)

type audios struct {
	audioContext      *audio.Context
	eatingSoundPlayer *audio.Player
	menuMusicPlayer   *audio.Player
	mainMusicPlayer   *audio.Player
}

func newAudios() (audios, error) {
	a := audios{}
	audioContext := audio.NewContext(44100)
	eatingSoundStream, err := mp3.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(assets.EatingSound))
	if err != nil {
		return a, err
	}
	eatingSoundPlayer, err := audioContext.NewPlayer(eatingSoundStream)
	if err != nil {
		return a, err
	}
	menuMusicStream, err := mp3.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(assets.MenuMusic))
	if err != nil {
		return a, err
	}
	menuLoop := audio.NewInfiniteLoop(menuMusicStream, menuMusicStream.Length())
	menuMusicPlayer, err := audioContext.NewPlayer(menuLoop)
	if err != nil {
		return a, err
	}
	mainMusicStream, err := mp3.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(assets.MainMusic))
	if err != nil {
		return a, err
	}
	mainLoop := audio.NewInfiniteLoop(mainMusicStream, mainMusicStream.Length())
	mainMusicPlayer, err := audioContext.NewPlayer(mainLoop)
	if err != nil {
		return a, err
	}
	menuMusicPlayer.SetVolume(0.5)
	mainMusicPlayer.SetVolume(0.5)
	return audios{
		audioContext:      audioContext,
		eatingSoundPlayer: eatingSoundPlayer,
		menuMusicPlayer:   menuMusicPlayer,
		mainMusicPlayer:   mainMusicPlayer,
	}, nil
}
