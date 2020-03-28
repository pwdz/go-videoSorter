package sortFuncs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sorter/requests"
	"strings"
)
var videoFormats = []string{"mp4","mkv","avi","m4v","m4p","mov","qt","ogg","wmv","mpg","mpv","webm"}
var unwantedString = []string{"720","1080","480","dvd"}
func SortVideo()  {
	err := filepath.Walk(".",processPath)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println("3333333333333333333333333333333333")
	//requests.
	//fmt.Println()
	//for _,value:= range requests.Get("title","friends"){
	//	fmt.Println(value)
	//}
}

func processPath(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	//buf, _ := ioutil.ReadFile(path)

	//if(filetype.IsVideo(buf)){
	for _,vFormat := range videoFormats{
		if strings.HasSuffix(strings.ToLower(info.Name()),vFormat) {
			fmt.Println("********************************************")
			fmt.Println(path,"[------]", strings.Trim(info.Name(),strings.TrimSuffix(info.Name(),vFormat)))
			videoName := strings.Replace(strings.TrimSuffix(info.Name(),vFormat),"."," ",-1)
			videoName = strings.Replace(videoName,"-"," ",-1)
			videoName = strings.Replace(videoName,"_"," ",-1)
			videoName = strings.Replace(videoName,")"," ",-1)
			videoName = strings.Replace(videoName,"("," ",-1)
			videoName = strings.Replace(videoName,"   "," ",-1)
			videoName = strings.Replace(videoName,"  "," ",-1)
			//mkdir()
			fmt.Println(videoName)
			findVideo(videoName)
			//parts := strings.Split(videoName," ")
			//fmt.Println(parts)
			break
		}
	}
	//}
	return nil
}
func findVideo(videoName string){

	parts := strings.Split(strings.Trim(videoName," ")," ")
	printSlice(parts)

	for index,val := range parts{
		//fmt.Println("Title Result:\n",requests.Get("title",val))
		res := requests.Get("search",val)
		a(res,parts,index)
		break
	}

}
func a(result []requests.Omdb,parts []string,startIndex int){
	maxCount :=0
	counter := 0
	for _,res := range  result{
		for i:=startIndex;i<len(parts);i++{
			if !isStringValid(parts[i]){
				continue
			}
			fmt.Println("[]:","|"+res.Title,"|"+parts[i])
			if strings.Contains(res.Title,parts[i]){
				counter++
			}else {
				if maxCount<counter{
					maxCount = counter
				}
				break
			}
		}
		if counter>maxCount{
			maxCount = counter
		}
		counter=0
	}
	fmt.Println(maxCount)
}
func printSlice(sss []string){
	fmt.Println("******************BEGIN*********************")
	for _,v:=range sss{

		if isStringValid(v){
			fmt.Println("|"+v+"|")

		}else{
			//fmt.Println("sssssssssssssssS")
		}
	}
	fmt.Println("********************END**********************")
}
func isStringValid(str string) bool{
	str = strings.ToLower(str)
	for _,u:= range unwantedString{
		if strings.Contains(str,u){
			return false
		}
	}
	return true
}
func mkdir(){

}