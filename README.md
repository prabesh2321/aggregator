# aggregator - concurrency example
clone this repo
cd into repo
go build
run the exe
hit it with a curl request
curl -d '[{"data":"hello", "time":1000}]' -H "Content-Type: application/json" -X POST http://localhost:8888 #prints "{"result": 10}"

