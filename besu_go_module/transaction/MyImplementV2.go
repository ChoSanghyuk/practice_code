// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transaction

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyImplementV2MetaData contains all meta data concerning the MyImplementV2 contract.
var MyImplementV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CUSTOM_OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000043575f80fd5b50620000546200006a60201b60201c565b620000646200006a60201b60201c565b620001d4565b5f6200007b6200016e60201b60201c565b9050805f0160089054906101000a900460ff1615620000c6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146200016b5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff604051620001629190620001b9565b60405180910390a15b50565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f67ffffffffffffffff82169050919050565b620001b38162000195565b82525050565b5f602082019050620001ce5f830184620001a8565b92915050565b608051611779620001fb5f395f8181610a5b01528181610ab00152610c8d01526117795ff3fe6080604052600436106100e7575f3560e01c80634f1ef2861161008957806391d148541161005857806391d14854146102b5578063a217fddf146102f1578063ad3cb1cc1461031b578063d547741f14610345576100e7565b80634f1ef2861461023157806352d1902d1461024d57806355241077146102775780638129fc1c1461029f576100e7565b806320965255116100c5578063209652551461017b578063248a9ca3146101a55780632f2ff15d146101e157806336568abe14610209576100e7565b806301ffc9a7146100eb57806303d390b5146101275780630d8e6e2c14610151575b5f80fd5b3480156100f6575f80fd5b50610111600480360381019061010c9190611186565b61036d565b60405161011e91906111cb565b60405180910390f35b348015610132575f80fd5b5061013b6103e6565b60405161014891906111fc565b60405180910390f35b34801561015c575f80fd5b5061016561040a565b6040516101729190611231565b60405180910390f35b348015610186575f80fd5b5061018f610412565b60405161019c9190611262565b60405180910390f35b3480156101b0575f80fd5b506101cb60048036038101906101c691906112a5565b61041a565b6040516101d891906111fc565b60405180910390f35b3480156101ec575f80fd5b506102076004803603810190610202919061132a565b610444565b005b348015610214575f80fd5b5061022f600480360381019061022a919061132a565b610466565b005b61024b600480360381019061024691906114a4565b6104e1565b005b348015610258575f80fd5b50610261610500565b60405161026e91906111fc565b60405180910390f35b348015610282575f80fd5b5061029d60048036038101906102989190611528565b610531565b005b3480156102aa575f80fd5b506102b361053a565b005b3480156102c0575f80fd5b506102db60048036038101906102d6919061132a565b6106ec565b6040516102e891906111cb565b60405180910390f35b3480156102fc575f80fd5b5061030561075d565b60405161031291906111fc565b60405180910390f35b348015610326575f80fd5b5061032f610763565b60405161033c91906115cd565b60405180910390f35b348015610350575f80fd5b5061036b6004803603810190610366919061132a565b61079c565b005b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103df57506103de826107be565b5b9050919050565b7ff5dda7ea6ff5c85469fa7647e83e14382e03b49f95c4d130d00ce0351a01ea4e81565b5f6002905090565b5f8054905090565b5f80610424610827565b9050805f015f8481526020019081526020015f2060010154915050919050565b61044d8261041a565b6104568161084e565b6104608383610862565b50505050565b61046e61095a565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146104d2576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104dc8282610961565b505050565b6104e9610a59565b6104f282610b3f565b6104fc8282610b6d565b5050565b5f610509610c8b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b805f8190555050565b5f610543610d12565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff1614801561058b5750825b90505f60018367ffffffffffffffff161480156105be57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156105cc575080155b15610603576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610650576001855f0160086101000a81548160ff0219169083151502179055505b610658610d39565b6106827ff5dda7ea6ff5c85469fa7647e83e14382e03b49f95c4d130d00ce0351a01ea4e33610862565b5061068b610d43565b83156106e5575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516106dc9190611642565b60405180910390a15b5050505050565b5f806106f6610827565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b5f801b81565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6107a58261041a565b6107ae8161084e565b6107b88383610961565b50505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b61085f8161085a61095a565b610d4d565b50565b5f8061086c610827565b905061087884846106ec565b61094f576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506108eb61095a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610954565b5f9150505b92915050565b5f33905090565b5f8061096b610827565b905061097784846106ec565b15610a4e575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506109ea61095a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001915050610a53565b5f9150505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480610b0657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610aed610d9e565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610b3d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7ff5dda7ea6ff5c85469fa7647e83e14382e03b49f95c4d130d00ce0351a01ea4e610b698161084e565b5050565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610bd557506040513d601f19601f82011682018060405250810190610bd2919061166f565b60015b610c1657816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610c0d91906116a9565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b8114610c7c57806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401610c7391906111fc565b60405180910390fd5b610c868383610df1565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610d10576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b610d41610e63565b565b610d4b610e63565b565b610d5782826106ec565b610d9a5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401610d919291906116c2565b60405180910390fd5b5050565b5f610dca7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b610ea3565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610dfa82610eac565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f81511115610e5657610e508282610f75565b50610e5f565b610e5e610ff5565b5b5050565b610e6b611031565b610ea1576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03610f0757806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610efe91906116a9565b60405180910390fd5b80610f337f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b610ea3565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f808473ffffffffffffffffffffffffffffffffffffffff1684604051610f9e919061172d565b5f60405180830381855af49150503d805f8114610fd6576040519150601f19603f3d011682016040523d82523d5f602084013e610fdb565b606091505b5091509150610feb85838361104f565b9250505092915050565b5f34111561102f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f61103a610d12565b5f0160089054906101000a900460ff16905090565b6060826110645761105f826110dc565b6110d4565b5f825114801561108a57505f8473ffffffffffffffffffffffffffffffffffffffff163b145b156110cc57836040517f9996b3150000000000000000000000000000000000000000000000000000000081526004016110c391906116a9565b60405180910390fd5b8190506110d5565b5b9392505050565b5f815111156110ee5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61116581611131565b811461116f575f80fd5b50565b5f813590506111808161115c565b92915050565b5f6020828403121561119b5761119a611129565b5b5f6111a884828501611172565b91505092915050565b5f8115159050919050565b6111c5816111b1565b82525050565b5f6020820190506111de5f8301846111bc565b92915050565b5f819050919050565b6111f6816111e4565b82525050565b5f60208201905061120f5f8301846111ed565b92915050565b5f61ffff82169050919050565b61122b81611215565b82525050565b5f6020820190506112445f830184611222565b92915050565b5f819050919050565b61125c8161124a565b82525050565b5f6020820190506112755f830184611253565b92915050565b611284816111e4565b811461128e575f80fd5b50565b5f8135905061129f8161127b565b92915050565b5f602082840312156112ba576112b9611129565b5b5f6112c784828501611291565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112f9826112d0565b9050919050565b611309816112ef565b8114611313575f80fd5b50565b5f8135905061132481611300565b92915050565b5f80604083850312156113405761133f611129565b5b5f61134d85828601611291565b925050602061135e85828601611316565b9150509250929050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6113b682611370565b810181811067ffffffffffffffff821117156113d5576113d4611380565b5b80604052505050565b5f6113e7611120565b90506113f382826113ad565b919050565b5f67ffffffffffffffff82111561141257611411611380565b5b61141b82611370565b9050602081019050919050565b828183375f83830152505050565b5f611448611443846113f8565b6113de565b9050828152602081018484840111156114645761146361136c565b5b61146f848285611428565b509392505050565b5f82601f83011261148b5761148a611368565b5b813561149b848260208601611436565b91505092915050565b5f80604083850312156114ba576114b9611129565b5b5f6114c785828601611316565b925050602083013567ffffffffffffffff8111156114e8576114e761112d565b5b6114f485828601611477565b9150509250929050565b6115078161124a565b8114611511575f80fd5b50565b5f81359050611522816114fe565b92915050565b5f6020828403121561153d5761153c611129565b5b5f61154a84828501611514565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561158a57808201518184015260208101905061156f565b5f8484015250505050565b5f61159f82611553565b6115a9818561155d565b93506115b981856020860161156d565b6115c281611370565b840191505092915050565b5f6020820190508181035f8301526115e58184611595565b905092915050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f61162c611627611622846115ed565b611609565b6115f6565b9050919050565b61163c81611612565b82525050565b5f6020820190506116555f830184611633565b92915050565b5f815190506116698161127b565b92915050565b5f6020828403121561168457611683611129565b5b5f6116918482850161165b565b91505092915050565b6116a3816112ef565b82525050565b5f6020820190506116bc5f83018461169a565b92915050565b5f6040820190506116d55f83018561169a565b6116e260208301846111ed565b9392505050565b5f81519050919050565b5f81905092915050565b5f611707826116e9565b61171181856116f3565b935061172181856020860161156d565b80840191505092915050565b5f61173882846116fd565b91508190509291505056fea264697066735822122062930804d3e8845670b964f10e6068ad42fc6831d364a72b396304470c1a07d064736f6c63430008140033",
}

