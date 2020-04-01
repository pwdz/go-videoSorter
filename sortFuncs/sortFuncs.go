package sortFuncs

import (
	"fmt"
	//"golang.org/x/text/search"

	//"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
	// "sorter/requests"
	"strconv"
	"strings"
)
var videoFormats = []string{"mp4","mkv","avi","m4v","m4p","mov","qt","ogg","wmv","mpg","mpv","webm"}
var unwantedString = []string{"720","1080","480","dvd","twoddl"}
var sYear,lYear = 1920,2200
type video struct{
	Name string
	Year,Episode,Season int
}
func SortVideo()  {
	err := filepath.Walk("D:/M&S",processPath)
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
			fmt.Println("\n********************************************")
			fmt.Println(path,"[------]", strings.Trim(info.Name(),strings.TrimSuffix(info.Name(),vFormat)))
			videoName := strings.Replace(strings.TrimSuffix(info.Name(),vFormat),"."," ",-1)
			videoName = strings.Replace(videoName,"-"," ",-1)
			videoName = strings.Replace(videoName,"_"," ",-1)
			videoName = strings.Replace(videoName,")"," ",-1)
			videoName = strings.Replace(videoName,"("," ",-1)
			videoName = strings.Replace(videoName,"   "," ",-1)
			videoName = strings.Replace(videoName,"  "," ",-1)
			//mkdir()
			// fmt.Println(videoName)
			video := findVideo(videoName)
			mkdir(video)
			//parts := strings.Split(videoName," ")
			//fmt.Println(parts)
			break
		}
	}
	//}
	return nil
}
func findVideo(videoName string)video{
	var v video
	parts := strings.Split(strings.Trim(videoName," ")," ")


	printSlice(parts)
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

						// fmt.Println("palang malang before",index)
						index=lastIndex
						// fmt.Println("palang malang after",index)
					}
					flag=isNameEmpty(v.Name)
				}
			} else {
				//fmt.Println("num bede aqa",num,"|name:"+name+"|")
				if v.Name == "" {
					//fmt.Println("0_____0")
					v.Name += strconv.Itoa(num)+" "
				} else if isYear(num) {
					//fmt.Println("akhe :|",num)
					v.Year = num
					flag=isNameEmpty(v.Name)
				} else {
					//fmt.Println("dafaq? in azonas :|--------------------------")
					seas,ep,num,lastIndex:=isSeasonAndEpisode(val,index,parts)
					if !setSeasonEpisode(seas,ep,&v.Season,&v.Episode){
						v.Name += strconv.Itoa(num)+" "


						// fmt.Println("palang malang before",index)


						// fmt.Println("palang malang afte",index)
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
	// fmt.Println(requests.Get("title",strings.Trim(name," ")))
	return v
}
/*
func a(result []requests.Omdb,parts []string,startIndex int){
	maxCount :=0
	counter := 0
	for _,res := range  result{
		for i:=startIndex;i<len(parts);i++{
			if !isValid(parts[i]){
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
}*/
func printSlice(sss []string){
	// fmt.Println("******************BEGIN*********************")
	for _,v:=range sss{

		if isValid(v){
			// fmt.Println("|"+v+"|")

		}else{
			//fmt.Println("sssssssssssssssS")
		}
	}
	// fmt.Println("********************END**********************")
}
func isNameEmpty(name string)bool{
	//fmt.Println("^_^",name)
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
	// fmt.Println("in season function:",str)
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
		//fmt.Println("in replace s:",str)
		var index int
		var ch int32
		for index,ch = range str{
			//fmt.Printf("%T",ch)
			//fmt.Println("in for:|",ch,string(ch))
			if s,err:=strconv.Atoi(string(ch));err==nil{
				//fmt.Println("in convert:",s,season)
				season*=10
				season+=s
			}else {
				if index==strings.Index(str,"e"){
					//fmt.Println("E dar:",str)
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
	//fmt.Println("in season func:",season,episode)
	return season,episode,nameNum,lastIndex
}
func isEpisode(parts []string,epIndex int)(int,int){//episode,lastIndex
	// fmt.Println(epIndex,len(parts))
	if epIndex>=len(parts)-1{
		return 0,epIndex
	}
	episode := 0
	lastIndex:=epIndex
	// fmt.Println("EPPPPPP2222222222222")
	str := strings.ToLower(parts[epIndex])
	// fmt.Println("EPPPPPP",str)
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
				//fmt.Println("@!@!@!@!@!@!@!@!@!@!@@",str)
				if epIndex+1 < len(parts) {
					// fmt.Println("@########################",parts[epIndex+1])
					if ep, err := strconv.Atoi(parts[epIndex+1]); err == nil {
						lastIndex = epIndex+1
						episode = ep
						//fmt.Println("$$$$$$$$$$$$$$$$$$",episode,ep)
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
func mkdir(v video){
	
}