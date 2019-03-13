package search

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/url"
)

func geniusSearchUrl(artist string, title string) (string, error) {
	v := url.Values{}
	v.Add("action", "lyrics")
	v.Add("q", artist+" "+title)
	v.Add("per_page", "5")
	var searchUrl = "https://genius.com/api/search/multi?" + v.Encode()
	url := ""
	response, err := client.Get(searchUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	status := gjson.Get(string(all), "meta.status").Int()
	if status != 200 {
		fmt.Println("no status", status)
		return url, errors.New("not found")
	}
	res := gjson.Get(string(all), `response.sections.#[type=="song"].hits.0.result.url`)
	if !res.Exists() {
		return "", errors.New("not found")
	}
	return res.String(), nil
}

func Genius(artist string, title string) (string, error) {
	songUrl, err := geniusSearchUrl(artist, title)
	if err != nil {
		return "", err
	}
	response, err := client.Get(songUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}
	var lyric = ""
	doc.Find(".lyrics p").Each(func(i int, s *goquery.Selection) {
		lyric = s.Text()
	})
	if lyric == "" {
		return "", errors.New("not found")
	}
	return lyric, nil
}
