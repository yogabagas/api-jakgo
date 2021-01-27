package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-github/api-jakgo/domain/model"
	"net/http"
)

func (j *JakGoImpl) JakGoRSU(ctx context.Context) (rsu *model.RSUResponse, err error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, j.urlRSU, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", j.auth)

	res, err := client.Do(req)
	if err != nil {
		return rsu, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return rsu, err
	}

	json.Unmarshal([]byte(body), &rsu)

	return
}
