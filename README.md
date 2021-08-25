### GoBlog
 ```
This is a blog write by go 
```

###### usage
```shell
git clone https://github.com/attitudefx7/goblog.git

cd goblog

cp .env.example .env

go build main.go

nohup ./main >/tmp/goblog.log 2>&1 &
```

