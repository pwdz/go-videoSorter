package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const api_key = "fdccdde8"
const omdb_api_path = "http://www.omdbapi.com/?"

type Omdb struct {
	Title string
	Year string
	Episode string
	Season string
	Genre string
	Writer string
	Actors string
	Plot string
	Language string
	Country string
	Awards string
	Poster string
	ImdbRating string
	Metascore string
	Type string
	ImdbID string

	Response bool
}
func checkErr(err error) {
	if err!=nil{
		fmt.Println("error msg: ",err)
		os.Exit(1)
	}
}
func Get(reqType string,title string) []Omdb {
	URL,params:= setBaseURL()
	setParams(&params,reqType,title)

	URL.RawQuery = params.Encode()
	fmt.Println(URL.String())
	response, err := http.Get(URL.String())
	checkErr(err)
	bytes := convertResponseToBytes(response)
	defer response.Body.Close()

	var result []Omdb
	switch reqType {
	case "search":
		result = parseSearchResult(bytes)
	case "title":
		result = parseTitleResult(bytes)
	}

	return result
}
func setBaseURL() (*url.URL,url.Values){
	URL, err := url.Parse(omdb_api_path)
	checkErr(err)

	URL.Path += "/"
	params := url.Values{}
	params.Add("apiKey", api_key)
	return URL,params
}
func setParams(params *url.Values,reqType string,title string){
	switch reqType{
	case "search":
		params.Add("s",title)
		params.Add("e","json")
	case "title":
		params.Add("t",title)
		params.Add("plot","short")
		params.Add("r","json")
	}
}
func convertResponseToBytes(response *http.Response) []byte{
	bytes,_ := ioutil.ReadAll(response.Body)
	return bytes
}
func parseTitleResult(response []byte) []Omdb{
	result := make([]Omdb,1)

	json.Unmarshal(response,&result[0])
	if result[0].Title!=""{
		result[0].Response=true
	}
	fmt.Println("[33333]",result[0].Response,result[0].Season)

	return result
}
func parseSearchResult(response []byte) []Omdb  {
	var searchJson map[string][]interface{}
	json.Unmarshal(response, &searchJson)
	searchRes := searchJson["Search"]

	result := make([]Omdb,len(searchRes))
	for index,video:= range searchRes{
		//each value contains a movie/serie data
		videoData := video.(map[string]interface{})

		result[index].Title = fmt.Sprintf("%v",videoData["Title"])
		result[index].Year = fmt.Sprintf("%v",videoData["Year"])
		result[index].ImdbID = fmt.Sprintf("%v",videoData["imdbID"])
		result[index].Type = fmt.Sprintf("%v",videoData["Type"])
		result[index].Poster = fmt.Sprintf("%v",videoData["Poster"])
		result[index].Response = true
	}
	return result
}
func DownloadFile(URL, path, fileName string) error {
	fileURL,_:=url.Parse(URL)
	parts := strings.Split(fileURL.Path,"/")
	lastPartOfURL:=strings.Split(parts[len(parts)-1],".")
	suffix := lastPartOfURL[len(lastPartOfURL)-1]

	response, err := http.Get(URL)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	file, err := os.Create(path+"/"+fileName+"."+suffix)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
