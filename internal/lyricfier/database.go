package lyricfier

import (
	"errors"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"os"
	"path/filepath"
	"strings"
)

var SongsBucket = []byte("songsBucket")
var GeneralBucket = []byte("general")
var SettingsKey = "settings"

func between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func SongKey(artist string, title string) string {
	return "__artist__" + artist + "__title__" + title
}

func SongKeyGetArtist(key []byte) string {
	stringKey := string(key)
	return between(stringKey, "__artist__", "__title__")
}

func SongKeyGetTitle(key []byte) string {
	stringKey := string(key)
	return after(stringKey, "__title__")
}

func Open() (*bolt.DB, error) {
	dir := GetDbPath()
	dbPath := filepath.Join(dir, "lyricfier.bbolt")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return bolt.Open(dbPath, os.ModePerm, nil)
}

func Read(bucket []byte, key string) ([]byte, error) {
	var db, err = Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var val []byte
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket)
		if bucket == nil {
			return errors.New("BUCKET_NOT_CREATED")
		}
		val = bucket.Get([]byte(key))
		return nil
	})
	if err != nil {
		fmt.Println("err_read", err)
		return nil, err
	}
	res := make([]byte, len(val), cap(val)+1)
	copy(res, val)
	return res, err
}

func Write(bucket []byte, key string, value []byte) error {
	var db, err = Open()
	if err != nil {
		fmt.Println("err_write", err)
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(key), value)
	})
	if err != nil {
		fmt.Println("err_write", err)
	}
	return err
}

func ListSongs(songs *[]*SongItem) error {
	var db, err = Open()
	if err != nil {
		fmt.Println("err_listing_songs", err)
		return err
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(SongsBucket)
		if bucket == nil {
			return errors.New("BUCKET_NOT_CREATED")
		}
		c := bucket.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			s := new(SongItem)
			s.Artist = SongKeyGetArtist(k)
			s.Title = SongKeyGetTitle(k)
			if s.Title != "" && s.Artist != "" {
				*songs = append(*songs, s)
			}
		}
		return nil
	})
	return err
}
