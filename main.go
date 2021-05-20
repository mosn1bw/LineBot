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
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
func crawlBlog(num int) *linebot.CarouselTemplate {
	template := linebot.NewCarouselTemplate()
	carouselCols := []*linebot.CarouselColumn{}
	// Request the HTML page.
	res, err := http.Get("https://a28283878.github.io/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".posts-wrapper article").Each(func(i int, s *goquery.Selection) {
		if i >= num {
			return
		}
		title := s.Find("a").Text()
		postURL, _ := s.Find("a").Attr("href")
		pictureURL, _ := s.Find("a").Find("div").Attr("style")
		pictureURL = pictureURL[strings.Index(pictureURL, "(")+1 : strings.Index(pictureURL, ")")]

		column := linebot.NewCarouselColumn(pictureURL, "文章", title, btn)

		carouselCols = append(carouselCols, column)
	})

	template = linebot.NewCarouselTemplate(carouselCols...)

	return template
	
	case "call":
		request, err := requestDao.FindByID(data[1])
		if err != nil {
			return errors.Wrap(err, "failed to find request")
		}

		if request.Finished || request.OperatorID != "" {
			_, err := bot.PushMessage(
				event.Source.UserID,
				linebot.NewTextMessage(
					"申し訳ありません, 他の方が先に申し込みました",
				),
			).Do()

			return err
		}

		request.OperatorID = event.Source.UserID

		request, err = requestDao.Update(request)
		if err != nil {
			return errors.Wrap(err, "failed to update request")
		}

		operators, err := operatorDao.FindAll()
		if err != nil {
			return errors.Wrap(err, "failed to find all operators")
		}

		to := make([]string, 0)
		for _, operator := range operators {
			if operator.UserID != request.OperatorID {
				to = append(to, operator.UserID)
			}
		}

		if _, err := bot.Multicast(
			to,
			linebot.NewTextMessage(
				"締め切りました",
			),
		).Do(); err != nil {
			errors.Wrap(err, "failed to multicast")
		}

		passenger, err := passengerDao.FindByUserID(request.PassengerID)
		if err != nil {
			return errors.Wrap(err, "failed to find passenger")
		}

		if _, err := bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(
				"ご協力ありがとうございます\n"+
					passenger.Name+"様名義で"+
					request.Address+"にタクシーを呼んでください\n"+
					"電話番号は03-5755-2151です\n"+
					"タクシーが到着するまでの時間も聞いてください",
			),
		).Do(); err != nil {
			return errors.Wrap(err, "failed to reply message to operator")
		}

		if _, err := bot.PushMessage(
			request.OperatorID,
			linebot.NewTemplateMessage(
				"時間の確認",
				linebot.NewButtonsTemplate(
					"",
					"時間の確認",
					"タクシーは何分で到着しますか",
					linebot.NewPostbackAction(
						"5分以内",
						"finish:a:"+request.ID,
						"5分以内",
						"",
					),
					linebot.NewPostbackAction(
						"10分以内",
						"finish:b:"+request.ID,
						"10分以内",
						"",
					),
					linebot.NewPostbackAction(
						"15分以内",
						"finish:c:"+request.ID,
						"15分以内",
						"",
					),
					linebot.NewPostbackAction(
						"15分以上",
						"finish:d:"+request.ID,
						"15分以上",
						"",
					),
				),
			).WithQuickReplies(linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton(
					"",
					linebot.NewPostbackAction(
						"配車に失敗",
						"error:"+request.ID,
						"配車に失敗",
						"",
					),
				),
			)),
		).Do(); err != nil {
			return errors.Wrap(err, "failed to push message to operator")
		}
	case "finish":
		request, err := requestDao.FindByID(data[2])
		if err != nil {
			return errors.Wrap(err, "failed to find request")
		}

		if request.Finished {
			return errors.Wrap(err, "already finished")
		}

		if _, err := app.PassengerBot().PushMessage(
			request.PassengerID,
			linebot.NewTextMessage("配車の手続きが完了しました. "+when(data[1])),
		).Do(); err != nil {
			return errors.Wrap(err, "failed to push message to passenger")
		}

		request.Finished = true

		if _, err := requestDao.Update(request); err != nil {
			return errors.Wrap(err, "failed to update request")
		}

	case "error":
		request, err := requestDao.FindByID(data[1])
		if err != nil {
			return errors.Wrap(err, "failed to find request")
		}

		if request.Finished {
			return errors.Wrap(err, "already finished")
		}

		if _, err := bot.PushMessage(
			request.OperatorID,
			linebot.NewTextMessage("残念です"),
		).Do(); err != nil {
			return errors.Wrap(err, "failed to push message to operator")
		}

		if _, err := app.PassengerBot().PushMessage(
			request.PassengerID,
			linebot.NewTextMessage("申し訳ありません, 配車に失敗しました"),
		).Do(); err != nil {
			return errors.Wrap(err, "failed to push message to passenger")
		}

		request.Finished = true

		if _, err := requestDao.Update(request); err != nil {
			return errors.Wrap(err, "failed to update request")
		}
	}

	return nil
}

