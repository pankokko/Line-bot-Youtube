package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/line/line-bot-sdk-go/linebot" //==========この行を追加==========
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}

func LineHandler(w http.ResponseWriter, r *http.Request) {

	cfg, _ := ini.Load("config.ini")

	accessToken := cfg.Section("line").Key("channel_access_token").MustString("")
	channelSecret := cfg.Section("line").Key("channel_secret").MustString("")

	bot, _ := linebot.New(
		channelSecret,
		accessToken,
	)

	body, _ := ioutil.ReadAll(r.Body)

	var webhook Webhook

	if err := json.Unmarshal(body, &webhook); err != nil {
		log.Print(err)
	}
	for _, event := range webhook.Events {
		switch event.Type {
			case linebot.EventTypeMessage:
				switch m := event.Message.(type) {
				case *linebot.TextMessage:
					spew.Dump(m)
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(m.Text)).Do()
				}
		}
	}


	//for _, event := range webhook.Events {
	//	switch event.Type {
	//	case linebot.EventTypeMessage: //type: messageのとき
	//		switch m := event.Message.(type) {
	//		case *linebot.TextMessage: //type: textのとき
	//			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(m.Text)).Do(); err != nil {
	//				log.Print(err)
	//				return events.APIGatewayProxyResponse{
	//					StatusCode: http.StatusInternalServerError,
	//					Body:       fmt.Sprintf(`{"message":"%s"}`+"\n", http.StatusText(http.StatusBadRequest)),
	//				}, nil
	//			}
	//		}
	//	}
	//}
	//return events.APIGatewayProxyResponse{
	//	StatusCode: 200,
	//}, nil

}

//func IndexHandler(w http.ResponseWriter, r *http.Request)  {
//	t, err := template.ParseFiles("Views/index.html")
//	if err != nil{
//		fmt.Println("error")
//	}
//	 if err := t.Execute(w,nil); err != nil{
//	 	fmt.Println("error")
//	 }
//
//}

func main() {
	http.HandleFunc("/", LineHandler)
	log.Fatalln(http.ListenAndServe(":8070", nil))
}
