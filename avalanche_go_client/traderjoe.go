package avalanche_go

import (
	"avalanche_go_client/codec"
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TraderJoeClient struct {
	pk         string
	lbPair     *codec.EvmContractCodec
	lbRouter   *codec.EvmContractCodec
	myAddr     common.Address
	routerAddr common.Address
}

func NewTraderJoeClient(pk string) (*TraderJoeClient, error) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		return nil, err
	}

	json, err := os.ReadFile("./abi/LBpair.json")
	if err != nil {
		return nil, err
	}
	lbPairABIAbi, err := abi.JSON(bytes.NewReader(json))
	if err != nil {
		return nil, err
	}

	lbPariCodec := codec.NewEvmCodec(client, common.HexToAddress("0x864d4e5Ee7318e97483DB7EB0912E09F161516EA"), &lbPairABIAbi)

	json, err = os.ReadFile("./abi/LBrouter.json")
	if err != nil {
		return nil, err
	}
	lbRouterABIAbi, err := abi.JSON(bytes.NewReader(json))
	if err != nil {
		return nil, err
	}
	lbRouterCodec := codec.NewEvmCodec(client, common.HexToAddress("0x18556DA13313f3532c54711497A8FedAC273220E"), &lbRouterABIAbi)

	// Read pk from os variable
	// pk := os.Getenv("PK")

	return &TraderJoeClient{
		pk:         pk,
		lbPair:     lbPariCodec,
		lbRouter:   lbRouterCodec,
		myAddr:     common.HexToAddress("0xb4dd4fb3D4bCED984cce972991fB100488b59223"),
		routerAddr: common.HexToAddress("0x18556DA13313f3532c54711497A8FedAC273220E"),
	}, nil
}

func (t *TraderJoeClient) MyBalanceOf(binId *big.Int) (*big.Int, *big.Int, error) {

	balance, err := t.balance(binId)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("%s Error", "MyBalanceOf"), err)
	}

	return t.balanceXy(binId, balance)
}

func (t *TraderJoeClient) MyBalanceOfBatch(binId *big.Int, r int64) (ids, amounts []*big.Int, amountX, amountY *big.Int, err error) {
	balances, err := t.balanceOfBatch(binId, r)
	if err != nil {
		return nil, nil, nil, nil, errors.Join(fmt.Errorf("%s Error", "MyBalanceOfBatch"), err)
	}

	ids = make([]*big.Int, 0, r)
	amounts = make([]*big.Int, 0, r)
	amountX = big.NewInt(0)
	amountY = big.NewInt(0)

	for i := int64(0); i < r; i++ {
		if balances[i].Cmp(big.NewInt(0)) == 0 {
			continue
		}

		id := big.NewInt(binId.Int64() + i)
		ids = append(ids, id)
		amounts = append(amounts, balances[i])
		x, y, err := t.balanceXy(id, balances[i])
		if err != nil {
			return nil, nil, nil, nil, errors.Join(fmt.Errorf("%s Error", "MyBalanceOfBatch"), err)
		}
		amountX.Add(amountX, x)
		amountY.Add(amountY, y)
	}

	return ids, amounts, amountX, amountY, nil
}

func (t *TraderJoeClient) RemoveLiquidity(ids []*big.Int, amounts []*big.Int, amountXMin *big.Int, amountYMin *big.Int, isNative bool) (*common.Hash, error) {

	tokenX, err := t.getTokenX()
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "RemoveLiquidity"), err)
	}

	tokenY, err := t.getTokenY()
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "RemoveLiquidity"), err)
	}
	binStep, err := t.getBinStep()
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "RemoveLiquidity"), err)
	}

	to := t.myAddr
	deadline := big.NewInt(time.Now().Add(time.Hour * 1).Unix())

	var hash common.Hash
	if isNative {
		// Navitve 일때 토큰 X에 대한 정보 필요 없음
		// amountMin 순서도 토큰Y 다음 Native 순서
		hash, err = t.lbRouter.Send(codec.High, nil, &t.myAddr, t.pk, "removeLiquidityNATIVE", tokenY, binStep, amountYMin, amountXMin, ids, amounts, to, deadline)
	} else {
		hash, err = t.lbRouter.Send(codec.Standard, nil, &t.myAddr, t.pk, "removeLiquidity", tokenX, tokenY, binStep, amountXMin, amountYMin, ids, amounts, to, deadline)
	}

	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "RemoveLiquidity"), err)
	}
	return &hash, err
}

