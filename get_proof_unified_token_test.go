package bridge

import (

	// "github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type UnifiedTokenTestSuite struct {
	suite.Suite
}

func (t *UnifiedTokenTestSuite) SetupTest() {

}

func TestUnifiedTokenTestSuite(t *testing.T) {
	suite.Run(t, new(UnifiedTokenTestSuite))
}

func (t *UnifiedTokenTestSuite) TestGetAndDecodeBurnProofUnifiedToken() {

	proof, err := GetAndDecodeBurnProofUnifiedToken(
		"http://51.89.21.38:11334",
		"c56d4a8233495ca1df64fe18868ac73c6f2960faf54e55d50b548b9800988585",
		0,
		3,
	)
	log.Println("proof, err", proof, err)
	t.Nil(err)
}
