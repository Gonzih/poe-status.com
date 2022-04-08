// +heroku goVersion go1.17
// +heroku install ./cmd/poe-status-server/

module github.com/Gonzih/poe-status.com

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.5
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tomsteele/go-nmap v0.0.0-20191202052157-3507e0b03523
	github.com/twitchtv/twirp v8.1.2+incompatible
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

go 1.16
