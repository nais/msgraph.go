// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AgreementFileRequestBuilder is request builder for AgreementFile
type AgreementFileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgreementFileRequest
func (b *AgreementFileRequestBuilder) Request() *AgreementFileRequest {
	return &AgreementFileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgreementFileRequest is request for AgreementFile
type AgreementFileRequest struct{ BaseRequest }

// Do performs HTTP request for AgreementFile
func (r *AgreementFileRequest) Do(method, path string, reqObj interface{}) (resObj *AgreementFile, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AgreementFile
func (r *AgreementFileRequest) Get() (*AgreementFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AgreementFile
func (r *AgreementFileRequest) Update(reqObj *AgreementFile) (*AgreementFile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AgreementFile
func (r *AgreementFileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}