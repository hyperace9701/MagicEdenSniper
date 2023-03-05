package sniper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/hyperace9701/magicedensniper/internal/models"
)

const storageTime = time.Duration(time.Second * 30)

var (
	mu    sync.Mutex
	cache = make(map[string]*models.Floor)
	cli   = &http.Client{}
)

func GetFloor(symbol string) (float64, error) {
	if symbol == "" {
		return 0, fmt.Errorf("empty string symbol")
	}

	if _, ok := cache[symbol]; ok && time.Since(cache[symbol].Time) < storageTime {

	} else {
		request, _ := http.NewRequest("GET", fmt.Sprintf("https://api-mainnet.magiceden.dev/v2/collections/%s/stats", symbol), nil)
		resp, err := cli.Do(request)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return 0, fmt.Errorf("Status code: %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		var floorResp models.FloorResponse
		json.Unmarshal(body, &floorResp)

		mu.Lock()
		cache[symbol] = &models.Floor{Value: floorResp.FloorPrice / 1e9, Time: time.Now()}
		mu.Unlock()
	}

	return cache[symbol].Value, nil
}
