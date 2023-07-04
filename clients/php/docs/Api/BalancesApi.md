# OpenAPI\Client\BalancesApi

All URIs are relative to https://api.xendit.co, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**getBalance()**](BalancesApi.md#getBalance) | **GET** /balance | Get Balance |


## `getBalance()`

```php
getBalance($account_type, $for_user_id, $x_idempotency_key): \OpenAPI\Client\Model\Balance
```

Get Balance

Get Balance for a specific account type

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure HTTP basic authorization: basicAuth
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()
              ->setUsername('YOUR_USERNAME')
              ->setPassword('YOUR_PASSWORD');


$apiInstance = new OpenAPI\Client\Api\BalancesApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$account_type = 'account_type_example'; // string | Account Type
$for_user_id = 'for_user_id_example'; // string | For User ID
$x_idempotency_key = 'x_idempotency_key_example'; // string | Idempotency Key

try {
    $result = $apiInstance->getBalance($account_type, $for_user_id, $x_idempotency_key);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling BalancesApi->getBalance: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **account_type** | **string**| Account Type | |
| **for_user_id** | **string**| For User ID | [optional] |
| **x_idempotency_key** | **string**| Idempotency Key | [optional] |

### Return type

[**\OpenAPI\Client\Model\Balance**](../Model/Balance.md)

### Authorization

[basicAuth](../../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
