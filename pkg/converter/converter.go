package converter

import (
	"csv/pkg/company"
	"encoding/base64"
	"fmt"

	"github.com/jszwec/csvutil"
)

type CSVConverter struct {
}

func (csvConverter *CSVConverter) Convert(base64Doc string) ([]company.Company, error) {
	var companies []company.Company
	decoded, _ := base64.StdEncoding.DecodeString(base64Doc)

	if err := csvutil.Unmarshal(decoded, &companies); err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return companies, nil
}
