# \AuthenticationApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2AuthSocialAuthorizeCodeGet**](AuthenticationApi.md#V2AuthSocialAuthorizeCodeGet) | **Get** /v2/auth/social/authorizeCode | Second Step in Social Authentication flow with JWT Token
[**V2AuthSocialAuthorizeTokenGet**](AuthenticationApi.md#V2AuthSocialAuthorizeTokenGet) | **Get** /v2/auth/social/authorizeToken | Native Social Authentication
[**V2AuthSocialLoginGet**](AuthenticationApi.md#V2AuthSocialLoginGet) | **Get** /v2/auth/social/login | First Setp in Social Authentication flow with JWT Token
[**V2Oauth2AccessTokenGet**](AuthenticationApi.md#V2Oauth2AccessTokenGet) | **Get** /v2/oauth2/access_token | Get a user access token
[**V2OauthAuthorizeGet**](AuthenticationApi.md#V2OauthAuthorizeGet) | **Get** /v2/oauth/authorize | Request Authorization Code


# **V2AuthSocialAuthorizeCodeGet**
> V2AuthSocialAuthorizeCodeGet($code, $provider)

Second Step in Social Authentication flow with JWT Token

 Here is the flow for how social authentication works with a JWT Token 1.**Client:** The client needs to open popup with social auth url (`https://app/quantimo.do/api/v2/auth/social/login?provider={provider}&redirectUrl={url}`) of server with `provider` and `redirectUrl`. (Url should be registered with our social apps. Facebook is fine with any redirect url with the same domain base url but Google needs exact redirect url.) 2.**Server:** The QM server will redirect user to that provider to get access. 3.**Client:** After successful or failed authentication, it will be redirected to given `redirectUrl` with code or error. 4.**Client:** The client needs to get that code and needs to send an Ajax request to server at `https://app.quantimo.do/api/v2/auth/social/authorizeCode?provider={provider}&code={authorizationCode}` 5.**Server:** The QM server will authorize that code from the social connection and will authenticate user and will retrieve user info. 6.**Server:** The QM server will try to find existing user by unique identity. If the user already exists then it will login. Otherwise, it will create new user and will then login. 7.**Server:** Once user is found/created, it will return a JWT token for that user in the response.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| Authorization code obtained from the provider. | 
 **provider** | **string**| The current options are &#x60;google&#x60; and &#x60;facebook&#x60;. | 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2AuthSocialAuthorizeTokenGet**
> V2AuthSocialAuthorizeTokenGet($accessToken, $provider, $refreshToken)

Native Social Authentication

If you are using native authentication via Facebook or Google SDKs then you should use the following flow. 1.**Client:** Using native authentication via your native mobile app, get an access token using the instructions provided by the Facebook SDK (https://developers.facebook.com/docs/facebook-login) or Google (https://developers.google.com/identity/protocols/OAuth2) 2.**Client:** Send an Ajax request with provider name and access token on `https://app.quantimo.do/api/v2/auth/social/authorizeToken?provider={provider}&accessToken={accessToken}&refreshToken={refreshToken}` (`refreshToken` is optional) 3.**Server:** Server will try to get user info and will find existing user by unique identity. If user exist then it will do a login for that or it will create new user and will do login 4.**Server:** Once user is found/created, it will return a JWT token for that user in response 5.**Client:** After getting the JWT token to get a QM access token follow these steps and include your JWT token in them as a header (Authorization: Bearer **{yourJWThere}**) or as a url parameter (https://app.quantimo.do/api/v2/oauth/authorize?token={yourJWThere}).


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token obtained from Google or FB native SDK | 
 **provider** | **string**| The current options are &#x60;google&#x60; and &#x60;facebook&#x60;. | 
 **refreshToken** | **string**| Optional refresh token obtained from Google or FB native SDK | [optional] 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2AuthSocialLoginGet**
> V2AuthSocialLoginGet($redirectUrl, $provider)

First Setp in Social Authentication flow with JWT Token

 Here is the flow for how social authentication works with a JWT Token 1.**Client:** The client needs to open popup with social auth url (`https://app/quantimo.do/api/v2/auth/social/login?provider={provider}&redirectUrl={url}`) of server with `provider` and `redirectUrl`. (Url should be registered with our social apps. Facebook and Twitter are fine with any redirect url with the same domain base url but Google needs exact redirect url.) 2.**Server:** The QM server will redirect user to that provider to get access. 3.**Client:** After successful or failed authentication, it will be redirected to given `redirectUrl` with code or error. 4.**Client:** The client needs to get that code and needs to send an Ajax request to server at `https://app.quantimo.do/api/v2/auth/social/authorizeCode?provider={provider}&code={authorizationCode}` 5.**Server:** The QM server will authorize that code from the social connection and will authenticate user and will retrieve user info. 6.**Server:** The QM server will try to find existing user by unique identity. If the user already exists then it will login. Otherwise, it will create new user and will then login. 7.**Server:** Once user is found/created, it will return a JWT token for that user in the response.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUrl** | **string**| The redirect URI is the URL within your client application that will receive the OAuth2 credentials. Url should be registered with our social apps. Facebook and Twitter are fine with any redirect url with the same domain base url but Google needs exact redirect url. | 
 **provider** | **string**| The current options are &#x60;google&#x60; and &#x60;facebook&#x60;. | 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2Oauth2AccessTokenGet**
> V2Oauth2AccessTokenGet($clientId, $clientSecret, $grantType, $code, $responseType, $scope, $redirectUri, $state)

Get a user access token

Client provides authorization token obtained from /api/v1/oauth2/authorize to this endpoint and receives an access token. Access token can then be used to query different API endpoints of QuantiModo. ### Request Access Token After user approves your access to the given scope form the https:/app.quantimo.do/v2/oauth2/authorize endpoint, you'll receive an authorization code to request an access token. This time make a `POST` request to `/api/v2/oauth/access_token` with parameters including: * `grant_type` Can be `authorization_code` or `refresh_token` since we are getting the `access_token` for the first time we don't have a `refresh_token` so this must be `authorization_code`. * `code` Authorization code you received with the previous request. * `redirect_uri` Your application's redirect url. ### Refreshing Access Token Access tokens expire at some point, to continue using our api you need to refresh them with `refresh_token` you received along with the `access_token`. To do this make a `POST` request to `/api/v2/oauth/access_token` with correct parameters, which are: * `grant_type` This time grant type must be `refresh_token` since we have it. * `clientId` Your application's client id. * `client_secret` Your application's client secret. * `refresh_token` The refresh token you received with the `access_token`. Every request you make to this endpoint will give you a new refresh token and make the old one expired. So you can keep getting new access tokens with new refresh tokens. ### Using Access Token Currently we support 2 ways for this, you can't use both at the same time. * Adding access token to the request header as `Authorization: Bearer {access_token}` * Adding to the url as a query parameter `?access_token={access_token}` You can read more about OAuth2 from [here](http://oauth.net/2/)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| This is the unique ID that QuantiModo uses to identify your application. Obtain a client id by emailing info@quantimo.do. | 
 **clientSecret** | **string**| This is the secret for your obtained clientId. QuantiModo uses this to validate that only your application uses the clientId. | 
 **grantType** | **string**| Grant Type can be &#39;authorization_code&#39; or &#39;refresh_token&#39; | 
 **code** | **string**| Authorization code you received with the previous request. | 
 **responseType** | **string**| If the value is code, launches a Basic flow, requiring a POST to the token endpoint to obtain the tokens. If the value is token id_token or id_token token, launches an Implicit flow, requiring the use of Javascript at the redirect URI to retrieve tokens from the URI #fragment. | [optional] 
 **scope** | **string**| Scopes include basic, readmeasurements, and writemeasurements. The \&quot;basic\&quot; scope allows you to read user info (displayname, email, etc). The \&quot;readmeasurements\&quot; scope allows one to read a user&#39;s data. The \&quot;writemeasurements\&quot; scope allows you to write user data. Separate multiple scopes by a space. | [optional] 
 **redirectUri** | **string**| The redirect URI is the URL within your client application that will receive the OAuth2 credentials. | [optional] 
 **state** | **string**| An opaque string that is round-tripped in the protocol; that is to say, it is returned as a URI parameter in the Basic flow, and in the URI | [optional] 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2OauthAuthorizeGet**
> V2OauthAuthorizeGet($clientId, $clientSecret, $responseType, $scope, $redirectUri, $state)

Request Authorization Code

You can implement OAuth2 authentication to your application using our **OAuth2** endpoints.  You need to redirect users to `/api/v2/oauth/authorize` endpoint to get an authorization code and include the parameters below.   This page will ask the user if they want to allow a client's application to submit or obtain data from their QM account. It will redirect the user to the url provided by the client application with the code as a query parameter or error in case of an error. See the /api/v2/oauth/access_token endpoint for the next steps.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| This is the unique ID that QuantiModo uses to identify your application. Obtain a client id by creating a free application at [https://app.quantimo.do/api/v2/apps](https://app.quantimo.do/api/v2/apps). | 
 **clientSecret** | **string**| This is the secret for your obtained clientId. QuantiModo uses this to validate that only your application uses the clientId.  Obtain this by creating a free application at [https://app.quantimo.do/api/v2/apps](https://app.quantimo.do/api/v2/apps). | 
 **responseType** | **string**| If the value is code, launches a Basic flow, requiring a POST to the token endpoint to obtain the tokens. If the value is token id_token or id_token token, launches an Implicit flow, requiring the use of Javascript at the redirect URI to retrieve tokens from the URI #fragment. | 
 **scope** | **string**| Scopes include basic, readmeasurements, and writemeasurements. The \&quot;basic\&quot; scope allows you to read user info (displayname, email, etc). The \&quot;readmeasurements\&quot; scope allows one to read a user&#39;s data. The \&quot;writemeasurements\&quot; scope allows you to write user data. Separate multiple scopes by a space. | 
 **redirectUri** | **string**| The redirect URI is the URL within your client application that will receive the OAuth2 credentials. | [optional] 
 **state** | **string**| An opaque string that is round-tripped in the protocol; that is to say, it is returned as a URI parameter in the Basic flow, and in the URI | [optional] 

### Return type

void (empty response body)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

