package baba

import (
	"fmt"
	"github.com/silentsokolov/go-vimeo/vimeo"
	"golang.org/x/oauth2"
	"strings"
)

func InitVimeoClient(vimeoToken string) (vimeoClient *vimeo.Client) {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: vimeoToken},
	)
	tokenContext := oauth2.NewClient(oauth2.NoContext, tokenSource)
	vimeoClient = vimeo.NewClient(tokenContext, nil)

	if !isValidToken(vimeoClient) {
		vimeoClient = nil
	}

	return
}

func isValidToken(vimeoClient *vimeo.Client) bool  {
	// make a simple API call to check token validity
	_, _, err := vimeoClient.CreativeCommons.List(vimeo.OptPage(1), vimeo.OptPerPage(1))
	return err == nil
}
func GetAlbumVideos(vimeoClient *vimeo.Client, community Community) (videos []Video) {

	optPerPage := 42
	optPage := 1
	optFields := "uri,name,link,release_time,user.uri,user.name,user.link" // description,

	videos = make([]Video, 1)
	for {
		//fmt.Printf("...optPage: %v", optPage)

		vimeoVideos, resp, err := vimeoClient.Users.AlbumListVideo(community.UserId, community.AlbumId, vimeo.OptPage(optPage), vimeo.OptPerPage(optPerPage), vimeo.OptFields{optFields})
		CheckError(err)

		fmt.Printf("Current page: %d\n", resp.Page)
		fmt.Printf("Next page: %s\n", resp.NextPage)
		fmt.Printf("Prev page: %s\n", resp.PrevPage)
		fmt.Printf("Total objects: %d\n", resp.Total)

		if len(strings.TrimSpace(resp.NextPage)) == 0 {
			break
		}
		optPage++
		/*if optPage > 2 {
			fmt.Println("manual break!!")
			break
		}*/

		for _,vimeoVideo := range vimeoVideos {

			video := Video{}
			video.Name = vimeoVideo.Name
			video.PersonName = parsePersonName(vimeoVideo.Name, community.Name)
			video.Link = vimeoVideo.Link
			video.ReleaseTime = vimeoVideo.ReleaseTime
			//video.Description = vimeoVideo.Description
			videos = append(videos, video)
		}
	}

	return
}

func parsePersonName(videoName string, communityName string) string {

	// FIXME: do err parsing check
	/*personName := ""
	if strings.Contains(strings.ToUpper(communityName), "NDC-OSLO") {
		personName = videoName[strings.Index(videoName, "-"):len(videoName)]
		personName = strings.Replace(personName,","," ", -1)
		personName = strings.TrimSpace(personName)
	} else {
		items := strings.Split(videoName, ":")
		personName = items[len(items)-1] // the last item is the personName
		personName = strings.Replace(personName,","," ", -1)
		personName = strings.TrimSpace(personName)
		if len(personName) == 0 {
			personName = "-"
		}
	}
	*/
	items := []string{}
	if strings.Contains(strings.ToUpper(communityName), "NDC-OSLO") {
		items = strings.Split(videoName, "- ")
	} else {
		items = strings.Split(videoName, ": ")
	}
	personName := items[len(items)-1] // the last item is the personName
	personName = strings.Replace(personName,",","", -1)
	personName = strings.TrimSpace(personName)
	if len(personName) == 0 {
		personName = "-"
	}

	return personName
}