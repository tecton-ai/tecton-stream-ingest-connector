# How to setup Integration

This uses the built-in AWS Lambda and Kinesis integration to call the Stream Ingest API. At a high level the Lambda is auto invoked when events arrive in AWS Kinesis. The Lambda simply parses the records and sends them to Tecton for low latency ingestion
https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html

### Build and zip the go code for AWS Lambda
```
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

### Create the AWS Lambda
##### Code Source

Use the zip file created above called `main.zip` and upload it

##### Runtime Settings

1. Runtime - `Go 1.x`
2. Handler - `main`
3. Architecture - `x86_64`
<img width="837" alt="Screen Shot 2023-05-31 at 11 48 05 AM" src="https://github.com/tecton-ai/tecton-stream-ingest-connector/assets/10210921/293d7cf7-f83b-4504-bb73-d3003643b7bc">

##### Environment Variable

Set the following enviroment variables appropriately:
1. `PUSH_SOURCE_NAME`
2. `WORKSPACE_NAME`
3. `TECTON_API_URL`
4. `TECTON_AUTH_TOKEN`
<img width="1428" alt="Screen Shot 2023-05-31 at 11 48 40 AM" src="https://github.com/tecton-ai/tecton-stream-ingest-connector/assets/10210921/35e2d607-ed5e-4546-87e8-b0012fe79398">


##### Setup Stream Trigger

1. Choose trigger as `Kinesis`
2. Choose stream name from drop down
3. Choose Batch size and Starting position appropriately: We recommend batch size as 5 and starting position as Earliest
<img width="810" alt="Screen Shot 2023-05-31 at 11 48 22 AM" src="https://github.com/tecton-ai/tecton-stream-ingest-connector/assets/10210921/b2d2824d-c8c7-4192-a6c2-8e36d547ea52">



**You can hit `Activate Trigger` when you have setup everything. The `Monitor` tab will provide relevant metrics and logs**
