
package forex

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestForex(t *testing.T) {
	stub := shim.NewMockStub("forex", new(ForexChaincode))
	uuid := uuid.New().String()

	const RATE = "0.88"
	const BASECURRENCY = "USD"
	const COUNTERCURRENCY = "GBP"

	stringArgs := []string{"createUpdateForexPair", BASECURRENCY, COUNTERCURRENCY, RATE}
	args := util.ArrayToChaincodeArgs(stringArgs)

	response := stub.MockInvoke(uuid, args)

	assert.EqualValues(t, shim.OK, response.GetStatus(), "failed to execute invocation")

	stringArgs = []string{"getForexPair", BASECURRENCY, COUNTERCURRENCY}
	args = util.ArrayToChaincodeArgs(stringArgs)

	response = stub.MockInvoke(uuid, args)
	assert.EqualValues(t, shim.OK, response.GetStatus(), "failed to execute invocation")

	responseForex := &forex{}
	err := json.Unmarshal(response.GetPayload(), responseForex)

	if err != nil {
		panic(err)
	}

	rateAsFloat, _ := strconv.ParseFloat(RATE, 64)

	assert.Equal(t, rateAsFloat, responseForex.Rate, "rate mismatch")

}
