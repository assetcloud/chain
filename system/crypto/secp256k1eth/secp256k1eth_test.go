package secp256k1eth

import (
	"crypto/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	//_ "github.com/assetcloud/chain/system/address"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func Test_All(t *testing.T) {
	testCrypto(t)
	testFromBytes(t)
}
func Test_Sign(t *testing.T) {
	//test key
	var privKeyBytes [32]byte
	rand.Read(privKeyBytes[:])
	var xpriv = PrivKeySecp256k1Eth(privKeyBytes)
	//chain 交易
	hextx := "0x0a0365766d12b74510bbe369180122f54460806040523480156200001157600080fd5b50604080518082018252600781526626bcaa37b5b2b760c91b6020808301918252835180850190945260038452624d544b60e81b9084015281519192916200005c91600091620000eb565b50805162000072906001906020840190620000eb565b5050506200008f620000896200009560201b60201c565b62000099565b620001ce565b3390565b600b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b828054620000f99062000191565b90600052602060002090601f0160209004810192826200011d576000855562000168565b82601f106200013857805160ff191683800117855562000168565b8280016001018555821562000168579182015b82811115620001685782518255916020019190600101906200014b565b50620001769291506200017a565b5090565b5b808211156200017657600081556001016200017b565b600281046001821680620001a657607f821691505b60208210811415620001c857634e487b7160e01b600052602260045260246000fd5b50919050565b61209780620001de6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806370a08231116100ad578063b88d4fde11610071578063b88d4fde14610256578063c87b56dd14610269578063e985e9c51461027c578063eacabe141461028f578063f2fde38b146102a25761012c565b806370a0823114610218578063715018a61461022b5780638da5cb5b1461023357806395d89b411461023b578063a22cb465146102435761012c565b806323b872dd116100f457806323b872dd146101b95780632f745c59146101cc57806342842e0e146101df5780634f6ccce7146101f25780636352211e146102055761012c565b806301ffc9a71461013157806306fdde031461015a578063081812fc1461016f578063095ea7b31461018f57806318160ddd146101a4575b600080fd5b61014461013f3660046117fc565b6102b5565b60405161015191906118f8565b60405180910390f35b6101626102c8565b6040516101519190611903565b61018261017d366004611834565b61035a565b60405161015191906118a7565b6101a261019d3660046117d3565b6103a6565b005b6101ac61043e565b6040516101519190611f27565b6101a26101c7366004611686565b610444565b6101ac6101da3660046117d3565b61047c565b6101a26101ed366004611686565b6104ce565b6101ac610200366004611834565b6104e9565b610182610213366004611834565b610544565b6101ac61022636600461163a565b610579565b6101a26105bd565b610182610608565b610162610617565b6101a261025136600461173a565b610626565b6101a26102643660046116c1565b61063c565b610162610277366004611834565b61067b565b61014461028a366004611654565b610686565b6101ac61029d366004611774565b6106b4565b6101a26102b036600461163a565b61080a565b60006102c08261087b565b90505b919050565b6060600080546102d790611f9f565b80601f016020809104026020016040519081016040528092919081815260200182805461030390611f9f565b80156103505780601f1061032557610100808354040283529160200191610350565b820191906000526020600020905b81548152906001019060200180831161033357829003601f168201915b5050505050905090565b6000610365826108a0565b61038a5760405162461bcd60e51b815260040161038190611d79565b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b60006103b182610544565b9050806001600160a01b0316836001600160a01b031614156103e55760405162461bcd60e51b815260040161038190611e49565b806001600160a01b03166103f76108bd565b6001600160a01b0316148061041357506104138161028a6108bd565b61042f5760405162461bcd60e51b815260040161038190611bff565b61043983836108c1565b505050565b60085490565b61045561044f6108bd565b8261092f565b6104715760405162461bcd60e51b815260040161038190611e8a565b6104398383836109b4565b600061048783610579565b82106104a55760405162461bcd60e51b815260040161038190611990565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6104398383836040518060200160405280600081525061063c565b60006104f361043e565b82106105115760405162461bcd60e51b815260040161038190611edb565b6008828154811061053257634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050919050565b6000818152600260205260408120546001600160a01b0316806102c05760405162461bcd60e51b815260040161038190611c5c565b60006001600160a01b0382166105a15760405162461bcd60e51b815260040161038190611bb6565b506001600160a01b031660009081526003602052604090205490565b6105c56108bd565b6001600160a01b03166105d6610608565b6001600160a01b0316146105fc5760405162461bcd60e51b815260040161038190611dc5565b6106066000610ae7565b565b600b546001600160a01b031690565b6060600180546102d790611f9f565b6106386106316108bd565b8383610b39565b5050565b61064d6106476108bd565b8361092f565b6106695760405162461bcd60e51b815260040161038190611e8a565b61067584848484610bdc565b50505050565b60606102c082610c0f565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60006106e1604051806040016040528060098152602001681c9958da5c1a595b9d60ba1b81525084610d28565b61070b60405180604001604052806008815260200167746f6b656e55524960c01b81525083610d6d565b610715600c610db2565b6000610721600c610dbb565b905061074e604051806040016040528060098152602001681b995dd25d195b525960ba1b81525082610dbf565b6107588482610e04565b6107a5604051806040016040528060088152602001673837b9b4ba34b7b760c11b8152506040518060400160405280600b81526020016a18599d195c8817db5a5b9d60aa1b815250610d6d565b6107af8184610eeb565b610803604051806040016040528060088152602001673837b9b4ba34b7b760c11b815250604051806040016040528060128152602001716166746572205f736574546f6b656e55524960701b815250610d6d565b9392505050565b6108126108bd565b6001600160a01b0316610823610608565b6001600160a01b0316146108495760405162461bcd60e51b815260040161038190611dc5565b6001600160a01b03811661086f5760405162461bcd60e51b815260040161038190611a2d565b61087881610ae7565b50565b60006001600160e01b0319821663780e9d6360e01b14806102c057506102c082610f2f565b6000908152600260205260409020546001600160a01b0316151590565b3390565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906108f682610544565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600061093a826108a0565b6109565760405162461bcd60e51b815260040161038190611b6a565b600061096183610544565b9050806001600160a01b0316846001600160a01b0316148061098857506109888185610686565b806109ac5750836001600160a01b03166109a18461035a565b6001600160a01b0316145b949350505050565b826001600160a01b03166109c782610544565b6001600160a01b0316146109ed5760405162461bcd60e51b815260040161038190611a73565b6001600160a01b038216610a135760405162461bcd60e51b815260040161038190611aef565b610a1e838383610f6f565b610a296000826108c1565b6001600160a01b0383166000908152600360205260408120805460019290610a52908490611f5c565b90915550506001600160a01b0382166000908152600360205260408120805460019290610a80908490611f30565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4610439838383610439565b600b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b03161415610b6b5760405162461bcd60e51b815260040161038190611b33565b6001600160a01b0383811660008181526005602090815260408083209487168084529490915290819020805460ff1916851515179055517f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3190610bcf9085906118f8565b60405180910390a3505050565b610be78484846109b4565b610bf384848484610f7a565b6106755760405162461bcd60e51b8152600401610381906119db565b6060610c1a826108a0565b610c365760405162461bcd60e51b815260040161038190611d28565b6000828152600a602052604081208054610c4f90611f9f565b80601f0160208091040260200160405190810160405280929190818152602001828054610c7b90611f9f565b8015610cc85780601f10610c9d57610100808354040283529160200191610cc8565b820191906000526020600020905b815481529060010190602001808311610cab57829003601f168201915b505050505090506000610cd9611095565b9050805160001415610ced575090506102c3565b815115610d1f578082604051602001610d07929190611878565b604051602081830303815290604052925050506102c3565b6109ac846110a7565b6106388282604051602401610d3e929190611916565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611129565b6106388282604051602401610d83929190611940565b60408051601f198184030181529190526020810180516001600160e01b0316634b5c427760e01b179052611129565b80546001019055565b5490565b6106388282604051602401610dd592919061196e565b60408051601f198184030181529190526020810180516001600160e01b03166309710a9d60e41b179052611129565b6001600160a01b038216610e2a5760405162461bcd60e51b815260040161038190611cf3565b610e33816108a0565b15610e505760405162461bcd60e51b815260040161038190611ab8565b610e5c60008383610f6f565b6001600160a01b0382166000908152600360205260408120805460019290610e85908490611f30565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a461063860008383610439565b610ef4826108a0565b610f105760405162461bcd60e51b815260040161038190611ca5565b6000828152600a60209081526040909120825161043992840190611514565b60006001600160e01b031982166380ac58cd60e01b1480610f6057506001600160e01b03198216635b5e139f60e01b145b806102c057506102c08261114a565b610439838383611163565b6000610f8e846001600160a01b03166111ec565b1561108a57836001600160a01b031663150b7a02610faa6108bd565b8786866040518563ffffffff1660e01b8152600401610fcc94939291906118bb565b602060405180830381600087803b158015610fe657600080fd5b505af1925050508015611016575060408051601f3d908101601f1916820190925261101391810190611818565b60015b611070573d808015611044576040519150601f19603f3d011682016040523d82523d6000602084013e611049565b606091505b5080516110685760405162461bcd60e51b8152600401610381906119db565b805181602001fd5b6001600160e01b031916630a85bd0160e11b1490506109ac565b506001949350505050565b60408051602081019091526000815290565b60606110b2826108a0565b6110ce5760405162461bcd60e51b815260040161038190611dfa565b60006110d8611095565b905060008151116110f85760405180602001604052806000815250610803565b80611102846111fb565b604051602001611113929190611878565b6040516020818303038152906040529392505050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6001600160e01b031981166301ffc9a760e01b14919050565b61116e838383610439565b6001600160a01b03831661118a5761118581611316565b6111ad565b816001600160a01b0316836001600160a01b0316146111ad576111ad838261135a565b6001600160a01b0382166111c9576111c4816113f7565b610439565b826001600160a01b0316826001600160a01b0316146104395761043982826114d0565b6001600160a01b03163b151590565b60608161122057506040805180820190915260018152600360fc1b60208201526102c3565b8160005b811561124a578061123481611fda565b91506112439050600a83611f48565b9150611224565b60008167ffffffffffffffff81111561127357634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801561129d576020820181803683370190505b5090505b84156109ac576112b2600183611f5c565b91506112bf600a86611ff5565b6112ca906030611f30565b60f81b8183815181106112ed57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535061130f600a86611f48565b94506112a1565b600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b6000600161136784610579565b6113719190611f5c565b6000838152600760205260409020549091508082146113c4576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061140990600190611f5c565b6000838152600960205260408120546008805493945090928490811061143f57634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050806008838154811061146e57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101929092558281526009909152604080822084905585825281205560088054806114b457634e487b7160e01b600052603160045260246000fd5b6001900381819060005260206000200160009055905550505050565b60006114db83610579565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b82805461152090611f9f565b90600052602060002090601f0160209004810192826115425760008555611588565b82601f1061155b57805160ff1916838001178555611588565b82800160010185558215611588579182015b8281111561158857825182559160200191906001019061156d565b50611594929150611598565b5090565b5b808211156115945760008155600101611599565b600067ffffffffffffffff808411156115c8576115c8612035565b604051601f8501601f19908116603f011681019082821181831017156115f0576115f0612035565b8160405280935085815286868601111561160957600080fd5b858560208301376000602087830101525050509392505050565b80356001600160a01b03811681146102c357600080fd5b60006020828403121561164b578081fd5b61080382611623565b60008060408385031215611666578081fd5b61166f83611623565b915061167d60208401611623565b90509250929050565b60008060006060848603121561169a578081fd5b6116a384611623565b92506116b160208501611623565b9150604084013590509250925092565b600080600080608085870312156116d6578081fd5b6116df85611623565b93506116ed60208601611623565b925060408501359150606085013567ffffffffffffffff81111561170f578182fd5b8501601f8101871361171f578182fd5b61172e878235602084016115ad565b91505092959194509250565b6000806040838503121561174c578182fd5b61175583611623565b915060208301358015158114611769578182fd5b809150509250929050565b60008060408385031215611786578182fd5b61178f83611623565b9150602083013567ffffffffffffffff8111156117aa578182fd5b8301601f810185136117ba578182fd5b6117c9858235602084016115ad565b9150509250929050565b600080604083850312156117e5578182fd5b6117ee83611623565b946020939093013593505050565b60006020828403121561180d578081fd5b81356108038161204b565b600060208284031215611829578081fd5b81516108038161204b565b600060208284031215611845578081fd5b5035919050565b60008151808452611864816020860160208601611f73565b601f01601f19169290920160200192915050565b6000835161188a818460208801611f73565b83519083019061189e818360208801611f73565b01949350505050565b6001600160a01b0391909116815260200190565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906118ee9083018461184c565b9695505050505050565b901515815260200190565b600060208252610803602083018461184c565b600060408252611929604083018561184c565b905060018060a01b03831660208301529392505050565b600060408252611953604083018561184c565b8281036020840152611965818561184c565b95945050505050565b600060408252611981604083018561184c565b90508260208301529392505050565b6020808252602b908201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560408201526a74206f6620626f756e647360a81b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b6020808252601c908201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604082015260600190565b60208082526024908201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646040820152637265737360e01b606082015260800190565b60208082526019908201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604082015260600190565b6020808252602c908201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860408201526b34b9ba32b73a103a37b5b2b760a11b606082015260800190565b60208082526029908201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616040820152683634b21037bbb732b960b91b606082015260800190565b60208082526038908201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760408201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000606082015260800190565b60208082526029908201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460408201526832b73a103a37b5b2b760b91b606082015260800190565b6020808252602e908201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60408201526d32bc34b9ba32b73a103a37b5b2b760911b606082015260800190565b6020808252818101527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604082015260600190565b60208082526031908201527f45524337323155524953746f726167653a2055524920717565727920666f72206040820152703737b732bc34b9ba32b73a103a37b5b2b760791b606082015260800190565b6020808252602c908201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860408201526b34b9ba32b73a103a37b5b2b760a11b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252602f908201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60408201526e3732bc34b9ba32b73a103a37b5b2b760891b606082015260800190565b60208082526021908201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656040820152603960f91b606082015260800190565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b6020808252602c908201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60408201526b7574206f6620626f756e647360a01b606082015260800190565b90815260200190565b60008219821115611f4357611f43612009565b500190565b600082611f5757611f5761201f565b500490565b600082821015611f6e57611f6e612009565b500390565b60005b83811015611f8e578181015183820152602001611f76565b838111156106755750506000910152565b600281046001821680611fb357607f821691505b60208210811415611fd457634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611fee57611fee612009565b5060010190565b6000826120045761200461201f565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b03198116811461087857600080fdfea2646970667358221220bdbace63d05b043493fe92c0459f8e1d823a78e5c1ed3a0f66550eef4ed0f14d64736f6c63430008010033320b45524332303a746f6b656e422a30783738303331326562663436323132643039653637623438653934663465643934636164333266666120bbe369309bb2b2b2d6a4e4fb733a2a3078373830333132656266343632313264303965363762343865393466346564393463616433326666615801"

	sig := xpriv.Sign(common.FromHex(hextx))
	t.Log("sign", "0x"+common.Bytes2Hex(sig.Bytes()))

	pub := xpriv.PubKey()
	ok := pub.VerifyBytes(common.FromHex(hextx), sig)
	t.Log("isOk:", ok)
	assert.Equal(t, true, ok)
	t.Log("publen:", len(pub.Bytes()))
	sigb := sig.Bytes()
	if sigb[0] != 0xff {
		sigb[0] = sigb[0] + 1 //修改签名数据进行验证测试
	} else {
		sigb[0] = sigb[0] - 1
	}

	nsig := SignatureSecp256k1Eth(sigb)
	ok = pub.VerifyBytes(common.FromHex(hextx), nsig)
	assert.Equal(t, false, ok)

	// 修改交易数据
	modifyHexTx := "0xabcd" + hextx[6:]
	ok = pub.VerifyBytes(common.FromHex(modifyHexTx), sig)
	assert.Equal(t, false, ok)
}

