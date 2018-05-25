aggregator - concurrency example <br />
clone this repo <br />
cd into repo <br />
go build <br />
run the exe <br />
hit it with a curl request <br />
curl -d '[{"data":"hello", "time":1000}]' -H "Content-Type:         application/json" -X POST http://localhost:8888                     #prints "{"result": 10}" <br />
or you can use a data.json file and use the following command <br />
curl -d "@data.json" -X POST http://localhost:8888 </br >
