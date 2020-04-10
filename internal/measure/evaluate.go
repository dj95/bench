package measure

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func (m *Measure) Evaluate() {
	// simple table with zero customizations
	tw := table.NewWriter()

	// set the table headers
	tw.AppendHeader(table.Row{"", "Min", "Avg", "Max"})

	// set the result into the rows
	tw.AppendRows([]table.Row{
		{
			"load",
			fmt.Sprintf("%.2f", m.result.LoadMin),
			fmt.Sprintf("%.2f", m.result.LoadAvg),
			fmt.Sprintf("%.2f", m.result.LoadMax),
		},
		{
			"memory (GB)",
			fmt.Sprintf("%.2f", m.result.MemMin),
			fmt.Sprintf("%.2f", m.result.MemAvg),
			fmt.Sprintf("%.2f", m.result.MemMax),
		},
	})

	// append a footer row
	tw.AppendFooter(table.Row{"Time (sec)", m.timeRun, "DataCount", m.dataCount})

	// use a ready-to-use style
	tw.SetStyle(table.StyleRounded)

	// customize the style and change some stuff
	tw.Style().Format.Header = text.FormatLower
	tw.Style().Format.Footer = text.FormatLower
	tw.Style().Options.SeparateColumns = false

	// render it
	fmt.Printf(":: result\n")
	fmt.Printf("%s\n", tw.Render())
}

func (m *Measure) EvaluateSystem() {
	var err error

	var hostData *host.InfoStat
	if hostData, err = host.Info(); err != nil {
		return
	}

	var cpuData []cpu.InfoStat
	if cpuData, err = cpu.Info(); err != nil {
		return
	}

	// simple table with zero customizations
	tw := table.NewWriter()

	// set the result into the rows
	tw.AppendRows([]table.Row{
		{"hostname", hostData.Hostname},
		{"os", hostData.Platform},
		{"kernel", hostData.KernelVersion},
		{"uptime (sec)", hostData.Uptime},
		{"", ""},
		{"cpu", cpuData[0].ModelName},
	})

	// use a ready-to-use style
	tw.SetStyle(table.StyleRounded)

	// customize the style and change some stuff
	tw.Style().Format.Header = text.FormatLower
	tw.Style().Format.Footer = text.FormatLower
	tw.Style().Options.SeparateColumns = false

	// render it
	fmt.Printf(":: system stats\n")
	fmt.Printf("%s\n\n", tw.Render())
}

func (m *Measure) evaluatePoint() {
	fmt.Printf(
		"load = {min: %.2f, avg: %.2f, max: %.2f} :: mem (GB) = {min: %.2f, avg: %.2f, max: %.2f}\n",
		m.result.LoadMin,
		m.result.LoadAvg,
		m.result.LoadMax,
		m.result.MemMin,
		m.result.MemAvg,
		m.result.MemMax,
	)
}
