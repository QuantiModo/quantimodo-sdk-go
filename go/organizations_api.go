/* 
 * QuantiModo
 *
 * QuantiModo makes it easy to retrieve normalized user data from a wide array of devices and applications. [Learn about QuantiModo](https://quantimo.do), check out our [docs](https://github.com/QuantiModo/docs) or contact us at [help.quantimo.do](https://help.quantimo.do). 
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type OrganizationsApi struct {
	Configuration *Configuration
}

func NewOrganizationsApi() *OrganizationsApi {
	configuration := NewConfiguration()
	return &OrganizationsApi{
		Configuration: configuration,
	}
}

func NewOrganizationsApiWithBasePath(basePath string) *OrganizationsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &OrganizationsApi{
		Configuration: configuration,
	}
}

/**
 * Get user tokens for existing users, create new users
 * Get user tokens for existing users, create new users
 *
 * @param organizationId Organization ID
 * @param body Provides organization token and user ID
 * @param accessToken User&#39;s OAuth2 access token
 * @param userId User&#39;s id
 * @return *UserTokenSuccessfulResponse
 */
func (a OrganizationsApi) V1OrganizationsOrganizationIdUsersPost(organizationId int32, body UserTokenRequest, accessToken string, userId int32) (*UserTokenSuccessfulResponse, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/v1/organizations/{organizationId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", fmt.Sprintf("%v", organizationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
		localVarQueryParams.Add("access_token", a.Configuration.APIClient.ParameterToString(accessToken, ""))
		localVarQueryParams.Add("userId", a.Configuration.APIClient.ParameterToString(userId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(UserTokenSuccessfulResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "V1OrganizationsOrganizationIdUsersPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

