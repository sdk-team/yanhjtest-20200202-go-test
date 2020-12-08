// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package V20200202

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateApiResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApiResponseBody) SetRequestId(v string) *UpdateApiResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApiResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApiResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApiResponse) GoString() string {
	return s.String()
}

func (s *UpdateApiResponse) SetHeaders(v map[string]*string) *UpdateApiResponse {
	s.Headers = v
	return s
}

func (s *UpdateApiResponse) SetBody(v *UpdateApiResponseBody) *UpdateApiResponse {
	s.Body = v
	return s
}

type SendCommondResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendCommondResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCommondResponseBody) GoString() string {
	return s.String()
}

func (s *SendCommondResponseBody) SetRequestId(v string) *SendCommondResponseBody {
	s.RequestId = &v
	return s
}

type SendCommondResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendCommondResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCommondResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCommondResponse) GoString() string {
	return s.String()
}

func (s *SendCommondResponse) SetHeaders(v map[string]*string) *SendCommondResponse {
	s.Headers = v
	return s
}

func (s *SendCommondResponse) SetBody(v *SendCommondResponseBody) *SendCommondResponse {
	s.Body = v
	return s
}

type GetAllProductRequest struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// pop产品
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s GetAllProductRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllProductRequest) GoString() string {
	return s.String()
}

func (s *GetAllProductRequest) SetRequestId(v string) *GetAllProductRequest {
	s.RequestId = &v
	return s
}

func (s *GetAllProductRequest) SetProduct(v string) *GetAllProductRequest {
	s.Product = &v
	return s
}

func (s *GetAllProductRequest) SetEnv(v string) *GetAllProductRequest {
	s.Env = &v
	return s
}

type GetAllProductResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 产品信息
	Data *GetAllProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllProductResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllProductResponseBody) SetCode(v string) *GetAllProductResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllProductResponseBody) SetData(v *GetAllProductResponseBodyData) *GetAllProductResponseBody {
	s.Data = v
	return s
}

func (s *GetAllProductResponseBody) SetRequestId(v string) *GetAllProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllProductResponseBody) SetSuccess(v bool) *GetAllProductResponseBody {
	s.Success = &v
	return s
}

type GetAllProductResponseBodyData struct {
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 域名
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// nameSpace
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// product
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAllProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllProductResponseBodyData) SetDescription(v string) *GetAllProductResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAllProductResponseBodyData) SetDomains(v []*string) *GetAllProductResponseBodyData {
	s.Domains = v
	return s
}

func (s *GetAllProductResponseBodyData) SetNameSpace(v string) *GetAllProductResponseBodyData {
	s.NameSpace = &v
	return s
}

func (s *GetAllProductResponseBodyData) SetProduct(v string) *GetAllProductResponseBodyData {
	s.Product = &v
	return s
}

func (s *GetAllProductResponseBodyData) SetType(v string) *GetAllProductResponseBodyData {
	s.Type = &v
	return s
}

type GetAllProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAllProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllProductResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllProductResponse) GoString() string {
	return s.String()
}

func (s *GetAllProductResponse) SetHeaders(v map[string]*string) *GetAllProductResponse {
	s.Headers = v
	return s
}

func (s *GetAllProductResponse) SetBody(v *GetAllProductResponseBody) *GetAllProductResponse {
	s.Body = v
	return s
}

type GetProductByNameRequest struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// market
	Market *int32 `json:"Market,omitempty" xml:"Market,omitempty"`
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// bucUid
	EmpId *int64 `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
}

func (s GetProductByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductByNameRequest) GoString() string {
	return s.String()
}

func (s *GetProductByNameRequest) SetRequestId(v string) *GetProductByNameRequest {
	s.RequestId = &v
	return s
}

func (s *GetProductByNameRequest) SetMarket(v int32) *GetProductByNameRequest {
	s.Market = &v
	return s
}

func (s *GetProductByNameRequest) SetName(v string) *GetProductByNameRequest {
	s.Name = &v
	return s
}

func (s *GetProductByNameRequest) SetEmpId(v int64) *GetProductByNameRequest {
	s.EmpId = &v
	return s
}

type GetProductByNameResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 产品信息
	Data *GetProductByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProductByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductByNameResponseBody) SetCode(v string) *GetProductByNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductByNameResponseBody) SetData(v *GetProductByNameResponseBodyData) *GetProductByNameResponseBody {
	s.Data = v
	return s
}

func (s *GetProductByNameResponseBody) SetRequestId(v string) *GetProductByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductByNameResponseBody) SetSuccess(v bool) *GetProductByNameResponseBody {
	s.Success = &v
	return s
}

type GetProductByNameResponseBodyData struct {
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 版本
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// domain
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetProductByNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductByNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductByNameResponseBodyData) SetTitle(v string) *GetProductByNameResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetProductByNameResponseBodyData) SetVersions(v []*string) *GetProductByNameResponseBodyData {
	s.Versions = v
	return s
}

func (s *GetProductByNameResponseBodyData) SetName(v string) *GetProductByNameResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductByNameResponseBodyData) SetDomain(v string) *GetProductByNameResponseBodyData {
	s.Domain = &v
	return s
}

type GetProductByNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductByNameResponse) GoString() string {
	return s.String()
}

func (s *GetProductByNameResponse) SetHeaders(v map[string]*string) *GetProductByNameResponse {
	s.Headers = v
	return s
}

func (s *GetProductByNameResponse) SetBody(v *GetProductByNameResponseBody) *GetProductByNameResponse {
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

func (client *Client) UpdateApiWithOptions(runtime *util.RuntimeOptions) (_result *UpdateApiResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &UpdateApiResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApi"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApi() (_result *UpdateApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApiResponse{}
	_body, _err := client.UpdateApiWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCommondWithOptions(runtime *util.RuntimeOptions) (_result *SendCommondResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &SendCommondResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendCommond"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCommond() (_result *SendCommondResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCommondResponse{}
	_body, _err := client.SendCommondWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllProductWithOptions(request *GetAllProductRequest, runtime *util.RuntimeOptions) (_result *GetAllProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAllProductResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAllProduct"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllProduct(request *GetAllProductRequest) (_result *GetAllProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllProductResponse{}
	_body, _err := client.GetAllProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductByNameWithOptions(request *GetProductByNameRequest, runtime *util.RuntimeOptions) (_result *GetProductByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetProductByNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProductByName"), tea.String("2020-02-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductByName(request *GetProductByNameRequest) (_result *GetProductByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductByNameResponse{}
	_body, _err := client.GetProductByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
