package cmd

import (
	"github.com/spf13/cobra"
	"sorter/requests"
	"sorter/sortFuncs"
)
var isVideo bool
var search bool
func init() {
	RootCmd.AddCommand(downloadImage)
	RootCmd.AddCommand(searchVideoInfo)
	downloadImage.Flags().BoolVarP(&isVideo, "video", "v", false, "When instead of URL you want to enter video name")
	searchVideoInfo.Flags().BoolVarP(&search, "search", "s", false, "Finds all movies/series containing the entered title")

}
var downloadImage = &cobra.Command{
	Use:   "download",
	Short: "Download the poster of a movie or from a URL",
	Long:  `Download the poster of a movie/serie by entering the name and using --video/-v flag
			or download any image from a url by just entering the url. 
			In both cases, destination directory for download result is needed`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if isVideo{
			//args[1]->name or URL
			//args[0]-> desDir
			result:=sortFuncs.OnlineSearch(args[1],"title")
			requests.DownloadFile(result[0].Poster,args[0],args[1])
		}else{
			requests.DownloadFile(args[1],args[0])
		}
	},
}
var searchVideoInfo = &cobra.Command{
	Use:   "search",
	Short: "Search for a movie/serie information",
	Long:  `Seach for a movie/serie information by entering the name of the video.
			If you know the title of the movie, enter the whole title for more precise result; Else
			enter the part of the title and --search/-s for seeing all the movies/series containing the title you entered`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//args[0]->title
		//args[1]->destonation
		var result []requests.Omdb
		if search{
			result=sortFuncs.OnlineSearch(args[0],"search")
		}else{
			result = sortFuncs.OnlineSearch(args[0],"title")
		}
		if len(args)>1{
			for _, value := range result {
				sortFuncs.CreateInfoFile(args[1], value, false)
			}
		}
		sortFuncs.PrintVideoInfo(result)
	},
}