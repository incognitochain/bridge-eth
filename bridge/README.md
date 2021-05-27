ETH bridge contracts

# Live deployment
```bash
npx hardhat deploy --network <network> --tags vault >> deploy-out.log
# network can be localhost, kovan or mainnet
```

# Fork-mainnet Deployment
```bash
./forked.sh
# save deployment logs to review token holdings after vault upgrade
FORK=true npx hardhat deploy --network localhost --tags vault,trade >> deploy-out.log
```

# Local Deployment
````bash
# start Ethereum endpoint at 0.0.0.0:8545 & deploy all contracts locally, including test-only contracts
# also start & connect Incognito Dev network
npx hardhat node --hostname 0.0.0.0
# in another shell, run Hardhat tests (require Incognito WebJS V2)
npx hardhat test --network localhost
```
