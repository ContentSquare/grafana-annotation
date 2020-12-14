package cmd

import (
	"github.com/contentsquare/grafana-annotation/pkg/poster"
	"github.com/spf13/cobra"
)

var tags []string

// cmdPlan
var cmdPost = &cobra.Command{
	Use:   "post",
	Short: "post one annotation",
	Run: func(cmd *cobra.Command, args []string) {
		poster.PostAnnotation(tags, args...)
	},
}

func init() {
	cmdPost.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "extra tags")
	rootCmd.AddCommand(cmdPost)

}
