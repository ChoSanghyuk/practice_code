
access list
https://dbadoy.tistory.com/25


getPriceFromId
binId를 넣으면 해당되는 price를 반환함. 단, 128.128-binary fixed-point number로 받음
이는, 원래 값이 정수 부분은 <<2^128 시키고, 그 뒤에 소수값을 붙였다는 의미로, 반환값에서 2^128를 나누어야 본래의 값이 나옴.
여기서, WAvax는 18-decimals이고, USDC는 6-decimals이기 때문에 10^12를 해주어야 해당 빈의 달러값이 나옴



AddLiquidity 할 때는 
Approve => AddLiquidity 였음
예시 TX
- approve : 0x7ac434aa3542b08899ea808991972a8ba721c71cd5745c89a91e14cf1c864483
- addLiquidity : 0x4f7ec30c20f29ac0584336fc0456157d3e3f0d94b4eeb6de4c201f0ff7408043

swap이 어려움
최적 swap 찾기 => 진행 순서인 것으로 보이지만... 흠

Swap 용어 비교
Mode	Fixed Value	Variable Value	Use Case
Swap Exact In	Input 	Output 	Spend exact, get at least X
Swap Exact Out	Output 	Input 	Buy exact, pay at most X

​	

### Swap



에시 TX : 0x7e1c222fa3695000b39aabdcfe7ab6e4f7d3695efc59dd24ce36d070800b9aef

To : 0x45A62B090DF48243F12A21897e7ed91863E2c86b

Inner ERC20 TX

| From | To | Amount | Token |
|------|----|--------|-------|
| 내 지갑 | ForwarderLogic	| 239.257496 | blank |
| ForwarderLogic | MagpieRouterV3_1	| 239.257496 | blank |
| MagpieRouterV3_1 | MagpieRouterCore	| 239.257496 | blank |
| MagpieRouterCore | StataTokenV2	| 239.257496 | blank |
| StataTokenV2 |Aave V3: aAvaAUSD |239.257496	| blank |
|[Null] | StataTokenV2	| 285.756996 | blank |
|[Null] | MagpieRouterCore	| 235.545683 |ERC-20: Wrapped Aave |
| MagpieRouterCore | Vault	| 235.545683 |ERC-20: Wrapped Aave |
| Vault | MagpieRouterCore	| 204.164301 |ERC20 *** |
| MagpieRouterCore |[Null]	| 204.164301 |ERC20 *** |
| StataTokenV2 |[Null]	| 234.57243 |Aave Avalanc |
|[Aave: aUSDT token V3] |MagpieRouterCore | 238.899095	| TetherToken |
| MagpieRouterCore | TransparentUpgradeableProxy	| 238.899095 |TetherToken |
| TransparentUpgradeableProxy | MagpieRouterCore	| 13.243492321234094979 |Wrapped AVAX |
| MagpieRouterCore | EOA	| 0.006120157108706369 |Wrapped AVAX |
| MagpieRouterCore | ForwarderLogic	| 13.23737216412538861 |Wrapped AVAX |
| ForwarderLogic |[LFJ: Joe Aggregator Router v2.2]	| 13.23737216412538861 |Wrapped AVAX |



| address                                    | DESC                            |
| ------------------------------------------ | ------------------------------- |
| 0xb4dd4fb3D4bCED984cce972991fB100488b59223 | 내 지갑                         |
| 0xC04f291347D21DC663f7646056db22bFf8CE8430 | ForwarderLogic                  |
| 0x3611B82c7B13e72b26eb0E9BE0613bEE7A45aC7c | MagpieRouterV3_1                |
| 0x29Ed0a2F22a92Ff84A7F196785Ca6b0D21AEeC62 | MagpieRouterCore                |
| 0x45cf39EeB437FA95Bb9b52c0105254a6BD25D01e | StataTokenV2                    |
| 0xbA1333333333a1BA1108E8412f11850A5C319bA9 | Vault                           |
| 0x59933c571d200dc6A7Fd1CDa22495dB442082E34 | StataTokenV2                    |
| 0x5520385bFcf07Ec87C4c53A7d8d65595Dff69FA4 | TransparentUpgradeableProxy     |
| 0xeDA49BcE2F38d284f839Be1f4f2E23e6C7cC7DBd | EOA                             |
| 0x45A62B090DF48243F12A21897e7ed91863E2c86b | LFJ: Joe Aggregator Router v2.2 |





### Swap ex2 (07.01 10USDC <=> WAVAX)

tx: 0xdfc5b31a8b999ef0c56cfd99dff3458201ccaf8f2de2e939b155b3893b118497

|i| From           | To             | Amount   | Token    |
| --- | -------------- | -------------- | -------- | -------- |
|0| 내 지갑        | ForwarderLogic | 10       | USD Coin |
|1| ForwarderLogic | 0x6E4141       | 10       | USDC     |
|2| PangolinV3Pool | 0x6E4141       | 0.5615.. | WAVAX    |
|3| 0x6E4141       | PangolinV3Pool | 10       | USDC     |
|4| 0x6E4141       | ForwarderLogic | 0.5615.. | WAVAX    |
|5| ForwarderLogic | 내 지갑        | 0.5615.. | WAVAX    |



