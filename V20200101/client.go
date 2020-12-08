// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package V20200101

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetVersionRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVersionRequest) GoString() string {
	return s.String()
}

func (s *GetVersionRequest) SetName(v string) *GetVersionRequest {
	s.Name = &v
	return s
}

type GetVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetVersionResponseBody) SetRequestId(v string) *GetVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetVersionResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVersionResponse) GoString() string {
	return s.String()
}

func (s *GetVersionResponse) SetHeaders(v map[string]*string) *GetVersionResponse {
	s.Headers = v
	return s
}

func (s *GetVersionResponse) SetBody(v *GetVersionResponseBody) *GetVersionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("yanhjtest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVersionWithOptions(request *GetVersionRequest, runtime *util.RuntimeOptions) (_result *GetVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVersion"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVersion(request *GetVersionRequest) (_result *GetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVersionResponse{}
	_body, _err := client.GetVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
