package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type SelfIntro struct {
	bot         *linebot.Client
	appBaseURL  string
	downloadDir string
}

func NewSelfIntro(channelSecret, channelToken string) (*SelfIntro, error) {
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		return nil, err
	}
	return &SelfIntro{
		bot:         bot,
		appBaseURL:  "nil",
		downloadDir: "nil",
	}, nil
}

func ParseMessage(sentence string) (string, string) {
	hmm := seg.CutSearch(sentence, true)
	for _, str := range hmm {
		if value, exist := keywords[str]; exist {
			return value, str
		}
	}
	return "default", "NULL"
}

func (s *SelfIntro) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := s.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		fmt.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := s.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Println(err)
				}
			case *linebot.StickerMessage:
				if err := s.handleSticker(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := s.handleJoin(event.ReplyToken, event.Source); err != nil {
				log.Println(err)
			}
		case linebot.EventTypeJoin:
			if err := s.handleJoin(event.ReplyToken, event.Source); err != nil {
				log.Println(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

func (s *SelfIntro) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	domain, keyword := ParseMessage(message.Text)
	switch domain {
	case "w5":
		works, err := readJSON("static/message/works.json")
		if err != nil {
			return err
		}
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(works))
		if err != nil {
			return err
		}
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(fmt.Sprintf("$$ 偵測到關鍵字 '%s'!\n 推斷你想要知道我的 '%s'！", keyword, domain)).AddEmoji(
				linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "098")).AddEmoji(
				linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "098")),
			linebot.NewFlexMessage("作品集介紹", contents),
		).Do(); err != nil {
			return err
		}
	case "w4":
		experience, err := readJSON("static/message/experience.json")
		if err != nil {
			return err
		}
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(experience))
		if err != nil {
			return err
		}
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(fmt.Sprintf("$$ 偵測到關鍵字 '%s'!\n 推斷你想要知道我的 '%s'！", keyword, domain)).AddEmoji(
				linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "098")).AddEmoji(
				linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "098")),
			linebot.NewFlexMessage("經歷介紹", contents),
		).Do(); err != nil {
			return err
		}
	case "w3":
		education, err := readJSON("static/message/education.json")
		if err != nil {
			return err
		}
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(education))
		if err != nil {
			return err
		}
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(fmt.Sprintf("$$ 偵測到關鍵字 '%s'!\n 推斷你想要知道我的 '%s'！", keyword, domain)).AddEmoji(
				linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "098")).AddEmoji(
				linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "098")),
			linebot.NewFlexMessage("學歷介紹", contents),
		).Do(); err != nil {
			return err
		}
	case "w2":
		profile, _ := s.bot.GetProfile(source.UserID).Do()
		intro, err := readJSON("static/message/intro.json")
		if err != nil {
			return err
		}
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(intro))
		if err != nil {
			return err
		}
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(fmt.Sprintf("$$ 歡迎 %s!!\n 按下方的按鈕來認識我吧！", profile.DisplayName)).AddEmoji(
				linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "098")).AddEmoji(
				linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "098")),
			linebot.NewFlexMessage("自我介紹", contents),
		).Do(); err != nil {
			return err
		}
	case "w1":
		skills, err := readJSON("static/message/skills.json")
		if err != nil {
			return err
		}
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(skills))
		if err != nil {
			return err
		}
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(fmt.Sprintf("$$ 偵測到關鍵字 '%s'!\n 推斷你想要知道我的 '%s'！", keyword, domain)).AddEmoji(
				linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "098")).AddEmoji(
				linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "098")),
			linebot.NewFlexMessage("技能介紹", contents),
		).Do(); err != nil {
			return err
		}
	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(fmt.Sprintf("$$ 謝謝您傳訊息給James!\n 可以按下方主選單問我問題或輸入'介紹自己'喔！")).AddEmoji(
				linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "094")).AddEmoji(
				linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "094")),
		).Do(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SelfIntro) handleSticker(message *linebot.StickerMessage, replyToken string) error {
	if _, err := s.bot.ReplyMessage(
		replyToken,
		linebot.NewStickerMessage(message.PackageID, message.StickerID),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (s *SelfIntro) replyText(replyToken, text string) error {
	if _, err := s.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (s *SelfIntro) handleJoin(replyToken string, source *linebot.EventSource) error {
	profile, _ := s.bot.GetProfile(source.UserID).Do()
	intro, err := readJSON("static/message/intro.json")
	if err != nil {
		return err
	}
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(intro))
	if err != nil {
		return err
	}
	if _, err := s.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(fmt.Sprintf("$$歡迎 %s!!\n 按下方的按鈕來認識我吧！", profile.DisplayName)).AddEmoji(
			linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "098")).AddEmoji(
			linebot.NewEmoji(1, "5ac1bfd5040ab15980c9b435", "098")),
		linebot.NewFlexMessage("自我介紹", contents),
	).Do(); err != nil {
		return err
	}
	return nil
}

func readJSON(file string) ([]byte, error) {
	jsonFile, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
