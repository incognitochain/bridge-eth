# set up your test private key & API endpoint if needed
# printf '%s' 00000000000000000000000000000000000000000000000000000000000000af > .deployerPrivateKey
# printf '%s' 00000000000000000000000000000000000000000000000000000000000000af > .vaultAdminKey
# printf '%s' zcQtm0rv46R4WOT7X8CQaMnAvFXtZZzz > .alchemyKey
# copy previous mainnet contract deployments
pushd deployments
rm -r localhost
# can choose which "mainnet" network to fork: mainnet/bscmainnet/plgmainnet/ftmmainnet
cp -r mainnet/ localhost/
printf '%s' 31337 > localhost/.chainId
popd

# use --no-reset to disable auto-clearing of deployments to read existing mainnet contracts that were copied
# can add --fork https://eth-mainnet.g.alchemy.com/v2/zcQtm0rv46R4WOT7X8CQaMnAvFXtH6ov # these are already specified in Hardhat config
# FORK options: eth/bsc/plg/ftm
FORK=eth npx hardhat node --tags fork,vault,upgrade --no-reset --hostname 0.0.0.0 --port 8545 

# manually check vault state after deploying on fork mainnet
# hh list-contracts --network localhost; hh show-committees --network localhost; hh query-var --name regulator --network localhost; hh query-var --name storageLayoutVersion --network localhost; hh query-var --name executor --network localhost