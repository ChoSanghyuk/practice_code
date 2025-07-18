"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.USDC_TOKEN = exports.WAVAX_TOKEN = void 0;
var sdk_core_1 = require("@uniswap/sdk-core");
exports.WAVAX_TOKEN = new sdk_core_1.Token(sdk_core_1.ChainId.AVALANCHE, '0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7', 18, 'WAVAX', 'Wrapped AVAX');
exports.USDC_TOKEN = new sdk_core_1.Token(sdk_core_1.ChainId.AVALANCHE, '0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e', 6, 'USDC', 'USD Coin');
