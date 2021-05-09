This application is a partial replacement of WordPress API. It takes a page content from WordPress database directly.

It supports only GET request for a page post_type.

# Before build
This application reuse DB_NAME, DB_USER, DB_PASSWORD and DB_HOST in wp-config.php. You need to change wpConfigPath in main.go
```go
wpConfigPath string = "/var/www/portfolio/wp/wp-config.php"
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

# API Request
GET /api/v1.0/pages?slug=page-slug
Response {title: "Page Title", content: "Page Content"}

