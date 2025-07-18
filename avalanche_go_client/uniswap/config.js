"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.universalRouter = exports.CurrentConfig = void 0;
var constants_1 = require("./constants");
var ethers_1 = require("ethers");
exports.CurrentConfig = {
    poolKey: {
        currency0: constants_1.WAVAX_TOKEN.address,
        currency1: constants_1.USDC_TOKEN.address,
        fee: 500,
        tickSpacing: 10,
        hooks: "0x0000000000000000000000000000000000000000",
    },
    zeroForOne: true,
    amountIn: ethers_1.ethers.utils.parseUnits('0.5', constants_1.WAVAX_TOKEN.decimals).toString(),
    amountOutMinimum: "6000000", // Change according to the slippage desired
    hookData: '0x00'
};
var UNIVERSAL_ROUTER_ADDRESS = "0x94b75331ae8d42c1b61065089b7d48fe14aa73b7"; // Change the Universal Router address as per the chain
var UNIVERSAL_ROUTER_ABI = [
    {
        inputs: [
            { internalType: "bytes", name: "commands", type: "bytes" },
            { internalType: "bytes[]", name: "inputs", type: "bytes[]" },
            { internalType: "uint256", name: "deadline", type: "uint256" },
        ],
        name: "execute",
        outputs: [],
        stateMutability: "payable",
        type: "function",
    },
];
var provider = new ethers_1.ethers.providers.JsonRpcProvider("https://api.avax.network/ext/bc/C/rpc");
var pk = process.env.PK;
if (!pk.startsWith('0x')) {
    pk = '0x' + pk;
}
var signer = new ethers_1.ethers.Wallet(pk, provider);
exports.universalRouter = new ethers_1.ethers.Contract(UNIVERSAL_ROUTER_ADDRESS, UNIVERSAL_ROUTER_ABI, signer);
