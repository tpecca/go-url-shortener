package models

type LinkEntry struct {
	OriginalURL string `bson:"original_url"`
	Hash        string `bson:"hash"`
}
