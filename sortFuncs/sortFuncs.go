package sortFuncs

import (
	"fmt"
	"sorter/requests"

	//"golang.org/x/text/search"

	//"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)
var(
	videoFormats = []string{"mp4","mkv","avi","m4v","m4p","mov","qt","ogg","wmv","mpg","mpv","webm","sub","srt"}
	unwantedString = []string{"720","1080","480","dvd","twoddl"}
	sYear,lYear = 1920,2200
	isOnline=false
	shouldDlImg=false
)
type video struct{
	Name string
	Year,Episode,Season int
}
func SortVideo(isOnlineSort,shouldDownloadImg bool) {
	isOnline = isOnlineSort
	shouldDlImg = shouldDownloadImg
	err := filepath.Walk("./",processPath)
	if err != nil {
		log.Println(err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	for _,vFormat := range videoFormats{
		if strings.HasSuffix(strings.ToLower(info.Name()),vFormat) {
			fmt.Println("\n********************************************")
			fmt.Println(path,"[------]", strings.Trim(info.Name(),strings.TrimSuffix(info.Name(),vFormat)))
			videoName := strings.Replace(strings.TrimSuffix(info.Name(),vFormat),"."," ",-1)
			videoName = strings.Replace(videoName,"-"," ",-1)
			videoName = strings.Replace(videoName,"_"," ",-1)
			videoName = strings.Replace(videoName,")"," ",-1)
			videoName = strings.Replace(videoName,"("," ",-1)
			videoName = strings.Replace(videoName,"   "," ",-1)
			videoName = strings.Replace(videoName,"  "," ",-1)
			video := extractVideoData(videoName)
			newPath := mkdir(video)
			mvFile(newPath,path)
			if isOnline{
				if fileExists(strings.Split(newPath,"Season")[0]+video.Title+".jpg"){
					createInfoFile(newPath,video)
				}
			}
			if shouldDlImg{
				if !fileExists(strings.Split(newPath,"Season")[0]+video.Title+".jpg") {
					if video.Type=="series"{
						newPath= strings.Split(newPath,"Season")[0]
					}
					requests.DownloadFile(video.Poster, newPath, video.Title)
				}
			}
			break
		}
	}
	return nil
}
func extractVideoData(videoName string)requests.Omdb{
	var v video
	parts := strings.Split(strings.Trim(videoName," ")," ")


	//printSlice(parts)
	v.Name=""
	v.Year=-1
	v.Season = 0
	v.Episode = 0
	flag :=true
	for index:=0;index<len(parts);index++{
		val := parts[index]
		if isValid(val) {
			if num := isNumber(val); num == -1 {
				if seas,ep,_,lastIndex:=isSeasonAndEpisode(val,index,parts);seas==0&&flag{
					//fmt.Println("*_*",flag,name)
					v.Name += val+ " "
				}else{
					if(setSeasonEpisode(seas,ep,&v.Season,&v.Episode)){
						index=lastIndex
					}
					flag=isNameEmpty(v.Name)
				}
			} else {
				if v.Name == "" {
					v.Name += strconv.Itoa(num)+" "
				} else if isYear(num) {
					v.Year = num
					flag=isNameEmpty(v.Name)
				} else {
					seas,ep,num,lastIndex:=isSeasonAndEpisode(val,index,parts)
					if !setSeasonEpisode(seas,ep,&v.Season,&v.Episode){
						v.Name += strconv.Itoa(num)+" "
					}else{

						index=lastIndex
					}
					flag=isNameEmpty(v.Name)
				}
			}
		}else{
			flag=isNameEmpty(v.Name)
		}
	}
	fmt.Println("name:|"+strings.Trim(v.Name," ")+"|year",v.Year,"|v.Season:",v.Season,"|v.Episode:",v.Episode)
	var res requests.Omdb
	res.Response=false
	if isOnline{
		temp:=onlineSearch(v,"title")
		res = temp[0]
	}
	setOmdbValues(v,&res)
	return res
}
func onlineSearch(video video,reqType string)[]requests.Omdb{
	result := make([]requests.Omdb,1)
	for index,value := range requests.Get(reqType,strings.Trim(video.Name," ")){
		//fmt.Println("%_%_%_%_%_%_%",value)
		result[index] = value
	}
	return result
}
func setOmdbValues(v video,res *requests.Omdb){
	if v.Season>0{
		res.Type = "series"
		res.Season = strconv.Itoa(v.Season)
		res.Episode = strconv.Itoa(v.Episode)
	}
	if !res.Response {
		res.Title = v.Name
		if isYear(v.Year) {
			res.Year = strconv.Itoa(v.Year)
		} else {
			res.Type = "movie"
		}
	}

	fmt.Println("in Checktype:",res.Season,v.Season)
}
func printSlice(sss []string){
	fmt.Println("******************BEGIN*********************")
	for _,v:=range sss{
		if isValid(v){
			fmt.Println("|"+v+"|")
		}else{
			fmt.Println("invalid part")
		}
	}
	fmt.Println("********************END**********************")
}
func isNameEmpty(name string)bool{
	if name=="" {
		return true
	}
	return false
}
func isValid(str string) bool{
	str = strings.ToLower(str)
	for _,u:= range unwantedString{
		if strings.Contains(str,u){
			return false
		}
	}
	return true
}
func isNumber(str string) int{
	if year,err:=strconv.Atoi(str);err==nil{
		//fmt.Println("year",year)
		//if 1950<year&&year<2150{
			return year
		//}
	}
	return -1
}
func isYear(num int)bool{
	if num>sYear && num<lYear {
		return true
	}
	return false
}
//check whether string is season and episode or not
func isSeasonAndEpisode(str string,strIndex int,parts []string)(int,int,int,int){//season,episode,nameNum,lastIndex
	str = strings.ToLower(str)
	season:=0
	episode:=0
	nameNum:=-1
	lastIndex:=strIndex
	if strings.Contains(str,"season"){
		str = strings.Replace(str,"season","",-1)
		if s,err:=strconv.Atoi(str);err==nil{
			season = s
			episode,lastIndex = isEpisode(parts,strIndex+1)
		} else{
			if s,err:= strconv.Atoi(parts[strIndex+1]); err==nil{
				season=s
				episode,lastIndex = isEpisode(parts,strIndex+2)
			}
		}
	}else if strings.Contains(str,"s"){
		str = strings.Replace(str,"s","",-1)
		var index int
		var ch int32
		for index,ch = range str{
			if s,err:=strconv.Atoi(string(ch));err==nil{
				season*=10
				season+=s
			}else {
				if index==strings.Index(str,"e"){
					episode,lastIndex = isEpisode(parts,strIndex)
				}
				break
			}
		}
		if str==""{
			if s,err:=strconv.Atoi(parts[strIndex+1]);err==nil {
				season = s
				episode,lastIndex = isEpisode(parts,strIndex+2)
			}
		} else if index==len(str)-1 && index!=strings.Index(str,"e"){
			episode,lastIndex = isEpisode(parts,strIndex+1)
		}
	}else if strings.Contains(str,"x"){
		subStr := strings.Split(str,"x")
		if len(subStr)==2{
			s :=isNumber(subStr[0])
			ep :=isNumber(subStr[1])
			if s!=-1 && ep!=-1{
				season = s
				episode = ep
			}
		}
	}else if isNumber(str)!=-1{
		if strIndex+1<len(parts){
			if isNumber(parts[strIndex+1])!=-1&&!isYear(isNumber(parts[strIndex+1])){
				season = isNumber(str)
				episode = isNumber(parts[strIndex+1])
				lastIndex=strIndex+1
			}else{
				nameNum = isNumber(str)
			}

		}
	}
	return season,episode,nameNum,lastIndex
}
func isEpisode(parts []string,epIndex int)(int,int){//episode,lastIndex
	if epIndex>=len(parts)-1{
		return 0,epIndex
	}
	episode := 0
	lastIndex:=epIndex
	str := strings.ToLower(parts[epIndex])
	if strings.Contains(str,"episode"){
		str = strings.Replace(str,"episode","",-1)
		if ep,err:=strconv.Atoi(str);err==nil{
			episode = ep
		} else{
			if epIndex+1<len(parts){
				if ep,err:= strconv.Atoi(parts[epIndex+1]); err==nil{
					lastIndex = epIndex+1
					episode=ep
				}
			}
			
		}
	}else if strings.Contains(str,"e"){
			index := strings.Index(str, "e")
			for index += 1; index < len(str); index++ {
				if ep, err := strconv.Atoi(string(str[index])); err == nil {
					episode *= 10
					episode += ep
				} else {
					break
				}
			}
			if episode == 0 {
				if epIndex+1 < len(parts) {
					if ep, err := strconv.Atoi(parts[epIndex+1]); err == nil {
						lastIndex = epIndex+1
						episode = ep
					}
				}
			}
	}
	return episode,lastIndex
}
func setSeasonEpisode(seas int,ep int,season *int,episode  *int)bool{
	if seas>0{
		*season=seas
	}
	if ep>0{
		*episode=ep
		return true
	}
	return false
}
func mkdir(v requests.Omdb)string{
	var newPath string
	//fmt.Println("in Mkdir:",v)
	if v.Type=="movie"{//it's a movie
		newPath = "./pwdSorter/Movies/"+v.Title
		if v.Response {
			newPath +=" ("+v.Year+")"
			newPath +=" IMDB("+v.ImdbRating+")"
		}else if isYear(isNumber(v.Year)){
			newPath += " ("+v.Year+")"
		}
	}else{//it's a serie
		newPath = "./pwdSorter/Series/"+v.Title+"$YEARIMDB/Season"+v.Season+ "/E"+v.Episode
		fmt.Println("_---------------------------",v.Season)
		if v.Response{
			newPath = strings.Replace(newPath,"$YEARIMDB","("+v.Year+")"+" IMDB("+v.ImdbRating+")",-1)
		}else if isYear(isNumber(v.Year)){
			newPath = strings.Replace(newPath,"$YEARIMDB","("+v.Year+")",-1)
		}else {
			newPath = strings.Replace(newPath,"$YEARIMDB","",-1)
		}
	}
	os.MkdirAll(newPath,0755)
	return newPath
}
func mvFile(newPath string,oldPath string){
	parts:=strings.Split(oldPath,"/")
	newPath+= "/"+parts[len(parts)-1]
	fmt.Println("newPath",newPath)
	err :=  os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("In mvFile:",err)
		return
	}
}
func createInfoFile(path string,info requests.Omdb){
	if info.Response{
		if info.Type=="series"{
			path= strings.Split(path,"Season")[0]
		}
		file,err := os.Create(path+"/"+info.Title+".txt")
		if err==nil{
			defer file.Close()
			file.WriteString("Title: "+info.Title+"\n")
			file.WriteString("Year: "+info.Year+"\n")
			file.WriteString("IMDB: "+info.ImdbRating+"\n")
			file.WriteString("Metascore: "+info.Metascore+"\n")
			file.WriteString("Genre: "+info.Genre+"\n")
			file.WriteString("Type: "+info.Type+"\n")
			file.WriteString("Writer: "+info.Writer+"\n")
			file.WriteString("Actors: "+info.Actors+"\n")
			file.WriteString("Plot: "+info.Plot+"\n")
			file.WriteString("Language: "+info.Language+"\n")
			file.WriteString("Awards: "+info.Awards+"\n")
			file.WriteString("Poster: "+info.Poster+"\n")
			file.WriteString("Country: "+info.Country+"\n")
		}

	}
}
func fileExists(path string)bool{
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}


