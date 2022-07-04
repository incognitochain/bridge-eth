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
		"d6454ea0c91a0ad06471e240e00ba5b2de4dad8cfe9f786aec67a5a29750cd0b",
		0,
		3,
	)
	log.Println("proof, err", proof, err)
	t.Nil(err)
}
