package store

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testStoreService = &StoreService{}

func init() {
	testStoreService = InitStore()
}

func Test_store_init(t *testing.T) {
	assert.NotNil(t, testStoreService.redisClient)
}

func Test_save_and_get(t *testing.T) {
	shortUrl := "short"
	longUrl := "long"
	SaveUrlMapping(shortUrl, longUrl)
	assert.Equal(t, longUrl, GetLongUrl(shortUrl))
}