// MyImplementV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MyImplementV2MetaData.ABI instead.
var MyImplementV2ABI = MyImplementV2MetaData.ABI

// MyImplementV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyImplementV2MetaData.Bin instead.
var MyImplementV2Bin = MyImplementV2MetaData.Bin

// DeployMyImplementV2 deploys a new Ethereum contract, binding an instance of MyImplementV2 to it.
func DeployMyImplementV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyImplementV2, error) {
	parsed, err := MyImplementV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyImplementV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyImplementV2{MyImplementV2Caller: MyImplementV2Caller{contract: contract}, MyImplementV2Transactor: MyImplementV2Transactor{contract: contract}, MyImplementV2Filterer: MyImplementV2Filterer{contract: contract}}, nil
}

// MyImplementV2 is an auto generated Go binding around an Ethereum contract.
type MyImplementV2 struct {
	MyImplementV2Caller     // Read-only binding to the contract
	MyImplementV2Transactor // Write-only binding to the contract
	MyImplementV2Filterer   // Log filterer for contract events
}

// MyImplementV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyImplementV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyImplementV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyImplementV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyImplementV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyImplementV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyImplementV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyImplementV2Session struct {
	Contract     *MyImplementV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyImplementV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyImplementV2CallerSession struct {
	Contract *MyImplementV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MyImplementV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyImplementV2TransactorSession struct {
	Contract     *MyImplementV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MyImplementV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyImplementV2Raw struct {
	Contract *MyImplementV2 // Generic contract binding to access the raw methods on
}

// MyImplementV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyImplementV2CallerRaw struct {
	Contract *MyImplementV2Caller // Generic read-only contract binding to access the raw methods on
}

// MyImplementV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyImplementV2TransactorRaw struct {
	Contract *MyImplementV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyImplementV2 creates a new instance of MyImplementV2, bound to a specific deployed contract.
func NewMyImplementV2(address common.Address, backend bind.ContractBackend) (*MyImplementV2, error) {
	contract, err := bindMyImplementV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2{MyImplementV2Caller: MyImplementV2Caller{contract: contract}, MyImplementV2Transactor: MyImplementV2Transactor{contract: contract}, MyImplementV2Filterer: MyImplementV2Filterer{contract: contract}}, nil
}

// NewMyImplementV2Caller creates a new read-only instance of MyImplementV2, bound to a specific deployed contract.
func NewMyImplementV2Caller(address common.Address, caller bind.ContractCaller) (*MyImplementV2Caller, error) {
	contract, err := bindMyImplementV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2Caller{contract: contract}, nil
}

// NewMyImplementV2Transactor creates a new write-only instance of MyImplementV2, bound to a specific deployed contract.
func NewMyImplementV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MyImplementV2Transactor, error) {
	contract, err := bindMyImplementV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2Transactor{contract: contract}, nil
}

