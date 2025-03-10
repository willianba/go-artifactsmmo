package actions_test

import (
	"testing"

	"github.com/0xN0x/go-artifactsmmo/models"
	"github.com/0xN0x/go-artifactsmmo/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccountBankSuite struct {
	suite.Suite
}

func (suite *AccountBankSuite) TestShouldGetBankDetails() {
	client := test.CreateMockedClient(&models.Bank{
		Slots:             10,
		Expansions:        1,
		NextExpansionCost: 100,
		Gold:              1000,
	})

	bankDetails, _ := client.GetBankDetails()

	assert.Equal(suite.T(), 10, bankDetails.Slots)
	assert.Equal(suite.T(), 1, bankDetails.Expansions)
	assert.Equal(suite.T(), 100, bankDetails.NextExpansionCost)
	assert.Equal(suite.T(), 1000, bankDetails.Gold)
}

func (suite *AccountBankSuite) TestShouldReturnErrorIfRequestIsMalformed() {
	client := test.CreateInvalidMockedClient(452, string(models.ErrBadToken))

	bankDetails, err := client.GetBankDetails()

	assert.Nil(suite.T(), bankDetails)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), string(models.ErrBadToken), err.Error())
}

func (suite *AccountBankSuite) TestShouldReturnBankItems() {
	client := test.CreateMockedClient(&[]models.SimpleItem{
		{
			Code:     "copper",
			Quantity: 5,
		},
	})

	bankItems, _ := client.GetBankItems("copper", 0, 0)

	assert.Equal(suite.T(), 1, len(*bankItems))
	assert.Equal(suite.T(), "copper", (*bankItems)[0].Code)
	assert.Equal(suite.T(), 5, (*bankItems)[0].Quantity)
}

func TestAccountBankSuite(t *testing.T) {
	suite.Run(t, new(AccountBankSuite))
}
