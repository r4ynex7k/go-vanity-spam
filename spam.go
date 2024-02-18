package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type s struct {
	client *http.Client
	m      map[string]*rs7k
	queue  map[string]string
}

func NewS() *s {
	return &s{
		client: &http.Client{},
		m:      make(map[string]*rs7k),
		queue:  make(map[string]string),
	}
}

func (s *s) UB() {
	const gld = "1201647687582367775"  // server id
	v := []string{"denenemqweqweqwe"}  // url

	tknBytes, err := ioutil.ReadFile("tkn.txt")
	if err != nil {
		fmt.Println("Error reading token file:", err)
		return
	}
	tkn := strings.Split(string(tknBytes), "\n")

	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		for _, vanity := range v {
			s.uu(vanity, gld, tkn)
		}
	}
}

func (s *s) Get(id string) *rs7k {
	k, exists := s.m[id]
	if !exists {
		k = NewRS7k(id)
		s.m[id] = k
	}
	return k
}

func (s *s) uu(v string, gld string, tkn []string) {
	randomIndex := rand.Intn(len(tkn))
	authorizationToken := strings.TrimSpace(tkn[randomIndex])

	url := fmt.Sprintf("https://discord.com/api/v7/guilds/%s/vanity-url", gld)
	data := map[string]string{"code": v}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", authorizationToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(v)

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("[DURUM] %s ALİNDİ\n", v)
		os.Exit(0)
	}
}

type rs7k struct {
	client *http.Client
	id     string
}

func NewRS7k(id string) *rs7k {
	return &rs7k{
		client: &http.Client{},
		id:     id,
	}
}

func main() {
	bot := NewS()
	bot.UB()
}
