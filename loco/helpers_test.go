package loco

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HelpersSuite struct {
	suite.Suite
}

func TestHelpersSuite(t *testing.T) {
	suite.Run(t, new(HelpersSuite))
}

func (s *HelpersSuite) SetupTest() {
}

func (s *HelpersSuite) SetupSuite() {
}

func (s *HelpersSuite) TestGenerateGeohashFromLatLng() {
	geohash := GenerateGeohashFromLatLng(-25.382708, -49.265506)
	assert.Equal(s.T(), -7720041938784775475, geohash)
}

func (s *HelpersSuite) TestGenerateHashKey() {
	hashkey := GenerateHashKey(-7720041938784775475, 5)
	assert.Equal(s.T(), -77200, hashkey)
}