type Path struct {
	PairBinSteps []*big.Int       // uint256[]
	Versions     []uint8          // uint8[] (enum)
	TokenPath    []common.Address // address[]
}

func (t *TraderJoeClient) SwapNATIVEForExactTokens(amountOut *big.Int, tokenAddress common.Address, deadline *big.Int) (*common.Hash, error) {

	pairBinSteps := []*big.Int{big.NewInt(10)}
	versions := []uint8{3}
	tokenPath := []common.Address{common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"), tokenAddress}

	path := Path{
		PairBinSteps: pairBinSteps,
		Versions:     versions,
		TokenPath:    tokenPath,
	}

	hash, err := t.lbRouter.Send(codec.Standard, nil, &t.myAddr, t.pk, "swapNATIVEForExactTokens", amountOut, path, t.myAddr, deadline)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "SwapNATIVEForExactTokens"), err)
	}

	return &hash, nil
}

func (t *TraderJoeClient) SwapExactTokensForNATIVE(amountIn *big.Int, amountOutMin *big.Int, tokenAddress common.Address, deadline *big.Int) (*common.Hash, error) {

	pairBinSteps := []*big.Int{big.NewInt(10)}
	versions := []uint8{2}
	tokenPath := []common.Address{tokenAddress, common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7")}

	path := Path{
		PairBinSteps: pairBinSteps,
		Versions:     versions,
		TokenPath:    tokenPath,
	}

	hash, err := t.lbRouter.Send(codec.High, nil, &t.myAddr, t.pk, "swapExactTokensForNATIVE", amountIn, amountOutMin, path, t.myAddr, deadline)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "SwapExactTokensForNATIVE"), err)
	}

	return &hash, nil
}

func (t *TraderJoeClient) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, tokenAddress common.Address, deadline *big.Int) (*common.Hash, error) {

	pairBinSteps := []*big.Int{big.NewInt(10)}
	versions := []uint8{3}
	tokenPath := []common.Address{tokenAddress, common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7")}

	path := Path{
		PairBinSteps: pairBinSteps,
		Versions:     versions,
		TokenPath:    tokenPath,
	}

	hash, err := t.lbRouter.Send(codec.High, nil, &t.myAddr, t.pk, "swapExactTokensForTokens", amountIn, amountOutMin, path, t.myAddr, deadline)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Error", "SwapExactTokensForTokens"), err)
	}

	return &hash, nil
}

func (t *TraderJoeClient) GetSwapIn(tokenAddress common.Address, amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {

	rtn, err := t.lbRouter.Call(nil, "getSwapIn", tokenAddress, amountIn, true)
	if err != nil {
		return nil, nil, nil, errors.Join(fmt.Errorf("%s Call Error", "GetSwapIn"), err)
	}

	return rtn[0].(*big.Int), rtn[1].(*big.Int), rtn[2].(*big.Int), nil

}

/*********************************************** inner method  ***********************************************************/

func (t *TraderJoeClient) balance(binId *big.Int) (*big.Int, error) {

	rtn, err := t.lbPair.Call(&t.myAddr, "balanceOf", t.myAddr, binId)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Call Error", "balanceOf"), err)
	}

	return rtn[0].(*big.Int), nil
}

func (t *TraderJoeClient) balanceOfBatch(binId *big.Int, r int64) ([]*big.Int, error) {

	accounts := make([]common.Address, r)
	ids := make([]*big.Int, r)

	startBinId := binId.Int64()
	for i := int64(0); i < r; i++ {
		accounts[i] = t.myAddr
		ids[i] = big.NewInt(startBinId + i)
	}

	rtn, err := t.lbPair.Call(&t.myAddr, "balanceOfBatch", accounts, ids)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Call Error", "balanceOfBatch"), err)
	}

	return rtn[0].([]*big.Int), nil
}

