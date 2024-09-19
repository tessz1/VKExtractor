package vkapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	AccessToken string
	Version     string
}

func NewClient(token string) *Client {
	return &Client{
		AccessToken: token,
		Version:     "5.199",
	}
}

func (c *Client) GetPosts(ownerID string, count int) ([]Post, error) {
	apiURL := "https://api.vk.com/method/wall.get"
	params := url.Values{}
	params.Set("owner_id", ownerID)
	params.Set("count", fmt.Sprintf("%d", count))
	params.Set("access_token", c.AccessToken)
	params.Set("v", c.Version)

	resp, err := http.Get(apiURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var vkResp VKResponse
	err = json.Unmarshal(body, &vkResp)
	if err != nil {
		return nil, err
	}
	return vkResp.Response.Items, nil
}

type VKResponse struct {
	Response struct {
		Count int    `json:"count"`
		Items []Post `json:"items"`
	} `json:"response"`
}

type Post struct {
	ID      int    `json:"id"`
	FromID  int    `json:"from_id"`
	OwnerID int    `json:"owner_id"`
	Date    int    `json:"date"`
	Text    string `json:"text"`
}
