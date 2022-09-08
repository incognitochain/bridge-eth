# copy previous mainnet contract deployments
pushd deployments
rm -r localhost
cp -r bsctest/ localhost/
printf '%s' 31337 > localhost/.chainId
popd
# --no-reset to disable auto-clearing of deployments to read existing mainnet contracts that were copied
npx hardhat node --no-deploy --no-reset --hostname 0.0.0.0 --port 8545
# --fork https://eth-mainnet.alchemyapi.io/v2/... --fork-block-number 12000000 # these are already specified in Hardhat config