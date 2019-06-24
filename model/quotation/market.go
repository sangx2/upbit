package quotation

import (
	"encoding/json"
	"io"
)

// Market 마켓 정보
type Market struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}

func MarketsFromJSON(r io.Reader) []*Market {
	var markets []*Market

	json.NewDecoder(r).Decode(&markets)

	return markets
}
