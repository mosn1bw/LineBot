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
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Constants
var timeFormat = "01/02 PM03:04:05"
var user_zchien = "U696bcb700dfc9254b27605374b86968b"
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
		"",
	}
var answers_ImageMessage = []string{
		"",
	}
var answers_StickerMessage = []string{
		"",
	}
var answers_VideoMessage = []string{
		"",
	}
var answers_AudioMessage = []string{
		"",
	}
var answers_LocationMessage = []string{
		"",
	}
var answers_ReplyCurseMessage = []string{
		"",
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
		log.Println("ساعت بوقت تهران:  " + nowString)
		bot.ReplyMessage(replyToken, linebot.NewTextMessage("ساعت بوقت تهران: " + nowString)).Do()
	} else if silent != True {
		log.Println("ساعت بوقت تهران:  " + nowString)
		bot.PushMessage(replyToken, linebot.NewTextMessage("ساعت بوقت تهران:  " + nowString)).Do()
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
// FlexContainerType type
type FlexContainerType string

// IntPtr is a helper function for using *int values
func IntPtr(v int) *int {
	return &v
}

// FlexContainerType constants
const (
	FlexContainerTypeBubble   FlexContainerType = "bubble"
	FlexContainerTypeCarousel FlexContainerType = "carousel"
)

// FlexComponentType type
type FlexComponentType string

// FlexComponentType constants
const (
	FlexComponentTypeBox       FlexComponentType = "box"
	FlexComponentTypeButton    FlexComponentType = "button"
	FlexComponentTypeFiller    FlexComponentType = "filler"
	FlexComponentTypeIcon      FlexComponentType = "icon"
	FlexComponentTypeImage     FlexComponentType = "image"
	FlexComponentTypeSeparator FlexComponentType = "separator"
	FlexComponentTypeSpacer    FlexComponentType = "spacer"
	FlexComponentTypeSpan      FlexComponentType = "span"
	FlexComponentTypeText      FlexComponentType = "text"
)

// FlexBubbleSizeType type
type FlexBubbleSizeType string

// FlexBubbleSizeType constants
const (
	FlexBubbleSizeTypeNano  FlexBubbleSizeType = "nano"
	FlexBubbleSizeTypeMicro FlexBubbleSizeType = "micro"
	FlexBubbleSizeTypeKilo  FlexBubbleSizeType = "kilo"
	FlexBubbleSizeTypeMega  FlexBubbleSizeType = "mega"
	FlexBubbleSizeTypeGiga  FlexBubbleSizeType = "giga"
)

// FlexBubbleDirectionType type
type FlexBubbleDirectionType string

// FlexBubbleDirectionType constants
const (
	FlexBubbleDirectionTypeLTR FlexBubbleDirectionType = "ltr"
	FlexBubbleDirectionTypeRTL FlexBubbleDirectionType = "rtl"
)

// FlexButtonStyleType type
type FlexButtonStyleType string

// FlexButtonStyleType constants
const (
	FlexButtonStyleTypeLink      FlexButtonStyleType = "link"
	FlexButtonStyleTypePrimary   FlexButtonStyleType = "primary"
	FlexButtonStyleTypeSecondary FlexButtonStyleType = "secondary"
)

// FlexButtonHeightType type
type FlexButtonHeightType string

// FlexButtonHeightType constants
const (
	FlexButtonHeightTypeMd FlexButtonHeightType = "md"
	FlexButtonHeightTypeSm FlexButtonHeightType = "sm"
)

// FlexIconAspectRatioType type
type FlexIconAspectRatioType string

// FlexIconAspectRatioType constants
const (
	FlexIconAspectRatioType1to1 FlexIconAspectRatioType = "1:1"
	FlexIconAspectRatioType2to1 FlexIconAspectRatioType = "2:1"
	FlexIconAspectRatioType3to1 FlexIconAspectRatioType = "3:1"
)

// FlexImageSizeType type
type FlexImageSizeType string

// FlexImageSizeType constants
const (
	FlexImageSizeTypeXxs  FlexImageSizeType = "xxs"
	FlexImageSizeTypeXs   FlexImageSizeType = "xs"
	FlexImageSizeTypeSm   FlexImageSizeType = "sm"
	FlexImageSizeTypeMd   FlexImageSizeType = "md"
	FlexImageSizeTypeLg   FlexImageSizeType = "lg"
	FlexImageSizeTypeXl   FlexImageSizeType = "xl"
	FlexImageSizeTypeXxl  FlexImageSizeType = "xxl"
	FlexImageSizeType3xl  FlexImageSizeType = "3xl"
	FlexImageSizeType4xl  FlexImageSizeType = "4xl"
	FlexImageSizeType5xl  FlexImageSizeType = "5xl"
	FlexImageSizeTypeFull FlexImageSizeType = "full"
)

// FlexImageAspectRatioType type
type FlexImageAspectRatioType string

// FlexImageAspectRatioType constants
const (
	FlexImageAspectRatioType1to1    FlexImageAspectRatioType = "1:1"
	FlexImageAspectRatioType1_51to1 FlexImageAspectRatioType = "1.51:1"
	FlexImageAspectRatioType1_91to1 FlexImageAspectRatioType = "1.91:1"
	FlexImageAspectRatioType4to3    FlexImageAspectRatioType = "4:3"
	FlexImageAspectRatioType16to9   FlexImageAspectRatioType = "16:9"
	FlexImageAspectRatioType20to13  FlexImageAspectRatioType = "20:13"
	FlexImageAspectRatioType2to1    FlexImageAspectRatioType = "2:1"
	FlexImageAspectRatioType3to1    FlexImageAspectRatioType = "3:1"
	FlexImageAspectRatioType3to4    FlexImageAspectRatioType = "3:4"
	FlexImageAspectRatioType9to16   FlexImageAspectRatioType = "9:16"
	FlexImageAspectRatioType1to2    FlexImageAspectRatioType = "1:2"
	FlexImageAspectRatioType1to3    FlexImageAspectRatioType = "1:3"
)

// FlexImageAspectModeType type
type FlexImageAspectModeType string

// FlexImageAspectModeType constants
const (
	FlexImageAspectModeTypeCover FlexImageAspectModeType = "cover"
	FlexImageAspectModeTypeFit   FlexImageAspectModeType = "fit"
)

// FlexBoxLayoutType type
type FlexBoxLayoutType string

// FlexBoxLayoutType constants
const (
	FlexBoxLayoutTypeHorizontal FlexBoxLayoutType = "horizontal"
	FlexBoxLayoutTypeVertical   FlexBoxLayoutType = "vertical"
	FlexBoxLayoutTypeBaseline   FlexBoxLayoutType = "baseline"
)

// FlexComponentPositionType type
type FlexComponentPositionType string

// FlexComponentPositionType constants
const (
	FlexComponentPositionTypeRelative FlexComponentPositionType = "relative"
	FlexComponentPositionTypeAbsolute FlexComponentPositionType = "absolute"
)

// FlexComponentSpacingType type
type FlexComponentSpacingType string

// FlexComponentSpacingType constants
const (
	FlexComponentSpacingTypeNone FlexComponentSpacingType = "none"
	FlexComponentSpacingTypeXs   FlexComponentSpacingType = "xs"
	FlexComponentSpacingTypeSm   FlexComponentSpacingType = "sm"
	FlexComponentSpacingTypeMd   FlexComponentSpacingType = "md"
	FlexComponentSpacingTypeLg   FlexComponentSpacingType = "lg"
	FlexComponentSpacingTypeXl   FlexComponentSpacingType = "xl"
	FlexComponentSpacingTypeXxl  FlexComponentSpacingType = "xxl"
)

// FlexComponentMarginType type
type FlexComponentMarginType string

// FlexComponentMarginType constants
const (
	FlexComponentMarginTypeNone FlexComponentMarginType = "none"
	FlexComponentMarginTypeXs   FlexComponentMarginType = "xs"
	FlexComponentMarginTypeSm   FlexComponentMarginType = "sm"
	FlexComponentMarginTypeMd   FlexComponentMarginType = "md"
	FlexComponentMarginTypeLg   FlexComponentMarginType = "lg"
	FlexComponentMarginTypeXl   FlexComponentMarginType = "xl"
	FlexComponentMarginTypeXxl  FlexComponentMarginType = "xxl"
)

// FlexComponentOffsetType type
type FlexComponentOffsetType string

// FlexComponentOffsetType constants
const (
	FlexComponentOffsetTypeNone FlexComponentOffsetType = "none"
	FlexComponentOffsetTypeXs   FlexComponentOffsetType = "xs"
	FlexComponentOffsetTypeSm   FlexComponentOffsetType = "sm"
	FlexComponentOffsetTypeMd   FlexComponentOffsetType = "md"
	FlexComponentOffsetTypeLg   FlexComponentOffsetType = "lg"
	FlexComponentOffsetTypeXl   FlexComponentOffsetType = "xl"
	FlexComponentOffsetTypeXxl  FlexComponentOffsetType = "xxl"
)

// FlexComponentPaddingType type
type FlexComponentPaddingType string

// FlexComponentPaddingType constants
const (
	FlexComponentPaddingTypeNone FlexComponentPaddingType = "none"
	FlexComponentPaddingTypeXs   FlexComponentPaddingType = "xs"
	FlexComponentPaddingTypeSm   FlexComponentPaddingType = "sm"
	FlexComponentPaddingTypeMd   FlexComponentPaddingType = "md"
	FlexComponentPaddingTypeLg   FlexComponentPaddingType = "lg"
	FlexComponentPaddingTypeXl   FlexComponentPaddingType = "xl"
	FlexComponentPaddingTypeXxl  FlexComponentPaddingType = "xxl"
)

// FlexComponentGravityType type
type FlexComponentGravityType string

// FlexComponentGravityType constants
const (
	FlexComponentGravityTypeTop    FlexComponentGravityType = "top"
	FlexComponentGravityTypeBottom FlexComponentGravityType = "bottom"
	FlexComponentGravityTypeCenter FlexComponentGravityType = "center"
)

// FlexComponentAlignType type
type FlexComponentAlignType string

// FlexComponentAlignType constants
const (
	FlexComponentAlignTypeStart  FlexComponentAlignType = "start"
	FlexComponentAlignTypeEnd    FlexComponentAlignType = "end"
	FlexComponentAlignTypeCenter FlexComponentAlignType = "center"
)

// FlexComponentCornerRadiusType type
type FlexComponentCornerRadiusType string

// FlexComponentCornerRadius constants
const (
	FlexComponentCornerRadiusTypeNone FlexComponentCornerRadiusType = "none"
	FlexComponentCornerRadiusTypeXs   FlexComponentCornerRadiusType = "xs"
	FlexComponentCornerRadiusTypeSm   FlexComponentCornerRadiusType = "sm"
	FlexComponentCornerRadiusTypeMd   FlexComponentCornerRadiusType = "md"
	FlexComponentCornerRadiusTypeLg   FlexComponentCornerRadiusType = "lg"
	FlexComponentCornerRadiusTypeXl   FlexComponentCornerRadiusType = "xl"
	FlexComponentCornerRadiusTypeXxl  FlexComponentCornerRadiusType = "xxl"
)

// FlexIconSizeType type
type FlexIconSizeType string

// FlexIconSizeType constants
const (
	FlexIconSizeTypeXxs FlexIconSizeType = "xxs"
	FlexIconSizeTypeXs  FlexIconSizeType = "xs"
	FlexIconSizeTypeSm  FlexIconSizeType = "sm"
	FlexIconSizeTypeMd  FlexIconSizeType = "md"
	FlexIconSizeTypeLg  FlexIconSizeType = "lg"
	FlexIconSizeTypeXl  FlexIconSizeType = "xl"
	FlexIconSizeTypeXxl FlexIconSizeType = "xxl"
	FlexIconSizeType3xl FlexIconSizeType = "3xl"
	FlexIconSizeType4xl FlexIconSizeType = "4xl"
	FlexIconSizeType5xl FlexIconSizeType = "5xl"
)

// FlexSpacerSizeType type
type FlexSpacerSizeType string

// FlexSpacerSizeType constants
const (
	FlexSpacerSizeTypeXs  FlexSpacerSizeType = "xs"
	FlexSpacerSizeTypeSm  FlexSpacerSizeType = "sm"
	FlexSpacerSizeTypeMd  FlexSpacerSizeType = "md"
	FlexSpacerSizeTypeLg  FlexSpacerSizeType = "lg"
	FlexSpacerSizeTypeXl  FlexSpacerSizeType = "xl"
	FlexSpacerSizeTypeXxl FlexSpacerSizeType = "xxl"
)

// FlexTextWeightType type
type FlexTextWeightType string

// FlexTextWeightType constants
const (
	FlexTextWeightTypeRegular FlexTextWeightType = "regular"
	FlexTextWeightTypeBold    FlexTextWeightType = "bold"
)

// FlexTextSizeType type
type FlexTextSizeType string

// FlexTextSizeType constants
const (
	FlexTextSizeTypeXxs FlexTextSizeType = "xxs"
	FlexTextSizeTypeXs  FlexTextSizeType = "xs"
	FlexTextSizeTypeSm  FlexTextSizeType = "sm"
	FlexTextSizeTypeMd  FlexTextSizeType = "md"
	FlexTextSizeTypeLg  FlexTextSizeType = "lg"
	FlexTextSizeTypeXl  FlexTextSizeType = "xl"
	FlexTextSizeTypeXxl FlexTextSizeType = "xxl"
	FlexTextSizeType3xl FlexTextSizeType = "3xl"
	FlexTextSizeType4xl FlexTextSizeType = "4xl"
	FlexTextSizeType5xl FlexTextSizeType = "5xl"
)

// FlexTextStyleType type
type FlexTextStyleType string

// FlexTextStyleType constants
const (
	FlexTextStyleTypeNormal FlexTextStyleType = "normal"
	FlexTextStyleTypeItalic FlexTextStyleType = "italic"
)

// FlexTextDecorationType type
type FlexTextDecorationType string

// FlexTextDecorationType constants
const (
	FlexTextDecorationTypeNone        FlexTextDecorationType = "none"
	FlexTextDecorationTypeUnderline   FlexTextDecorationType = "underline"
	FlexTextDecorationTypeLineThrough FlexTextDecorationType = "line-through"
)

// FlexComponentJustifyContentType type
type FlexComponentJustifyContentType string

// FlexComponentJustifyContentType constants
const (
	FlexComponentJustifyContentTypeFlexStart    FlexComponentJustifyContentType = "flex-start"
	FlexComponentJustifyContentTypeFlexEnd      FlexComponentJustifyContentType = "flex-end"
	FlexComponentJustifyContentTypeCenter       FlexComponentJustifyContentType = "center"
	FlexComponentJustifyContentTypeSpaceBetween FlexComponentJustifyContentType = "space-between"
	FlexComponentJustifyContentTypeSpaceAround  FlexComponentJustifyContentType = "space-around"
	FlexComponentJustifyContentTypeSpaceEvenly  FlexComponentJustifyContentType = "space-evenly"
)

// FlexComponentAlignItemsType type
type FlexComponentAlignItemsType string

// FlexComponentAlignItemsType constants
const (
	FlexComponentAlignItemsTypeFlexStart FlexComponentAlignItemsType = "flex-start"
	FlexComponentAlignItemsTypeFlexEnd   FlexComponentAlignItemsType = "flex-end"
	FlexComponentAlignItemsTypeCenter    FlexComponentAlignItemsType = "center"
)

// FlexComponentAdjustModeType type
type FlexComponentAdjustModeType string

// FlexComponentAdjustModeType constants
const (
	FlexComponentAdjustModeTypeShrinkToFit FlexComponentAdjustModeType = "shrink-to-fit"
)

// FlexBoxBackgroundType type
type FlexBoxBackgroundType string

// FlexBoxBackgroundType constants
const (
	FlexBoxBackgroundTypeLinearGradient FlexBoxBackgroundType = "linearGradient"
)

// FlexContainer interface
type FlexContainer interface {
	FlexContainer()
}

// BubbleContainer type
type BubbleContainer struct {
	Type      FlexContainerType
	Size      FlexBubbleSizeType
	Direction FlexBubbleDirectionType
	Header    *BoxComponent
	Hero      FlexComponent
	Body      *BoxComponent
	Footer    *BoxComponent
	Styles    *BubbleStyle
}

// MarshalJSON method of BubbleContainer
func (c *BubbleContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      FlexContainerType       `json:"type"`
		Size      FlexBubbleSizeType      `json:"size,omitempty"`
		Direction FlexBubbleDirectionType `json:"direction,omitempty"`
		Header    *BoxComponent           `json:"header,omitempty"`
		Hero      FlexComponent           `json:"hero,omitempty"`
		Body      *BoxComponent           `json:"body,omitempty"`
		Footer    *BoxComponent           `json:"footer,omitempty"`
		Styles    *BubbleStyle            `json:"styles,omitempty"`
	}{
		Type:      FlexContainerTypeBubble,
		Size:      c.Size,
		Direction: c.Direction,
		Header:    c.Header,
		Hero:      c.Hero,
		Body:      c.Body,
		Footer:    c.Footer,
		Styles:    c.Styles,
	})
}

