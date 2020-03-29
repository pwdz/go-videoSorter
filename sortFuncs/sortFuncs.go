package sortFuncs

import (
	"fmt"
	"golang.org/x/text/search"

	//"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"sorter/requests"
	"strconv"
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
	name :=""
	year:=-1
	season := -1
	episode := -1
	for _,val := range parts{
		if isStringValid(val)&&isStringNumber(val)==-1{
			fmt.Println("konde",isStringNumber(val))
			if year==-1 && season==-1 && episode==-1{
				name += val+" "
			}else{

			}
		}else if name==""{
			name += val+" "
		} else {
			y:=isStringNumber(val)
			if y>=1950 && y<2200 {
				year = y
			}else if 0<y&&y<100{
				if(season==-1) {
					season = y
				}else {
					episode = y
				}
			}
			//break
		}
		//fmt.Println("Title Result:\n",requests.Get("title",name))
		//res := requests.Get("search",val)
		//a(res,parts,index)
		//break
	}
	fmt.Println("name:|"+strings.Trim(name," ")+"|year",year,"|season:",season,"|episode:",episode)
	fmt.Println(requests.Get("title",strings.Trim(name," ")))

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
func isStringNumber(str string) int{
	if year,err:=strconv.Atoi(str);err==nil{
		fmt.Println("year",year)
		//if 1950<year&&year<2150{
			return year
		//}
	}
	return -1
}
func isSearsonOrEpisode(str string)(int,int){//s,e
	str = strings.ToLower(str)
	season:=-1
	episode:=-1
	if strings.Contains(str,"season"){
		str = strings.Replace(str,"season","",-1)
		if num,err:=strconv.Atoi(str);err==nil{

		}
	}else if strings.Contains(str,"s"){
		str = strings.Replace(str,"s","",-1)
		if num,err:=strconv.Atoi(str);err==nil{
			if num<100 && num>0 {
				season = num
			}
		}else if str==""{

		}
	}
	return -1
}
func isEpisode(str string)int{
	return -1
}
func mkdir(){

}