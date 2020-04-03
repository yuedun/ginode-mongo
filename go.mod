module github.com/yuedun/ginode-mongo

go 1.14

require (
	github.com/appleboy/gin-jwt/v2 v2.6.3
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/klauspost/compress v1.10.3 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
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
