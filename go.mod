module github.com/status-im/status-protocol-bandwidth-test

require (
	github.com/Microsoft/go-winio v0.4.12 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/dhui/dktest v0.3.1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/ethereum/go-ethereum v1.9.5
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/status-im/status-go v0.34.0-beta.5.0.20191105190644-bb5023d53c0f
	github.com/status-im/status-protocol-go v0.4.5-0.20191105190240-1c906e10ed65
	github.com/stretchr/objx v0.2.0 // indirect
	google.golang.org/genproto v0.0.0-20190701230453-710ae3a149df // indirect
	google.golang.org/grpc v1.22.0 // indirect
)

replace github.com/ethereum/go-ethereum v1.9.5 => github.com/status-im/go-ethereum v1.9.5-status.4

go 1.13

replace github.com/NaySoftware/go-fcm => github.com/status-im/go-fcm v1.0.0-status

replace github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.4.2

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20190717161051-705d9623b7c1
