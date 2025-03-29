## Commands

```
solc --standard-json ./LoadTest.input.json | jq '.contracts["LoadTest.sol"].LoadTest.abi' > LoadTest.abi


solc --standard-json ./LoadTest.input.json | jq '.contracts["LoadTest.sol"].LoadTest.abi' > LoadTest.abi
solc --standard-json ./LoadTest.input.json | jq -r '.contracts["LoadTest.sol"].LoadTest.evm.bytecode.object' > LoadTest.bin
abigen --bin=./LoadTest.bin --abi=./LoadTest.abi --pkg=contract --out=LoadTest.go
```