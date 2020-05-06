package main

import (
	"encoding/hex"

	ec "github.com/ethereum/go-ethereum/common"
)

// getCommitteeHardcoded is for deploying scripts
func getCommitteeHardcoded() *committees {
	beaconComm := []string{
		"0x77b9F481f979e16Cf4234866C0803c2e65106862",
		"0x8F4aEd5Adb0347eF2Db3bDce9f0DF2747D0107E8",
		"0xcf836142D459B1257ed52aef34E62F6F4e7eF894",
		"0x4C265F5eb8Eb68A1eA99626cD838558836438f80",
		"0xc7CC1EE53eF28CC551e97476F0dB01596D945fE0",
		"0xFd119B9DBEb478154E3650F14f55db3787E1bd38",
		"0x780329f064F0BE00FdbDeA4bA8A5b04C7AB2866c",
	}
	bridgeComm := []string{
		"0xaced3cf99897a55057B7513b8505c22DaF9378D9",
		"0x8CbDC490c4188b721e210622027a54E14f27CA7F",
		"0x74d0C0A0A5f89527b4e2850EA09d4F6cE9BBb3bB",
		"0x6964D5c5A7C1503E2228852d1EC115c0e7a20593",
		"0x803c90C23a8a34a676B57CaF0372026C988B416d",
		"0x17E21A7a018046ab3cAE7Aab4215087a2497a7D7",
		"0x5bA8281b5BE1F864E52B3ef8FcEF80560e41005C",
		"0x9Ee3002A85701ae62B16e92e0d8F2044D79a35B6",
		"0xb7eF123cc555cA977aE2fbB5A3690ce57628C664",
		"0x211880118421A814Da0151A4bd06be703DB3654e",
		"0x6A25Ed4Ef6Fa034c895D5721D73dBEC5163Ad4f9",
		"0x1A7232f56F4D71e794D8Bbfc5fa5991d544e1C9f",
		"0xcFcFc3A2CC9DFdF98aC075441E45818C7A70a29e",
		"0x99118446796dFa58d8327834347806711f67Cb79",
		"0x8986acdde31E4519acFcabb139Fd2A2B1da274b2",
		"0xE59C59D87f52D39B1BB8136966e0E1817D7a845A",
		"0x604589220D909878ebDC906d0b33b433Fc3cc0a3",
		"0xf069494c92A85DD31FE6850D8EfE6F2398Ea072c",
		"0x93E1b517726d05c235AE3AF53fa84C34d400Cae9",
		"0x6284C7FD0F623E705d0e0a2D4621299B98eA3895",
		"0xf57Ac7832b1C8F7f5C3E228eF7811D58647A70BF",
		"0x8fa98CBa06b199922E9Acc5749F25FF549e5eEbd",
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
