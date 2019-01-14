package main

import (
	"bufio"
	"fmt"
	"github.com/silentsokolov/go-vimeo/vimeo"
	"golang.org/x/oauth2"
	"io"
	"os"
	"strings"
	"time"
)

type Video struct {
	PersonName string `json:",omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ReleaseTime time.Time `json:"release_time,omitempty"`
}
type JavaZone struct {
	UserId  string
	Albums  map[string]string
	AlbumId string
}


func main() {

	/*
	OK! ask user input for
		OK! vimeoToken as text or load from file
		choose Javazone albums: 2016, 2017, 2018

	OK! collect data from Vimeo
	OK! write as a CSV file
	*/

	// --- vimeoToken ---
	// try to get vimeToken from env-vars
	vimeoToken := os.Getenv("VIMEO_TOKEN")
	//fmt.Printf("token from ENV: %v \n", vimeoToken)
	vimeoToken = strings.TrimSpace(vimeoToken)

	if len(vimeoToken) == 0 || len(vimeoToken) < 20 { // if not ask the user to enter it manually
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your VIMEO token: ")
		vimeoToken, _ := reader.ReadString('\n')
		vimeoToken = strings.TrimSpace(vimeoToken)
		if len(vimeoToken) == 0 || len(vimeoToken) < 20 {
			fmt.Println("!!!MUST provide vimeoToken to be able to use Vimeo API!!!")
			fmt.Println("get it from https://developer.vimeo.com/api/start ")
			panic("!!!missing vimeoToken!!!")
		}
	}
	vimeoClient := initVimeoClient(vimeoToken)

	javaZone := JavaZone{}
	javaZone.UserId = "7540193"
	// TODO: NDC channelId 1411895

	javaZone.Albums = getAlbums()
	javaZone.AlbumId = javaZone.Albums["JavaZone-2018"]
	fmt.Println(javaZone)

	for albumName, albumId := range javaZone.Albums {
		javaZone.AlbumId = albumId
		videos := getAlbumVideos(vimeoClient, javaZone)
		fmt.Printf("--> %v | videos #: %d \n", albumName, len(videos))

		if len(videos) > 0 {
			writeToFile(getAlbumName(javaZone), videos)
		}
	}

}

func initVimeoClient(vimeoToken string) (vimeoClient *vimeo.Client) {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: vimeoToken},
	)
	tokenContext := oauth2.NewClient(oauth2.NoContext, tokenSource)
	vimeoClient = vimeo.NewClient(tokenContext, nil)

	// FIXME: make a simple API call to validate the token

	return
}

func csvHeader() (lineHeader string){

	lineHeader = ""
	lineHeader += "PersonName, Name, Link, ReleaseTime, "//Description, "
	lineHeader += "GitHub, DuckDuckGo, Twitter, LinkedIn"
	return
}
func composeLine(video Video) (line string) {

	duckduckgo := "https://duckduckgo.com/?q="
	github := "https://github.com/search?type=Users&q="
	twitter := "https://twitter.com/search?f=users&q="
	linkedin := "https://www.linkedin.com/search/results/people/?keywords="

	video.Name = strings.Replace(video.Name,",","", -1)
	video.Name = strings.Replace(video.Name,";","", -1)

	line = ""
	line += video.PersonName + ", "
	line += video.Name + ", "
	line += video.Link + ", "
	line += video.ReleaseTime.String() + ", "
	//line += strings.Replace(video.Description,","," ", -1) + ", "

	searchTerm := strings.Replace(video.PersonName," ","%20", -1)
	//searchTerm = strings.Replace(video.PersonName,","," ", -1)
	github += searchTerm
	line += github + ", "
	duckduckgo += searchTerm
	line += duckduckgo + ", "
	twitter += searchTerm
	line += twitter + ", "
	linkedin += searchTerm
	line += linkedin + ", "

	//fmt.Println(line)

	return
}
func writeToFile(albumName string, videos []Video) {

	file, err := os.Create("./"+albumName + ".csv")
	checkError(err)
	defer file.Close()

	_, err = io.WriteString(file, strings.TrimSpace(csvHeader()) + "\n")
	checkError(err)

	for _, video := range videos {

		if len(strings.TrimSpace(video.Name)) == 0 || strings.Contains(video.Name, "{"){
			continue
		}
		line := composeLine(video)
		_, err := io.WriteString(file, strings.TrimSpace(line) + "\n")
		checkError(err)
	}
}

func getNdcOsloChannels() (channels map[string]string) {
	// Oslo 2018 https://vimeo.com/channels/1411895/videos
	channels = make(map[string]string)
	channels["NDC-Oslo-2018"] = "1411895"
	channels["NDC-Oslo-2017"] = "1264603"
	channels["NDC-Oslo-2016"] = "ndcoslo2016"

	return
}

func getAlbums() (albums map[string]string){

	// TODO: load album list from API
	albums = make(map[string]string)
	albums["JavaZone-2018"] = "5419780"
	//albums["JavaZone-2017"] = "4766821"
	//albums["JavaZone-2016"] = "4133413"
	// albums["JavaZone-2015"] = "3556815"
	// albums["JavaZone-2014"] = "3031533"

	return
}

func getAlbumName(javaZone JavaZone) (albumName string) {

	for key,val := range javaZone.Albums {
		if javaZone.AlbumId == val {
			albumName = key
			break
		}
	}

	return
}

func getAlbumVideos(vimeoClient *vimeo.Client, javaZone JavaZone) (videos []Video) {

	optPerPage := 42
	optPage := 1
	optFields := "uri,name,link,release_time,user.uri,user.name,user.link" // description,

	/*
	// make use of go-lib for vimeo
	fmt.Println("--- using VIMEO Lib ---")
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: vimeoToken},
	)
	tokenContext := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := vimeo.NewClient(tokenContext, nil)
	*/

	videos = make([]Video, 1)
	nextPage := 1
	for nextPage>0 {
		//fmt.Printf("...optPage: %v", optPage)

		vimeoVideos, resp, err := vimeoClient.Users.AlbumListVideo(javaZone.UserId, javaZone.AlbumId, vimeo.OptPage(optPage), vimeo.OptPerPage(optPerPage), vimeo.OptFields{optFields})
		checkError(err)

		fmt.Printf("Current page: %d\n", resp.Page)
		fmt.Printf("Next page: %s\n", resp.NextPage)
		fmt.Printf("Prev page: %s\n", resp.PrevPage)
		fmt.Printf("Total objects: %d\n", resp.Total)

		if len(strings.TrimSpace(resp.NextPage)) == 0 {
			nextPage = 0
		}
		optPage++

		/*if optPage > 2 {
			nextPage = 0
			fmt.Println("manual break!!")
		}*/

		for _,vimeoVideo := range vimeoVideos {

			video := Video{}
			video.Name = vimeoVideo.Name
			video.PersonName = parsePersonName(vimeoVideo.Name)
			video.Link = vimeoVideo.Link
			video.ReleaseTime = vimeoVideo.ReleaseTime
			//video.Description = vimeoVideo.Description
			videos = append(videos, video)
		}
	}

	return
}

func parsePersonName(videoName string) string {
	items := strings.Split(videoName, ":")
	// FIXME: do err parsing check
	personName := items[len(items)-1] // the last item is the personName
	personName = strings.Replace(personName,","," ", -1)
	personName = strings.TrimSpace(personName)
	if len(personName) == 0 {
		personName = "-"
	}
	return personName
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
