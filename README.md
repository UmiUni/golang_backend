# Golang_backend Production APIs:

* (Most Updated Swagger API):
http://api.referhelper.com
* (Outdated README API) User login/registration:
https://github.com/Jogchat/golang_backend/blob/master/user_registration_api.md


# Why use Gin Framework for Golang
  Gin uses HttpRouter github repo, all the benefits of HttpRouter are listed in the project README: 
  https://github.com/julienschmidt/httprouter
 
# Why use Swaggo/swagger Framework
* Swaggo has a nice gin-middleware that using Swagger2.0 UI for API doc generating purpose. 
* Swaggo is an open source org created by a Taiwaness and a Japanese geeks aimed at providing swagger API support for different popular open source golang framework.
https://github.com/swaggo
* Comments should follow Declarative Comments Format:
https://github.com/Jogchat/golang_backend/blob/master/DeclarativeCommentsFormat.md


# GET STARTED

## Install Golang
For Ubuntu:
```
sudo apt-get update
sudo apt-get -y upgrade

find latestes version in:
https://golang.org/dl/

sudo curl -O https://storage.googleapis.com/golang/go1.10.3.linux-amd64.tar.gz
sudo tar -xvf go1.10.3.linux-amd64.tar.gz

```
* reference: https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79

## Set GOPATH
```
sudo mv DOWNLOADED_GO_DIR /usr/local
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
source ~/.profile
mkdir -p go/src/code.jogchat.internal
```
* for mac add this two lines to `~./bash_profile` and then `source ~/.bash_profile`
```
export GOPATH=$HOME/gocode
export PATH="$PATH:$GOPATH/bin"
```
## Clone two repositories to `$GOPATH/src`
```
git clone https://github.com/Jogchat/go-schemaless.git
git clone https://github.com/Jogchat/golang_backend.git
```
 
## Clone other dependency repos to `$GOPATH/src`
* git clone repos under $GOPATH/src/code.jogchat.internal
```
https://github.com/Jogchat/dgrijalva-jwt-go
https://github.com/Jogchat/dgryski-go-shardedkv
https://github.com/Jogchat/dgryski-go-metro
```

* download swaggo/swagger repos and swag executable
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles

```
* link  swaggo/swag executable, note that swag is installed in go's bin folder
```
export PATH=$PATH:/root/go/bin
```
* Then 
```source ~/.bash_profile```
* generate swag doc in our repo

```
swag init
```

* git clone repos under $GOPATH/src/github.com
```
https://github.com/go-sql-driver/mysql.git
```
## To update Go packages:
```
go get 
```

## Start server
under golang_backend, do `go run main.go`

# Jogchat React+Golang server

```
Rest API

POST ip:3001/signup
Input: JSON typed:
{
  Username:"test",
  Password:"test_password",
}
Successs: return true
Fail: return error
Description:

````


```
http://138.68.227.175:3001
```

# Backbone reference for this project:
* [odewahn/react-golang-auth](https://github.com/odewahn/react-golang-auth)
* [Building a Web App With Go, Gin and React](https://hakaselogs.me/2018-04-20/building-a-web-app-with-go-gin-and-react)

# Run this program:
```
go run main.go
```

# aws configuration
## login to console
username: jogchat@gmail.com 
password: jogchat_AWS
## config credential
Only one User is created, called admin. For aws a User is created with certain access. A user is associated with credentials so as to config a program to act as the User with the certain access. The credentials are stored in ~/.aws/credentials at the backend server. 

# For Future Microservice Migration Reference
Jogchat scalable authentication backend architecture
Backend:
Why we use Go? language built for solving scaling problem.
cons: learning curve high. less devs use it.
pros: Google and Uber use it to solve scaling/high-throughput problem.
Why and how Uber use Go services to scale:
https://eng.uber.com/go-geofence/


To scale use microservice RPC service to service talk:
https://ewanvalentine.io/microservices-in-golang-part-1/
https://github.com/ewanvalentine/shippy

Use global UUID for any object


Some good Udemy course:
https://www.udemy.com/go-programming-language/

Uber business logic open source Workflow cadence:
https://github.com/uber/cadence

Uber open source:
https://uber.github.io/orgs.html

Authenticate google login using Go:
https://skarlso.github.io/2016/06/12/google-signin-with-go/
https://cloud.google.com/go/getting-started/authenticate-users


Technologies used: 
google protobuf, rpc call, UUID, mongodb/mysql/postgres, maybe apache thrift? , Uber cadence workflow, docker virtual


Middle layer:
gRPC, kafka


Front-end:
front end grpc call to backend:
https://cloud.google.com/solutions/mobile/mobile-compute-engine-grpc

For payment systems:
We can outsourcing braintree and stripe
for financial PCI compliance


BREAKING DOWN 'PCI Compliance'. Credit card companies want merchants to handle cardholder information in a secure manner in order to reduce the likelihood of cardholders having their sensitivefinancial data stolen. 


grafana:
https://github.com/grafana/grafana



twemproxy (pronounced "two-em-proxy"), aka nutcracker is a fast and lightweight proxy for memcached and redisprotocol. It was built primarily to reduce the number of connections to the caching servers on the backend. This, together with protocol pipelining and sharding enables you to horizontally scale your distributed caching architecture.
https://github.com/twitter/twemproxy


Golang setup:
https://stackoverflow.com/questions/33774950/execute-gofmt-on-file-save-in-intellij

Use google protobuffer3 to serialize, deserialize data for sending messages

Use Go fmt:
https://stackoverflow.com/questions/33774950/execute-gofmt-on-file-save-in-intellij

always use UUID:
http://www.mysqltutorial.org/mysql-uuid/
https://eng.uber.com/mezzanine-migration/



https://grafana.com/
https://www.splunk.com/
monitoring


microservice tutorial
https://ewanvalentine.io/microservices-in-golang-part-1/


https://www.confluent.io/blog/apache-kafka-for-service-architectures/

service discovery:
http://callistaenterprise.se/blogg/teknik/2017/04/24/go-blog-series-part7/


https://dzone.com/articles/go-microservices-blog-series-part-1
http://callistaenterprise.se/blogg/teknik/2017/03/09/go-blog-series-part5/
http://callistaenterprise.se/blogg/teknik/2015/04/10/building-microservices-with-spring-cloud-and-netflix-oss-part-1/
http://callistaenterprise.se/blogg/teknik/2015/05/20/blog-series-building-microservices/



https://medium.freecodecamp.org/how-to-build-a-web-app-with-go-gin-and-react-cffdc473576



When we are dealing with Money, we don't want to use floats ever. It can lead to rounding errors since floats can't be represented exactly in some cases. We want to use the package "github.com/shopspring/decimal".. check out the repo for examples of where this is used.
You will want to use dec, err := decimal.NewFromString(amountStr) to get the percentage to a decimal type and then can perform calculations with that

https://socketloop.com/tutorials/golang-accurate-and-reliable-decimal-calculations


References:
react-golang-auth https://github.com/odewahn/react-golang-auth/
Password authentication and storage in Go (Golang) 
https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/
https://github.com/sohamkamani/go-password-auth-example


Why we use UUID as primary key?
https://www.clever-cloud.com/blog/engineering/2015/05/20/why-auto-increment-is-a-terrible-idea/
https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/

How to setup golang?
Just add the following lines to ~/.bashrc and this will persist. However, you can use other paths you like as GOPATH instead of $HOME/go in my sample.






