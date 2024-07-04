# Service
## Service API

| API              | method | path | request body                            |
|------------------|--------|------|-----------------------------------------|
| getBlockNumber   | POST   |getBlockNumber      | { "id": 2 }                             |
| getBlockByNumber | POST   |getBlockByNumber      | { "id": 2, "blockNumber": "0x134e82a" } |
| healthCheck      | GET    | health_check     | NA                                      |

## for local test API
```shell
go run main.go handler.go rpcClient.go
```

```shell
curl --location 'http://localhost:8080/getBlockNumber' \
--header 'Content-Type: application/json' \
--data '{
    "id": 28
}'

{"id":28,"jsonrpc":"2.0","result":"0x38389df"}


 curl --location --request GET 'http://localhost:8080/health_check' \
--header 'Content-Type: application/json' \
--data '{
    "id": 2
    "blockNumber":
}'

ok

curl --location 'http://localhost:8080/getBlockByNumber' \
--header 'Content-Type: application/json' \
--data '{
    "id": 2,
    "blockNumber": "0x134e82a"
}'
```

## unit test
```shell
➜   go test -v

=== RUN   TestGetBlockNumber
--- PASS: TestGetBlockNumber (0.00s)
=== RUN   TestGetBlockByNumber
--- PASS: TestGetBlockByNumber (0.00s)
=== RUN   TestHealthCheck
--- PASS: TestHealthCheck (0.00s)
PASS
ok      blockchain-client       0.315s
```

## docker build and upload to docker hub
You can build and upload to your docker hub space
```shell
docker build -t feiyang233/proxy-client:v1 .

docker login

docker push feiyang233/proxy-client:v1
```
You can get this image from https://hub.docker.com/r/feiyang233/proxy-client/tags

# Terraform 
## before start
Update the credentials under terraform.tfvars

```terraform
aws_access_key = "AWS_ACCESS_KEY"
aws_secret_key = "AWS_SECRET_KEY"
aws_region     = "ap-southeast-1"
```


## terraform init issue
if you got issue like
```shell
│ Error: Incompatible provider version
│ 
│ Provider registry.terraform.io/hashicorp/template v2.2.0 does not have a package available for your current platform, darwin_arm64.
│ 
│ Provider releases are separate from Terraform CLI releases, so not all providers are available for all platforms. Other versions of this provider may have different platforms supported.
╵
```
refer this https://discuss.hashicorp.com/t/template-v2-2-0-does-not-have-a-package-available-mac-m1/35099/7
Try below command in your M-chip mac
```shell
brew install kreuzwerker/taps/m1-terraform-provider-helper
m1-terraform-provider-helper activate
m1-terraform-provider-helper install hashicorp/template -v v2.2.0
```

## fmt and validate
```shell
➜   terraform -v
Terraform v1.9.1
on darwin_arm64
+ provider registry.terraform.io/hashicorp/aws v5.57.0
+ provider registry.terraform.io/hashicorp/template v2.2.0

➜   terraform init 
Initializing the backend...
Initializing provider plugins...
- Reusing previous version of hashicorp/aws from the dependency lock file
- Reusing previous version of hashicorp/template from the dependency lock file
- Using previously-installed hashicorp/aws v5.57.0
- Using previously-installed hashicorp/template v2.2.0

Terraform has been successfully initialized!


➜   terraform fmt
➜   terraform validate 
Success! The configuration is valid.

```

## terraform plan
Can check the full output at tfplan.txt file
```shell
terraform plan

Terraform will perform the following actions:

  # aws_alb.main will be created
  + resource "aws_alb" "main" {
      + arn                                                          = (known after apply)
      + arn_suffix                                                   = (known after apply)
      + client_keep_alive                                            = 3600
      + desync_mitigation_mode                                       = "defensive"
      + dns_name                                                     = (known after apply)
      + drop_invalid_header_fields                                   = false
      + enable_deletion_protection                                   = false
      + enable_http2                                                 = true
      + enable_tls_version_and_cipher_suite_headers                  = false
      + enable_waf_fail_open                                         = false
      + enable_xff_client_port                                       = false
      + enforce_security_group_inbound_rules_on_private_link_traffic = (known after apply)
      + id                                                           = (known after apply)
      + idle_timeout                                                 = 60
      + internal                                                     = (known after apply)
      + ip_address_type                                              = (known after apply)
      + load_balancer_type                                           = "application"
      + name                                                         = "proxy-client-load-balancer"
      + name_prefix                                                  = (known after apply)
      + preserve_host_header                                         = false
      + security_groups                                              = (known after apply)
      + subnets                                                      = (known after apply)
      + tags_all                                                     = (known after apply)
      + vpc_id                                                       = (known after apply)
      + xff_header_processing_mode                                   = "append"
      + zone_id                                                      = (known after apply)

      + subnet_mapping (known after apply)
    }


```
## terrform apply
```shell
Changes to Outputs:
  + alb_hostname = (known after apply)
```
After get ALB DNS domain, we can test the API in AWS env

## What can improve
In this experiment, we use docker hub as free registry. https://hub.docker.com/r/feiyang233/proxy-client/tags  
We can change to ECR, which is better when pull image from AWS