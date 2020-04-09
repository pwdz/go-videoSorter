package cmd

import (
	"github.com/spf13/cobra"
	"sorter/sortFuncs"
)

var createTextFile,downloadImg bool
func init() {
	RootCmd.AddCommand(sortVideos)
	sortVideos.Flags().BoolVarP(&createTextFile, "createTxt", "c", false, "For each movie/serie creates a text file containing information aboud it")
	sortVideos.Flags().BoolVarP(&downloadImg, "dlImage", "d", false, "For each movie/serie downloads the poster of it")

}
var sortVideos = &cobra.Command{
	Use:   "sort",
	Short: "sorting all the series/movies existing in a parent directory",
	Long:  `sort all the series/movies exitsting in a parent directory into
			ordered directories based on video names`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sortFuncs.SortVideo(createTextFile,downloadImg,args)
	},
}
