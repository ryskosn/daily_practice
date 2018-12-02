package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/upamune/go-esa/esa"
)

type Esa struct {
	Token    string `toml:"token"`
	Team     string `toml:"team"`
	Category string `toml:"category"`
}

type Config struct {
	// Slack Slack `toml:"slack"`
	Esa Esa `toml:"esa"`
}

func (c *Config) loadConfig(filename string) error {
	_, err := toml.DecodeFile(filename, c)
	return err
}

func (c *Config) postToEsa(message string) error {
	client := esa.NewClient(c.Esa.Token)
	post := esa.Post{}

	today := time.Now()
	todayStr := today.Format("2006/01/02")
	post.Name = c.Esa.Category + "/" + todayStr

	header := "foobar header"
	body := message

	post.BodyMd += header + body + "\n"

	_, err := client.Post.Create(c.Esa.Team, post)
	return err
}

func (c *Config) getPosts(category string) *esa.PostsResponse {
	client := esa.NewClient(c.Esa.Token)
	query := url.Values{}
	query.Add("in", category)
	res, err := client.Post.GetPosts("rrr", query)
	if err != nil {
		return nil
	}
	return res
}

func main() {
	config := Config{}
	configFile := "config.toml"
	err := config.loadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	// config.postToEsa("hello")
	res := config.getPosts("memo")
	fmt.Printf("%s\n", "foo")
	fmt.Println(res)
}
