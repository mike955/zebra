# zebra

gateway 
  - server: 5080
  - prometheus: 5081
flake 
  - server: 5180
  - prometheus: 5181
account 
  - server: 5280
  - prometheus: 5281
age 
  - server: 5380
  - prometheus: 5381
email 
  - server: 5480
  - prometheus: 5481
bank 
  - server: 5580
  - prometheus: 5581
cellphone 
  - server: 5680
  - prometheus: 5681


1.pmem 的使用是否一定要创建 region 和 namespace，然后把 pmem 关联到 region 和 namespace？
2.创建 region 和 namespace 并且与 pmem 关联后，是否就能直接使用 pmem 了，如果不行，还需要做什么操作或配置？


protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ flake.proto

protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ account.proto

protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ age.proto

protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ email.proto

protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ bank.proto

protoc --proto_path=./api/.  --go-grpc_out=./api/ --go_out=./api/ cellphone.proto






sum(rate(grpc_server_started_total{grpc_service="flake.FlakeService"}[1m])) by (grpc_service)

sum(rate(grpc_server_started_total{grpc_service="account.AccountService"}[1m])) by (grpc_service)

sum(rate(grpc_server_started_total{grpc_service="age.AgeService"}[1m])) by (grpc_service)

sum(rate(grpc_server_started_total{grpc_service="email.EmailService"}[1m])) by (grpc_service)

sum(rate(grpc_server_started_total{grpc_service="bank.BankService"}[1m])) by (grpc_service)

sum(rate(grpc_server_started_total{grpc_service="cellphone.CellphoneService"}[1m])) by (grpc_service)






docker run --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -d docker.elastic.co/elasticsearch/elasticsearch:7.12.0