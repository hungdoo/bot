# Run

```
BOT_ENV="dev|prod" go run main.go

or 

env GOOS=target-OS GOARCH=target-architecture go build main.go && BOT_ENV="dev|prod" ./main

```
```
GOOS - Target Operating System	GOARCH - Target Platform
android	          arm
darwin	          386
darwin	          amd64
darwin	          arm
darwin	          arm64
dragonfly	          amd64
freebsd	          386
freebsd	          amd64
freebsd	          arm
linux	          386
linux	          amd64
linux	          arm
linux	          arm64
linux	          ppc64
linux	          ppc64le
linux	          mips
linux	          mipsle
linux	          mips64
linux	          mips64le
netbsd	          386
netbsd	          amd64
netbsd	          arm
openbsd	          386
openbsd	          amd64
openbsd	          arm
plan9	          386
plan9	          amd64
solaris	          amd64
windows	          386
windows	          amd64
```

# Example commands
```
rpc, contractAddr, method, params<;>, valueIdx, marginStr, precisionStr

// Track SOLID price against oxSOLID
add callContract-solid https://rpc.ftm.tools/ 0xa38cd27185a464914d3046f0ab9d43356b34829d getAmountOut(uint256,address,address)(uint256,bool) 1000000000000000000000;0x888ef71766ca594ded1f0fa3ae64ed2941740a20;0xda0053f0befcbcac208a3f867bb243716734d809 0 1 18

// Track Geist health
add callContract-geist https://rpc.ankr.com/fantom 0x9FAD24f572045c7869117160A571B2e50b10d068 getUserAccountData(address)(uint256,uint256,uint256,uint256,uint256,uint256) 0x9BD3e72f2F1CA05Ad8D4489ec870Bf2478b10397 5 10 18
```