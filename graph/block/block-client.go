package block

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mcdexio/mai3-trade-mining-watcher/common/logging"
	utils "github.com/mcdexio/mai3-trade-mining-watcher/utils/http"
	"strconv"
)

type Interface interface {
	GetBlockNumberWithTS(timestamp int64) (int64, error)
}

type Client struct {
	logger logging.Logger
	client *utils.Client
}

func NewClient(logger logging.Logger, url string) *Client {
	logger.Info("New block graph client with url %s", url)
	return &Client{
		logger: logger,
		client: utils.NewHttpClient(utils.DefaultTransport, logger, url),
	}
}

type Block struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	Timestamp string `json:"timestamp"`
}

// GetBlockNumberWithTS which is the closest but less than or equal to timestamp
func (b *Client) GetBlockNumberWithTS(timestamp int64) (int64, error) {
	b.logger.Debug("get block number which is the closest but <= @ts:%d", timestamp)
	query := `{
		blocks(
			first:1, orderBy: number, orderDirection: asc, 
			where: {timestamp_gt: %d}
		) {
			id
			number
			timestamp
		}
	}`
	var response struct {
		Data struct {
			Blocks []*Block
		}
	}
	// return err when can't get block number in three times
	if err := b.queryGraph(&response, query, timestamp); err != nil {
		return -1, err
	}

	if len(response.Data.Blocks) != 1 {
		return -1, fmt.Errorf("length of block response: expect=1, actual=%v, timestamp=%v",
			len(response.Data.Blocks), timestamp)
	}
	bn := response.Data.Blocks[0].Number
	number, err := strconv.Atoi(bn)
	if err != nil {
		return -1, fmt.Errorf("fail to get block number %s from string err=%s", bn, err)
	}
	return int64(number - 1), nil
}

// queryGraph return err if failed to get response from graph in three times
func (b *Client) queryGraph(resp interface{}, query string, args ...interface{}) error {
	var params struct {
		Query string `json:"query"`
	}
	params.Query = fmt.Sprintf(query, args...)
	for i := 0; i < 3; i++ {
		err, code, res := b.client.Post(nil, params, nil)
		if err != nil {
			b.logger.Error("fail to post http params=%+v err=%s", params, err)
			continue
		} else if code/100 != 2 {
			b.logger.Error("unexpected http response=%v", code)
			continue
		}
		err = json.Unmarshal(res, &resp)
		if err != nil {
			b.logger.Error("fail to unmarshal err=%s", err)
			continue
		}
		// success
		return nil
	}
	return errors.New("fail to query block graph in three times")
}