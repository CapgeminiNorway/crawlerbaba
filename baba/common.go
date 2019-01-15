package baba

import (
	"fmt"
	"time"
)

type Video struct {
	PersonName string `json:",omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ReleaseTime time.Time `json:"release_time,omitempty"`
}
type Community struct {
	Name string
	UserId  string
	Albums  map[string]string
	AlbumId string
}

func DisplayTokenWarning() {
	//fmt.Println("!!!vimeoToken is NOT valid!!!")
	fmt.Println("!!!you MUST provide vimeoToken to be able to use Vimeo API!!!")
	fmt.Println("get it from https://developer.vimeo.com/api/start ")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
