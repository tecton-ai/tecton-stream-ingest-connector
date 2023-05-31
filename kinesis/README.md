# How to setup Integration

### Build and zip the go code for AWS Lambda
```
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

### Create the AWS Lambda
**Code Source**
Use the zip file created above called `main.zip` and upload it

**Runtime Settings**
1. Runtime - `Go 1.x`
2. Handler - `main`
3. Architectire - `x86_64`

**Environment Variable**
Set the following enviroment variables appropriately
1. `PUSH_SOURCE_NAME`
2. `WORKSPACE_NAME`
3. `TECTON_API_URL`
4. `TECTON_AUTH_TOKEN`


**Setup Stream Trigger**
1. Choose trigger as `Kinesis`
2. Choose stream name from drop down
3. Choose Batch size and Starting position appropriately: We recommend batch size as 5 and starting position as Earliest

You can hit `Activate Trigger` when you have setup everything. The `Monitor` tab will provide relevant metrics and logs