| address                                    | DESC                    |
| ------------------------------------------ | ----------------------- |
| 0xb4dd4fb3D4bCED984cce972991fB100488b59223 | 내 지갑                 |
| 0xC04f291347D21DC663f7646056db22bFf8CE8430 | ForwarderLogic          |
| 0x6E4141d33021b52C91c28608403db4A0FFB50Ec6 | ???                     |
| 0x11476e10eB79ddfFa6F2585BE526d2bd840C3E20 | PangolinV3Pool          |
| 0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e | USDC                    |
| 0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7 | WAVAX                   |
| 0x6131b5fae19ea4f9d964eac0408e4408b66337b5 | MetaAggregationRouterV2 |





0x6131b5fae19ea4f9d964eac0408e4408b66337b5

0x6E4141d33021b52C91c28608403db4A0FFB50Ec6





## ts code tx vs uniswap tx



### uniswap tx - success

```
{
          "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
          "blockNumber": "0x3e1a517",
          "contractAddress": "",
          "cumulativeGasUsed": "0x60915",
          "effectiveGasPrice": "0x6945271a",
          "from": "0xb4dd4fb3d4bced984cce972991fb100488b59223",
          "gasUsed": "0x2f891",
          "logs": [
            {
              "address": "0x000000000022d473030f116ddee9f6b43ac78ba3",
              "topics": [
                "0xc6a377bfc4eb120024a8ac08eef205be16b817020812c73223e81d1bdb9708ec",
                "0x000000000000000000000000b4dd4fb3d4bced984cce972991fb100488b59223",
                "0x000000000000000000000000b31f66aa3c1e785363f0875a1b74e27b85fd66c7",
                "0x00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7"
              ],
              "data": "0x000000000000000000000000ffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000689440c60000000000000000000000000000000000000000000000000000000000000000",
              "blockNumber": "0x3e1a517",
              "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
              "transactionIndex": "0x2",
              "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
              "logIndex": "0x4",
              "removed": false
            },
            {
              "address": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
              "topics": [
                "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                "0x000000000000000000000000fae3f424a0a47706811521e3ee268f00cfb5c45e",
                "0x00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7"
              ],
              "data": "0x0000000000000000000000000000000000000000000000000000000000880273",
              "blockNumber": "0x3e1a517",
              "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
              "transactionIndex": "0x2",
              "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
              "logIndex": "0x5",
              "removed": false
            },
            {
              "address": "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7",
              "topics": [
                "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                "0x000000000000000000000000b4dd4fb3d4bced984cce972991fb100488b59223",
                "0x000000000000000000000000fae3f424a0a47706811521e3ee268f00cfb5c45e"
              ],
              "data": "0x00000000000000000000000000000000000000000000000006f05b59d3b20000",
              "blockNumber": "0x3e1a517",
              "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
              "transactionIndex": "0x2",
              "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
              "logIndex": "0x6",
              "removed": false
            },
            {
              "address": "0xfae3f424a0a47706811521e3ee268f00cfb5c45e",
              "topics": [
                "0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67",
                "0x00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7",
                "0x00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7"
              ],
              "data": "0x00000000000000000000000000000000000000000000000006f05b59d3b20000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff77fd8d0000000000000000000000000000000000000000000046dac3a140e04fc5095b0000000000000000000000000000000000000000000000000d625a275985d486fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3929",
              "blockNumber": "0x3e1a517",
              "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
              "transactionIndex": "0x2",
              "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
              "logIndex": "0x7",
              "removed": false
            },
            {
              "address": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
              "topics": [
                "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                "0x00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7",
                "0x0000000000000000000000007ffc3dbf3b2b50ff3a1d5523bc24bb5043837b14"
              ],
              "data": "0x000000000000000000000000000000000000000000000000000000000000570b",
              "blockNumber": "0x3e1a517",
              "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
              "transactionIndex": "0x2",
              "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
              "logIndex": "0x8",
              "removed": false
            },
            {
              "address": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
              "topics": [
                "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
                "0x00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7",
                "0x000000000000000000000000b4dd4fb3d4bced984cce972991fb100488b59223"
              ],
              "data": "0x000000000000000000000000000000000000000000000000000000000087ab68",
              "blockNumber": "0x3e1a517",
              "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
              "transactionIndex": "0x2",
              "blockHash": "0x8ca6f0da018729c9187f4e041f91b0c2dd63a497c606e940c96ef3c2485a5dc1",
              "logIndex": "0x9",
              "removed": false
            }
          ],
          "logsBloom": "0x00010000000000000000000000004081000000040000000020000000000100000000000000000040004000000000000000010040000020000000020000080000000080000040010800000008000000000001000000080000000000020000000000000010000000000000000000000000000000000000000080000010000800000000004000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000400000000000001000000000000000000000002000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000402",
          "revertReason": "",
          "status": "0x1",
          "to": "0x94b75331ae8d42c1b61065089b7d48fe14aa73b7",
          "transactionHash": "0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a",
          "transactionIndex": "0x2",
          "type": "0x2"
        }
```



