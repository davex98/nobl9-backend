package infrastracture

import (
	"context"
	"errors"
	"fmt"
	"github.com/davex98/nobl9-backend/random-generator/application"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type service struct {
	client                http.Client
	maxConcurrentRequests int
}

func NewRandomService(client http.Client) application.RandomService {
	requests := os.Getenv("CONCURRENT_REQUESTS")
	requestsNumber, err := strconv.Atoi(requests)
	if err != nil {
		panic("invalid CONCURRENT_REQUESTS")
	}
	srv := service{client: client, maxConcurrentRequests: requestsNumber}

	return srv
}

func (s service) GetRandomNumbers(ctx context.Context, length, requestsNumber int) ([]application.RandomNumbersResponse, error) {
	var randomNumbers []application.RandomNumbersResponse
	var result *multierror.Error
	sem := semaphore.NewWeighted(int64(s.maxConcurrentRequests))
	mutex := &sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 1; i <= requestsNumber; i++ {
		wg.Add(1)
		if err := sem.Acquire(ctx, 1); err != nil {
			mutex.Lock()
			result = multierror.Append(result, err)
			mutex.Unlock()
		}

		go func() {
			url, err := s.getRandomNumber(ctx, length)
			mutex.Lock()
			defer mutex.Unlock()
			result = multierror.Append(result, err)
			randomNumbers = append(randomNumbers, url)
			wg.Done()
			sem.Release(1)
		}()
	}
	wg.Wait()
	if err := result.Unwrap(); err != nil {
		return nil, err
	}

	return randomNumbers, nil
}

func (s service) getRandomNumber(ctx context.Context, length int) (application.RandomNumbersResponse, error) {
	url := fmt.Sprintf("https://www.random.org/integers/?num=%v&min=1&max=10&format=plain&col=1&base=10", length)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := s.client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		statusMsg := http.StatusText(resp.StatusCode)
		msg := fmt.Sprintf("could not get data from www.random.org, status code: %v: %s", resp.StatusCode, statusMsg)
		fetchError := errors.New(msg)
		return application.RandomNumbersResponse{}, fetchError
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return application.RandomNumbersResponse{}, err
	}
	bodyString := string(bodyBytes)
	numbers := strings.Fields(bodyString)

	toNumbers, err := StringSliceToNumbers(numbers)
	if err != nil {
		return application.RandomNumbersResponse{}, err
	}

	return application.RandomNumbersResponse{
		Data: toNumbers,
	}, nil
}

func StringSliceToNumbers(input []string) ([]int, error) {
	var output []int
	for _, i := range input {
		if i == "" {
			continue
		}
		parseInt, err := strconv.ParseInt(i, 10, 32)
		if err != nil {
			return nil, err
		}
		output = append(output, int(parseInt))
	}
	return output, nil
}
