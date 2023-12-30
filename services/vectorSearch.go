package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/cast"
)

func GetVectorSearch(productName string) (arr []float64) {
	baseURL := "http://localhost:8000/v1/getVector"
	queryParams := url.Values{}
	queryParams.Set("productName", cast.ToString(productName))

	urlWithParams := baseURL + "?" + queryParams.Encode()

	response, err := http.Get(urlWithParams)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		panic(err)
	}
	embeddings, ok := responseBody["product_embeddings"].([]interface{})
	if !ok {
		panic("Unable to extract product_embeddings")
	}

	for _, val := range embeddings {
		arr = append(arr, cast.ToFloat64(val))
	}
	return arr
}
