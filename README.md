This application is a partial replacement of WordPress API. It takes a page content from WordPress database directly.

Note: This application supports only GET request, and a "page" post_type of WordPress. I intended to use this behind a web server. This program doesn't handle CORS and pre-flight (OPTION) response.

# Before build
This application reuse DB_NAME, DB_USER, DB_PASSWORD and DB_HOST in wp-config.php. You need to change wpConfigPath in main.go
```script
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

# Test
go test ./...
```

# API request sample
GET /api/v1.0/pages?slug=page-slug
Response {title: {rendered: "Page Title"}, content: {rendered: "Page Content"}}

# apache2 proxy_http sample
e.g.)
API request
	http://custom-domain.com/api/v1.0/pages?slug=summary

site config
```bash
 ProxyPass /api/v1.0 http://localhost:8080/api/v1.0
 ProxyPassReverse /api/v1.0 http://localhost:8080/api/v1.0
```
