# Qudash

![](https://github.com/zillani/img/blob/master/go/pacman.gif)
## Table of Contents
1. [Pre-requisites](#Pre-requisites)
2. [Installation](#Installation)
   1. [Local](#local)
   2. [Docker](#docker)
   3. [K8s](#k8s)
3. [Access](#Access)
4. [Developer-guide](#Developer-guide)
    1. [Bee-cli](#Bee-cli)
    2. [Beego from scratch](#Beego-from-scratch)
    3. [Modules](#modules)
        1. [Convert to module](#convert-to-module)
        2. [vendor](#vendor)
    4. [Docker](#docker)
    5. [k8s](#k8s)
        1. [kompose](#kompose)
    6. [Issues](#Issues)
       1. [import redis as sideeffect](#import-redis-as-sideeffect)
       2. [unable to find schema on goland](#unable-to-find-schema-on-goland)
       3. [post not working](#post-not-working)
       4. [issue with dependencies](#issue-with-dependencies)
       5. [derekparker delve issue](#derekparker-delve-issue)
       6. [Port already running?](#Port-already-running?)
       9. [annotations problem](#annotations-problem)
5. [Misc](#Misc)
   1. [CLI guide](#CLI-guide)
   2. [Run](#Run)
   3. [db2struct](#db2struct)
   4. [yaml to json](#yaml-to-json)
   5. [convert json to go struct](#convert-json-to-go-struct)
   6. [generate swagger api](#generate-swagger-api)
   7. [disable go modules while build](#disable-go-modules-while-build)
   8. [how to read params](#how-to-read-params)
   10. [session guide](#session-guide)
   11. [logout from all devices?](#logout-from-all-devices?)
   12. [Disable template render ](#Disable-template-render )
   13. [HTML tags not recognized?](#HTML-tags-not-recognized?)
   14. [Install via docker](#Install-via-docker)
   15. [access-mysql](#access-mysql)
   16. [Router not loading](#router-not-loading)
   17. [fixing the mess](#fixing-the-mess)


### Pre-requisites
- `redis` should be running & accessible from port `6379`  [install-redis-for-windows](#https://github.com/dmajkic/redis/downloads)
- `mysql` should be running & accessible from port `3306`

### Installation
#### Local
Install locally using the command below, make sure
```bash
go build
go mod download
GO111MODULE=off go get github.com/beego/bee
bee run --gendoc=true --downdoc=true
```
#### Docker
```bash
docker run -d zshaik/zillani/qudash -p 8080:8080
```
#### k8s
// to be added

### Access
For chat app, you can access on, 
```bash
localhost:8080
```
For swagger,
```bash
localhost:8080/swagger
```
For admin,
```bash
localhost:8088
```

### Developer guide
#### Bee cli
```bash
go get -u github.com/beego/bee
```
_note: For some reason, the `go get` is only working from goland terminal_

#### Beego from scratch
For fullstack projects, use
```bash
bee new
``` 
for APIs, this the best,
```bash
bee api myapp
```

### Modules

#### Convert to module
[switch-to-gomod](https://vsupalov.com/switch-to-go-modules/)
```bash
go mod init github.com/zillani/qudash
go get
```
__How to use dockerfile with go mod__
[medium](https://medium.com/@petomalina/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888)

__Dockerize go app__
[dockerize-go](#https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2)
[awesome-compose](https://github.com/docker/awesome-compose)

#### vendor 
Creates all dependencies under vendor directory
```bash
go mod vendor
```

#### CLI guide
[beego-cli](https://beego.me/docs/quickstart/new.md)
[using docker](https://ncona.com/2017/10/introduction-to-beego/)

#### Run
copy the files to `%GOPATH%/src/github.com`
```bash
go run main.go
```
or you can simply do,
```bash
bee run --gendoc=true --downdoc=true
```
you can access the application on port 8080
and admin on 8088, these are configured under `conf/app.cnf`

```bash
curl http://localhost:8080
curl http://localhost:8088
```

#### db2struct
just use `util/db2Struct.go`

#### yaml to json
[yaml2json](https://www.json2yaml.com/convert-yaml-to-json)

#### convert json to go struct
[jsontogo](https://mholt.github.io/json-to-go/)
[jsontogo-github](https://github.com/mholt/json-to-go)

#### generate swagger api
Hint: first create api app using beego, then copy paste annotations on controllers appropriately,
then generate this doc.
```bash
bee run -downdoc=true -gendoc=true
```

#### disable go modules while build
```bash
GO111MODULE=off go build .
```

#### how to read params
[read-params](https://beego.me/docs/mvc/controller/params.md)

#### annotations problem
If you are doing api development like `v1/user` or `v1beta/user` then every method in the controller 
should have the annotations, check `objectcontroller` for reference 
Check Annotation for Controller section below,
[this is very important](https://beego.me/blog/beego_api)

#### session guide
[sessions-beego](https://mrwaggel.be/post/golang-beego-how-to-use-sessions/)

#### logout from all devices?
check if you can just clear cookie for a particular session
[cookies](https://thinkingeek.com/2018/05/31/setting-and-deleting-cookies-in-go/)

#### Disable template render 
`autorender = false`

#### HTML tags not recognized?
remove react-templates plugin

#### Install via docker
```bash
docker-compose up .
```
then do, 
docker build .

### Docker

#### access-mysql
Get docker machine ip and access it,
```bash
docker-machine ip
mysql -h 192.168.99.100 -P 8983 --protocol=tcp -u zshaik -proot
```
#### k8s
#### kompose
```bash
kompose.exe --file docker-compose.yml convert -o k8s
```
### Issues

#### import redis as sideeffect
This is required,
```bash
_ "github.com/astaxie/beego/session/redis"
```
#### unable to find schema on goland
On goland if you are unable to find schema, on the database dialect menu, 
just click on `+` icon and select cli from the cli, you can select the schema

#### post not working
need to enable `CopyRequestBody = true` in `conf`
[beego-issues](https://github.com/astaxie/beego/issues/2938)

#### issue with dependencies
If there is any kind of issue with dependencies, just delete `vendor`, `go.sum`, `go.mod`
and run by disabling go modules like below,

```bash
GO111MODULE=off go build .
```
#### derekparker delve issue
```bash
rm -rf go.mod
go mod init github.com/zillani/qudash
go get github.com/go-delve/delve/cmd/dlv
```
and replace with, 
```bash
github.com/go-delve/delve => github.com/derekparker/delve v1.3.2 // indirect
```
OR
simply delete bee import in main

#### Port already running?
For windows users,
```bash
netstat -ano | findstr :8080
taskkill /PID 20380 /F
```

#### Routers not loading
Check if the main method has this, 
```bash
	_ "github.com/zillani/qudash/routers"
```

#### fixing the mess
If the whole project goes a mess! then,
```bash
create the api project using bee and copy file one by one 
from old project to the new project
```

#### why go modules & how?
go modules creates all your deps under `$GOPATH/pkg` just like `.m2` fo maven
so you will files will index to that location instead of `$GOPATH/src`
```bash
go mod init
go mod downlaod
go get
```