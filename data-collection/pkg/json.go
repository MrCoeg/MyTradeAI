package pkg

import (
	"data-collection/domain"
	"encoding/json"
	"io"
)

func GetJSONFromUrl(container *domain.CryptoCoins, url string) {

	res := CreateConnection(url)
	defer res.Close()
	body, err := io.ReadAll(res)
	LogIfError(err)
	err = json.Unmarshal(body, container)
	LogIfError(err)

}
