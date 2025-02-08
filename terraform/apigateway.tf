resource "aws_api_gateway_rest_api" "curricular_and_academic_api" {
  name        = "Curricular API"
  description = "An API for exposing teacher, class, course data"
}

resource "aws_api_gateway_method_settings" "curricular_and_academic_api" {
  rest_api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id
  stage_name  = aws_api_gateway_stage.curricular_and_academic_api.stage_name
  method_path = "*/*"

  settings {
    logging_level = "INFO"
  }
}

resource "aws_api_gateway_resource" "students" {
  rest_api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id
  parent_id   = aws_api_gateway_rest_api.curricular_and_academic_api.root_resource_id
  path_part   = "students"
}

resource "aws_api_gateway_method" "students" {
  rest_api_id      = aws_api_gateway_rest_api.curricular_and_academic_api.id
  resource_id      = aws_api_gateway_resource.students.id
  http_method      = "GET"
  authorization    = "NONE"
  api_key_required = true
}

resource "aws_api_gateway_integration" "get_students_integration" {
  rest_api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id
  resource_id = aws_api_gateway_resource.students.id
  http_method = aws_api_gateway_method.students.http_method

  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.curricular_api.invoke_arn
}

resource "aws_api_gateway_resource" "student_emplid" {
  rest_api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id
  parent_id   = aws_api_gateway_resource.students.id
  path_part   = "{emplid}"
}

resource "aws_api_gateway_method" "get_student" {
  rest_api_id      = aws_api_gateway_rest_api.curricular_and_academic_api.id
  resource_id      = aws_api_gateway_resource.student_emplid.id
  http_method      = "GET"
  authorization    = "NONE"
  api_key_required = true
}

resource "aws_api_gateway_integration" "get_student_integration" {
  rest_api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id
  resource_id = aws_api_gateway_resource.student_emplid.id
  http_method = aws_api_gateway_method.get_student.http_method

  integration_http_method = "GET"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.curricular_api.invoke_arn
}

resource "aws_api_gateway_deployment" "curricular_and_academic_api" {
  rest_api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id

  lifecycle {
    create_before_destroy = true
  }

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_resource.students,
      aws_api_gateway_method.get_student,
      aws_api_gateway_integration.get_student_integration,
    ]))
  }
}

resource "aws_api_gateway_stage" "curricular_and_academic_api" {
  deployment_id = aws_api_gateway_deployment.curricular_and_academic_api.id
  rest_api_id   = aws_api_gateway_rest_api.curricular_and_academic_api.id
  stage_name    = "curricular-and-academic-api"
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.curricular_api.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.curricular_and_academic_api.execution_arn}/*/*"
}

resource "aws_api_gateway_usage_plan" "curricular_and_academic_api" {
  name = "curricular_and_academic_api"

  depends_on = [
    aws_api_gateway_method_settings.curricular_and_academic_api,
    aws_api_gateway_stage.curricular_and_academic_api
  ]

  api_stages {
    api_id = aws_api_gateway_rest_api.curricular_and_academic_api.id
    stage  = aws_api_gateway_stage.curricular_and_academic_api.stage_name
  }
}

resource "aws_api_gateway_api_key" "curricular_and_academic_api" {
  name    = "CurricularAcademicApiKey"
  enabled = true
}

resource "aws_api_gateway_usage_plan_key" "curricular_and_academic_api" {
  key_id        = aws_api_gateway_api_key.curricular_and_academic_api.id
  key_type      = "API_KEY"
  usage_plan_id = aws_api_gateway_usage_plan.curricular_and_academic_api.id
}
