package main

import (
	"encoding/hex"

	ec "github.com/ethereum/go-ethereum/common"
)

// getCommitteeHardcoded is for deploying scripts
func getCommitteeHardcoded() *committees {
	beaconComm := []string{
		"0x3cD69B1A595B7A9589391538d29ee7663326e4d3",
		"0xc687470342f4E80ECEf6bBd25e276266d40b8429",
		"0x2A40c96b41AdEc5641F28eF923e270B73e29bb53",
		"0x131B772A9ADe1793F000024eAb23b77bEd3BFe64",
	}
	bridgeComm := []string{
		"0x28655822DAf6c4B32303B06e875F92dC6e242cE4",
		"0xD2902ab2F5dF2b17C5A5aa380f511F04a2542E10",
		"0xB67376ad63EAdC22f05efE428e93f09D4f13B4fD",
		"0x40bAA64EAFbD355f5427d127979f377cfA48cc10",
	}
	beacons, bridges := toAddresses(beaconComm, bridgeComm)
	return &committees{
		beacons: beacons,
		bridges: bridges,
	}
}

// getFixedCommittee is for unittest
func getFixedCommittee() *committees {
	beaconCommPrivs := []string{
		"5a417f54357fff96fe4c2a9cafd322ed72b52bf046beb69a9730a26181088489",
		"b9cd32581922f447acb1cfd148069fc40cbbce1e8badb84c4b509486e6f713ce",
		"22e23970b853407e16ccb174443f27c37dbbea05729aba546ee649e0aef2d9cb",
		"4d16dadc89656fbda140e2fe467631ddac3ed9cc326cef3a8f1b1bd5f3cfd155",
	}
	beaconComm := []string{
		"0xA5301a0d25103967bf0e29db1576cba3408fD9bB",
		"0x9BC0faE7BB432828759B6e391e0cC99995057791",
		"0x6cbc2937FEe477bbda360A842EeEbF92c2FAb613",
		"0xcabF3DB93eB48a61d41486AcC9281B6240411403",
	}
	beaconPrivs := make([][]byte, len(beaconCommPrivs))
	for i, p := range beaconCommPrivs {
		priv, _ := hex.DecodeString(p)
		beaconPrivs[i] = priv
	}

	bridgeComm := []string{
		"0x3c78124783E8e39D1E084FdDD0E097334ba2D945",
		"0x76E34d8a527961286E55532620Af5b84F3C6538F",
		"0x68686dB6874588D2404155D00A73F82a50FDd190",
		"0x1533ac4d2922C150551f2F5dc2b0c1eDE382b890",
	}
	bridgeCommPrivs := []string{
		"3560e649ce326a2eb9fbb59fba4b29e10fb064627f61487aecc8b92afbb127dd",
		"b71af1a7e2ca74400187cbf2333ab1f20e9b39517347fb655ffa309d1b51b2b0",
		"07f91f98513c203103f8d44683ce47920d1aea0eaf1cb86a373be835374d1490",
		"7412e24d4ac1796866c44a0d5b966f8db1c3022bba8afd370a09dc49a14efeb4",
	}

	bridgePrivs := make([][]byte, len(bridgeCommPrivs))
	for i, p := range bridgeCommPrivs {
		priv, _ := hex.DecodeString(p)
		bridgePrivs[i] = priv
	}

	beacons, bridges := toAddresses(beaconComm, bridgeComm)
	return &committees{
		beacons:     beacons,
		beaconPrivs: beaconPrivs,
		bridges:     bridges,
		bridgePrivs: bridgePrivs,
	}
}

func toAddresses(beaconComm, bridgeComm []string) ([]ec.Address, []ec.Address) {
	beacons := make([]ec.Address, len(beaconComm))
	for i, p := range beaconComm {
		beacons[i] = ec.HexToAddress(p)
	}

	bridges := make([]ec.Address, len(bridgeComm))
	for i, p := range bridgeComm {
		bridges[i] = ec.HexToAddress(p)
	}
	return beacons, bridges
}
