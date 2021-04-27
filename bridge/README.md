ETH bridge contracts

# Live deployment
```bash
RESTORE=4.5 npx hardhat deploy --network <network> --tags vault >> deploy-out.log
# network can be localhost, kovan or mainnet
```

# Fork-mainnet Deployment
```bash
./forked.sh
# save deployment logs to review token holdings after vault upgrade
FORK=true npx hardhat deploy --network localhost --tags vault,trade >> deploy-out.log

# FORK=true npx hardhat deploy --network localhost --tags dev-chain # start DEV chain instance
# FORK=true npx hardhat test --network localhost # run Hardhat tests (require Incognito WebJS V2). When testing in non-local environments, edit the providers in Hardhat config (providers are needed when running Hardhat tests, not when deploying)
```