// NewMyImplementV2Filterer creates a new log filterer instance of MyImplementV2, bound to a specific deployed contract.
func NewMyImplementV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MyImplementV2Filterer, error) {
	contract, err := bindMyImplementV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2Filterer{contract: contract}, nil
}

// bindMyImplementV2 binds a generic wrapper to an already deployed contract.
func bindMyImplementV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyImplementV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyImplementV2 *MyImplementV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyImplementV2.Contract.MyImplementV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyImplementV2 *MyImplementV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV2.Contract.MyImplementV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyImplementV2 *MyImplementV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyImplementV2.Contract.MyImplementV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyImplementV2 *MyImplementV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyImplementV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyImplementV2 *MyImplementV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyImplementV2 *MyImplementV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyImplementV2.Contract.contract.Transact(opts, method, params...)
}

// CUSTOMOWNERROLE is a free data retrieval call binding the contract method 0x03d390b5.
//
// Solidity: function CUSTOM_OWNER_ROLE() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Caller) CUSTOMOWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "CUSTOM_OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CUSTOMOWNERROLE is a free data retrieval call binding the contract method 0x03d390b5.
//
// Solidity: function CUSTOM_OWNER_ROLE() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Session) CUSTOMOWNERROLE() ([32]byte, error) {
	return _MyImplementV2.Contract.CUSTOMOWNERROLE(&_MyImplementV2.CallOpts)
}

