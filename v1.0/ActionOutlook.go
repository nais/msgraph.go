// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nais/msgraph.go/jsonx"
)

// MasterCategories returns request builder for OutlookCategory collection
func (b *OutlookUserRequestBuilder) MasterCategories() *OutlookUserMasterCategoriesCollectionRequestBuilder {
	bb := &OutlookUserMasterCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/masterCategories"
	return bb
}

// OutlookUserMasterCategoriesCollectionRequestBuilder is request builder for OutlookCategory collection
type OutlookUserMasterCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OutlookCategory collection
func (b *OutlookUserMasterCategoriesCollectionRequestBuilder) Request() *OutlookUserMasterCategoriesCollectionRequest {
	return &OutlookUserMasterCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OutlookCategory item
func (b *OutlookUserMasterCategoriesCollectionRequestBuilder) ID(id string) *OutlookCategoryRequestBuilder {
	bb := &OutlookCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutlookUserMasterCategoriesCollectionRequest is request for OutlookCategory collection
type OutlookUserMasterCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OutlookCategory, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OutlookCategory
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []OutlookCategory
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for OutlookCategory collection, max N pages
func (r *OutlookUserMasterCategoriesCollectionRequest) GetN(ctx context.Context, n int) ([]OutlookCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Get(ctx context.Context) ([]OutlookCategory, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Add(ctx context.Context, reqObj *OutlookCategory) (resObj *OutlookCategory, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
