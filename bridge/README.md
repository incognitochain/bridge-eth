# Incognito bridge contracts for EVM networks

```bash
npm install
```

## Live deployment
```bash
npx hardhat deploy --network <network> --tags vault >> deploy-out.log
# network can be localhost, kovan or mainnet
```

## Fork-mainnet Deployment
```bash
./forked.sh
# save deployment logs to review token holdings after vault upgrade
FORK=true npx hardhat deploy --network localhost --tags vault,trade >> deploy-out.log
```

## Local Deployment
```bash
# start Ethereum endpoint at 0.0.0.0:8545 & deploy all contracts locally, including test-only contracts
npx hardhat node
```

## Run tests (require Incognito)
```bash
# use local network; skip when using a live Incognito network
npm run local-network; sleep 30; npm run start:dev # also setup funds

# run test files (or all tests if none specified)
npx hardhat test --network localhost [ <test-file-name> ... ]
```