// CUSTOMOWNERROLE is a free data retrieval call binding the contract method 0x03d390b5.
//
// Solidity: function CUSTOM_OWNER_ROLE() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2CallerSession) CUSTOMOWNERROLE() ([32]byte, error) {
	return _MyImplementV2.Contract.CUSTOMOWNERROLE(&_MyImplementV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _MyImplementV2.Contract.DEFAULTADMINROLE(&_MyImplementV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MyImplementV2.Contract.DEFAULTADMINROLE(&_MyImplementV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MyImplementV2 *MyImplementV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MyImplementV2 *MyImplementV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _MyImplementV2.Contract.UPGRADEINTERFACEVERSION(&_MyImplementV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MyImplementV2 *MyImplementV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MyImplementV2.Contract.UPGRADEINTERFACEVERSION(&_MyImplementV2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MyImplementV2.Contract.GetRoleAdmin(&_MyImplementV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MyImplementV2 *MyImplementV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MyImplementV2.Contract.GetRoleAdmin(&_MyImplementV2.CallOpts, role)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_MyImplementV2 *MyImplementV2Caller) GetValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "getValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_MyImplementV2 *MyImplementV2Session) GetValue() (*big.Int, error) {
	return _MyImplementV2.Contract.GetValue(&_MyImplementV2.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() view returns(uint256)
func (_MyImplementV2 *MyImplementV2CallerSession) GetValue() (*big.Int, error) {
	return _MyImplementV2.Contract.GetValue(&_MyImplementV2.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint16)
func (_MyImplementV2 *MyImplementV2Caller) GetVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint16)
func (_MyImplementV2 *MyImplementV2Session) GetVersion() (uint16, error) {
	return _MyImplementV2.Contract.GetVersion(&_MyImplementV2.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint16)
func (_MyImplementV2 *MyImplementV2CallerSession) GetVersion() (uint16, error) {
	return _MyImplementV2.Contract.GetVersion(&_MyImplementV2.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyImplementV2 *MyImplementV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyImplementV2 *MyImplementV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MyImplementV2.Contract.HasRole(&_MyImplementV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MyImplementV2 *MyImplementV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MyImplementV2.Contract.HasRole(&_MyImplementV2.CallOpts, role, account)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2Session) ProxiableUUID() ([32]byte, error) {
	return _MyImplementV2.Contract.ProxiableUUID(&_MyImplementV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MyImplementV2 *MyImplementV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _MyImplementV2.Contract.ProxiableUUID(&_MyImplementV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyImplementV2 *MyImplementV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MyImplementV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyImplementV2 *MyImplementV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyImplementV2.Contract.SupportsInterface(&_MyImplementV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyImplementV2 *MyImplementV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyImplementV2.Contract.SupportsInterface(&_MyImplementV2.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyImplementV2 *MyImplementV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MyImplementV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyImplementV2 *MyImplementV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MyImplementV2.Contract.GrantRole(&_MyImplementV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MyImplementV2 *MyImplementV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MyImplementV2.Contract.GrantRole(&_MyImplementV2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyImplementV2 *MyImplementV2Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyImplementV2.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyImplementV2 *MyImplementV2Session) Initialize() (*types.Transaction, error) {
	return _MyImplementV2.Contract.Initialize(&_MyImplementV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_MyImplementV2 *MyImplementV2TransactorSession) Initialize() (*types.Transaction, error) {
	return _MyImplementV2.Contract.Initialize(&_MyImplementV2.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MyImplementV2 *MyImplementV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MyImplementV2.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MyImplementV2 *MyImplementV2Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MyImplementV2.Contract.RenounceRole(&_MyImplementV2.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MyImplementV2 *MyImplementV2TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MyImplementV2.Contract.RenounceRole(&_MyImplementV2.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyImplementV2 *MyImplementV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MyImplementV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyImplementV2 *MyImplementV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MyImplementV2.Contract.RevokeRole(&_MyImplementV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MyImplementV2 *MyImplementV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MyImplementV2.Contract.RevokeRole(&_MyImplementV2.TransactOpts, role, account)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 newValue) returns()
func (_MyImplementV2 *MyImplementV2Transactor) SetValue(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _MyImplementV2.contract.Transact(opts, "setValue", newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 newValue) returns()
func (_MyImplementV2 *MyImplementV2Session) SetValue(newValue *big.Int) (*types.Transaction, error) {
	return _MyImplementV2.Contract.SetValue(&_MyImplementV2.TransactOpts, newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x55241077.
//
// Solidity: function setValue(uint256 newValue) returns()
func (_MyImplementV2 *MyImplementV2TransactorSession) SetValue(newValue *big.Int) (*types.Transaction, error) {
	return _MyImplementV2.Contract.SetValue(&_MyImplementV2.TransactOpts, newValue)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MyImplementV2 *MyImplementV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MyImplementV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MyImplementV2 *MyImplementV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MyImplementV2.Contract.UpgradeToAndCall(&_MyImplementV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MyImplementV2 *MyImplementV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MyImplementV2.Contract.UpgradeToAndCall(&_MyImplementV2.TransactOpts, newImplementation, data)
}

// MyImplementV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MyImplementV2 contract.
type MyImplementV2InitializedIterator struct {
	Event *MyImplementV2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV2Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV2Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV2Initialized represents a Initialized event raised by the MyImplementV2 contract.
type MyImplementV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyImplementV2 *MyImplementV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*MyImplementV2InitializedIterator, error) {

	logs, sub, err := _MyImplementV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MyImplementV2InitializedIterator{contract: _MyImplementV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyImplementV2 *MyImplementV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MyImplementV2Initialized) (event.Subscription, error) {

	logs, sub, err := _MyImplementV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV2Initialized)
				if err := _MyImplementV2.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MyImplementV2 *MyImplementV2Filterer) ParseInitialized(log types.Log) (*MyImplementV2Initialized, error) {
	event := new(MyImplementV2Initialized)
	if err := _MyImplementV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyImplementV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MyImplementV2 contract.
type MyImplementV2RoleAdminChangedIterator struct {
	Event *MyImplementV2RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV2RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV2RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV2RoleAdminChanged represents a RoleAdminChanged event raised by the MyImplementV2 contract.
type MyImplementV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MyImplementV2 *MyImplementV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MyImplementV2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MyImplementV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2RoleAdminChangedIterator{contract: _MyImplementV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MyImplementV2 *MyImplementV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MyImplementV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MyImplementV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV2RoleAdminChanged)
				if err := _MyImplementV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MyImplementV2 *MyImplementV2Filterer) ParseRoleAdminChanged(log types.Log) (*MyImplementV2RoleAdminChanged, error) {
	event := new(MyImplementV2RoleAdminChanged)
	if err := _MyImplementV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyImplementV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MyImplementV2 contract.
type MyImplementV2RoleGrantedIterator struct {
	Event *MyImplementV2RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV2RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV2RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV2RoleGranted represents a RoleGranted event raised by the MyImplementV2 contract.
type MyImplementV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyImplementV2 *MyImplementV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MyImplementV2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyImplementV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2RoleGrantedIterator{contract: _MyImplementV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyImplementV2 *MyImplementV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MyImplementV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyImplementV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV2RoleGranted)
				if err := _MyImplementV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyImplementV2 *MyImplementV2Filterer) ParseRoleGranted(log types.Log) (*MyImplementV2RoleGranted, error) {
	event := new(MyImplementV2RoleGranted)
	if err := _MyImplementV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyImplementV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MyImplementV2 contract.
type MyImplementV2RoleRevokedIterator struct {
	Event *MyImplementV2RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV2RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV2RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV2RoleRevoked represents a RoleRevoked event raised by the MyImplementV2 contract.
type MyImplementV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyImplementV2 *MyImplementV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MyImplementV2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyImplementV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2RoleRevokedIterator{contract: _MyImplementV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyImplementV2 *MyImplementV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MyImplementV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyImplementV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV2RoleRevoked)
				if err := _MyImplementV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MyImplementV2 *MyImplementV2Filterer) ParseRoleRevoked(log types.Log) (*MyImplementV2RoleRevoked, error) {
	event := new(MyImplementV2RoleRevoked)
	if err := _MyImplementV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyImplementV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MyImplementV2 contract.
type MyImplementV2UpgradedIterator struct {
	Event *MyImplementV2Upgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyImplementV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyImplementV2Upgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyImplementV2Upgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyImplementV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyImplementV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyImplementV2Upgraded represents a Upgraded event raised by the MyImplementV2 contract.
type MyImplementV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyImplementV2 *MyImplementV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MyImplementV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MyImplementV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MyImplementV2UpgradedIterator{contract: _MyImplementV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyImplementV2 *MyImplementV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MyImplementV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MyImplementV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyImplementV2Upgraded)
				if err := _MyImplementV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MyImplementV2 *MyImplementV2Filterer) ParseUpgraded(log types.Log) (*MyImplementV2Upgraded, error) {
	event := new(MyImplementV2Upgraded)
	if err := _MyImplementV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
