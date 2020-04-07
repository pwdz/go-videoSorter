package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"sorter/sortFuncs"
)

var createTextFile,downloadImg bool
func init() {
	RootCmd.AddCommand(sortVideos)
	flag.BoolVarP(&createTextFile, "createTxt", "c", false, "For each movie/serie creates a text file containing information aboud it")
	flag.BoolVarP(&downloadImg, "dlImage", "d", false, "For each movie/serie downloads the poster of it")
	flag.Parse()
	fmt.Println("kire khar",createTextFile,downloadImg)
}
var sortVideos = &cobra.Command{
	Use:   "sort",
	Short: "sort sort sort",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args,createTextFile)
		if len(args)>=1{
			sortFuncs.SortVideo(createTextFile,downloadImg,args[0],args[1])
		}else {
			fmt.Println("Source path is not given!")
		}
	},
}
