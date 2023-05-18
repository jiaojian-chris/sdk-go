package v20160304
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type DescribeAddressesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}


type GetLinesRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *GetLinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLinesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetLinesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type GetLinesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	LineSet []struct {
		LineId *string `json:"LineId"`
		LineType *string `json:"LineType"`
		LineName *string `json:"LineName"`
		IpVersion *string `json:"IpVersion"`
	} `json:"LineSet"`
}

func (r *GetLinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetLinesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesRequest struct {
    *ksyunhttp.BaseRequest
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    AllocationId []*string `json:"AllocationId,omitempty" name:"AllocationId"`
    Filter []*DescribeAddressesFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
    State *string `json:"State,omitempty" name:"State"`
    IpVersion *string `json:"IpVersion,omitempty" name:"IpVersion"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAddressesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAddressesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	AddressesSet []struct {
		LineId *string `json:"LineId"`
		BandWidth *int `json:"BandWidth"`
		CreateTime *string `json:"CreateTime"`
		State *string `json:"State"`
		AllocationId *string `json:"AllocationId"`
		PublicIp *string `json:"PublicIp"`
		InstanceType *string `json:"InstanceType"`
		InstanceId *string `json:"InstanceId"`
		IpVersion *string `json:"IpVersion"`
		NetworkInterfaceId *string `json:"NetworkInterfaceId"`
		NetworkInterfaceType *string `json:"NetworkInterfaceType"`
		PrivateIpAddress *string `json:"PrivateIpAddress"`
		BandWidthShareId *string `json:"BandWidthShareId"`
		ProjectId *string `json:"ProjectId"`
		Mode *string `json:"Mode"`
	} `json:"AddressesSet"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAddressesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressRequest struct {
    *ksyunhttp.BaseRequest
    LineId *string `json:"LineId,omitempty" name:"LineId"`
    BandWidth *int `json:"BandWidth,omitempty" name:"BandWidth"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
    PurchaseTime *int `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
}

func (r *AllocateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AllocateAddressRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AllocateAddressResponse struct {
    *ksyunhttp.BaseResponse
    AllocationId *string `json:"AllocationId" name:"AllocationId"`
    PublicIp *string `json:"PublicIp" name:"PublicIp"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    IpVersion *string `json:"IpVersion" name:"IpVersion"`
}

func (r *AllocateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AllocateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressRequest struct {
    *ksyunhttp.BaseRequest
    AllocationId *string `json:"AllocationId,omitempty" name:"AllocationId"`
}

func (r *ReleaseAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ReleaseAddressRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseAddressResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *ReleaseAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressRequest struct {
    *ksyunhttp.BaseRequest
    AllocationId *string `json:"AllocationId,omitempty" name:"AllocationId"`
    InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
    NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" name:"NetworkInterfaceId"`
    Mode *string `json:"Mode,omitempty" name:"Mode"`
    PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateAddressRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AssociateAddressResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressRequest struct {
    *ksyunhttp.BaseRequest
    AllocationId *string `json:"AllocationId,omitempty" name:"AllocationId"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateAddressRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateAddressResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressRequest struct {
    *ksyunhttp.BaseRequest
    AllocationId *string `json:"AllocationId,omitempty" name:"AllocationId"`
    BandWidth *int `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *ModifyAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyAddressRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAddressResponse struct {
    *ksyunhttp.BaseResponse
    LineId *string `json:"LineId" name:"LineId"`
    BandWidth *int `json:"BandWidth" name:"BandWidth"`
    CreateTime *string `json:"CreateTime" name:"CreateTime"`
    State *string `json:"State" name:"State"`
    AllocationId *string `json:"AllocationId" name:"AllocationId"`
    PublicIp *string `json:"PublicIp" name:"PublicIp"`
    RequestId *string `json:"RequestId" name:"RequestId"`
    InstanceType *string `json:"InstanceType" name:"InstanceType"`
    InstanceId *string `json:"InstanceId" name:"InstanceId"`
    IpVersion *string `json:"IpVersion" name:"IpVersion"`
    NetworkInterfaceId *string `json:"NetworkInterfaceId" name:"NetworkInterfaceId"`
}

func (r *ModifyAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAddressResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

