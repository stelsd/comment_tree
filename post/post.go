package post

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Post struct {
	ID      int    `json:"id"`
	Body    string `json:"body"`
	Replies []Post `json:"replies"`
}

func (p *Post) Apply(f func(p *Post)) {
	f(p)
	for i := range p.Replies {
		p.Replies[i].Apply(f)
	}
}

func loadBody(id int) string {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	var url = "https://25.ms/posts/" + strconv.Itoa(id)
	resp, err := httpClient.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	restBody, _ := io.ReadAll(resp.Body)
	var post Post
	json.Unmarshal(restBody, &post)
	return post.Body
}

func (p *Post) FillBodies() {
	bodies := Bodies{make(map[int]string), &sync.RWMutex{}}

	p.Apply(func(p *Post) {
		bodies.SetValue(p.ID, "")
	})

	wg := &sync.WaitGroup{}
	ids := bodies.GetIds()

	wg.Add(len(ids))
	for _, id := range ids {
		go func(id int) {
			defer wg.Done()
			body := loadBody(id)
			bodies.SetValue(id, body)
		}(id)
	}
	wg.Wait()

	p.Apply(func(p *Post) {
		p.Body = bodies.GetValue(p.ID)
	})
}
