package placeholder

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

type Placeholder struct {
	baseURL *url.URL
	client  *http.Client
}

func NewToDos(path string) (*Placeholder, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load todos - invalid path")
	}

	t := &Placeholder{
		baseURL: u,
		client:  &http.Client{},
	}

	return t, nil
}

func (p *Placeholder) GetTodos(ctx context.Context, todos [][]string) ([]*TODO, error) {
	tmp := map[int]*TODO{}
	keys := make([]int, 0)

	for key, row := range todos {
		if key == 0 {
			if _, err := strconv.Atoi(row[0]); err != nil {
				continue
			}
		}

		result, err := p.get(ctx, row[0])
		if err != nil {
			return nil, err
		}
		if !result.IsValid() {
			log.Printf("invalid id request:%v\n", row[0])
			continue
		}

		result.Sanitise()
		tmp[result.ID] = result
		keys = append(keys, result.ID)
	}

	sort.Ints(keys)
	arr := make([]*TODO, 0)
	for _, val := range keys {
		arr = append(arr, tmp[val])
	}

	return arr, nil
}

func (p *Placeholder) get(ctx context.Context, id string) (*TODO, error) {
	u := p.baseURL.String() + fmt.Sprintf("/%s", id)
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create new request")
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to make request")
	}

	defer resp.Body.Close()

	todo, err := ToDoFromJSON(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to map response body to struct")
	}

	return todo, nil
}
