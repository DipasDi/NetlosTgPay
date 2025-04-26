package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"telegarm/config"
)

func CreatePayURL(name string, productID int) config.JsonResponseEasyDonate {
	var DonatCreateURL config.JsonResponseEasyDonate

	product := map[int]int{
		productID: 1,
	}

	body, _ := json.Marshal(product)

	url := fmt.Sprintf(`https://easydonate.ru/api/v3/shop/payment/create?customer=%v&email=Discord@gmail.com&server_id=17277&success_url=https://netlos.ru/&products=`+string(body), name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Shop-Key", config.TokenPayEasyDonate)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(err.Error())
	}

	decorder := json.NewDecoder(bytes.NewReader(body))
	err = decorder.Decode(&DonatCreateURL)
	if err != nil {
		log.Println(err)
	}
	// Выводим ответ от сервера
	fmt.Println(DonatCreateURL)

	return DonatCreateURL
}
