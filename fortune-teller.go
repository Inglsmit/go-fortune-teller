package main

import (
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var TOKEN = os.Getenv("TG_TOKEN")

var (
	// set pointer
	// [?] what is pointer
	bot                *tgbotapi.BotAPI
	chatId             int64
	fortuneTellerNames = [3]string{"stiven", "stiv", "king"}
	answers            = []string{
		"The most important things are the hardest to say. (The Green Mile)",
		"Monsters are real, and ghosts are real too. They live inside us, and sometimes, they win. (The Shining)",
		"Fear can hold you prisoner. Hope can set you free. (The Shawshank Redemption - novella in Different Seasons)",
		"We make up horrors to help us cope with the real ones. (Danse Macabre)",
		"The trust of the innocent is the liar’s most useful tool. (It: Chapter Two)",
		"Time takes it all, whether you want it to or not. (The Green Mile)",
		"Get busy living, or get busy dying. (The Shawshank Redemption - novella in Different Seasons)",
		"The more you leave out, the more you highlight what you leave in. (On Writing: A Memoir of the Craft)",
		"It is the tale, not he who tells it. (The Wind Through the Keyhole)",
		"Sometimes the things in our heads are far worse than anything they could put in books or on film. (The Stand)",
		"We lie best when we lie to ourselves. (It)",
		"You can't deny laughter; when it comes, it plops down in your favorite chair and stays as long as it wants. (Under the Dome)",
		"The mind can calculate, but the spirit yearns, and the heart knows what the heart knows. (Bag of Bones)",
		"The scariest moment is always just before you start. (On Writing: A Memoir of the Craft)",
		"Some birds are not meant to be caged, that's all. Their feathers are too bright, their songs too sweet and wild. (Different Seasons)",
		"Sooner or later, even the fastest runners have to stand and fight. (The Dark Tower)",
		"We fall from womb to tomb, from one blackness and toward another, remembering little of the one and knowing nothing of the other. (Duma Key)",
		"There's nothing like stories on a windy night when folks have found a warm place in a cold world. (The Eyes of the Dragon)",
		"We never know which lives we influence, or when, or why. (11/22/63)",
		"What happened was roughly what he expected to happen, but the sharper, brighter outline of reality took a little while to come into focus. (The Stand)",
		"The place where you made your stand never mattered. Only that you were there... and still on your feet. (The Stand)",
		"The only real requirement is the ability to remember every scar. (Duma Key)",
		"The world had teeth and it could bite you with them anytime it wanted. (The Girl Who Loved Tom Gordon)",
		"The world was full of monsters with friendly faces. (The Stand)",
		"Memory is the basis of every journey. (The Wind Through the Keyhole)",
		"The beauty of religious mania is that it has the power to explain everything. (Gerald's Game)",
		"It’s funny how things work out, how time can make the best of friends fade away, how things can change but still stay the same. (11/22/63)",
		"All good things must end, but all bad things can begin again. (The Green Mile)",
		"Time is a face on the water. (Duma Key)",
		"The mind can calculate, but the spirit yearns, and the heart knows what the heart knows. (Bag of Bones)",
	}
)

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	//TODO
	msConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msConfig)
}

// [?] what params in func does mean?
// what *mean. Repeat basics knowledge
func isMessageFOrFortuneTeller(update *tgbotapi.Update) bool {
	// check if this is text msg
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}

func getFortuneTellerAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getFortuneTellerAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Ask me question, call me by my name. Answer will be \"YES\" or \"NO\"" +
				"Example: \"King, do I really want change my work?\" ")
		}
		// [?] what means &param ?
		if isMessageFOrFortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
