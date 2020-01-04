module github.com/yuedun/ginode

go 1.13

require (
	cloud.google.com/go v0.41.0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190710001350-29e7b2419f38 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/sys v0.0.0-20200103143344-a1369afcdac7 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.39.1-0.20190528154449-61166ef30553

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a

	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522

	golang.org/x/image => github.com/golang/image v0.0.0-20191214001246-9130b4cfad52

	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422

	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20191210151939-1a1fef82734d

	golang.org/x/net => github.com/golang/net v0.0.0-20190318221613-d196dffd7c2b

	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190523182746-aaccbc9213b0

	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6

	golang.org/x/sys => github.com/golang/sys v0.0.0-20190318195719-6c81ef8f67ca

	golang.org/x/text => github.com/golang/text v0.3.0

	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4

	golang.org/x/tools => github.com/golang/tools v0.0.0-20190529010454-aa71c3f32488

	google.golang.org/api => github.com/googleapis/google-api-go-client v0.15.0

	google.golang.org/appengine => github.com/golang/appengine v1.6.1-0.20190515044707-311d3c5cf937

	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190522204451-c2c4e71fbf69

	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.0

)
