// Package commandline Specify the commandline handling.
package commandline

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitializeFlags Initializes the commandline flags.
func InitializeFlags() {
	// create the commandline flags
	pflag.IntP("interval", "i", 1, "choose the interval in which data points should be aggregated (in seconds)")
	pflag.BoolP("follow", "f", false, "print every data point collected")
	pflag.BoolP("system", "s", false, "include system stats in the results")
	pflag.BoolP("help", "h", false, "print the help")

	// parse the pflags
	pflag.Parse()

	// bind the pflags
	viper.BindPFlags(pflag.CommandLine)
}

// PrintHelp Print the help and the default tag
func PrintHelp() {
	fmt.Printf("Usage: bench [OPTION]\n")
	fmt.Printf("Create load statistics over a short amount of time.\n")
	fmt.Printf("Example: bench -fi 2\n")
	fmt.Printf("Options:\n")
	pflag.PrintDefaults()
	fmt.Printf("\n")
	fmt.Printf("Report bugs at: https://github.com/dj95/bench/issues\n")
	fmt.Printf("Homepage: https://github.com/dj95/bench\n")
}