// func (t *TraderJoeClient) balanceOfBatchOld(binId int64, r int64) ([]*big.Int, error) {
// 	accounts := make([]common.Address, 2*r+1)
// 	ids := make([]*big.Int, 2*r+1)
// 	for i := int64(0); i < 2*r+1; i++ {
// 		accounts[i] = t.myAddr
// 		ids[i] = big.NewInt(binId - r + i)
// 	}
// 	rtn, err := t.ec.Call(&t.myAddr, "balanceOfBatch", accounts, ids)
// 	if err != nil {
// 		return nil, errors.Join(fmt.Errorf("%s Call Error", "balanceOfBatch"), err)
// 	}
// 	return rtn[0].([]*big.Int), nil
// }

func (t *TraderJoeClient) activeId() (*big.Int, error) {

	rtn, err := t.lbPair.Call(nil, "getActiveId")
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Call Error", "getActiveId"), err)
	}

	return rtn[0].(*big.Int), nil
}

func (t *TraderJoeClient) isApprovedForAll() (bool, error) {

	rtn, err := t.lbPair.Call(nil, "isApprovedForAll", t.myAddr, t.routerAddr)
	if err != nil {
		return false, errors.Join(fmt.Errorf("%s Call Error", "isApprovedForAll"), err)
	}

	return rtn[0].(bool), nil
}

func (t *TraderJoeClient) getTokenX() (common.Address, error) {
	rtn, err := t.lbPair.Call(nil, "getTokenX")
	if err != nil {
		return common.Address{}, errors.Join(fmt.Errorf("%s Call Error", "getTokenX"), err)
	}

	return rtn[0].(common.Address), nil
}

func (t *TraderJoeClient) getTokenY() (common.Address, error) {
	rtn, err := t.lbPair.Call(nil, "getTokenY")
	if err != nil {
		return common.Address{}, errors.Join(fmt.Errorf("%s Call Error", "getTokenY"), err)
	}

	return rtn[0].(common.Address), nil
}

func (t *TraderJoeClient) getBinStep() (uint16, error) {
	rtn, err := t.lbPair.Call(nil, "getBinStep")
	if err != nil {
		return 0, errors.Join(fmt.Errorf("%s Call Error", "getBinStep"), err)
	}
	return rtn[0].(uint16), nil
}

func (t *TraderJoeClient) getTargetId(current float64, target float64) (*big.Int, error) {

	// gap := target - current
	// gapPercent := gap / current * 100
	// binGapCount := gapPercent / 0.1

	gapCount, err := findN(current, target, 1.001)
	if err != nil {
		return nil, err
	}

	activeId, err := t.activeId()
	if err != nil {
		return nil, err
	}

	return activeId.Add(activeId, big.NewInt(int64(gapCount))), err
}

func (t *TraderJoeClient) getPriceFromId(id *big.Int) (*big.Int, error) {
	rtn, err := t.lbPair.Call(nil, "getPriceFromId", id)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("%s Call Error", "getPriceFromId"), err)
	}
	return rtn[0].(*big.Int), nil
}

func (t *TraderJoeClient) balanceXy(binId *big.Int, balance *big.Int) (myBalanceX, myBalanceY *big.Int, err error) {

	bins, err := t.lbPair.Call(nil, "getBin", binId)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("%s Error", "MyBalanceOf"), err)
	}

	binReserveX, binReserveY := bins[0].(*big.Int), bins[1].(*big.Int)

	totalSupplyRtn, err := t.lbPair.Call(nil, "totalSupply", binId)
	if err != nil {
		return nil, nil, errors.Join(fmt.Errorf("%s Error", "MyBalanceOf"), err)
	}
	totalSupply := totalSupplyRtn[0].(*big.Int)

	myBalanceX = big.NewInt(0).Div(big.NewInt(0).Mul(balance, binReserveX), totalSupply)
	myBalanceY = big.NewInt(0).Div(big.NewInt(0).Mul(balance, binReserveY), totalSupply)

	return myBalanceX, myBalanceY, nil
}

/**************** inner function ***************/
func findN(init, target, r float64) (int, error) {
	if init == 0 || r <= 0 {
		return 0, fmt.Errorf("invalid input: a must not be zero and r must be positive")
	}

	n := math.Log10(target/(init*r)) / math.Log10(r)
	n += 1

	fmt.Println("n :", n)
	// 소수점 오차로 인해 약간의 오차가 생길 수 있어 반올림
	return int(math.Round(n)), nil
}
