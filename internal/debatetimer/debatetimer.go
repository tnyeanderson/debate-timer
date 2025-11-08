package debatetimer

import (
	"fmt"
	"os"
	"slices"
	"time"
)

// GetSpeakerName returns the speaker name from the environment variable based
// on the speakerNumber, or empty string if not set.
func GetSpeakerName(speakerNumber int) string {
	env := fmt.Sprintf("DEBATETIMER_SPEAKER_%v", speakerNumber)
	return os.Getenv(env)
}

// GetSpeakerNameDefault returns the speaker name from the environment variable
// based on the speakerNumber, or the default speaker name if not set.
func GetSpeakerNameDefault(speakerNumber int) string {
	if name := GetSpeakerName(speakerNumber); name != "" {
		return name
	}
	return fmt.Sprintf("Speaker %v", speakerNumber)
}

// SpeakerTimer is the timer for a given speaker.
type SpeakerTimer struct {
	total time.Duration
	times []time.Duration
}

// MeanSpeakingTime returns the mean duration of each time a speaker spoke.
func (s *SpeakerTimer) MeanSpeakingTime() time.Duration {
	sum := time.Duration(0)
	for _, duration := range s.times {
		sum += duration
	}
	return time.Duration(sum.Nanoseconds() / int64(len(s.times)))
}

// MedianSpeakingTime returns the median duration of each time a speaker spoke.
func (s *SpeakerTimer) MedianSpeakingTime() time.Duration {
	length := len(s.times)
	if length == 0 {
		return time.Duration(0)
	}
	if length == 1 {
		return s.times[0]
	}
	sorted := make([]time.Duration, length)
	copy(sorted, s.times)
	slices.Sort(sorted)
	if length%2 == 0 {
		middleIndex := length / 2
		return sorted[middleIndex]
	}
	lower := length / 2
	higher := (length + 1) / 2
	return (sorted[lower] + sorted[higher]) / 2
}

// DebateTimer contains the speaker timers and current state for the overall
// debate timer.
type DebateTimer struct {
	currentSpeakerNumber int
	currentSpeakerStart  time.Time
	speakerTimers        map[int]SpeakerTimer
}

// Report returns a Report for the DebateTimer.
func (d DebateTimer) Report() (*Report, error) {
	d.endTimer()
	r := Report{}
	for speakerNumber, speakerTimer := range d.speakerTimers {
		mean := speakerTimer.MeanSpeakingTime()
		median := speakerTimer.MedianSpeakingTime()
		r = append(r, ReportEntry{
			Name:   GetSpeakerNameDefault(speakerNumber),
			Count:  len(speakerTimer.times),
			Total:  speakerTimer.total,
			Mean:   mean,
			Median: median,
		})
	}
	return &r, nil
}

// StartTimer will stop the timer for the current speaker, and begin the timer
// for speakerNumber.
func (d *DebateTimer) StartTimer(speakerNumber int) error {
	if speakerNumber < 1 || speakerNumber > 9 {
		return NewErrorUnsupportedSpeaker(fmt.Sprint(speakerNumber))
	}
	if speakerNumber == d.currentSpeakerNumber {
		return NewErrorAlreadySpeaking(speakerNumber)
	}
	d.endTimer()
	d.currentSpeakerStart = time.Now()
	d.currentSpeakerNumber = speakerNumber
	return nil
}

// Pause will end all speaker timers.
func (d *DebateTimer) Pause() error {
	d.endTimer()
	return nil
}

func (d *DebateTimer) endTimer() {
	if d.currentSpeakerNumber != 0 {
		speakTime := time.Since(d.currentSpeakerStart)
		d.appendSpeakerTime(d.currentSpeakerNumber, speakTime)
	}
	d.currentSpeakerStart = time.Now()
	d.currentSpeakerNumber = 0
}

func (d *DebateTimer) appendSpeakerTime(speakerNumber int, speakTime time.Duration) {
	if d.speakerTimers == nil {
		d.speakerTimers = map[int]SpeakerTimer{}
	}

	if s, speakerExists := d.speakerTimers[speakerNumber]; speakerExists {
		s.times = append(s.times, speakTime)
		s.total += speakTime
		d.speakerTimers[speakerNumber] = s
		return
	}

	d.speakerTimers[speakerNumber] = SpeakerTimer{
		total: time.Duration(speakTime),
		times: []time.Duration{speakTime},
	}
}
