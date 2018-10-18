package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simplifier",
	Short: "Translates text to ELI5",
	Long: `Simplifier is a tool for generating "ELI5-ized"
pieces of text from more complicated inputs.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("api-key", "k", "", "Big Huge Thesaurus API key")
	viper.BindPFlags(rootCmd.PersistentFlags())
}

// initConfig reads ENV variables if set.
func initConfig() {
	viper.BindEnv("api-key", "BHT_API_KEY")
}
