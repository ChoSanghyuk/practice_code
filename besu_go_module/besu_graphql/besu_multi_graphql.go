package besu_graphql

var _ = `
	query getCall($blockNumber: Long, $callDatas: [CallData!]!) {
		block(number: $blockNumber){
			calls : call(data: $callDatas){
						data, 
						gasUsed, 
						status
				}
		}
	}
`
