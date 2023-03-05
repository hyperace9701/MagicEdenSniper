package sniper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hyperace9701/magicedensniper/internal/models"
)

func GetNFTData(tokenMint string) (*models.NFT, error) {

	request, _ := http.NewRequest("GET", fmt.Sprintf("https://api-mainnet.magiceden.dev/v2/tokens/%s", tokenMint), nil)
	resp, err := cli.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var nftResp models.NFT
	json.Unmarshal(body, &nftResp)

	return &nftResp, nil
}
