# refer Using Terraform and Fargate to create Amazon’s ECS
https://medium.com/@olayinkasamuel44/using-terraform-and-fargate-to-create-amazons-ecs-e3308c1b9166

# before start
Update the credentials under terraform.tfvars

```terraform
aws_access_key = "AWS_ACCESS_KEY"
aws_secret_key = "AWS_SECRET_KEY"
aws_region     = "xxxxx"
```


# terraform init issue
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

# fmt and validate
```shell
➜  terraform terraform -v
Terraform v1.9.1
on darwin_arm64
+ provider registry.terraform.io/hashicorp/aws v5.57.0
+ provider registry.terraform.io/hashicorp/template v2.2.0
➜  terraform terraform fmt
➜  terraform terraform validate 
Success! The configuration is valid.

➜  terraform 
```