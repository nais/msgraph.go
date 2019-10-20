// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationDeviceStateSummaryRequestBuilder is request builder for DeviceConfigurationDeviceStateSummary
type DeviceConfigurationDeviceStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationDeviceStateSummaryRequest
func (b *DeviceConfigurationDeviceStateSummaryRequestBuilder) Request() *DeviceConfigurationDeviceStateSummaryRequest {
	return &DeviceConfigurationDeviceStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationDeviceStateSummaryRequest is request for DeviceConfigurationDeviceStateSummary
type DeviceConfigurationDeviceStateSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceConfigurationDeviceStateSummary
func (r *DeviceConfigurationDeviceStateSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceConfigurationDeviceStateSummary, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceConfigurationDeviceStateSummary
func (r *DeviceConfigurationDeviceStateSummaryRequest) Get() (*DeviceConfigurationDeviceStateSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceConfigurationDeviceStateSummary
func (r *DeviceConfigurationDeviceStateSummaryRequest) Update(reqObj *DeviceConfigurationDeviceStateSummary) (*DeviceConfigurationDeviceStateSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceConfigurationDeviceStateSummary
func (r *DeviceConfigurationDeviceStateSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}