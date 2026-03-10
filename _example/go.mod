module example.com/log-mux-demo

go 1.26.0

replace github.com/taigrr/log-mux => ../

require (
	github.com/taigrr/log-mux v0.0.0-00010101000000-000000000000
	github.com/taigrr/log-socket/v2 v2.4.0
)

require github.com/gorilla/websocket v1.5.3 // indirect
