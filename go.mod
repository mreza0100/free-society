module microServiceBoilerplate

go 1.15

replace github.com/mreza0100/golog => ../github.com/mreza0100/golog

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/mreza0100/golog v0.0.0-20210127200816-11d3354e8a97
	github.com/nats-io/nats-server/v2 v2.1.9 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/vektah/gqlparser/v2 v2.1.0
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.20.12
)
