package sortFuncs

import (
	"fmt"
	"path"
	"sorter/requests"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)
var(
	videoFormats = []string{"mp4","mkv","avi","m4v","m4p","mov","qt","wmv","mpg","mpv","webm","sub","srt"}
	unwantedString = []string{"720","1080","480","dvd","twoddl"}
	sYear,lYear = 1920,2200
	isOnline=false
	shouldCreateTxt=false
	shouldDlImg=false
	sourceDir,desDir string
)
type Video struct{
	Name string
	Year,Episode,Season int
}
func SortVideo(shouldCreateText,shouldDownloadImg bool,paths []string) {
	setSrceAndDes(paths)
	setFlags(shouldCreateText,shouldDownloadImg)

	err := filepath.Walk(sourceDir,processPath)
	if err != nil {
		log.Println(err)
	}
}
func setSrceAndDes(paths []string){
	sourceDir = paths[0]
	if len(paths)>1{
		desDir = paths[1]
	}else{
		desDir = sourceDir
	}
	if strings.LastIndex(desDir,"/")!=len(desDir)-1{
		desDir+="/"
	}
	if !fileExists(desDir){
		os.MkdirAll(desDir,0755)
	}
}
func setFlags(shouldCreateText,shouldDownloadImg bool){
	shouldDlImg = shouldDownloadImg
	shouldCreateTxt = shouldCreateText

	isOnline = shouldCreateTxt || shouldDownloadImg
}
var existingPaths = make(map[string]string)
func processPath(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	for _,vFormat := range videoFormats{
		if strings.HasSuffix(strings.ToLower(info.Name()),vFormat) {
			fmt.Println("\n********************************************")
			videoName := pathPartializer(info.Name(),vFormat)

			video := extractVideoData(videoName)
			basePath, ok := existingPaths[video.Title]
			var newPath string
			if !ok{
				var newBasePath string
				newBasePath,newPath = mkdir(video)
				existingPaths[video.Title] = newBasePath
			}else if video.Type=="series"{
				_,newPath = mkdir(video,basePath)
			}

			fmt.Println("current path:",path)

			mvFile(newPath,path)
			if shouldCreateTxt{
				if !fileExists(strings.Split(newPath,"Season")[0]+video.Title+".txt") {
					CreateInfoFile(newPath,video,true)
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
			fmt.Println("Done.")
			break
		}
	}
	return nil
}
func pathPartializer(name,suffix string)string{
	videoName := strings.Replace(strings.TrimSuffix(name,suffix),"."," ",-1)
	videoName = strings.Replace(videoName,"-"," ",-1)
	videoName = strings.Replace(videoName,"_"," ",-1)
	videoName = strings.Replace(videoName,")"," ",-1)
	videoName = strings.Replace(videoName,"("," ",-1)
	videoName = strings.Replace(videoName,"   "," ",-1)
	videoName = strings.Replace(videoName,"  "," ",-1)
	return videoName
}
func extractVideoData(videoName string)requests.Omdb{
	var v Video
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
	//fmt.Println("name:|"+strings.Trim(v.Name," ")+"|year",v.Year,"|v.Season:",v.Season,"|v.Episode:",v.Episode,"|type:")
	var res requests.Omdb
	res.Response=false
	if isOnline{
		temp:=OnlineSearch(v.Name,"title")
		res = temp[0]
	}
	setOmdbValues(v,&res)
	return res
}
func OnlineSearch(name string,reqType string)[]requests.Omdb{
	return requests.Get(reqType,strings.Trim(name," "))
}
func setOmdbValues(v Video,res *requests.Omdb){
	if v.Season>0{
		res.Type = "series"
		res.Season = strconv.Itoa(v.Season)
		res.Episode = strconv.Itoa(v.Episode)
	}else{
		res.Type = "movie"
	}
	if !res.Response {
		res.Title = v.Name
		if isYear(v.Year) {
			res.Year = strconv.Itoa(v.Year)
		}
	}
}
func PrintVideoInfo(videos []requests.Omdb){
	fmt.Println("by PWDZ")
	fmt.Println("https://github.com/pwdz")
	fmt.Println("============O_o============")
	if (len(videos)==0 || (len(videos)>0 && !videos[0].Response)){
		fmt.Println("No result found! :(")
		return
	}
	for _,v:=range videos{
		fmt.Println("Title: " + v.Title)
		fmt.Println("Year: " + v.Year)
		fmt.Println("IMDB: " + v.ImdbRating)
		fmt.Println("Metascore: " + v.Metascore)
		fmt.Println("Genre: " + v.Genre)
		fmt.Println("Type: " + v.Type)
		fmt.Println("Writer: " + v.Writer)
		fmt.Println("Actors: " + v.Actors)
		fmt.Println("Plot: " + v.Plot)
		fmt.Println("Language: " + v.Language)
		fmt.Println("Awards: " + v.Awards)
		fmt.Println("Poster: " + v.Poster)
		fmt.Println("Country: " + v.Country)
		fmt.Println("***************************")
	}
	fmt.Println("============o_O============")
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
func mkdir(v requests.Omdb,basePath...string)(string,string){
	var newFullPath,newBasePath string
	if v.Type=="movie"{//it's a movie
		newFullPath = desDir+"pwdSorter/Movies/"+strings.Trim(v.Title," ")
		if isYear(isNumber(v.Year)){
			newFullPath += " ("+v.Year+")"
		}
		newBasePath = newFullPath
	}else{//it's a serie
		if len(basePath)>0{
			newBasePath = basePath[0]
		}else {
			newBasePath = desDir + "pwdSorter/Series/" + strings.Trim(v.Title," ")+ "$YEAR"
			if isYear(isNumber(v.Year)){
				newBasePath = strings.Replace(newBasePath,"$YEAR","("+v.Year+")",-1)
			}else {
				newBasePath = strings.Replace(newBasePath,"$YEAR","",-1)
			}
		}
		newFullPath = path.Join(newBasePath,"/Season"+v.Season+ "/E"+v.Episode)
	}
	if !fileExists(newFullPath){
		fmt.Println("Mkdir:",newFullPath)
		os.MkdirAll(newFullPath,0755)
	}
	return newBasePath,newFullPath
}
func mvFile(newPath string,oldPath string){
	fmt.Println("Moving video...")
	parts:=strings.Split(oldPath,"/")
	newPath+= "/"+parts[len(parts)-1]

	fmt.Println("new path:",newPath)
	err :=  os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("In mvFile:",err)
		return
	}
}
func CreateInfoFile(path string,info requests.Omdb,canChangePath bool){
	if info.Response{
		fmt.Println("Creating",info.Title+".txt")
		if info.Type=="series" && canChangePath{
			path= strings.Split(path,"Season")[0]
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path,0755)
		}
		file,err := os.Create(path+"/"+info.Title+".txt")
		if err==nil{
			defer file.Close()
			file.WriteString("by PWDZ\n")
			file.WriteString("https://github.com/pwdz\n")
			file.WriteString("\n============O_o============\n")
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
			file.WriteString("============o_O============\n")
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