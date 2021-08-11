package concurrency_api_patterns

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		Url  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

type Result struct {
	First  *Response `json:"first,omitempty"`
	Second *Response `json:"second,omitempty"`
	Third  *Response `json:"third,omitempty"`
}

func GetResult() {

}

func getResponse(ctx context.Context, id int, res *Response) <-chan error {
	out := make(chan error)
	go func() {
		defer close(out)

		req, err := http.NewRequestWithContext(
			ctx,
			"GET",
			fmt.Sprintf("https://reqres.in/api/users/%d", id),
			nil,
		)
		if err != nil {
			out <- err
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			out <- err
			return
		} else if resp.StatusCode != http.StatusOK {
			out <- fmt.Errorf("%d %s", resp.StatusCode, resp.Status)
		}

		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
			out <- err
			return
		}
	}()
	return out
}
