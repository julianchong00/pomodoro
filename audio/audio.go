package audio

import (
	"os"

	"emperror.dev/errors"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

type AudioStream struct {
	Streamer beep.StreamSeekCloser
	Format   beep.Format
}

func NewAudioStream(filePath string) (AudioStream, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return AudioStream{}, errors.Wrap(err, "failed to load audio file")
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		return AudioStream{}, errors.Wrap(err, "failed to decode audio file into audio stream")
	}
	defer streamer.Close()

	return AudioStream{Streamer: streamer, Format: format}, nil
}
