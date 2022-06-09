package v20200101
import (
    "context"
    "fmt"
    "github.com/kingsoftcloud/sdk-go/ksyun/common"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

const APIVersion = "2020-01-01"

type Client struct {
    common.Client
}

func NewClient(credential common.Credentials, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByPayMode")
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (string) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (string) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeBillSummaryByPayModeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProduct")
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (string) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (string) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeBillSummaryByProductResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
    request = &DescribeBillSummaryByProjectRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeBillSummaryByProject")
    return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
    response = &DescribeBillSummaryByProjectResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (string) {
    return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (string) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeBillSummaryByProjectResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeInstanceSummaryBillsRequest() (request *DescribeInstanceSummaryBillsRequest) {
    request = &DescribeInstanceSummaryBillsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeInstanceSummaryBills")
    return
}

func NewDescribeInstanceSummaryBillsResponse() (response *DescribeInstanceSummaryBillsResponse) {
    response = &DescribeInstanceSummaryBillsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeInstanceSummaryBills(request *DescribeInstanceSummaryBillsRequest) (string) {
    return c.DescribeInstanceSummaryBillsWithContext(context.Background(), request)
}

func (c *Client) DescribeInstanceSummaryBillsWithContext(ctx context.Context, request *DescribeInstanceSummaryBillsRequest) (string) {
    if request == nil {
        request = NewDescribeInstanceSummaryBillsRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeInstanceSummaryBillsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeProductCodeRequest() (request *DescribeProductCodeRequest) {
    request = &DescribeProductCodeRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeProductCode")
    return
}

func NewDescribeProductCodeResponse() (response *DescribeProductCodeResponse) {
    response = &DescribeProductCodeResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeProductCode(request *DescribeProductCodeRequest) (string) {
    return c.DescribeProductCodeWithContext(context.Background(), request)
}

func (c *Client) DescribeProductCodeWithContext(ctx context.Context, request *DescribeProductCodeRequest) (string) {
    if request == nil {
        request = NewDescribeProductCodeRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeProductCodeResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}
func NewDescribeSplitItemBillDetailsRequest() (request *DescribeSplitItemBillDetailsRequest) {
    request = &DescribeSplitItemBillDetailsRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bill-union", APIVersion, "DescribeSplitItemBillDetails")
    return
}

func NewDescribeSplitItemBillDetailsResponse() (response *DescribeSplitItemBillDetailsResponse) {
    response = &DescribeSplitItemBillDetailsResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) DescribeSplitItemBillDetails(request *DescribeSplitItemBillDetailsRequest) (string) {
    return c.DescribeSplitItemBillDetailsWithContext(context.Background(), request)
}

func (c *Client) DescribeSplitItemBillDetailsWithContext(ctx context.Context, request *DescribeSplitItemBillDetailsRequest) (string) {
    if request == nil {
        request = NewDescribeSplitItemBillDetailsRequest()
    }
    request.SetContext(ctx)

    response := NewDescribeSplitItemBillDetailsResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


