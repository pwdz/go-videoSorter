package sortFuncs

import (
	"fmt"
	"github.com/h2non/filetype"
	"io/ioutil"
	"os"
	"sorter/requests"
	"strings"
)
var videoFormats = []string{"mp4","mkv","avi","m4v","m4p","mov","qt","ogg","wmv","mpg","mpv","webm"}

func SortVideo()  {
	//err := filepath.Walk(".",processPath)
	//if err != nil {
	//	log.Println(err)
	//}

	fmt.Println("3333333333333333333333333333333333")
	fmt.Println(requests.Get("search","friends"))
}

func processPath(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	buf, _ := ioutil.ReadFile(path)

	if(filetype.IsVideo(buf)){
		for _,vFormat := range videoFormats{
			if strings.HasSuffix(strings.ToLower(info.Name()),vFormat) {
				fmt.Println(path,"[------]", strings.Trim(info.Name(),strings.TrimSuffix(info.Name(),vFormat)))
				videoName := strings.Replace(strings.TrimSuffix(info.Name(),vFormat),"."," ",-1)
				mkdir()
				fmt.Println(videoName)
				break
			}
		}
	}
	return nil
}

func mkdir(){

}