package mempool

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

func (m Mempool) Do(ctx context.Context, method, url, params string, v interface{}) (err error) {
	reqURL := m.url
	reqURL.Path = url
	reqURL.RawQuery = params

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := m.cl.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Resp with error: %d", resp.StatusCode)
	}

	err = handleBatchResponse(req.Context(), resp, v)
	if err != nil {
		return err
	}

	return nil
}

func handleBatchResponse(ctx context.Context, resp *http.Response, v interface{}) error {

	typ := reflect.TypeOf(v)
	dec := json.NewDecoder(resp.Body)

	cases := []reflect.SelectCase{
		reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(v),
		},
		reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ctx.Done()),
		},
	}

	for {
		chunkVal := reflect.New(typ.Elem())

		if err := dec.Decode(chunkVal.Interface()); err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				// Tezos doesn't output the trailing zero lenght chunk leading to io.ErrUnexpectedEOF
				break
			}
			return err
		}

		cases[0].Send = chunkVal.Elem()
		if chosen, _, _ := reflect.Select(cases); chosen == 1 {
			return ctx.Err()
		}
	}

	return nil
}
