package pkg

import (
	"encoding/json"
	"io"
	"trading-ai/domain"
)

func GetJSONFromUrl(container *domain.CryptoCoins, url string) {

	res := CreateConnection(url)
	defer res.Close()
	body, err := io.ReadAll(res)
	LogIfError(err)
	err = json.Unmarshal(body, container)
	LogIfError(err)

}
