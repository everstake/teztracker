package ipfs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type NuncNFTDescription struct {
	Name               string
	Description        string
	Tags               []string
	Category           string `json:"category"`
	Symbol             string
	ArtifactUri        string
	DisplayUri         string
	ThumbnailUri       string
	Creators           []string
	Formats            []Format
	Decimals           int64
	IsBooleanAmount    bool
	ShouldPreferSymbol bool
}

type Format struct {
	Uri      string `json:"uri"`
	MimeType string `json:"mimeType"`
}

type ipfs struct {
	url url.URL
	cl  *http.Client
}

func NewIPFSClient(ipfsURL string) (*ipfs, error) {

	url, err := url.Parse(ipfsURL)
	if err != nil {
		return nil, err
	}

	return &ipfs{url: *url, cl: http.DefaultClient}, nil
}

func (ip ipfs) GetIPFSMetadata(ipfsID string) (desc NuncNFTDescription, err error) {

	reqURL := ip.url
	reqURL.Path = fmt.Sprintf("%s/%s", reqURL.Path, ipfsID)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return desc, err
	}

	resp, err := ip.cl.Do(req)
	if err != nil {
		return desc, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if resp.StatusCode == http.StatusNoContent {
		return desc, nil
	}

	if resp.StatusCode != http.StatusOK {
		return desc, fmt.Errorf("Resp with error: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&desc)
	if err != nil {
		return desc, err
	}

	return desc, nil
}
