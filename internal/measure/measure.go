// Package measure Measure cpu and memory performance in the background
package measure

import (
	"time"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type Measure struct {
	shouldRun bool
	result    *Result
	dataCount float64
	timeRun   int64
	options   *MeasureOptions
}

type MeasureOptions struct {
	Interval    float64
	PrintPoints bool
}

type Result struct {
	LoadMax float64
	LoadAvg float64
	LoadMin float64

	MemMax float64
	MemAvg float64
	MemMin float64
}

// New Create a new Measure object
func New(options *MeasureOptions) *Measure {
	return &Measure{
		shouldRun: false,
		result:    nil,
		dataCount: 0.0,
		options:   options,
	}
}

// Start Start the measurement in the background
func (m *Measure) Start() {
	m.shouldRun = true

	go m.loop()
}

// Stop Stop measuring data
func (m *Measure) Stop() *Result {
	m.shouldRun = false

	return m.result
}

func (m *Measure) loop() {
	// initialize the values
	m.updateResult()

	// run as long as it does not get stopped
	for m.shouldRun {
		// wait one second
		time.Sleep(time.Duration(m.options.Interval) * time.Second)

		// save the time, how long this program is running
		m.timeRun = m.timeRun + int64(m.options.Interval)

		// update the current data set
		m.updateResult()
	}
}

func (m *Measure) updateResult() {
	var err error

	// get the current cpu load average
	var cpu *load.AvgStat
	if cpu, err = load.Avg(); err != nil {
		return
	}

	// get the current memory usage
	var ram *mem.VirtualMemoryStat
	if ram, err = mem.VirtualMemory(); err != nil {
		return
	}

	// increase the sample size
	m.dataCount = m.dataCount + 1.0

	if m.result == nil {
		m.result = &Result{
			LoadMin: cpu.Load1,
			LoadAvg: cpu.Load1,
			LoadMax: cpu.Load1,

			MemMin: float64(ram.Used) / 1073741824.0,
			MemAvg: float64(ram.Used) / 1073741824.0,
			MemMax: float64(ram.Used) / 1073741824.0,
		}
	}

	// recalculate the cpu usage
	m.result.LoadMin = min(m.result.LoadMin, cpu.Load1)
	m.result.LoadAvg = m.movingAverage(m.result.LoadAvg, cpu.Load1)
	m.result.LoadMax = max(m.result.LoadMax, cpu.Load1)

	// recalculate the memory usage
	m.result.MemMin = min(m.result.MemMin, float64(ram.Used)/1073741824.0)
	m.result.MemAvg = m.movingAverage(m.result.MemAvg, float64(ram.Used)/1073741824.0)
	m.result.MemMax = max(m.result.MemMax, float64(ram.Used)/1073741824.0)

	// print the current data set, if the flag was given
	if m.options.PrintPoints {
		m.evaluatePoint()
	}
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}

	return b
}

func min(a, b float64) float64 {
	if a <= b {
		return a
	}

	return b
}

func (m *Measure) movingAverage(average, sample float64) float64 {
	average = average - (average / m.dataCount)

	return average + (sample / m.dataCount)
}
