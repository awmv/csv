package converter

import (
	"csv/pkg/config"
	"testing"
)

func TestConvert(t *testing.T) {
	converter := &CSVConverter{}
	_, err := converter.Convert(config.Config.Base64)
	if err != nil {
		t.Error("Failed to convert: ", err)
	}
}
