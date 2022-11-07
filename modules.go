package trojan

import (
	_ "github.com/pengcu/caddy-trojan/admin"
	_ "github.com/pengcu/caddy-trojan/app"
	_ "github.com/pengcu/caddy-trojan/handler"
	_ "github.com/pengcu/caddy-trojan/listener"
	_ "github.com/pengcu/caddy-trojan/websocket"
	_ "github.com/pengcu/caddy-trojan/grpc"
)
