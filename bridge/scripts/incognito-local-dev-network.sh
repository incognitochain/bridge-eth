npm ci pm2
mkdir -p src && cd src
if [ ! -d incognito-highway ]
then
    git clone https://github.com/incognitochain/incognito-highway --branch local/v3-dcs
fi
cd incognito-highway
# go1.13 or above
go mod tidy
go build -o highway
npx pm2 start --name highway ./run.sh -- lc1
cd ..
if [ ! -d incognito-chain ]
then
    git clone https://github.com/incognitochain/incognito-chain --branch dev/papp-new-flow-local-cfg --depth 1
fi
cd incognito-chain
go mod tidy
go build -o incognito
npx pm2 start --name b0 ./run_node.sh -- beacon-0
npx pm2 start --name b1 ./run_node.sh -- beacon-1
npx pm2 start --name b2 ./run_node.sh -- beacon-2
npx pm2 start --name b3 ./run_node.sh -- beacon-3
npx pm2 start --name n00 ./run_node.sh -- shard0-0
npx pm2 start --name n01 ./run_node.sh -- shard0-1
npx pm2 start --name n02 ./run_node.sh -- shard0-2
npx pm2 start --name n03 ./run_node.sh -- shard0-3
npx pm2 start --name n10 ./run_node.sh -- shard1-0
npx pm2 start --name n11 ./run_node.sh -- shard1-1
npx pm2 start --name n12 ./run_node.sh -- shard1-2
npx pm2 start --name n13 ./run_node.sh -- shard1-3
npx pm2 start --name fn ./run_node.sh -- fullnode
npx pm2 list
cd ../..