## Tx data 분석



tx hash : 0x0ee26793629f51b044c9c95771b240f48a75d823429969e349c740f30401e69a

desc : 0.5 WAVAX => 8.89 USDC 



```bash
0x3593564c # execute
0000000000000000000000000000000000000000000000000000000000000060 # bytes   calldata 시작 위치 = 96
00000000000000000000000000000000000000000000000000000000000000a0 # bytes[] inputs 시작 위치 = 160 
00000000000000000000000000000000000000000000000000000000686cbb08 # deadline = 1751956232 (20250708 15:30:32)
# calldata 시작
0000000000000000000000000000000000000000000000000000000000000004 # length of commands = 4
0a00060400000000000000000000000000000000000000000000000000000000 # commands content, padded to 32 bytes
# 0a(PERMIT2_PERMIT) + 00(V3_SWAP_EXACT_IN) + 06(PAY_PORTION) + 04(SWEEP)

# inputs 시작
0000000000000000000000000000000000000000000000000000000000000004 # length of inputs = 4

# inputs들 시작 위치
0000000000000000000000000000000000000000000000000000000000000080 # input 1 위치. 128
0000000000000000000000000000000000000000000000000000000000000200 # input 2 위치. 512
0000000000000000000000000000000000000000000000000000000000000320 # input 3 위치. 800
00000000000000000000000000000000000000000000000000000000000003a0 # input 4 위치. 928

# input 1 - PERMIT2_PERMIT
0000000000000000000000000000000000000000000000000000000000000160 # input 1 길이. 352 (32*11)
000000000000000000000000b31f66aa3c1e785363f0875a1b74e27b85fd66c7 # PermitSingle.PermitDetails.token = WAVAX
000000000000000000000000ffffffffffffffffffffffffffffffffffffffff # PermitSingle.PermitDetails.amount
00000000000000000000000000000000000000000000000000000000689440c6 # PermitSingle.PermitDetails.expiration
0000000000000000000000000000000000000000000000000000000000000000 # PermitSingle.PermitDetails.nonce
00000000000000000000000094b75331ae8d42c1b61065089b7d48fe14aa73b7 # PermitSingle.spender = UniversalRouter
00000000000000000000000000000000000000000000000000000000686cbace # PermitSingle.sigDeadline = 1751956174
# signature The owner's signature over the permit data
00000000000000000000000000000000000000000000000000000000000000e0 
0000000000000000000000000000000000000000000000000000000000000041
efdda88d0d53eae44b4e6eb2aafb98cfa52ca2657741d620b91f3f951f6a1b99
7ea2ebf023289c059ce1b5480b40538a5222c5766d091624d97be396ee4c4d11
1c00000000000000000000000000000000000000000000000000000000000000

# input 2 - V3_SWAP_EXACT_IN
0000000000000000000000000000000000000000000000000000000000000100 # input 2 길이. 256 (32*8)
0000000000000000000000000000000000000000000000000000000000000002 # recipient(address) = address(this)
00000000000000000000000000000000000000000000000006f05b59d3b20000 # amount = 0.5 * 10^18
0000000000000000000000000000000000000000000000000000000000000000 # amountOutMin = 0 
00000000000000000000000000000000000000000000000000000000000000a0 # path 시작 위치 160. (단, 동적부분에서 160)
0000000000000000000000000000000000000000000000000000000000000001 # payerIsUser
000000000000000000000000000000000000000000000000000000000000002b # path 길이 = 43
b31f66aa3c1e785363f0875a1b74e27b85fd66c7 # WAVAX(address)
0001f4 # 500
b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e # USDC(address)
000000000000000000000000000000000000000000

# input 3 - PAY_PORTION
0000000000000000000000000000000000000000000000000000000000000060 # 길이
000000000000000000000000b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e # token(address)
0000000000000000000000007ffc3dbf3b2b50ff3a1d5523bc24bb5043837b14 # recipient(address)
0000000000000000000000000000000000000000000000000000000000000019 # bips(uint256) // fees 0.25%

# input 4 - SWEEP
0000000000000000000000000000000000000000000000000000000000000060
000000000000000000000000b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e # token(address)
000000000000000000000000b4dd4fb3d4bced984cce972991fb100488b59223 # recipient
000000000000000000000000000000000000000000000000000000000086fdea # amountMin
0c
```

0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7





```solidity
'permit(address,((address,uint160,uint48,uint48),address,uint256),bytes)',
msgSender(),
struct PermitSingle {
    // the permit data for a single token alownce
    PermitDetails details;
    // address permissioned on the allowed tokens
    address spender;
    // deadline on the permit signature
    uint256 sigDeadline;
}

struct PermitDetails {
    // ERC20 token address
    address token;
    // the maximum amount allowed to spend
    uint160 amount;
    // timestamp at which a spender's token allowances become invalid
    uint48 expiration;
    // an incrementing value indexed per owner,token,and spender for each signature
    uint48 nonce;
}
```







