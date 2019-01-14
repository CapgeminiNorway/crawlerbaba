package baba

import (
	"io"
	"os"
	"strings"
)

func WriteToFile(albumName string, videos []Video) {

	file, err := os.Create("./"+albumName + ".csv")
	CheckError(err)
	defer file.Close()

	_, err = io.WriteString(file, strings.TrimSpace(csvHeader()) + "\n")
	CheckError(err)

	for _, video := range videos {

		if len(strings.TrimSpace(video.Name)) == 0 || strings.Contains(video.Name, "{"){
			continue
		}
		line := composeLine(video)
		_, err := io.WriteString(file, strings.TrimSpace(line) + "\n")
		CheckError(err)
	}
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
