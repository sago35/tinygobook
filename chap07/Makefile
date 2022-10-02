update_test: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

access_test: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

init_test: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

tcp-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

udp-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

ntp: FORCE
	tinygo flash -target wioterminal $(ARGS) tinygo.org/x/drivers/examples/rtl8720dn/ntpclient

mqttpub-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

mqttsub-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

http-client-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

http-server-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

https-client-tinygo: FORCE
	tinygo flash -target wioterminal $(ARGS) ./$@

smoketest: FORCE
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) tinygo.org/x/drivers/examples/rtl8720dn/ntpclient
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./access_test
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./http-client-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./http-server-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./https-client-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./init_test
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./mqttpub-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./mqttsub-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./tcp-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./udp-tinygo
	tinygo build -o /tmp/out.uf2 -target wioterminal $(ARGS) ./update_test
	go build -o /tmp/out ./http-client-go
	go build -o /tmp/out ./http-server-go
	go build -o /tmp/out ./mqttpub-go
	go build -o /tmp/out ./mqttsub-go
	go build -o /tmp/out ./tcp-go/server.go
	go build -o /tmp/out ./tcp-go/client.go
	go build -o /tmp/out ./udp-go/server.go
	go build -o /tmp/out ./udp-go/client.go

FORCE:
