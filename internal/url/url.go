package url

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"

type URL struct {
	ShortURL string `json:"shorturl"`
	LongURL  string `json:"longurl"`
}
