// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/dhowden/tag"
	"github.com/remko/go-mkvparse"
	"log"
	"os"
	"sort"
	"sorter/sortFuncs"
	"time"
)
//var a [4]string = {}

//const movieTypes =  []string{""}
type MyParser struct {
	currentTagGlobal bool
	currentTagName   *string
	currentTagValue  *string
	title            *string
	tags             map[string]string
}

func (p *MyParser) HandleMasterBegin(id mkvparse.ElementID, info mkvparse.ElementInfo) (bool, error) {
	if id == mkvparse.TagElement {
		p.currentTagGlobal = true
	} else if id == mkvparse.SimpleTagElement {
		p.currentTagName = nil
		p.currentTagValue = nil
	}
	return true, nil
}

func (p *MyParser) HandleMasterEnd(id mkvparse.ElementID, info mkvparse.ElementInfo) error {
	if id == mkvparse.SimpleTagElement && p.currentTagGlobal && p.currentTagName != nil && p.currentTagValue != nil {
		p.tags[*p.currentTagName] = *p.currentTagValue
	}
	return nil
}

func (p *MyParser) HandleString(id mkvparse.ElementID, value string, info mkvparse.ElementInfo) error {
	if id == mkvparse.TagNameElement {
		p.currentTagName = &value
	} else if id == mkvparse.TagStringElement {
		p.currentTagValue = &value
	} else if id == mkvparse.TitleElement {
		p.title = &value
	}
	return nil
}

func (p *MyParser) HandleInteger(id mkvparse.ElementID, value int64, info mkvparse.ElementInfo) error {
	if (id == mkvparse.TagTrackUIDElement || id == mkvparse.TagEditionUIDElement || id == mkvparse.TagChapterUIDElement || id == mkvparse.TagAttachmentUIDElement) && value != 0 {
		p.currentTagGlobal = false
	}
	return nil
}

func (p *MyParser) HandleFloat(id mkvparse.ElementID, value float64, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleDate(id mkvparse.ElementID, value time.Time, info mkvparse.ElementInfo) error {
	return nil
}

func (p *MyParser) HandleBinary(id mkvparse.ElementID, value []byte, info mkvparse.ElementInfo) error {
	return nil
}

func main() {
	//cmd.Execute()
	fmt.Println("here u go:")
	//stat, _ := os.Stat("/home/mmd/Desktop/first-season/Friends.S01.E10.480p.mkv")
	//fmt.Println(stat.Name(),stat.IsDir(),stat.Mode(),stat.ModTime())
	sortFuncs.SortVideo()
	fmt.Println("end")
	//buf, _ := ioutil.ReadFile("./Friends.S01.E01.480p.mkv")
	//
	//kind, _ := filetype.Match(buf)
	//if kind == filetype.Unknown {
	//	fmt.Println("Unknown file type")
	//	return
	//}
	fmt.Println("kon-----------------------")
/*	Lib,err := metadata.Init(1000000,"metadata.db","https://developers.themoviedb.org/3/", "TVRAGE-API-KEY", "TVDB-API-KEY")
	if err == nil {
		data, err := Lib.GetMetadata("Argo","movie")
		if err == nil {
			fmt.Println("kon-----------------------")
			fmt.Println(data)
		}else{
			fmt.Println("goooooooooooooooooooooooooooooz",err)
		}
	}else {
		fmt.Println("||||||||||||||||||||||||||")
	}*/

	//fmt.Println("koon+++++++++++++++++++++++")
	//fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
	//fmt.Println(kind.MIME.Type)
	//fmt.Println(kind.MIME.Subtype)


	fmt.Println("maaaaaaaaaaaaaan1")
	f,_ := os.Open("./012 String Values.mp4")
	m, err := tag.ReadFrom(f)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}
	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).
	log.Print(m.FileType())  // The title of the track (see Metadata interface for more details).
	log.Print(m.Disc())  // The title of the track (see Metadata interface for more details).
	log.Print(m.Comment())  // The title of the track (see Metadata interface for more details).
	log.Print(m)  // The title of the track (see Metadata interface for more details).
	fmt.Println("maaaaaaaaaaaaaan")

	//file, err := os.Open("Friends.S01.E01.480p.mkv")
	file, err := os.Open("Legend.2015.720p.Farsi.Dubbed.mkv")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	defer file.Close()
	handler := MyParser{
		tags: make(map[string]string),
	}
	err = mkvparse.ParseSections(file, &handler, mkvparse.InfoElement, mkvparse.TagsElement)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	// Print (sorted) tags
	if handler.title != nil {
		fmt.Printf("- title: %q\n", *handler.title)
	}
	fmt.Println("wtf?")
	var tagNames []string
	for tagName := range handler.tags {
		tagNames = append(tagNames, tagName)
	}
	sort.Strings(tagNames)
	for _, tagName := range tagNames {
		fmt.Printf("- %s: %q\n", tagName, handler.tags[tagName])
	}
	//------------------------------------------------------
	/*file, err := os.Open("Friends.S01.E01.480p.mkv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	defer file.Close()

	doc := mkv.InitDocument(file)
	err = doc.GetElementContent()
	err = doc.ParseAll(func(el mkv.Element) {
		fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}*/
/*	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		return
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata("./Legend.2015.720p.Farsi.Dubbed.mkv")

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			fmt.Printf("[%v] %v\n", k, v)
		}
	}*/



}

