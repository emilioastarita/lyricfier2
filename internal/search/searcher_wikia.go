package search

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

var wikiaReExtractSongUrl = regexp.MustCompile(`'url':'(https?:\/\/lyrics.wikia.com\/.+)'`)

func wikiaExtractSong(body string) string {
	res := wikiaReExtractSongUrl.FindStringSubmatch(body)
	if len(res) != 2 {
		return ""
	}
	return res[1]
}

func wikiaSearchSongUrl(artist string, title string) (string, error) {
	v := url.Values{}
	v.Add("action", "lyrics")
	v.Add("artist", artist)
	v.Add("song", title)
	v.Add("fmt", "json")
	v.Add("func", "getSong")
	var searchUrl = "http://lyrics.wikia.com/api.php?" + v.Encode()
	fmt.Println("Lookup", searchUrl)
	response, err := client.Get(searchUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	urlSong := wikiaExtractSong(string(responseData))

	if urlSong == "" {
		fmt.Println(string(responseData))
		return "", errors.New("not found")
	}

	return urlSong, nil
}

func Wikia(artist string, title string) (string, error) {
	songUrl, err := wikiaSearchSongUrl(artist, title)
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
	doc.Find(".lyricbox").Each(func(i int, s *goquery.Selection) {
		songUrl = s.Find(".lyrics-spotify").Text()
		s1, _ := s.Html()
		re1 := regexp.MustCompile(`(<br>)|(<br ?\/>)`)
		placeholder := "!NEWLINE!"
		re2 := regexp.MustCompile(placeholder)
		d, _ := goquery.NewDocumentFromReader(strings.NewReader(re1.ReplaceAllString(s1, placeholder)))
		lyric = re2.ReplaceAllString(d.Text(), "\n")
	})
	if lyric == "" {
		return "", errors.New("not found")
	}
	return lyric, nil
}
