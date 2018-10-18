package cmd

import (
	"fmt"

	"github.com/pbar1/simplifier/pkg/thesaurus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// thesaurusCmd represents the thesaurus command
var thesaurusCmd = &cobra.Command{
	Use:   "thesaurus [word]",
	Short: "Search thesaurus for a word or phrase",
	Long: `Searches Big Huge Thesaurus for a word or phrase
and prints the results.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("api-key")
		results := thesaurus.Thesaurus(apiKey, args[0])
		for i, res := range results {
			fmt.Println(i, res.PartOfSpeech, res.Category, res.Word)
		}
	},
}

func init() {
	rootCmd.AddCommand(thesaurusCmd)
}
