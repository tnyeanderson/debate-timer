# debate-timer

Simple, terminal-based debate speaker timer with summary of total, mean, and
median time per speaker.

## Installation

```sh
go install github.com/tnyeanderson/debate-timer@latest
```

## Example

```text
$ DEBATETIMER_SPEAKER_1=bob debate-timer
Press a number to begin timing that speaker.
Press p to pause all timers.
Press q to quit and print the report.
Speaker 1 is bob
---
bob is now speaking
Speaker 2 is now speaking
Speaker 3 is now speaking
Speaker 2 is now speaking
bob is now speaking
bob is already speaking
Speaker 3 is now speaking
Speaker 4 is now speaking
Speaker 5 is now speaking
Speaker 2 is now speaking
bob is now speaking
--- Speaker 2 ---
Total: 5.136101871s
Count: 3
Mean: 1.712033957s
Median: 1.791949766s
--- Speaker 3 ---
Total: 3.472029046s
Count: 2
Mean: 1.736014523s
Median: 1.803986071s
--- Speaker 4 ---
Total: 3.395621ms
Count: 1
Mean: 3.395621ms
Median: 3.395621ms
--- Speaker 5 ---
Total: 1.912394797s
Count: 1
Mean: 1.912394797s
Median: 1.912394797s
--- bob ---
Total: 6.728157538s
Count: 3
Mean: 2.242719179s
Median: 2.698081657s
```
