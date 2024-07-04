# logs.tf

# Set up CloudWatch group and log stream and retain logs for 30 days
resource "aws_cloudwatch_log_group" "proxy_log_group" {
  name              = "/ecs/proxy-client"
  retention_in_days = 30

  tags = {
    Name = "proxy-log-group"
  }
}

resource "aws_cloudwatch_log_stream" "proxy_log_stream" {
  name           = "proxy-log-stream"
  log_group_name = aws_cloudwatch_log_group.proxy_log_group.name
}