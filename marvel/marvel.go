package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const (
	endpoint = "https://gateway.marvel.com/v1/public"
)

type Setup struct {
	publickKey string
	privateKey string
	timeout    time.Duration
}

func New(publicKey, privateKey string, timeout time.Duration) Setup {
	return Setup{
		publickKey: publicKey,
		privateKey: privateKey,
		timeout:    timeout,
	}
}

func makeURLQuery(s Setup) string {
	ts := time.Now().UTC().Unix()
	query := fmt.Sprintf("%d%s%s", ts, s.privateKey, s.publickKey)
	hash := md5.Sum([]byte(query))
	hashToString := hex.EncodeToString(hash[:])

	return fmt.Sprintf("?ts=%d&apikey=%s&hash=%s", ts, s.publickKey, hashToString)
}

func (s Setup) GetComicsList() ([]Comic, error) {
	currentEndpoint := fmt.Sprintf("%s/comics%s", endpoint, makeURLQuery(s))

	c, err := NewClient(s.timeout)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Get(currentEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var comic ComicDataWrapper
	if err = json.Unmarshal(body, &comic); err != nil {
		return nil, err
	}

	return comic.Data.Results, nil
}

func (s Setup) GetComic(comicId int) ([]Comic, error) {
	currentEndpoint := fmt.Sprintf("%s/comics/%d%s", endpoint, comicId, makeURLQuery(s))

	c, err := NewClient(s.timeout)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Get(currentEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var comic ComicDataWrapper
	if err = json.Unmarshal(body, &comic); err != nil {
		return nil, err
	}

	return comic.Data.Results, nil
}
