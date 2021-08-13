package concurrency_api_patterns

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
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
	// getResult will stop immediately if the http.Request is canceled
	result := Result{
		First:  &Response{},
		Second: &Response{},
		Third:  &Response{},
	}
	ctx := context.Background()
	if err := <-getResult(ctx, &result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
	fmt.Println(result)
}

// getResult returns the result of many concurrent API calls
func getResult(ctx context.Context, result *Result) <-chan error {
	out := make(chan error)

	go func() {
		// Correct memory management
		defer close(out)

		// The cancel func will allow us to stop all pending requests if one
		// fails
		ctx, cancel := context.WithCancel(ctx)

		// Merge allows us to recieve the all of errors returned from all of
		// the calls to `getPieces` in a single `<-chan error`.
		// If no errors are returned, Merge will wait until all of the
		// `<-chan error`s close before proceeding
		for err := range merge(
			getResponse(ctx, 1, result.First),
			getResponse(ctx, 2, result.Second),
			getResponse(ctx, 3, result.Third),
		) {
			if err != nil {

				// Cancel all pending requests
				cancel()

				// Surface the error to the caller
				out <- err
				return
			}
		}
	}()
	return out
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

		buf := new(strings.Builder)
		defer resp.Body.Close()
		if _, err := io.Copy(buf, resp.Body); err != nil {
			out <- err
			return
		}

		err = json.Unmarshal([]byte(buf.String()), res)
		if err != nil {
			out <- err
			return
		}
	}()
	return out
}

// Merge fans multiple error channels in to a single error channel
func merge(errChans ...<-chan error) <-chan error {
	mergedChan := make(chan error)

	// Create a WaitGroup that waits for all of the errChans to close
	var wg sync.WaitGroup
	wg.Add(len(errChans))
	go func() {
		// When all of the errChans are closed, close the mergedChan
		wg.Wait()
		close(mergedChan)
	}()

	for i := range errChans {
		go func(errChan <-chan error) {
			// Wait for each errChan to close
			for err := range errChan {
				if err != nil {
					// Fan the contents of each errChan into the mergedChan
					mergedChan <- err
				}
			}
			// Tell the WaitGroup that one of the errChans is closed
			wg.Done()
		}(errChans[i])
	}

	return mergedChan
}
