# debate-timer

Simple, terminal-based debate speaker timer with summary of total, mean, and
median time per speaker.

## Installation

```sh
go install github.com/tnyeanderson/debate-timer@latest
```

## Example

```text
$ debate-timer
Press a number to begin timing that speaker
speaker 1 is now speaking
speaker 3 is now speaking
speaker 2 is now speaking
speaker 5 is now speaking
speaker 4 is now speaking
speaker 3 is now speaking
speaker 4 is now speaking
--- Speaker 1 ---
Total: 1.951903604s
Mean: 1.951903604s
Median: 1.951903604s
--- Speaker 3 ---
Total: 2.607855823s
Mean: 1.303927911s
Median: 1.560106516s
--- Speaker 2 ---
Total: 939.784277ms
Mean: 939.784277ms
Median: 939.784277ms
--- Speaker 5 ---
Total: 928.34699ms
Mean: 928.34699ms
Median: 928.34699ms
--- Speaker 4 ---
Total: 3.356031926s
Mean: 1.678015963s
Median: 2.035948426s
```
