package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

// jsonPage represents one page of a paginated result.
type jsonPage struct {
	Count    int
	Previous string
	Next     string
	Results  []struct {
		ID   int
		Name string
	}
}

// result collects the results for all the pages of one repo.
type result struct {
	Count int
	Found int
	Tags  []string
}

var (
	baseURL = "https://registry.hub.docker.com/v2/repositories/%s/tags/"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n    %s <repo> [repo]...\n", os.Args[0])
		return
	}
	repos := os.Args[1:]

	for _, r := range repos {
		t, err := explore(r)
		if err != nil {
			fmt.Printf("%s:\n    Error: %s\n\n", r, err)
		} else {
			sort.Strings(t.Tags)
			fmt.Printf("%s (%d/%d):\n    %s\n\n", r, t.Found, t.Count, strings.Join(t.Tags, " "))
		}
	}
}

func explore(repo string) (*result, error) {
	if !strings.Contains(repo, "/") {
		repo = "library/" + repo
	}
	nextURL := fmt.Sprintf(baseURL, strings.ToLower(repo))
	t := &result{}

	for nextURL != "" {
		resp, err := http.Get(nextURL)
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()

		p := jsonPage{}
		err = json.Unmarshal(data, &p)
		if err != nil {
			return nil, err
		}
		if p.Count > t.Count {
			t.Count = p.Count
		}

		for _, v := range p.Results {
			t.Tags = append(t.Tags, v.Name)
			t.Found++
		}

		nextURL = p.Next
	}

	return t, nil
}