// UnmarshalJSON method of BubbleContainer
func (c *BubbleContainer) UnmarshalJSON(b []byte) error {
	type alias BubbleContainer
	a := struct {
		Hero json.RawMessage `json:"hero,omitempty"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	if a.Hero == nil {
		return nil
	}

	var raw rawFlexComponent
	if err := json.Unmarshal(a.Hero, &raw); err != nil {
		return err
	}
	c.Hero = raw.Component

	return nil
}

// CarouselContainer type
type CarouselContainer struct {
	Type     FlexContainerType
	Contents []*BubbleContainer
}

// MarshalJSON method of CarouselContainer
func (c *CarouselContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     FlexContainerType  `json:"type"`
		Contents []*BubbleContainer `json:"contents"`
	}{
		Type:     FlexContainerTypeCarousel,
		Contents: c.Contents,
	})
}

// FlexContainer implements FlexContainer interface
func (*BubbleContainer) FlexContainer() {}

// FlexContainer implements FlexContainer interface
func (*CarouselContainer) FlexContainer() {}

// BubbleStyle type
type BubbleStyle struct {
	Header *BlockStyle `json:"header,omitempty"`
	Hero   *BlockStyle `json:"hero,omitempty"`
	Body   *BlockStyle `json:"body,omitempty"`
	Footer *BlockStyle `json:"footer,omitempty"`
}

// BlockStyle type
type BlockStyle struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Separator       bool   `json:"separator,omitempty"`
	SeparatorColor  string `json:"separatorColor,omitempty"`
}

// FlexComponent interface
type FlexComponent interface {
	FlexComponent()
}

// BoxComponent type
type BoxComponent struct {
	Type            FlexComponentType
	Layout          FlexBoxLayoutType
	Position        FlexComponentPositionType
	Contents        []FlexComponent
	Flex            *int
	Spacing         FlexComponentSpacingType
	Margin          FlexComponentMarginType
	Width           string
	Height          string
	CornerRadius    FlexComponentCornerRadiusType
	BackgroundColor string
	BorderColor     string
	JustifyContent  FlexComponentJustifyContentType
	AlignItems      FlexComponentAlignItemsType
	Background      *BoxBackground
	OffsetTop       FlexComponentOffsetType
	OffsetBottom    FlexComponentOffsetType
	OffsetStart     FlexComponentOffsetType
	OffsetEnd       FlexComponentOffsetType
	PaddingAll      FlexComponentPaddingType
	PaddingTop      FlexComponentPaddingType
	PaddingBottom   FlexComponentPaddingType
	PaddingStart    FlexComponentPaddingType
	PaddingEnd      FlexComponentPaddingType
}

// MarshalJSON method of BoxComponent
func (c *BoxComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type            FlexComponentType               `json:"type"`
		Layout          FlexBoxLayoutType               `json:"layout"`
		Position        FlexComponentPositionType       `json:"position,omitempty"`
		Contents        []FlexComponent                 `json:"contents"`
		Flex            *int                            `json:"flex,omitempty"`
		Spacing         FlexComponentSpacingType        `json:"spacing,omitempty"`
		Margin          FlexComponentMarginType         `json:"margin,omitempty"`
		Width           string                          `json:"width,omitempty"`
		Height          string                          `json:"height,omitempty"`
		CornerRadius    FlexComponentCornerRadiusType   `json:"cornerRadius,omitempty"`
		BackgroundColor string                          `json:"backgroundColor,omitempty"`
		BorderColor     string                          `json:"borderColor,omitempty"`
		JustifyContent  FlexComponentJustifyContentType `json:"justifyContent,omitempty"`
		AlignItems      FlexComponentAlignItemsType     `json:"alignItems,omitempty"`
		Background      *BoxBackground                  `json:"background,omitempty"`
		OffsetTop       FlexComponentOffsetType         `json:"offsetTop,omitempty"`
		OffsetBottom    FlexComponentOffsetType         `json:"offsetBottom,omitempty"`
		OffsetStart     FlexComponentOffsetType         `json:"offsetStart,omitempty"`
		OffsetEnd       FlexComponentOffsetType         `json:"offsetEnd,omitempty"`
		PaddingAll      FlexComponentPaddingType        `json:"paddingAll,omitempty"`
		PaddingTop      FlexComponentPaddingType        `json:"paddingTop,omitempty"`
		PaddingBottom   FlexComponentPaddingType        `json:"paddingBottom,omitempty"`
		PaddingStart    FlexComponentPaddingType        `json:"paddingStart,omitempty"`
		PaddingEnd      FlexComponentPaddingType        `json:"paddingEnd,omitempty"`
	}{
		Type:            FlexComponentTypeBox,
		Layout:          c.Layout,
		Position:        c.Position,
		Contents:        c.Contents,
		Flex:            c.Flex,
		Spacing:         c.Spacing,
		Margin:          c.Margin,
		Width:           c.Width,
		Height:          c.Height,
		CornerRadius:    c.CornerRadius,
		BackgroundColor: c.BackgroundColor,
		BorderColor:     c.BorderColor,
		Action:          c.Action,
		JustifyContent:  c.JustifyContent,
		AlignItems:      c.AlignItems,
		Background:      c.Background,
		OffsetTop:       c.OffsetTop,
		OffsetBottom:    c.OffsetBottom,
		OffsetStart:     c.OffsetStart,
		OffsetEnd:       c.OffsetEnd,
		PaddingAll:      c.PaddingAll,
		PaddingTop:      c.PaddingTop,
		PaddingBottom:   c.PaddingBottom,
		PaddingStart:    c.PaddingStart,
		PaddingEnd:      c.PaddingEnd,
	})
}

// BoxBackground type
type BoxBackground struct {
	Type           FlexBoxBackgroundType `json:"type,omitempty"`
	Angle          string                `json:"angle,omitempty"`
	StartColor     string                `json:"startColor,omitempty"`
	EndColor       string                `json:"endColor,omitempty"`
	CenterColor    string                `json:"centerColor,omitempty"`
	CenterPosition string                `json:"centerPosition,omitempty"`
}

// ButtonComponent type
type ButtonComponent struct {
	Type         FlexComponentType
	Position     FlexComponentPositionType
	Flex         *int
	Margin       FlexComponentMarginType
	Height       FlexButtonHeightType
	Style        FlexButtonStyleType
	Color        string
	Gravity      FlexComponentGravityType
	AdjustMode   FlexComponentAdjustModeType
	OffsetTop    FlexComponentOffsetType
	OffsetBottom FlexComponentOffsetType
	OffsetStart  FlexComponentOffsetType
	OffsetEnd    FlexComponentOffsetType
}

// MarshalJSON method of ButtonComponent
func (c *ButtonComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type         FlexComponentType           `json:"type"`
		Position     FlexComponentPositionType   `json:"position,omitempty"`
		Flex         *int                        `json:"flex,omitempty"`
		Margin       FlexComponentMarginType     `json:"margin,omitempty"`
		Height       FlexButtonHeightType        `json:"height,omitempty"`
		Style        FlexButtonStyleType         `json:"style,omitempty"`
		Color        string                      `json:"color,omitempty"`
		Gravity      FlexComponentGravityType    `json:"gravity,omitempty"`
		AdjustMode   FlexComponentAdjustModeType `json:"adjustMode,omitempty"`
		OffsetTop    FlexComponentOffsetType     `json:"offsetTop,omitempty"`
		OffsetBottom FlexComponentOffsetType     `json:"offsetBottom,omitempty"`
		OffsetStart  FlexComponentOffsetType     `json:"offsetStart,omitempty"`
		OffsetEnd    FlexComponentOffsetType     `json:"offsetEnd,omitempty"`
	}{
		Type:         FlexComponentTypeButton,
		Action:       c.Action,
		Position:     c.Position,
		Flex:         c.Flex,
		Margin:       c.Margin,
		Height:       c.Height,
		Style:        c.Style,
		Color:        c.Color,
		Gravity:      c.Gravity,
		AdjustMode:   c.AdjustMode,
		OffsetTop:    c.OffsetTop,
		OffsetBottom: c.OffsetBottom,
		OffsetStart:  c.OffsetStart,
		OffsetEnd:    c.OffsetEnd,
	})
}

// FillerComponent type
type FillerComponent struct {
	Type FlexComponentType
	Flex *int
}

// MarshalJSON method of FillerComponent
func (c *FillerComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type FlexComponentType `json:"type"`
		Flex *int              `json:"flex,omitempty"`
	}{
		Type: FlexComponentTypeFiller,
		Flex: c.Flex,
	})
}

// IconComponent type
type IconComponent struct {
	Type        FlexComponentType
	URL         string
	Margin      FlexComponentMarginType
	Size        FlexIconSizeType
	AspectRatio FlexIconAspectRatioType
}

// MarshalJSON method of IconComponent
func (c *IconComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type        FlexComponentType       `json:"type"`
		URL         string                  `json:"url"`
		Margin      FlexComponentMarginType `json:"margin,omitempty"`
		Size        FlexIconSizeType        `json:"size,omitempty"`
		AspectRatio FlexIconAspectRatioType `json:"aspectRatio,omitempty"`
	}{
		Type:        FlexComponentTypeIcon,
		URL:         c.URL,
		Margin:      c.Margin,
		Size:        c.Size,
		AspectRatio: c.AspectRatio,
	})
}

// ImageComponent type
type ImageComponent struct {
	Type            FlexComponentType
	URL             string
	Position        FlexComponentPositionType
	Flex            *int
	Margin          FlexComponentMarginType
	Align           FlexComponentAlignType
	Gravity         FlexComponentGravityType
	Size            FlexImageSizeType
	AspectRatio     FlexImageAspectRatioType
	AspectMode      FlexImageAspectModeType
	BackgroundColor string
	Animated        bool
	OffsetTop       FlexComponentOffsetType
	OffsetBottom    FlexComponentOffsetType
	OffsetStart     FlexComponentOffsetType
	OffsetEnd       FlexComponentOffsetType
}

// MarshalJSON method of ImageComponent
func (c *ImageComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type            FlexComponentType         `json:"type"`
		URL             string                    `json:"url"`
		Position        FlexComponentPositionType `json:"position,omitempty"`
		Flex            *int                      `json:"flex,omitempty"`
		Margin          FlexComponentMarginType   `json:"margin,omitempty"`
		Align           FlexComponentAlignType    `json:"align,omitempty"`
		Gravity         FlexComponentGravityType  `json:"gravity,omitempty"`
		Size            FlexImageSizeType         `json:"size,omitempty"`
		AspectRatio     FlexImageAspectRatioType  `json:"aspectRatio,omitempty"`
		AspectMode      FlexImageAspectModeType   `json:"aspectMode,omitempty"`
		BackgroundColor string                    `json:"backgroundColor,omitempty"`
		Animated        bool                      `json:"animated,omitempty"`
		OffsetTop       FlexComponentOffsetType   `json:"offsetTop,omitempty"`
		OffsetBottom    FlexComponentOffsetType   `json:"offsetBottom,omitempty"`
		OffsetStart     FlexComponentOffsetType   `json:"offsetStart,omitempty"`
		OffsetEnd       FlexComponentOffsetType   `json:"offsetEnd,omitempty"`
	}{
		Type:            FlexComponentTypeImage,
		URL:             c.URL,
		Position:        c.Position,
		Flex:            c.Flex,
		Margin:          c.Margin,
		Align:           c.Align,
		Gravity:         c.Gravity,
		Size:            c.Size,
		AspectRatio:     c.AspectRatio,
		AspectMode:      c.AspectMode,
		BackgroundColor: c.BackgroundColor,
		Action:          c.Action,
		Animated:        c.Animated,
		OffsetTop:       c.OffsetTop,
		OffsetBottom:    c.OffsetBottom,
		OffsetStart:     c.OffsetStart,
		OffsetEnd:       c.OffsetEnd,
	})
}

// SeparatorComponent type
type SeparatorComponent struct {
	Type   FlexComponentType
	Margin FlexComponentMarginType
	Color  string
}

// MarshalJSON method of SeparatorComponent
func (c *SeparatorComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   FlexComponentType       `json:"type"`
		Margin FlexComponentMarginType `json:"margin,omitempty"`
		Color  string                  `json:"color,omitempty"`
	}{
		Type:   FlexComponentTypeSeparator,
		Margin: c.Margin,
		Color:  c.Color,
	})
}

// SpacerComponent type
type SpacerComponent struct {
	Type FlexComponentType
	Size FlexSpacerSizeType
}

// MarshalJSON method of SpacerComponent
func (c *SpacerComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type FlexComponentType  `json:"type"`
		Size FlexSpacerSizeType `json:"size,omitempty"`
	}{
		Type: FlexComponentTypeSpacer,
		Size: c.Size,
	})
}

// SpanComponent type
type SpanComponent struct {
	Type       FlexComponentType
	Text       string
	Size       FlexTextSizeType
	Weight     FlexTextWeightType
	Color      string
	Style      FlexTextStyleType
	Decoration FlexTextDecorationType
}

// MarshalJSON method of SpanComponent
func (c *SpanComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type       FlexComponentType      `json:"type"`
		Text       string                 `json:"text,omitempty"`
		Size       FlexTextSizeType       `json:"size,omitempty"`
		Weight     FlexTextWeightType     `json:"weight,omitempty"`
		Color      string                 `json:"color,omitempty"`
		Style      FlexTextStyleType      `json:"style,omitempty"`
		Decoration FlexTextDecorationType `json:"decoration,omitempty"`
	}{
		Type:       FlexComponentTypeSpan,
		Text:       c.Text,
		Size:       c.Size,
		Weight:     c.Weight,
		Color:      c.Color,
		Style:      c.Style,
		Decoration: c.Decoration,
	})
}

// TextComponent type
type TextComponent struct {
	Type         FlexComponentType
	Text         string
	Position     FlexComponentPositionType
	Contents     []*SpanComponent
	Flex         *int
	Margin       FlexComponentMarginType
	Size         FlexTextSizeType
	Align        FlexComponentAlignType
	Gravity      FlexComponentGravityType
	Wrap         bool
	Weight       FlexTextWeightType
	Color        string
	Style        FlexTextStyleType
	Decoration   FlexTextDecorationType
	MaxLines     *int
	AdjustMode   FlexComponentAdjustModeType
	OffsetTop    FlexComponentOffsetType
	OffsetBottom FlexComponentOffsetType
	OffsetStart  FlexComponentOffsetType
	OffsetEnd    FlexComponentOffsetType
}

// MarshalJSON method of TextComponent
func (c *TextComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type         FlexComponentType           `json:"type"`
		Text         string                      `json:"text,omitempty"`
		Position     FlexComponentPositionType   `json:"position,omitempty"`
		Contents     []*SpanComponent            `json:"contents,omitempty"`
		Flex         *int                        `json:"flex,omitempty"`
		Margin       FlexComponentMarginType     `json:"margin,omitempty"`
		Size         FlexTextSizeType            `json:"size,omitempty"`
		Align        FlexComponentAlignType      `json:"align,omitempty"`
		Gravity      FlexComponentGravityType    `json:"gravity,omitempty"`
		Wrap         bool                        `json:"wrap,omitempty"`
		Weight       FlexTextWeightType          `json:"weight,omitempty"`
		Color        string                      `json:"color,omitempty"`
		Style        FlexTextStyleType           `json:"style,omitempty"`
		Decoration   FlexTextDecorationType      `json:"decoration,omitempty"`
		MaxLines     *int                        `json:"maxLines,omitempty"`
		AdjustMode   FlexComponentAdjustModeType `json:"adjustMode,omitempty"`
		OffsetTop    FlexComponentOffsetType     `json:"offsetTop,omitempty"`
		OffsetBottom FlexComponentOffsetType     `json:"offsetBottom,omitempty"`
		OffsetStart  FlexComponentOffsetType     `json:"offsetStart,omitempty"`
		OffsetEnd    FlexComponentOffsetType     `json:"offsetEnd,omitempty"`
	}{
		Type:         FlexComponentTypeText,
		Text:         c.Text,
		Position:     c.Position,
		Contents:     c.Contents,
		Flex:         c.Flex,
		Margin:       c.Margin,
		Size:         c.Size,
		Align:        c.Align,
		Gravity:      c.Gravity,
		Wrap:         c.Wrap,
		Weight:       c.Weight,
		Color:        c.Color,
		Action:       c.Action,
		Style:        c.Style,
		Decoration:   c.Decoration,
		MaxLines:     c.MaxLines,
		AdjustMode:   c.AdjustMode,
		OffsetTop:    c.OffsetTop,
		OffsetBottom: c.OffsetBottom,
		OffsetStart:  c.OffsetStart,
		OffsetEnd:    c.OffsetEnd,
	})
}

// FlexComponent implements FlexComponent interface
func (*BoxComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*ButtonComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*FillerComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*IconComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*ImageComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*SeparatorComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*SpacerComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*SpanComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*TextComponent) FlexComponent() {}
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
	case "a1":
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
	case "a2":
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
	case "a3":
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
	case "image a4":
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
	case "datetime1":
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
		if event.Type == linebot.EventTypeMessage {
			_, silent := silentMap[sourceId]
			
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				log.Print("ReplyToken[" + replyToken + "] TextMessage: ID(" + message.ID + "), Text(" + message.Text  + "), current silent status=" + strconv.FormatBool(silent) )
				//if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
				//	log.Print(err)
				//}
				
				if strings.Contains(message.Text, "test") {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("success")).Do()
				} else if "groupid"  == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(string(source.GroupID))).Do()
				} else if "help" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("help\n・[image:画像url]=從圖片網址發送圖片\n・[speed]=測回話速度\n・[groupid]=發送GroupID\n・[roomid]=發送RoomID\n・[byebye]=取消訂閱\n・[about]=作者\n・[me]=發送發件人信息\n・[test]=test bowwow是否正常\n・[now]=現在時間\n・[mid]=mid\n・[sticker]=隨機圖片\n\n[其他機能]\n位置測試\n捉貼圖ID\n加入時發送消息")).Do()
				} else if "mid" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(source.UserID)).Do()
				} else if "mee6" == message.Text  {
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()
				} else if "mee5" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()
					bot.ReplyMessage(replyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg","https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg")).Do()
				} else if "roomid" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(source.RoomID)).Do()
				} else if "hidden" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("hidden")).Do()
				} else if "bowwow" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg","https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg")).Do()}
					if err != nil {
						log.Fatal(err)
				} else if "me" == message.Text  {
					mid := source.UserID
					p, err := bot.GetProfile(mid).Do()
					if err != nil {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("新增同意"))
					}

					bot.ReplyMessage(replyToken, linebot.NewTextMessage("mid:"+mid+"\nname:"+p.DisplayName+"\nstatusMessage:"+p.StatusMessage)).Do()
				} else if res := strings.Contains(message.Text, "hello"); res == True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("hello!"), linebot.NewTextMessage("my name is bowwow")).Do()
				} else if res := strings.Contains(message.Text, "image:"); res == True {
					image_url := strings.Replace(message.Text, "image:", "", -1)
					bot.ReplyMessage(replyToken, linebot.NewImageMessage(image_url, image_url)).Do()
				} else if "1" == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9..0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.\n\n💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀"), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9..0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.\n\n💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀")).Do()
				} else if  "time" == message.Text {
					tellTime(replyToken, True)
				} else if  "4" == message.Text {
					silentMap[sourceId] = false
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"),linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"),linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"), linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"), linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H") ,linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H")).Do()
				} else if "profile" == message.Text {
					if source.UserID != "" {
						profile, err := bot.GetProfile(source.UserID).Do()
						if err != nil {
							log.Print(err)
						} else if _, err := bot.ReplyMessage(
							replyToken,
							linebot.NewTextMessage("Display name: "+profile.DisplayName + ", Status message: "+profile.StatusMessage)).Do(); err != nil {
								log.Print(err)
						}
					} else {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("Bot can't use profile API without user ID")).Do()
					}
				} else if "buttons" == message.Text {
					imageURL := baseURL + "/static/buttons/1040.jpg"
					//log.Print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "+imageURL)
					template := linebot.NewButtonsTemplate(
						imageURL, "My button sample", "Hello, my button",
						linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
						linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
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
					imageURL := baseURL + "/static/buttons/1040.jpg"
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
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
				} else if "/bye" == message.Text {
					if rand.Intn(100) > 70 {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("請神容易送神難, 我偏不要, 嘿嘿")).Do()
					} else {
						switch source.Type {
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
				} else if "無恥" == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_ReplyCurseMessage[rand.Intn(len(answers_ReplyCurseMessage))])).Do()
				} else if silentMap[sourceId] != True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_TextMessage[rand.Intn(len(answers_TextMessage))])).Do()
				}
			case *linebot.ImageMessage :
				log.Print("ReplyToken[" + replyToken + "] ImageMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), PreviewImageURL(" + message.PreviewImageURL + ")" )
				if silent != True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_ImageMessage[rand.Intn(len(answers_ImageMessage))])).Do()
				}
			case *linebot.VideoMessage :
				log.Print("ReplyToken[" + replyToken + "] VideoMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), PreviewImageURL(" + message.PreviewImageURL + ")" )
				if silent != True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_VideoMessage[rand.Intn(len(answers_VideoMessage))])).Do()
				}
			case *linebot.AudioMessage :
				log.Print("ReplyToken[" + replyToken + "] AudioMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), Duration(" + strconv.Itoa(message.Duration) + ")" )
				if silent != True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_AudioMessage[rand.Intn(len(answers_AudioMessage))])).Do()
				}
			case *linebot.LocationMessage:
				log.Print("ReplyToken[" + replyToken + "] LocationMessage[" + message.ID + "] Title(" + message.Title  + "), Address(" + message.Address + "), Latitude(" + strconv.FormatFloat(message.Latitude, 'f', -1, 64) + "), Longitude(" + strconv.FormatFloat(message.Longitude, 'f', -1, 64) + ")" )
				if silent != True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_LocationMessage[rand.Intn(len(answers_LocationMessage))])).Do()
				}
			case *linebot.StickerMessage :
				log.Print("ReplyToken[" + replyToken + "] StickerMessage[" + message.ID + "] PackageID(" + message.PackageID + "), StickerID(" + message.StickerID + ")" )
				if silent != True {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_StickerMessage[rand.Intn(len(answers_StickerMessage))])).Do()
				}
			}
		} else if event.Type == linebot.EventTypePostback {
		} else if event.Type == linebot.EventTypeBeacon {
		}
	}
	
}