func testCrypto(t *testing.T) {
	require := require.New(t)
	c := &Driver{}
	priv, err := c.GenKey()
	t.Log("privkey", common.Bytes2Hex(priv.Bytes()))
	require.Nil(err)
	t.Logf("priv:%X, len:%d", priv.Bytes(), len(priv.Bytes()))

	pub := priv.PubKey()
	require.NotNil(pub)
	t.Logf("pub:%X, len:%d", pub.Bytes(), len(pub.Bytes()))

	msg := []byte("hello world")
	signature := priv.Sign(msg)
	t.Logf("sign:%X, len:%d", signature.Bytes(), len(signature.Bytes()))

	ok := pub.VerifyBytes(msg, signature)
	require.Equal(true, ok)
}

func testFromBytes(t *testing.T) {
	require := require.New(t)

	c := &Driver{}

	priv, err := c.GenKey()
	require.Nil(err)

	priv2, err := c.PrivKeyFromBytes(priv.Bytes())
	require.Nil(err)
	require.Equal(true, priv.Equals(priv2))

	s1 := string(priv.Bytes())
	s2 := string(priv2.Bytes())
	require.Equal(0, strings.Compare(s1, s2))

	pub := priv.PubKey()
	require.NotNil(pub)

	pub2, err := c.PubKeyFromBytes(pub.Bytes())
	require.Nil(err)
	require.Equal(true, pub.Equals(pub2))
	t.Log("pub str:", common.Bytes2Hex(pub.Bytes()), "len:", len(pub.Bytes()))
	s1 = string(pub.Bytes())
	s2 = string(pub2.Bytes())
	require.Equal(0, strings.Compare(s1, s2))

	var msg = []byte("hello world")
	sign1 := priv.Sign(msg)
	sign2 := priv2.Sign(msg)

	sign3, err := c.SignatureFromBytes(sign1.Bytes())
	require.Nil(err)
	require.Equal(true, sign3.Equals(sign1))

	require.Equal(true, pub.VerifyBytes(msg, sign1))
	require.Equal(true, pub2.VerifyBytes(msg, sign1))
	require.Equal(true, pub.VerifyBytes(msg, sign2))
	require.Equal(true, pub2.VerifyBytes(msg, sign2))
	require.Equal(true, pub.VerifyBytes(msg, sign3))
	require.Equal(true, pub2.VerifyBytes(msg, sign3))
}
