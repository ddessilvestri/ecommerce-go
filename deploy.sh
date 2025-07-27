#!/bin/bash

# 🚀 Variables (EDIT these for your Lambda)
LAMBDA_FUNCTION_NAME="YourLambdaFunctionName"

echo "🔨 Building Go binary for Lambda..."
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

if [ $? -ne 0 ]; then
  echo "❌ Build failed"
  exit 1
fi

echo "📦 Packaging into function.zip..."
rm -f function.zip
zip function.zip bootstrap

# echo "☁️ Deploying to AWS Lambda: $LAMBDA_FUNCTION_NAME..."
# aws lambda update-function-code \
#   --function-name "$LAMBDA_FUNCTION_NAME" \
#   --zip-file fileb://function.zip \
#   --region us-east-1

if [ $? -eq 0 ]; then
  echo "✅ Deployment successful!"
else
  echo "❌ Deployment failed"
fi