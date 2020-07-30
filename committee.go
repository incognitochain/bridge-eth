package main

import (
	"encoding/hex"

	ec "github.com/ethereum/go-ethereum/common"
)

// getCommitteeHardcoded is for deploying scripts
func getCommitteeHardcoded() *committees {
	beaconComm := []string{
		"0xD7d93b7fa42b60b6076f3017fCA99b69257A912D",
		"0xf25ee30cfed2d2768C51A6Eb6787890C1c364cA4",
		"0x0D8c517557f3edE116988DD7EC0bAF83b96fe0Cb",
		"0xc225fcd5CE8Ad42863182Ab71acb6abD9C4ddCbE",
		"0x785164C41DA00C4E463fe35Ace1E2d2A95332d5d",
		"0xcF85b4F7FF4e92818756937F594AC77dC49d66A0",
		"0xF40857dA979984f51067911D4A67Dd1a2e77bA62",
	}
	bridgeComm := []string{
		"0xC81Fc20A03F1B33fd21ED9D35B24266adbAA2769",
		"0xC932B446033649d422425EB486A924c7F85D6351",
		"0x55AbF902256357059e175A255E82D9A1Df6eBB47",
		"0x97AE27DB0a9AEf131C11Ec9C626E68c8Cb294537",
		"0xc9A69F8496E4da75E5f7e09eaA43ecf010FD48A2",
		"0x0C98c326184adDE224d4883d3336559D5d694374",
		"0x6945d7e7C516b6eBa12E15472EA62c38b5B7F82b",
		"0x5f5BC3f5D6d05363Fa80cC0Ba07dFDd8A912f3ff",
		"0x6272080D1904069b5E8eBD813C0464F125854147",
		"0xb53f9552Ed1Ce9538D42C0b4C1863FD1568B8798",
		"0x712FdBCF0BebB4aB1d5B58A60035057AF426e77C",
		"0x1110eAdF442b9dE4Ff91c917eD8DfbDBc0Cbe479",
		"0xd738f2DCFBDa6D2fF031d0868700bd2DfB438b2D",
		"0x4e6B4EE11fc5f235F142b320A5C4218cF1B2871d",
		"0xd4837977fEBF24BEa3A179e52710509eAb2eEd61",
		"0xE60d61cA58Cd138BaC6Ff645b4d986E7474D55c6",
		"0x7789735cAB631E039E42425d71840076574bA9F6",
		"0x8C73B75d50Fe9a88102ECEc646F03B9D1E3042cd",
		"0x1Bb9C495C90d4D6FdddfFde18b9048D12D702c48",
		"0xD05aFd28b70d1326A4472BA4446915C62143C16c",
		"0x9f3747d34EAF789a84AaE82754Da2fcE4E89395F",
		"0xC81F55dB612E03b79a10088de34d0951f10575d8",
		"0xf2F1612147beF1bB70BAB25896D4A36c0652e955",
		"0x6CB80CaA3F18c9b781a11Dba65b3C4703c849907",
		"0x5C5431b45e586EfCe1Fb0912E7A1576C724B256e",
		"0x754F68FDaB4f1baF105A0187e07C9a28f83E3EAb",
		"0x583b800d8a4D064061b298fbA034AB1B0a300402",
		"0x4656D963F1A7bc7743F5c7e51B19e887ad153F8e",
		"0x8454ACaf3be55795dD01c6A6E61E23B817d0527c",
		"0xfE690FAcc4489df6D3540233B78dCa30A3b40078",
		"0x48D6929408bC5a3393bc17F9646bd74a8d904719",
		"0x78c307266BFaB059Bf9ac93207135fF58e3698E5",
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
