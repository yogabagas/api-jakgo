package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-github/api-jakgo/domain/model"
	"net/http"
)

func (j *JakGoImpl) JakGoKelurahan(ctx context.Context) (kel *model.KelurahanResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, j.urlKel, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", j.auth)

	res, err := client.Do(req)
	if err != nil {
		return kel, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return kel, err
	}

	json.Unmarshal([]byte(body), &kel)

	return
}
