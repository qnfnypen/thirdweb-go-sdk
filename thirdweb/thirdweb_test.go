package thirdweb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

var (
	ipfsStore *IpfsStorage
	ctx       = context.Background()
)

func TestMain(m *testing.M) {
	secret := ""
	chainID := deriveClientId(secret)
	gatewayURL := fmt.Sprintf("https://%s.ipfscdn.io/ipfs/", chainID)

	ipfsStore = newIpfsStorage(secret, gatewayURL, http.DefaultClient)

	m.Run()
}

func TestUploadJSON(t *testing.T) {
	maps := map[string]interface{}{
		"test": 123,
	}

	body, _ := json.Marshal(&maps)
	uri, err := ipfsStore.UploadJSON(ctx, body, "11dd50a6-fb1f-435a-a86a-bb0a8b991689", "", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(uri)
}
