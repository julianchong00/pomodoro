package timer

import (
	"fmt"
	"math"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/julianchong00/pomodoro/config"
)

const (
	workTmpl = `{{ red "Work Period" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}} {{string . "remaining_time"}}`
	restTmpl = `{{ red "Rest Period" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}} {{string . "remaining_time"}}`

	remainingTimeElement = "remaining_time"
)

// Start timer in background and make sound when duration runs out
func StartTimer(cfg *config.TimerConfig) error {

	runProgressBar(cfg.WorkingDuration, true)

	runProgressBar(cfg.RestingDuration, false)

	return nil
}

func runProgressBar(duration time.Duration, isWorkPeriod bool) {
	durationSeconds := duration.Seconds()

	var tmpl string
	if isWorkPeriod {
		tmpl = workTmpl
	} else {
		tmpl = restTmpl
	}

	bar := pb.ProgressBarTemplate(tmpl).Start64(int64(durationSeconds))

	for i := 0; i < int(durationSeconds); i++ {
		remainingTime := int64(durationSeconds) - bar.Current()
		bar.Set(remainingTimeElement, formatTime(remainingTime))
		bar.Increment()
		time.Sleep(time.Second)
	}

	bar.Finish()
}

func formatTime(seconds int64) string {
	if seconds > 60 {
		minutes := math.Floor(float64(seconds) / 60)
		seconds := seconds % 60
		return fmt.Sprintf("%vM%vS", minutes, seconds)
	}

	return fmt.Sprintf("%vS", seconds)
}
