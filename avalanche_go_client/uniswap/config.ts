import { SwapExactInSingle } from '@uniswap/v4-sdk'
import { USDC_TOKEN, WAVAX_TOKEN } from './constants'
import { ethers } from 'ethers'

export const CurrentConfig: SwapExactInSingle = {
    poolKey: {
        currency0: WAVAX_TOKEN.address,
        currency1: USDC_TOKEN.address,
        fee: 500,
        tickSpacing: 10,
        hooks: "0x0000000000000000000000000000000000000000",
    },
    zeroForOne: true,
    amountIn: ethers.utils.parseUnits('0.5', WAVAX_TOKEN.decimals).toString(), 
    amountOutMinimum: "6000000", // Change according to the slippage desired
    hookData: '0x00'
}

const UNIVERSAL_ROUTER_ADDRESS = "0x94b75331ae8d42c1b61065089b7d48fe14aa73b7" // Change the Universal Router address as per the chain

const UNIVERSAL_ROUTER_ABI = [
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
]

const provider = new ethers.providers.JsonRpcProvider("https://api.avax.network/ext/bc/C/rpc");
let pk = process.env.PK as string;
if (!pk.startsWith('0x')) {
  pk = '0x' + pk;
}
const signer = new ethers.Wallet(
  pk,
  provider
);
export const universalRouter = new ethers.Contract(
    UNIVERSAL_ROUTER_ADDRESS,
    UNIVERSAL_ROUTER_ABI,
    signer
)

