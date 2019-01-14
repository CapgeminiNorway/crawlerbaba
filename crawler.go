package main

import (
	"./baba"
	"bufio"
	"fmt"
	"github.com/silentsokolov/go-vimeo/vimeo"
	"os"
	"strings"
)

func main() {

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
	vimeoClient := baba.InitVimeoClient(vimeoToken)

	loadVideosJavaZone(vimeoClient)
	loadVideosNDCOslo(vimeoClient)
}

func loadVideosJavaZone(vimeoClient *vimeo.Client) {
	community := baba.Community{}
	community.UserId = "7540193"
	community.Name = "JavaZone"

	community.Albums = albums("JavaZone")
	community.AlbumId = community.Albums["JavaZone-2018"]
	fmt.Println(community)

	for albumName, albumId := range community.Albums {
		community.AlbumId = albumId
		videos := baba.GetAlbumVideos(vimeoClient, community)
		fmt.Printf("--> %v | videos #: %d \n", albumName, len(videos))

		if len(videos) > 0 {
			baba.WriteToFile(albumName, videos)
		}
	}
}

func loadVideosNDCOslo(vimeoClient *vimeo.Client) {
	community := baba.Community{}
	community.UserId = "12026726"
	community.Name = "NDC-Oslo"

	community.Albums = albums("NDC-Oslo")
	community.AlbumId = community.Albums["NDC-Oslo-2018"]
	fmt.Println(community)

	for albumName, albumId := range community.Albums {
		community.AlbumId = albumId
		videos := baba.GetAlbumVideos(vimeoClient, community)
		fmt.Printf("--> %v | videos #: %d \n", albumName, len(videos))

		if len(videos) > 0 {
			baba.WriteToFile(albumName, videos)
		}
	}
}

func albums(which string) (albums map[string]string){

	// TODO: load album list from API
	albums = make(map[string]string)
	if strings.Contains(strings.ToUpper(which), "NDC-OSLO") {
		albums["NDC-Oslo-2018"] = "5477311"
		albums["NDC-Oslo-2017"] = "5491392"
		albums["NDC-Oslo-2016"] = "5506154"
	} else { // default javazone

		albums["JavaZone-2018"] = "5419780"
		albums["JavaZone-2017"] = "4766821"
		albums["JavaZone-2016"] = "4133413"
		//albums["JavaZone-2015"] = "3556815"
		//albums["JavaZone-2014"] = "3031533"
	}


	return
}

func albumName(javaZone baba.Community) (albumName string) {

	for key,val := range javaZone.Albums {
		if javaZone.AlbumId == val {
			albumName = key
			break
		}
	}

	return
}



