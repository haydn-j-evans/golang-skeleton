package featureflags

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type featureFlags struct {
	FeatureFlagMap map[string]featureFlagModel
	mutex          sync.RWMutex
}

type featureFlagModel struct {
	code            string `json:"code"`
	enabled         bool   `json:"enabled"`
	frontendRelated bool   `json:"frontendRelated"`
}

func InitFeatureFLags() *featureFlags {

	flags := &featureFlags{}

	newMap := make(map[string]featureFlagModel)

	flags.FeatureFlagMap = newMap

	return flags
}

func WatchFeatureFlags(featureflags *featureFlags) {

	url := "https://stg-storefront-cache-plus.pub.az-we.dglecom.net/dglwebservices/v2/dgl-DE/featureflags"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("blah")
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	featureflags.mutex.Lock()

	json.Unmarshal(body, &featureflags)

	featureflags.mutex.Unlock()

	fmt.Println(featureflags)

}
