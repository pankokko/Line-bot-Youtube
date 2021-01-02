package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

var (
	query           = flag.String("query", "rabbit", "Search term")
	maxResults      = flag.Int64("max-results", 25, "Max YouTube results")
	mediaType       = flag.String("type", "video", "classify content")
	videoCategoryId = flag.String("video-category-id", "15", "video categoryId")
)

const BASEURL = "https://www.youtube.com/watch?v="


func fetchYoutubeMovieUrl() string {
	flag.Parse() // コマンドラインから受け取った引数を解析する。　go run youtube.go -query test の場合
	//flag.String("query", "test", "Search term") 第二引数がtestに入れ替わる　
	spew.Dump(*query)

	cfg, _ := ini.Load("config.ini")
	youtubeApiKey := cfg.Section("youtube").Key("api_key").MustString("")

	client := &http.Client{
		Transport: &transport.APIKey{Key: youtubeApiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	snippet := []string{"snippet"}

	call := service.Search.List(snippet).
		Q(*query).
		MaxResults(*maxResults).
		VideoCategoryId(*videoCategoryId).
		Type(*mediaType)

	response, _ := call.Do()

	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	videoIds := filterVideoIds("Videos", videos)

	url := makeYoutubeUrl(videoIds)
	return url
}

func filterVideoIds(sectionName string, matches map[string]string) []string {
	fmt.Printf("%v:\n", sectionName)
	var videoIds []string

	for id, title := range matches {
		videoIds = append(videoIds, id)
		fmt.Printf("[%v] %v\n", id, title)
	}
	return videoIds
}

func makeYoutubeUrl(videoIds []string) string {
	videoUrl := BASEURL + videoIds[0]
	return videoUrl
}
