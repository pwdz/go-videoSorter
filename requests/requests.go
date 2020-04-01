package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const api_key = "fdccdde8"
const omdb_api_path = "http://www.omdbapi.com/?"

type Omdb struct {
	Title,Year string
	//Response bool
}
func checkErr(err error) {
	if err!=nil{
		fmt.Println("error msg: ",err)
		os.Exit(1)
	}
}
func Get(reqType string,title string) []Omdb {

	fmt.Println("2222222@")
	URL,params:= setBaseURL()
	fmt.Println("111111111111")
	setParams(&params,reqType,title)

	URL.RawQuery = params.Encode()
	fmt.Println(URL.String())
	response, err := http.Get(URL.String())
	checkErr(err)
	bytes := convertResponseToBytes(response)

	//fmt.Println(string(bytes))

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
	fmt.Println(result[0].Title,result[0].Year)

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
	}
	return result
}
