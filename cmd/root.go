/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"time"

	"github.com/julianchong00/pomodoro-timer/config"
	"github.com/julianchong00/pomodoro-timer/timer"
	"github.com/spf13/cobra"
)

const (
	workDurationFlag = "work"
	restDurationFlag = "rest"

	defaultWorkingDuration = time.Minute * 30
	defaultRestingDuration = time.Minute * 10
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomodoro-timer",
	Short: "CLI tool which simulates a pomodoro timer",
	Long: `A CLI tool which creates a pomodoro timer with configurable working and resting durations in minutes.

For example:

pomodoro -w 45 -r 5
The command above will start a timer which have a 45 minute working block and a 5 minute rest block.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: CreateTimer,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomodoro-timer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().DurationP(workDurationFlag, "w", defaultWorkingDuration, "Duration of working period")
	rootCmd.Flags().DurationP(restDurationFlag, "r", defaultRestingDuration, "Duration of resting period")
}

func CreateTimer(cmd *cobra.Command, args []string) {
	workDuration, err := cmd.Flags().GetDuration(workDurationFlag)
	if err != nil {
		log.Panic(err)
	}

	restDuration, err := cmd.Flags().GetDuration(restDurationFlag)
	if err != nil {
		log.Panic(err)
	}

	options := make([]func(*config.TimerConfig) error, 0)
	if workDuration >= time.Minute {
		options = append(options, config.Work(workDuration))
	}
	if restDuration >= time.Minute {
		options = append(options, config.Rest(restDuration))
	}

	cfg, err := config.NewConfig(options...)
	if err != nil {
		log.Panic(err)
	}

	err = timer.StartTimer(cfg)
	if err != nil {
		log.Panic(err)
	}
}
