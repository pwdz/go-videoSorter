#go-videoSorter

go-videoSorter is a CLI applications for **sorting**, **downloading posters** and **gathering information** about movies/series.
##Contents
[Overview](#overview)  
[Installation](#installation)  
[Sort](#sort)  
[Download Image](#download-image)  
[Search](#searching-for-information)
##Overview

go-videoSorter is written using [Cobra](https://github.com/spf13/cobra) library and uses [OMDB](http://www.omdbapi.com/) API in order to fetch data about movies/series and 
can download posters and create text files containing information aboud movies/series. Infos such as
 Writer, Actors, Awards and more.    
 
 At first I wanted to sort videos based on each video's **metadata**. But it could take quite a while since reading each file could take a second or two. Also in some cases, metadatas didn't contain usful information such as title of the video which is the base resource for sorting.  
 Hence I turned to videoNames for sorting.  
##Installation
##Sort
sort command, sorts all the videos existing in the given **source directory** and all of its **subdirectories**
base on the video **title** and **if series**, also **seasons** and **episodes**.  
It works based on the **name of videos**, so in case of irrelevant video namings, sorting series might not work properly.  

The **sort result**, if destinationDir given, will be in ```destionationDir/GoVideoSort/``` if not, will be in ```sourceDir/GoVideoSort/```.    

sort command has the following **optional** flags:  
***--dlImage/-d*** Downloads **poster** of each go-videoSorter automatically while sorting and places it under the folder of the sorted video.  
***--createTxt/-c***  Creates a text file which has information about the go-videoSorter.    

An example of directory tree before and after sort using both flags.  
Before sort:
```
SourceDir
├── code.java
├── ...
├── temp1
|   └── abc                   
|       └── Silicon.Valley.S01E01.BDRip.x264-DEMAND.mkv  
├── temp2 
│   ├── 1917.2019.720p.DVDScr.Farsi.Dubbed.mkv           
│   ├── temp3  
│   │   ├── Se7en.1995.1080p.Farsi.Dubbed.mkv 
|   |   └── Friends.S01E04.720p.mkv
│   ├── Friends.1.2.480p.mkv
│   └── Friends.S01.E03.480p.mkv 
├── ...
└── Friends.1x1.480p.mkv
```
After sort:
```
destinationDir
└── GoVideoSort
    ├── Movies
    |   ├── 1917 (2019)
    |   |   ├── 1917.2019.720p.DVDScr.Farsi.Dubbed.mkv  
    |   |   ├── 1917.jpg
    |   |   └── 1917.txt
    |   └── Se7en (1995)
    |       ├── Se7en.1995.1080p.Farsi.Dubbed.mkv
    |       ...
    └── Series
        ├── Friends
        |   ├── Season1
        |   |   ├── E1
        |   |   |   └── Friends.1x1.480p.mkv
        |   |   ...
        |   ├── Friends.jpg
        |   └── Friends.txt
        └── Silicon Valley
            ├── ...
            ...
```

 For example for movie 1917, using --dlImage/-d will place 1917.jpg under `1917/` folder:  
![1917](./1917.jpg?raw=true)

Also using --createTxt/-c will place 1917.txt by this format under `1917/`
```cassandraql
https://github.com/pwdz 
============O_o============
Title: 1917
Year: 2019
IMDB: 8.4
Metascore: 78
Genre: Drama, War
Type: movie
Writer: Sam Mendes, Krysty Wilson-Cairns
Actors: Dean-Charles Chapman, George MacKay, Daniel Mays, Colin Firth
Plot: April 6th, 1917. As a regiment assembles to wage war deep in enemy territory, two soldiers are assigned to race     against time and deliver a message that will stop 1,600 men from walking straight into a deadly trap.
Language: English, French, German
Awards: Won 3 Oscars. Another 108 wins & 158 nominations.
Poster: https://m.media-amazon.com/images/M/MV5BOTdmNTFjNDEtNzg0My00ZjkxLTg1ZDAtZTdkMDc2ZmFiNWQ1XkEyXkFqcGdeQXVyNTAzN    zgwNTg@._V1_SX300.jpg
Country: USA, UK, India, Spain, Canada
============o_O============
```
NOTICE: Using flags will result in calling API so sort process depending on the count of the existing videos would take some time.  

sort command usage:   
```
govideo sort sourceDirectory [DestinationDirectory] [--dlImage/-d][--createTxt/-c]
```
###Download Image
download command can be used in two ways, first by the [videoName --video/-v] and second by giving it the URL of any image after download it will place it under destionationDir.  
download command usage:   
```
govideo download destinationDir [videoName --video/-v]/[URL]
```
###Searching for Information
search command finds the info of the given videoName and prints the results. If giving destionation, besides printing result, text files of the result will be created under destionationDir.  
If using --search/-s, it will search for all the movies and series that contain the given videoName.  

search result will be just like the 1917.txt example above.  

search command usage:  
```
govideo search videoName [destonationDir] [--search/-s]
```
