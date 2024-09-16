package main

import (
	"ParserVK/pkg/vkapi"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	vkAPIURL  = "https://api.vk.com/method/wall.get"
	vkVersion = "5.199"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	groupID := os.Getenv("GROUP_ID")
	accessToken := os.Getenv("TOKEN")
	if groupID == "" || accessToken == "" {
		log.Fatal("GROUP_ID или TOKEN не определен")
	}
	client := vkapi.NewClient(accessToken)

	posts, err := client.GetPosts(groupID, 20)
	if err != nil {
		log.Fatalf("Ошибка при получении постов: %v", err)
	}

	for _, post := range posts {
		fmt.Printf("Пост ID: %d\nТекст: %s\n\n", post.ID, post.Text)
	}
}
