# copy previous mainnet contract deployments
pushd deployments
rm -r localhost
cp -r mainnet/ localhost/
printf '%s' 31337 > localhost/.chainId
popd
# --no-reset to disable auto-clearing of deployments to read existing mainnet contracts that were copied
FORK=true npx hardhat node --show-accounts --no-deploy --no-reset 
# --fork https://eth-mainnet.alchemyapi.io/v2/... --fork-block-number 12000000 # these are already specified in Hardhat config