# Dependencies management
[https://blog.golang.org/using-go-modules#TOC_4.]
```shell
# Dependencies list
go list -m all

# add/upgrade library
go get github.com/aws/aws-sdk-go
```

# Build
```shell
go build bitbucket.org/y16i/backend-go

# To set breakpoint
go build -gcflags="all=-N -l"

# Run
./backend-go

# Watch
watcher backend-go
```

