// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nais/msgraph.go/jsonx"
)

// ThreatAssessmentRequests returns request builder for ThreatAssessmentRequestObject collection
func (b *InformationProtectionRequestBuilder) ThreatAssessmentRequests() *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder {
	bb := &InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/threatAssessmentRequests"
	return bb
}

// InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder is request builder for ThreatAssessmentRequestObject collection
type InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ThreatAssessmentRequestObject collection
func (b *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder) Request() *InformationProtectionThreatAssessmentRequestsCollectionRequest {
	return &InformationProtectionThreatAssessmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ThreatAssessmentRequestObject item
func (b *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder) ID(id string) *ThreatAssessmentRequestObjectRequestBuilder {
	bb := &ThreatAssessmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InformationProtectionThreatAssessmentRequestsCollectionRequest is request for ThreatAssessmentRequestObject collection
type InformationProtectionThreatAssessmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ThreatAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ThreatAssessmentRequestObject, error) {
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
	var values []ThreatAssessmentRequestObject
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
			value  []ThreatAssessmentRequestObject
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

// GetN performs GET request for ThreatAssessmentRequestObject collection, max N pages
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]ThreatAssessmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ThreatAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) Get(ctx context.Context) ([]ThreatAssessmentRequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ThreatAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *ThreatAssessmentRequestObject) (resObj *ThreatAssessmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
