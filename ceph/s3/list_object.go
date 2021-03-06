package s3

import (
	"encoding/xml"
	"io/ioutil"

	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

//ListObject ...
func (s *S3Config) ListObject() (*http.Response, error) {
	res, err := s.Send(http.MethodGet, fmt.Sprintf("%s", s.Bucket))
	if err != nil {
		return res, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		body := S3Error{}
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		xml.Unmarshal(bodyBytes, &body)

		return res, errors.New(body.Code)
	}
	return res, err
}
