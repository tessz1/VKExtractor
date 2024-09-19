package ui

import (
	"ParserVK/pkg/vkapi"
	"fmt"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

func RunUI() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accessToken := os.Getenv("TOKEN")
	if accessToken == "" {
		log.Fatal("TOKEN environment variable is not set")
	}

	a := app.New()
	w := a.NewWindow("VK Post Fetcher")

	groupIDEntry := widget.NewEntry()
	groupIDEntry.SetPlaceHolder("Enter Group ID")
	countEntry := widget.NewEntry()
	countEntry.SetPlaceHolder("Enter number of Posts")

	result := widget.NewMultiLineEntry()
	result.Disable()
	result.Wrapping = fyne.TextWrapWord
	var posts []vkapi.Post
	currentPage := 0
	postsPerPage := 20

	nextButton := widget.NewButton("Next", func() {
		if len(posts) == 0 {
			result.SetText("No posts to display")
			return
		}

		start := currentPage * postsPerPage
		if start >= len(posts) {
			start = 0
			currentPage = 0
		}
		end := start + postsPerPage
		if end > len(posts) {
			end = len(posts)
		}

		resultText := ""
		for _, post := range posts[start:end] {
			resultText += fmt.Sprintf("Post ID: %d\nText: %s\n\n", post.ID, post.Text)
		}
		result.SetText(resultText)

		currentPage++

		if currentPage*postsPerPage >= len(posts) {
			currentPage = 0
		}
	})

	fetchButton := widget.NewButton("Fetch Posts", func() {
		groupID := groupIDEntry.Text
		count, err := strconv.Atoi(countEntry.Text)
		if err != nil {
			result.SetText("Invalid number of posts")
			return
		}

		client := vkapi.NewClient(accessToken)
		posts, err = client.GetPosts(groupID, count)
		if err != nil || len(posts) == 0 {
			result.SetText(fmt.Sprintf("Error: %v", err))
			return
		}

		currentPage = 0
		nextButton.OnTapped()
	})

	scrollContainer := container.NewScroll(result)
	scrollContainer.SetMinSize(fyne.NewSize(600, 600))

	content := container.NewVBox(
		widget.NewLabel("VK Post Fetcher"),
		groupIDEntry,
		countEntry,
		fetchButton,
		scrollContainer,
		nextButton,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 1000))
	w.ShowAndRun()
}
