import { Token, ChainId } from '@uniswap/sdk-core'

export const WAVAX_TOKEN = new Token(
  ChainId.AVALANCHE,
  '0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7',
  18,
  'WAVAX',
  'Wrapped AVAX'
)

export const USDC_TOKEN = new Token(
  ChainId.AVALANCHE,
  '0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e',
  6,
  'USDC',
  'USD Coin'
)
