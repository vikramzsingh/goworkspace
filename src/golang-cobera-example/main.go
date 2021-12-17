package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	persistRootFlag bool
	localRootFlag bool
	times int
	rootCmd = &cobra.Command{
		Use: "example", // command name
		Short: "An example of go and cobra", // short declarition of command
		Long: `This is simple example of cobra program. It will have several subcommand and flags.  (its summary of command)`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command")
		},
	}

	echoCmd = &cobra.Command{
		Use: "echo [strings to echo]",
		Short: "prints given string to stdout",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	timesCmd = &cobra.Command{
		Use: "times [strings to echo]",
		Short: "prints given string to stdout multiple times",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistant root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	timesCmd.Flags().IntVarP(&times, "timesFlag", "t", 1, "number of times to echo to stdout")

	// adding commands
	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
