module github.com/yuedun/ginode

go 1.12

require (
	cloud.google.com/go v0.41.0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190710001350-29e7b2419f38 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542 // indirect
)

replace golang.org/x/net v0.0.0-20190628185345-da137c7871d7 => github.com/golang/net v0.0.0-20190628185345-da137c7871d7

replace golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4

replace golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542 => github.com/golang/sys v0.0.0-20190710143415-6ec70d6a5542
