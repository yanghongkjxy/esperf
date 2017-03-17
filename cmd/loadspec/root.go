package loadspec

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

const (
	timeLayout = "2006-01-02 15:04:05.999"
)

var (
	host       string
	index      string
	types      string
	searchType string
	randGen    = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var RootCmd = &cobra.Command{
	Use:   "loadspec",
	Short: "Generates loadspecs for esperf",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&host, "host", "http://localhost:9200", "Elastic search HTTP address")
	RootCmd.PersistentFlags().StringVar(&index, "index", "wikipediax", "Index to perform queries")
	RootCmd.PersistentFlags().StringVar(&types, "type", "", "Index type to be acted upon")
	RootCmd.PersistentFlags().StringVar(&searchType, "search_type", "", "Type of search, for search queries.")
	RootCmd.AddCommand(parseSlowlogCmd)
	RootCmd.AddCommand(genLoadspec)
}