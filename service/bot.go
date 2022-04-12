package service

import (
	"context"
	"log"
	"time"

	pb "github.com/Shahboz4131/bot-service/genproto"
	l "github.com/Shahboz4131/bot-service/pkg/logger"


	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotService struct {
	// storage storage.IStorage
	logger l.Logger
}

// NewBotService ...
func NewBotService(log l.Logger) *BotService {
	// func NewTaskService(storage storage.IStorage, log l.Logger) *TaskService {
	return &BotService{
		// storage: storage,
		logger: log,
	}
}

const token = "5216887760:AAGMtcpjpieOkG4_eryPwTavQWzURUgPtYU"

var low []string
var medium []string
var high []string

func (s *BotService) GetMessage(ctx context.Context, req *pb.Message) (*pb.EmptyRes, error) {

	if req.Priority == "low" {
		low = append(low, req.Text)
	} else if req.Priority == "medium" {
		medium = append(medium, req.Text)
	} else if req.Priority == "high" {
		high = append(high, req.Text)
	}

	go SendMessage()

	return &pb.EmptyRes{}, nil
}

func SendMessage() {

	time.Sleep(5000 * time.Millisecond)

	var a = true

	for a {
		if len(high) == 0 {
			break
		}
		SendBot(high[0])

		high = high[1:]
	}

	for a {
		if len(medium) == 0 {
			break
		}
		SendBot(medium[0])

		medium = medium[1:]
	}

	for a {
		if len(low) == 0 {
			break
		}
		SendBot(low[0])

		low = low[1:]
	}
}

func SendBot(message string) {

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	chatid := int64(977890590)

	time.Sleep(10000 * time.Millisecond)

	msg := tgbotapi.NewMessage(chatid, message)

	bot.Send(msg)
}
