package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/viper"

	"github.com/dj95/bench/internal/commandline"
	"github.com/dj95/bench/internal/measure"
)

func init() {
	commandline.InitializeFlags()

	if viper.GetBool("help") {
		commandline.PrintHelp()

		os.Exit(0)
	}
}

func main() {
	fmt.Println(":: starting bench")

	// configure the measurement
	measurement := measure.New(&measure.MeasureOptions{
		Interval:    viper.GetFloat64("interval"),
		PrintPoints: viper.GetBool("follow"),
	})

	// start capturing data
	measurement.Start()

	// create the channel that later blocks until the signal
	// is retrieved
	sig := make(chan os.Signal, 1)

	// register that the signal should be send through the
	// channel
	signal.Notify(sig, os.Interrupt)

	// block until a signal was received
	<-sig

	// remove the Ctrl-C character from the console
	fmt.Print("\r  \n")

	// stop capturing data
	measurement.Stop()

	// print system stats if the flag is set
	if viper.GetBool("system") {
		measurement.EvaluateSystem()
	}

	// pretty print a table
	measurement.Evaluate()

	// exit gracefully
	os.Exit(0)
}
