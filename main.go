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
// https://github.com/line/line-bot-sdk-go/tree/master/linebot

package main

import (
	"strconv"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Constants
var timeFormat = "01/02 PM03:04:05"
var user_mosen = "u2023c2d6c4de3dc7c266f3f07cfabdcc"
var user_yaoming = "U3aaab6c6248bb38f194134948c60f757"
var user_jackal = "U3effab06ddf5bcf0b46c1c60bcd39ef5"
var user_shane = "U2ade7ac4456cb3ca99ffdf9d7257329a"

// Global Settings
var channelSecret = os.Getenv("CHANNEL_SECRET")
var channelToken = os.Getenv("CHANNEL_TOKEN")
//var baseURL = os.Getenv("APP_BASE_URL")
var baseURL = "https://line-talking-bot-go.herokuapp.com"
var endpointBase = os.Getenv("ENDPOINT_BASE")
var tellTimeInterval int = 15
var answers_TextMessage = []string{
	}

var silentMap = make(map[string]bool) // [UserID/GroupID/RoomID]:bool

//var echoMap = make(map[string]bool)

var loc, _ = time.LoadLocation("Asia/Tehran")
var bot *linebot.Client


func tellTime(replyToken string, doTell bool){
	var silent = false
	now := time.Now().In(loc)
	nowString := now.Format(timeFormat)
	
	if doTell {
		log.Println("Time time (Tehran): " + nowString)
		bot.ReplyMessage(replyToken, linebot.NewTextMessage("Time time (Tehran): " + nowString)).Do()
	} else if silent != true {
		log.Println("Automatic report(Tehran): " + nowString)
		bot.PushMessage(replyToken, linebot.NewTextMessage("Automatic report(Tehran): " + nowString)).Do()
	} else {
		log.Println("tell time misfired")
	}
}

func tellTimeJob(sourceId string) {
	for {
		time.Sleep(time.Duration(tellTimeInterval) * time.Minute)
		now := time.Now().In(loc)
		log.Println("time to tell time to : " + sourceId + ", " + now.Format(timeFormat))
		tellTime(sourceId, false)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	/*
	go func() {
		tellTimeJob(user_mosen);
	}()
	go func() {
		for {
			now := time.Now().In(loc)
			log.Println("keep alive at : " + now.Format(timeFormat))
			//http.Get("https://line-talking-bot-go.herokuapp.com")
			time.Sleep(time.Duration(rand.Int31n(29)) * time.Minute)
		}
	}()
	*/

	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	log.Print("URL:"  + r.URL.String())
	
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		var replyToken = event.ReplyToken

		var source = event.Source //EventSource		
		var userId = source.UserID
		var groupId = source.GroupID
		var roomId = source.RoomID
		log.Print("callbackHandler to source UserID/GroupID/RoomID: " + userId + "/" + groupId + "/" + roomId)
		
		var sourceId = roomId
		if sourceId == "" {
			sourceId = groupId
			if sourceId == "" {
				sourceId = userId
			}
		}
		
		if event.Type == linebot.EventTypeMessage {
			_, silent := silentMap[sourceId]
			
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				log.Print("ReplyToken[" + replyToken + "] TextMessage: ID(" + message.ID + "), Text(" + message.Text  + "), current silent status=" + strconv.FormatBool(silent) )
				//if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
				//	log.Print(err)
				//}
				
				if source.UserID != "" && source.UserID != user_mosen {
					profile, err := bot.GetProfile(source.UserID).Do()
					if err != nil {
						log.Print(err)
					} else if _, err := bot.PushMessage(user_mosen, linebot.NewTextMessage(profile.DisplayName + ": "+message.Text)).Do(); err != nil {
							log.Print(err)
					}
				}
				
				if strings.Contains(message.Text, "你閉嘴") {
					silentMap[sourceId] = true
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
                }    
				if strings.Contains(message.Text, "1") {
					silentMap[sourceId] = true
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("Bot can't use profile API without user ID")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("Bot can't use profile API without user ID")).Do()
				} else if strings.Contains(message.Text, "time") {
					tellTime(replyToken, true)
				} else if "say" == message.Text {
					silentMap[sourceId] = false
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("Microphone test，1、2、3... OK")).Do()
				} else if "buttons" == message.Text {
					imageURL := "https://lh3.googleusercontent.com/-xHqQP4wTZDU/YBq5AgqjvCI/AAAAAAAAL6c/TmVGaX4tgIk07K5bZIPDtV9Ct49xEwaxwCK8BGAsYHg/s512/2021-02-03.gif"
					//log.Print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "+imageURL)
					template := linebot.NewButtonsTemplate(
						imageURL, "My button sample", "Hello, my button",
						linebot.NewURITemplateAction("Go to line.me", "line://ti/p/~M_BW"),
						linebot.NewURITemplateAction("Say hello1", "https://lh3.googleusercontent.com/-xHqQP4wTZDU/YBq5AgqjvCI/AAAAAAAAL6c/TmVGaX4tgIk07K5bZIPDtV9Ct49xEwaxwCK8BGAsYHg/s512/2021-02-03.gif"),
						linebot.NewPostbackTemplateAction("言 hello2", "hello2", "hello こんにちは"),
						linebot.NewMessageTemplateAction("Say message", "Rice=米"),
					)
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Buttons alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "confirm" == message.Text {
					template := linebot.NewConfirmTemplate(
						"Do it?",
						linebot.NewMessageTemplateAction("Yes", "Yes!"),
						linebot.NewMessageTemplateAction("No", "No!"),
					)
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Confirm alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "carousel" == message.Text {
					imageURL := "https://lh3.googleusercontent.com/-buBdz24kuAQ/XzzphunjcDI/AAAAAAAAIVI/FJXAP-jE3X0PlpcuwiyHeHBJepS8JU1sgCK8BGAsYHg/s512/2020-08-19.png"
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewURITemplateAction("Go to line.me", "line://ti/p/~M_BW"),
							linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
						),
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
							linebot.NewMessageTemplateAction("Say message", "Rice=米"),
						),
					)
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Carousel alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "imagemap" == message.Text {
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewImagemapMessage(
							baseURL + "/static/rich",
							"Imagemap alt text",
							linebot.ImagemapBaseSize{1040, 1040},
							linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
							linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
						),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "你滾開" == message.Text {
					if rand.Intn(100) > 70 {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("請神容易送神難, 我偏不要, 嘿嘿")).Do()
					} else {
						switch source.Type 
						case linebot.EventSourceTypeUser:
							bot.ReplyMessage(replyToken, linebot.NewTextMessage("我想走, 但是我走不了...")).Do()
						case linebot.EventSourceTypeGroup:
							bot.ReplyMessage(replyToken, linebot.NewTextMessage("我揮一揮衣袖 不帶走一片雲彩")).Do()
							bot.LeaveGroup(source.GroupID).Do()
						case linebot.EventSourceTypeRoom:
							bot.ReplyMessage(replyToken, linebot.NewTextMessage("我揮一揮衣袖 不帶走一片雲彩")).Do()
							bot.LeaveRoom(source.RoomID).Do()
						}
					}
				} else if "image carousel"== message.Text {
					imageURL := app.appBaseURL + "/static/buttons/1040.jpg"
					template := linebot.NewImageCarouselTemplate(
						linebot.NewImageCarouselColumn(
							imageURL,
							linebot.NewURIAction("Go to LINE", "https://line.me"),
						),
						linebot.NewImageCarouselColumn(
							imageURL,
							linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", ""),
						),
						linebot.NewImageCarouselColumn(
							imageURL,
							linebot.NewMessageAction("Say message", "Rice=米"),
						),
						linebot.NewImageCarouselColumn(
							imageURL,
							linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
						),
					)
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Image carousel alt text", template),
					).Do(); err != nil {
						return err
					}
				
				} else if "datetime" == message.Text {
					template := linebot.NewButtonsTemplate(
						"", "", "Select date / time !",
						linebot.NewDatetimePickerAction("date", "DATE", "date", "", "", ""),
						linebot.NewDatetimePickerAction("time", "TIME", "time", "", "", ""),
						linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
					)
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Datetime pickers alt text", template),
					).Do(); err != nil {
						return err
					}
				} else if "flex" == message.Text {
					// {
					//   "type": "bubble",
					//   "body": {
					//     "type": "box",
					//     "layout": "horizontal",
					//     "contents": [
					//       {
					//         "type": "text",
					//         "text": "Hello,"
					//       },
					//       {
					//         "type": "text",
					//         "text": "World!"
					//       }
					//     ]
					//   }
					// }
					contents := &linebot.BubbleContainer{
						Type: linebot.FlexContainerTypeBubble,
						Body: &linebot.BoxComponent{
							Type:   linebot.FlexComponentTypeBox,
							Layout: linebot.FlexBoxLayoutTypeHorizontal,
							Contents: []linebot.FlexComponent,
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: "Hello,",
								},
								&linebot.TextComponent{
									Type: linebot.FlexComponentTypeText,
									Text: "World!",
								},
							},
						},
					}
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewFlexMessage("Flex message alt text", contents),
					).Do(); err != nil {
						return err
					}
				} else if "flex carousel" == message.Text {
					// {
					//   "type": "carousel",
					//   "contents": [
					//     {
					//       "type": "bubble",
					//       "body": {
					//         "type": "box",
					//         "layout": "vertical",
					//         "contents": [
					//           {
					//             "type": "text",
					//             "text": "First bubble"
					//           }
					//         ]
					//       }
					//     },
					//     {
					//       "type": "bubble",
					//       "body": {
					//         "type": "box",
					//         "layout": "vertical",
					//         "contents": [
					//           {
					//             "type": "text",
					//             "text": "Second bubble"
					//           }
					//         ]
					//       }
					//     }
					//   ]
					// }
					contents := &linebot.CarouselContainer{
						Type: linebot.FlexContainerTypeCarousel,
						Contents: []*linebot.BubbleContainer{
							{
								Type: linebot.FlexContainerTypeBubble,
								Body: &linebot.BoxComponent{
									Type:   linebot.FlexComponentTypeBox,
									Layout: linebot.FlexBoxLayoutTypeVertical,
									Contents: []linebot.FlexComponent,
										&linebot.TextComponent{
											Type: linebot.FlexComponentTypeText,
											Text: "First bubble",
										},
									},
								},
							},
								Type: linebot.FlexContainerTypeBubble,
								Body: &linebot.BoxComponent{
									Type:   linebot.FlexComponentTypeBox,
									Layout: linebot.FlexBoxLayoutTypeVertical,
									Contents: []linebot.FlexComponent{,
										&linebot.TextComponent{
											Type: linebot.FlexComponentTypeText,
											Text: "Second bubble",
										}
									}
								}
							}
						}
					}
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewFlexMessage("Flex message alt text", contents),
					).Do(); err != nil {
						return err
					}
				} else if "flex 2" == message.Text {
					jsonString := `{
					"type": "bubble",
					"hero": {
					"type": "image",
					"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
					"size": "full",
					"aspectRatio": "20:13",
					"aspectMode": "cover",
					"action": {
					"type": "uri",
					"uri": "http://linecorp.com/"
 			     }
 			 },
			  "body": {
    			"type": "box",
			    "layout": "vertical",
 			   "contents": [
  			    {
   			     "type": "text",
    	         "text": "Brown Cafe",
    	         "weight": "bold",
    	         "size": "xl"
    	       },
    	       {
    	         "type": "box",
     	        "layout": "baseline",
    	         "margin": "md",
    	         "contents": [
    	           {
    	             "type": "icon",
    	             "size": "sm",
    	             "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
     	          },
     	          {
     	            "type": "icon",
     	            "size": "sm",
     	            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
     	          },
     	          {
     	            "type": "icon",
    	             "size": "sm",
    	             "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
       	        },
       	        {
       	          "type": "icon",
       	          "size": "sm",
       	          "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
       	        },
       	        {
        	      "type": "icon",
         	      "size": "sm",
        	      "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
                },
       	        {
       	          "type": "text",
        	      "text": "4.0",
         	      "size": "sm",
        	      "color": "#999999",
                  "margin": "md",
                  "flex": 0
   	            }
   	          ]
   	        },
    	       {
                  "type": "box",
                  "layout": "vertical",
                  "margin": "lg",
                  "spacing": "sm",
                  "contents": [
    	           {
                  "type": "box",
       	          "layout": "baseline",
                  "spacing": "sm",
                  "contents": [
       	            {
       	              "type": "text",
       	              "text": "Place",
        	             "color": "#aaaaaa",
        	             "size": "sm",
        	             "flex": 1
        	           },
         	          {
                      "type": "text",
                      "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
                      "wrap": True,
                      "color": "#666666",
           	          "size": "sm",
                      "flex": 5
       	            }
        	         ]
         	      },
    			      {
    			        "type": "box",
    			        "layout": "baseline",
    			        "spacing": "sm",
    			        "contents": [
    			          {
    			            "type": "text",
    			            "text": "Time",
    			            "color": "#aaaaaa",
    			            "size": "sm",
    			            "flex": 1
    			          },
    			          {
    			            "type": "text",
    			            "text": "10:00 - 23:00",
    			            "wrap": True,
    			            "color": "#666666",
    			            "size": "sm",
    			            "flex": 5
    			          }
    			        ]
    			      }
    			    ]
    			  }
    			]
			  },
			    "footer": {
    			"type": "box",
    			"layout": "vertical",
    			"spacing": "sm",
    			"contents": [
    			  {
    			    "type": "button",
    			    "style": "link",
    			    "height": "sm",
    			    "action": {
    			      "type": "uri",
    			      "label": "CALL",
    			      "uri": "https://linecorp.com"
    			    }
    			  },
    			  {
    			    "type": "button",
    			    "style": "link",
    			    "height": "sm",
    			    "action": {
    			      "type": "uri",
    			      "label": "WEBSITE",
    			      "uri": "https://linecorp.com",
    			      "altUri": {
    			        "desktop": "https://line.me/ja/download"
    			      }
    			    }
    			  },
    			  {
    			    "type": "spacer",
    			    "size": "sm"
    			  }
    			],
    			"flex": 0
 			    }
		        	}`
					contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
					if err != nil {
						return err
					}
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewFlexMessage("Flex message alt text", contents),
					).Do(); err != nil {
						return err
					}
					
				} else if "imagemap" == message.Text {
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewImagemapMessage(
							app.appBaseURL+"/static/rich",
							"Imagemap alt text",
							linebot.ImagemapBaseSize{Width: 1040, Height: 1040},
							linebot.NewURIImagemapAction("LINE Store Manga", "https://store.line.me/family/manga/en", linebot.ImagemapArea{X: 0, Y: 0, Width: 520, Height: 520}),
							linebot.NewURIImagemapAction("LINE Store Music", "https://store.line.me/family/music/en", linebot.ImagemapArea{X: 520, Y: 0, Width: 520, Height: 520}),
							linebot.NewURIImagemapAction("LINE Store Play", "https://store.line.me/family/play/en", linebot.ImagemapArea{X: 0, Y: 520, Width: 520, Height: 520}),
							linebot.NewMessageImagemapAction("URANAI!", "URANAI!", linebot.ImagemapArea{X: 520, Y: 520, Width: 520, Height: 520}),
						),
					).Do(); err != nil {
						return err
					}
				} else if "imagemap v" == message.Text {
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewImagemapMessage(
							app.appBaseURL+"/static/rich",
							"Imagemap with video alt text",
							linebot.ImagemapBaseSize{Width: 1040, Height: 1040},
							linebot.NewURIImagemapAction("LINE Store Manga", "https://store.line.me/family/manga/en", linebot.ImagemapArea{X: 0, Y: 0, Width: 520, Height: 520}),
							linebot.NewURIImagemapAction("LINE Store Music", "https://store.line.me/family/music/en", linebot.ImagemapArea{X: 520, Y: 0, Width: 520, Height: 520}),
							linebot.NewURIImagemapAction("LINE Store Play", "https://store.line.me/family/play/en", linebot.ImagemapArea{X: 0, Y: 520, Width: 520, Height: 520}),
							linebot.NewMessageImagemapAction("URANAI!", "URANAI!", linebot.ImagemapArea{X: 520, Y: 520, Width: 520, Height: 520}),
									).WithVideo(&linebot.ImagemapVideo{
							OriginalContentURL: app.appBaseURL + "/static/imagemap/video.mp4",
							PreviewImageURL:    app.appBaseURL + "/static/imagemap/preview.jpg",
							Area:               linebot.ImagemapArea{X: 280, Y: 385, Width: 480, Height: 270},
							ExternalLink:       &linebot.ImagemapVideoExternalLink{LinkURI: "https://line.me", Label: "LINE"},
						}),
					).Do(); err != nil {
						return err
					}
				} else if "quick" == message.Text {
					if _, err := app.bot.ReplyMessage(
						replyToken,
						linebot.NewTextMessage("Select your favorite food category or send me your location!").
							WithQuickReplies(linebot.NewQuickReplyItems(
								linebot.NewQuickReplyButton(
									app.appBaseURL+"/static/quick/sushi.png",
									linebot.NewMessageAction("Sushi", "Sushi")),
								linebot.NewQuickReplyButton(
									app.appBaseURL+"/static/quick/tempura.png",
									linebot.NewMessageAction("Tempura", "Tempura")),
								linebot.NewQuickReplyButton(
									"",
									linebot.NewLocationAction("Send location")),
							)),
					).Do(); err != nil {
						return err
					}
				} else if "bye" == message.Text {
					switch source.Type {
					case linebot.EventSourceTypeUser:
						return app.replyText(replyToken, "Bot can't leave from 1:1 chat")
					case linebot.EventSourceTypeGroup:
						if err := app.replyText(replyToken, "Leaving group"); err != nil {
							return err
						}
						if _, err := app.bot.LeaveGroup(source.GroupID).Do(); err != nil {
							return app.replyText(replyToken, err.Error())
						}
					case linebot.EventSourceTypeRoom:
						if err := app.replyText(replyToken, "Leaving room"); err != nil {
							return err
						}
						if _, err := app.bot.LeaveRoom(source.RoomID).Do(); err != nil {
							return app.replyText(replyToken, err.Error())
						}
					}
				default:
					log.Printf("Echo message to %s: %s", replyToken, message.Text)
					if _, err := app.bot.ReplyMessage(
						replyToken,{
						linebot.NewTextMessage(message.Text),
					).Do(); err != nil {
						return err
					}
                }   
				return nil														

				} else if "無恥" == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_ReplyCurseMessage[rand.Intn(len(answers_ReplyCurseMessage))])).Do()
				} else if silentMap[sourceId] != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_TextMessage[rand.Intn(len(answers_TextMessage))])).Do()
				}
			case *linebot.ImageMessage :
				log.Print("ReplyToken[" + replyToken + "] ImageMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), PreviewImageURL(" + message.PreviewImageURL + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_ImageMessage[rand.Intn(len(answers_ImageMessage))])).Do(){
				}
			case *linebot.VideoMessage :
				log.Print("ReplyToken[" + replyToken + "] VideoMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), PreviewImageURL(" + message.PreviewImageURL + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_VideoMessage[rand.Intn(len(answers_VideoMessage))])).Do(){
				}
			case *linebot.AudioMessage :
				log.Print("ReplyToken[" + replyToken + "] AudioMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), Duration(" + strconv.Itoa(message.Duration) + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_AudioMessage[rand.Intn(len(answers_AudioMessage))])).Do(){
				}
			case *linebot.LocationMessage:
				log.Print("ReplyToken[" + replyToken + "] LocationMessage[" + message.ID + "] Title(" + message.Title  + "), Address(" + message.Address + "), Latitude(" + strconv.FormatFloat(message.Latitude, 'f', -1, 64) + "), Longitude(" + strconv.FormatFloat(message.Longitude, 'f', -1, 64) + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_LocationMessage[rand.Intn(len(answers_LocationMessage))])).Do(){
				}
			case *linebot.StickerMessage :
				log.Print("ReplyToken[" + replyToken + "] StickerMessage[" + message.ID + "] PackageID(" + message.PackageID + "), StickerID(" + message.StickerID + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_StickerMessage[rand.Intn(len(answers_StickerMessage))])).Do(){
				}
			}
		} else if event.Type == linebot.EventTypePostback {
		} else if event.Type == linebot.EventTypeBeacon {
		}
	}
	
}
