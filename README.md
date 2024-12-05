# Go client for Forward Email

- [Forward Email API Documentation](https://forwardemail.net/en/email-api)
- [Forward Email Terraform Provider](https://github.com/the-infra-company/terraform-provider-forwardemail)

### How to install

```shell
$ go get github.com/the-infra-company/forwardemail-api-go
```

### Basic usage

```go
import "github.com/the-infra-company/forwardemail-api-go/forwardemail"

client := forwardemail.NewClient(forwardemail.ClientOptions{
    ApiKey: key,
})

account, err := client.GetAccount()
```

