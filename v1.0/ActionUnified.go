// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nais/msgraph.go/jsonx"
)

// AppScope is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) AppScope() *AppScopeRequestBuilder {
	bb := &AppScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appScope"
	return bb
}

// DirectoryScope is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) DirectoryScope() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryScope"
	return bb
}

// Principal is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// InheritsPermissionsFrom returns request builder for UnifiedRoleDefinition collection
func (b *UnifiedRoleDefinitionRequestBuilder) InheritsPermissionsFrom() *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder {
	bb := &UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/inheritsPermissionsFrom"
	return bb
}

// UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder is request builder for UnifiedRoleDefinition collection
type UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UnifiedRoleDefinition collection
func (b *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder) Request() *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest {
	return &UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UnifiedRoleDefinition item
func (b *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequestBuilder) ID(id string) *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest is request for UnifiedRoleDefinition collection
type UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UnifiedRoleDefinition collection
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UnifiedRoleDefinition, error) {
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
	var values []UnifiedRoleDefinition
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
			value  []UnifiedRoleDefinition
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

// GetN performs GET request for UnifiedRoleDefinition collection, max N pages
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) GetN(ctx context.Context, n int) ([]UnifiedRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UnifiedRoleDefinition collection
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) Get(ctx context.Context) ([]UnifiedRoleDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UnifiedRoleDefinition collection
func (r *UnifiedRoleDefinitionInheritsPermissionsFromCollectionRequest) Add(ctx context.Context, reqObj *UnifiedRoleDefinition) (resObj *UnifiedRoleDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
