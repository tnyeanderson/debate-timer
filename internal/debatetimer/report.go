package debatetimer

import (
	"fmt"
	"time"
)

// ReportEntry is an entry in the debate timer Report.
type ReportEntry struct {
	Name   string
	Total  time.Duration
	Mean   time.Duration
	Median time.Duration
	Count  int
}

// Report is the generated report from a debate timer.
type Report []ReportEntry

// String returns a pretty-printed output of the report.
func (r *Report) String() string {
	out := ""
	for _, reportEntry := range *r {
		out += fmt.Sprintf("--- %v ---\nTotal: %v\nCount: %v\nMean: %v\nMedian: %v\n", reportEntry.Name, reportEntry.Total, reportEntry.Count, reportEntry.Mean, reportEntry.Median)
	}
	return out
}
