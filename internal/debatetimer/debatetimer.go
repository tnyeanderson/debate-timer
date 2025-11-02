package debatetimer

import (
	"fmt"
	"slices"
	"time"
)

type ErrorUnsupportedSpeaker struct {
	input string
}

func NewErrorUnsupportedSpeaker(input string) ErrorUnsupportedSpeaker {
	return ErrorUnsupportedSpeaker{input}
}

func (e ErrorUnsupportedSpeaker) Error() string {
	return fmt.Sprintf("unsupported speaker %s, only speaker numbers 1-9 are supported", e.input)
}

type SpeakerTimer struct {
	total time.Duration
	times []time.Duration
}

func (s *SpeakerTimer) MeanSpeakingTime() time.Duration {
	sum := time.Duration(0)
	for _, duration := range s.times {
		sum += duration
	}
	return time.Duration(sum.Nanoseconds() / int64(len(s.times)))
}

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

type DebateTimer struct {
	currentSpeakerNumber int
	currentSpeakerStart  time.Time
	speakerTimers        map[int]SpeakerTimer
}

func (d DebateTimer) Report() string {
	d.endTimer()
	if len(d.speakerTimers) == 0 {
		return "no speakers found"
	}
	out := ""
	for speakerNumber, speakerTimer := range d.speakerTimers {
		mean := speakerTimer.MeanSpeakingTime()
		median := speakerTimer.MedianSpeakingTime()
		out += fmt.Sprintf("--- Speaker %v ---\nTotal: %v\nMean: %v\nMedian: %v\n", speakerNumber, speakerTimer.total, mean, median)
	}
	return out
}

func (d *DebateTimer) StartTimer(speakerNumber int) error {
	if speakerNumber == 0 || speakerNumber > 9 {
		return ErrorUnsupportedSpeaker{fmt.Sprint(speakerNumber)}
	}
	d.endTimer()
	d.currentSpeakerStart = time.Now()
	d.currentSpeakerNumber = speakerNumber
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
