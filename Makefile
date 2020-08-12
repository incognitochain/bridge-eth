all: beacon bridge burn erc20

beacon: build
	go test -run=TestSimulatedSwapBeacon

bridge: build
	go test -run=TestSimulatedSwapBridge

burn: build
	go test -run=TestSimulatedBurn

erc20: build
	go test -run=TestSimulatedErc20

build: bridge/incognito_proxy/incognito_proxy.go bridge/vault/vault.go bridge/locker/locker.go bridge/pause/pause.go erc20/fail/fail.go erc20/dless/dless.go bridge/kbntrade/kbntrade.go bridge/kbnmultitrade/kbnmultitrade.go bridge/reentrancer/reentrancer.go

.PHONY: all beacon bridge burn erc20 build

erc20/ERC20.go: erc20/ERC20.vy
	./gengo.sh erc20/ERC20.vy erc20

bridge/incognito_proxy/incognito_proxy.go: bridge/contracts/incognito_proxy.sol
	./gengo.sh bridge/contracts/incognito_proxy.sol bridge/incognito_proxy

bridge/vault/vault.go: bridge/contracts/vault.sol
	./gengo.sh bridge/contracts/vault.sol bridge/vault

bridge/locker/locker.go: bridge/contracts/locker.sol
	./gengo.sh bridge/contracts/locker.sol bridge/locker

bridge/kbntrade/kbntrade.go: bridge/contracts/kbn_trade.sol
	./gengo.sh bridge/contracts/kbn_trade.sol bridge/kbntrade

bridge/kbnmultitrade/kbnmultitrade.go: bridge/contracts/simpleMultiKnbTrade.sol
	./gengo.sh bridge/contracts/simpleMultiKnbTrade.sol bridge/kbnmultitrade

bridge/pause/pause.go: bridge/contracts/pause.sol
	./gengo.sh bridge/contracts/pause.sol bridge/pause

bridge/reentrancer/reentrancer.go: bridge/contracts/reentrancer.sol
	./gengo.sh bridge/contracts/reentrancer.sol bridge/reentrancer

erc20/fail/fail.go: erc20/fail/fail.sol
	./gengo.sh erc20/fail/fail.sol erc20/fail

erc20/dless/dless.go: erc20/dless/dless.sol
	./gengo.sh erc20/dless/dless.sol erc20/dless

