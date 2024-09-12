package bot

import (
	"bufio"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
	"strings"
	"time"
)

func SaveUserID(chatID int64) error {
	file, err := os.OpenFile("user_ids.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("%d\n", chatID))
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func GetAllUserIDs() ([]int64, error) {
	file, err := os.Open("user_ids.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var userIDs []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var id int64
		_, err := fmt.Sscanf(scanner.Text(), "%d", &id)
		if err == nil {
			userIDs = append(userIDs, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}

func BotLog(message string) {
	fmt.Println(message, "----------------------------------------------------------------")
	botToken := "7499621695:AAHObWgsPI9uSaCq1p1E1PRK7zmfb4Gu3OU" // Bot tokeningizni kiriting
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userIDs, err := GetAllUserIDs()
	if err != nil {
		fmt.Println("Error reading user IDs:", err)
		return
	}

	for _, chatID := range userIDs {
		_, err := bot.SendMessage(tu.Message(tu.ID(chatID), message))
		if err != nil {
			fmt.Printf("Error sending message to %d: %v\n", chatID, err)
		} else {
			fmt.Printf("Message sent to %d\n", chatID)
		}
	}

	timeout := time.After(5 * time.Second)
	updateChan, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

loop:
	for {
		select {
		case update := <-updateChan:
			if update.Message != nil {
				mas := update.Message.Text
				chatID := update.Message.Chat.ID

				if strings.ToLower(mas) == "/start" {
					bot.SendMessage(tu.Message(tu.ID(chatID), "salom dodi dan"))

					err := SaveUserID(chatID)
					if err != nil {
						fmt.Println("Error saving user ID:", err)
					} else {
						fmt.Printf("User ID %d saved.\n", chatID)
					}
				}
			}
		case <-timeout:
			fmt.Println("10 seconds passed, stopping bot.")
			break loop
		}
	}
}