func when(c string) string {
	switch c {
	case "a":
		return "5分以内にタクシーが到着します, お待ちください"
	case "b":
		return "10分以内にタクシーが到着します, お待ちください"
	case "c":
		return "15分以内にタクシーが到着します, お待ちください"
	case "d":
		return "タクシーの到着には15分以上かかります, しばらくお待ちください"
	}

	return ""
}

func (app *KitchenSink) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "profile":
		if source.UserID != "" {
			profile, err := app.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return app.replyText(replyToken, err.Error())
			}
			if _, err := app.bot.ReplyMessage(
				replyToken,
				linebot.NewTextMessage("Display name: "+profile.DisplayName),
				linebot.NewTextMessage("Status message: "+profile.StatusMessage),
			).Do(); err != nil {
				return err
			}
		} else {
			return app.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	case "buttons":
		imageURL := app.appBaseURL + "/assets/buttons/1040.jpg"
		template := linebot.NewButtonsTemplate(
			imageURL, "My button sample", "Hello, my button",
			linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
			linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", "", "hello こんにちは"),
			linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
			linebot.NewMessageTemplateAction("Say message", "Rice=米"),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Buttons alt text", template),
		).Do(); err != nil {
			return err
		}
	case "confirm":
		template := linebot.NewConfirmTemplate(
			"Do it?",
			linebot.NewMessageTemplateAction("Yes", "Yes!"),
			linebot.NewMessageTemplateAction("No", "No!"),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Confirm alt text", template),
		).Do(); err != nil {
			return err
		}
	case "carousel":
		imageURL := app.appBaseURL + "/assets/buttons/1040.jpg"
		template := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
				linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
				linebot.NewMessageTemplateAction("Say message", "Rice=米"),
			),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Carousel alt text", template),
		).Do(); err != nil {
			return err
		}
	case "image carousel":
		imageURL := app.appBaseURL + "/assets/buttons/1040.jpg"
		template := linebot.NewImageCarouselTemplate(
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewURITemplateAction("Go to LINE", "https://line.me"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewMessageTemplateAction("Say message", "Rice=米"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewDatetimePickerTemplateAction("datetime", "DATETIME", "datetime", "", "", ""),
			),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Image carousel alt text", template),
		).Do(); err != nil {
			return err
		}
	case "datetime":
		template := linebot.NewButtonsTemplate(
			"", "", "Select date / time !",
			linebot.NewDatetimePickerTemplateAction("date", "DATE", "date", "", "", ""),
			linebot.NewDatetimePickerTemplateAction("time", "TIME", "time", "", "", ""),
			linebot.NewDatetimePickerTemplateAction("datetime", "DATETIME", "datetime", "", "", ""),
		)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Datetime pickers alt text", template),
		).Do(); err != nil {
			return err
		}
	case "flex":
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
				Contents: []linebot.FlexComponent{
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
	case "flex carousel":
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
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "First bubble",
							},
						},
					},
				},
				&linebot.BubbleContainer{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "Second bubble",
							},
						},
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
	case "flex json":
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
                "wrap": true,
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
                "wrap": true,
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
          "uri": "https://linecorp.com"
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
	case "imagemap":
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewImagemapMessage(
				app.appBaseURL+"/assets/rich",
				"Imagemap alt text",
				linebot.ImagemapBaseSize{1040, 1040},
				linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
				linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
				linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
				linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
			),
		).Do(); err != nil {
			return err
		}
	case "bye":
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
			replyToken,
			linebot.NewTextMessage(message.Text),
		).Do(); err != nil {
			return err
		}
	}
	return nil
}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
