terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
}

provider "aws" {
  profile = ""
  region  = ""
}

//----------------------------------------------------------------------------------------------------------------------
//  VARIABLES
//----------------------------------------------------------------------------------------------------------------------
variable "slack-bot-release-app-name" {
  description = "Lambda Notify Releases"
  default = "slack-release-service"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "dev"
}

locals {
  app_id = "${lower(var.slack-bot-release-app-name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

//----------------------------------------------------------------------------------------------------------------------
//  DYNAMODB
//----------------------------------------------------------------------------------------------------------------------

resource "aws_dynamodb_table" "slack-bot" {
  name           = "Releases"
  billing_mode   = "PAY_PER_REQUEST"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "id"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name        = "SlackBot"
    Environment = "dev"
  }
}

//----------------------------------------------------------------------------------------------------------------------
//  ROLES
//----------------------------------------------------------------------------------------------------------------------

resource "random_id" "unique_suffix" {
  byte_length = 2
}

resource "aws_iam_role" "lambda_exec" {
  name_prefix = local.app_id

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

resource "aws_iam_role_policy_attachment" "lambda_AmazonDynamoDBFullAccess" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
}

resource "aws_iam_role_policy_attachment" "lambda_CloudWatchFullAccess" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchFullAccess"
}

//----------------------------------------------------------------------------------------------------------------------
//  LAMBDAS
//----------------------------------------------------------------------------------------------------------------------
data "archive_file" "slack-bot-release" {
  output_path = "../lambda/bin/slack-bot-release.zip"
  source_file = "../lambda/bin/slack-bot-release-aws-lambda"
  type = "zip"
}

resource "aws_lambda_function" "slack-bot-release-aws-lambda" {
  filename = data.archive_file.slack-bot-release.output_path
  function_name = local.app_id
  handler = "slack-bot-release-aws-lambda"
  role = aws_iam_role.lambda_exec.arn
  source_code_hash = base64sha256(data.archive_file.slack-bot-release.output_path)
  runtime = "go1.x"
}

//----------------------------------------------------------------------------------------------------------------------
//  SNS
//----------------------------------------------------------------------------------------------------------------------
resource "aws_sns_topic" "slack-release" {
  name = "slackBotReleaseAwsSnsTopic"
  display_name = "SlackRelease"
}

resource "aws_sns_topic_subscription" "slack-release-suscription" {
  topic_arn = aws_sns_topic.slack-release.arn
  protocol  = "email"
  endpoint  = "gmartinezgranella@gmail.com"
}

//----------------------------------------------------------------------------------------------------------------------
//  ALARM CLOUDWATCH
//----------------------------------------------------------------------------------------------------------------------

resource "aws_cloudwatch_metric_alarm" "slack-release-alarm" {
  alarm_name                = "slack-release-alarm"
  comparison_operator       = "GreaterThanThreshold"
  evaluation_periods        = "4"
  metric_name               = "Errors"
  namespace                 = "AWS/Lambda"
  period                    = "60"
  statistic                 = "Sum"
  alarm_description         = "This metric detect count errors on lambda"
  actions_enabled     = "true"
  alarm_actions       = [aws_sns_topic.slack-release.arn]
  ok_actions          = [aws_sns_topic.slack-release.arn]
  dimensions = {
    FunctionName = local.app_id
  }
}