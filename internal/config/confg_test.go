package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig_ShouldUnmarshalConfigFileAndReturnInstance(t *testing.T){
	config := NewConfig("test_config.json")
	assert.Len(t,config.GetSections(),1)
	assert.Equal(t,config.GetSections(),[]Section{
		{
			Name: "A",
			NumberOfSeats: 10,
		},
	})
	assert.Equal(t,float32(20),config.GetTicketCost())
}