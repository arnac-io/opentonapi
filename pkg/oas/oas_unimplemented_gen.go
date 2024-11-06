// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AccountDnsBackResolve implements accountDnsBackResolve operation.
//
// Get account's domains.
//
// GET /v2/accounts/{account_id}/dns/backresolve
func (UnimplementedHandler) AccountDnsBackResolve(ctx context.Context, params AccountDnsBackResolveParams) (r *DomainNames, _ error) {
	return r, ht.ErrNotImplemented
}

// AddressParse implements addressParse operation.
//
// Parse address and display in all formats.
//
// GET /v2/address/{account_id}/parse
func (UnimplementedHandler) AddressParse(ctx context.Context, params AddressParseParams) (r *AddressParseOK, _ error) {
	return r, ht.ErrNotImplemented
}

// BlockchainAccountInspect implements blockchainAccountInspect operation.
//
// Blockchain account inspect.
//
// GET /v2/blockchain/accounts/{account_id}/inspect
func (UnimplementedHandler) BlockchainAccountInspect(ctx context.Context, params BlockchainAccountInspectParams) (r *BlockchainAccountInspect, _ error) {
	return r, ht.ErrNotImplemented
}

// DecodeMessage implements decodeMessage operation.
//
// Decode a given message. Only external incoming messages can be decoded currently.
//
// POST /v2/message/decode
func (UnimplementedHandler) DecodeMessage(ctx context.Context, req *DecodeMessageReq) (r *DecodedMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// DnsResolve implements dnsResolve operation.
//
// DNS resolve for domain name.
//
// GET /v2/dns/{domain_name}/resolve
func (UnimplementedHandler) DnsResolve(ctx context.Context, params DnsResolveParams) (r *DnsRecord, _ error) {
	return r, ht.ErrNotImplemented
}

// ExecGetMethodForBlockchainAccount implements execGetMethodForBlockchainAccount operation.
//
// Execute get method for account.
//
// GET /v2/blockchain/accounts/{account_id}/methods/{method_name}
func (UnimplementedHandler) ExecGetMethodForBlockchainAccount(ctx context.Context, params ExecGetMethodForBlockchainAccountParams) (r *MethodExecutionResult, _ error) {
	return r, ht.ErrNotImplemented
}

// GaslessConfig implements gaslessConfig operation.
//
// Returns configuration of gasless transfers.
//
// GET /v2/gasless/config
func (UnimplementedHandler) GaslessConfig(ctx context.Context) (r *GaslessConfig, _ error) {
	return r, ht.ErrNotImplemented
}

// GaslessEstimate implements gaslessEstimate operation.
//
// Estimates the cost of the given messages and returns a payload to sign.
//
// POST /v2/gasless/estimate/{master_id}
func (UnimplementedHandler) GaslessEstimate(ctx context.Context, req *GaslessEstimateReq, params GaslessEstimateParams) (r *SignRawParams, _ error) {
	return r, ht.ErrNotImplemented
}

// GaslessSend implements gaslessSend operation.
//
// Submits the signed gasless transaction message to the network.
//
// POST /v2/gasless/send
func (UnimplementedHandler) GaslessSend(ctx context.Context, req *GaslessSendReq) error {
	return ht.ErrNotImplemented
}

// GetAccount implements getAccount operation.
//
// Get human-friendly information about an account without low-level details.
//
// GET /v2/accounts/{account_id}
func (UnimplementedHandler) GetAccount(ctx context.Context, params GetAccountParams) (r *Account, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountDiff implements getAccountDiff operation.
//
// Get account's balance change.
//
// GET /v2/accounts/{account_id}/diff
func (UnimplementedHandler) GetAccountDiff(ctx context.Context, params GetAccountDiffParams) (r *GetAccountDiffOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountDnsExpiring implements getAccountDnsExpiring operation.
//
// Get expiring account .ton dns.
//
// GET /v2/accounts/{account_id}/dns/expiring
func (UnimplementedHandler) GetAccountDnsExpiring(ctx context.Context, params GetAccountDnsExpiringParams) (r *DnsExpiring, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountEvent implements getAccountEvent operation.
//
// Get event for an account by event_id.
//
// GET /v2/accounts/{account_id}/events/{event_id}
func (UnimplementedHandler) GetAccountEvent(ctx context.Context, params GetAccountEventParams) (r *AccountEvent, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountEvents implements getAccountEvents operation.
//
// Get events for an account. Each event is built on top of a trace which is a series of transactions
// caused by one inbound message. TonAPI looks for known patterns inside the trace and splits the
// trace into actions, where a single action represents a meaningful high-level operation like a
// Jetton Transfer or an NFT Purchase. Actions are expected to be shown to users. It is advised not
// to build any logic on top of actions because actions can be changed at any time.
//
// GET /v2/accounts/{account_id}/events
func (UnimplementedHandler) GetAccountEvents(ctx context.Context, params GetAccountEventsParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountInfoByStateInit implements getAccountInfoByStateInit operation.
//
// Get account info by state init.
//
// POST /v2/tonconnect/stateinit
func (UnimplementedHandler) GetAccountInfoByStateInit(ctx context.Context, req *GetAccountInfoByStateInitReq) (r *AccountInfoByStateInit, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountInscriptions implements getAccountInscriptions operation.
//
// Get all inscriptions by owner address. It's experimental API and can be dropped in the future.
//
// GET /v2/experimental/accounts/{account_id}/inscriptions
func (UnimplementedHandler) GetAccountInscriptions(ctx context.Context, params GetAccountInscriptionsParams) (r *InscriptionBalances, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountInscriptionsHistory implements getAccountInscriptionsHistory operation.
//
// Get the transfer inscriptions history for account. It's experimental API and can be dropped in the
// future.
//
// GET /v2/experimental/accounts/{account_id}/inscriptions/history
func (UnimplementedHandler) GetAccountInscriptionsHistory(ctx context.Context, params GetAccountInscriptionsHistoryParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountInscriptionsHistoryByTicker implements getAccountInscriptionsHistoryByTicker operation.
//
// Get the transfer inscriptions history for account. It's experimental API and can be dropped in the
// future.
//
// GET /v2/experimental/accounts/{account_id}/inscriptions/{ticker}/history
func (UnimplementedHandler) GetAccountInscriptionsHistoryByTicker(ctx context.Context, params GetAccountInscriptionsHistoryByTickerParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountJettonBalance implements getAccountJettonBalance operation.
//
// Get Jetton balance by owner address.
//
// GET /v2/accounts/{account_id}/jettons/{jetton_id}
func (UnimplementedHandler) GetAccountJettonBalance(ctx context.Context, params GetAccountJettonBalanceParams) (r *JettonBalance, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountJettonHistoryByID implements getAccountJettonHistoryByID operation.
//
// Get the transfer jetton history for account and jetton.
//
// GET /v2/accounts/{account_id}/jettons/{jetton_id}/history
func (UnimplementedHandler) GetAccountJettonHistoryByID(ctx context.Context, params GetAccountJettonHistoryByIDParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountJettonsBalances implements getAccountJettonsBalances operation.
//
// Get all Jettons balances by owner address.
//
// GET /v2/accounts/{account_id}/jettons
func (UnimplementedHandler) GetAccountJettonsBalances(ctx context.Context, params GetAccountJettonsBalancesParams) (r *JettonsBalances, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountJettonsHistory implements getAccountJettonsHistory operation.
//
// Get the transfer jettons history for account.
//
// GET /v2/accounts/{account_id}/jettons/history
func (UnimplementedHandler) GetAccountJettonsHistory(ctx context.Context, params GetAccountJettonsHistoryParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountMultisigs implements getAccountMultisigs operation.
//
// Get account's multisigs.
//
// GET /v2/accounts/{account_id}/multisigs
func (UnimplementedHandler) GetAccountMultisigs(ctx context.Context, params GetAccountMultisigsParams) (r *Multisigs, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountNftHistory implements getAccountNftHistory operation.
//
// Get the transfer nft history.
//
// GET /v2/accounts/{account_id}/nfts/history
func (UnimplementedHandler) GetAccountNftHistory(ctx context.Context, params GetAccountNftHistoryParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountNftItems implements getAccountNftItems operation.
//
// Get all NFT items by owner address.
//
// GET /v2/accounts/{account_id}/nfts
func (UnimplementedHandler) GetAccountNftItems(ctx context.Context, params GetAccountNftItemsParams) (r *NftItems, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountNominatorsPools implements getAccountNominatorsPools operation.
//
// All pools where account participates.
//
// GET /v2/staking/nominator/{account_id}/pools
func (UnimplementedHandler) GetAccountNominatorsPools(ctx context.Context, params GetAccountNominatorsPoolsParams) (r *AccountStaking, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountPublicKey implements getAccountPublicKey operation.
//
// Get public key by account id.
//
// GET /v2/accounts/{account_id}/publickey
func (UnimplementedHandler) GetAccountPublicKey(ctx context.Context, params GetAccountPublicKeyParams) (r *GetAccountPublicKeyOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountSeqno implements getAccountSeqno operation.
//
// Get account seqno.
//
// GET /v2/wallet/{account_id}/seqno
func (UnimplementedHandler) GetAccountSeqno(ctx context.Context, params GetAccountSeqnoParams) (r *Seqno, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountSubscriptions implements getAccountSubscriptions operation.
//
// Get all subscriptions by wallet address.
//
// GET /v2/accounts/{account_id}/subscriptions
func (UnimplementedHandler) GetAccountSubscriptions(ctx context.Context, params GetAccountSubscriptionsParams) (r *Subscriptions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccountTraces implements getAccountTraces operation.
//
// Get traces for account.
//
// GET /v2/accounts/{account_id}/traces
func (UnimplementedHandler) GetAccountTraces(ctx context.Context, params GetAccountTracesParams) (r *TraceIDs, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAccounts implements getAccounts operation.
//
// Get human-friendly information about several accounts without low-level details.
//
// POST /v2/accounts/_bulk
func (UnimplementedHandler) GetAccounts(ctx context.Context, req OptGetAccountsReq, params GetAccountsParams) (r *Accounts, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAllAuctions implements getAllAuctions operation.
//
// Get all auctions.
//
// GET /v2/dns/auctions
func (UnimplementedHandler) GetAllAuctions(ctx context.Context, params GetAllAuctionsParams) (r *Auctions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAllRawShardsInfo implements getAllRawShardsInfo operation.
//
// Get all raw shards info.
//
// GET /v2/liteserver/get_all_shards_info/{block_id}
func (UnimplementedHandler) GetAllRawShardsInfo(ctx context.Context, params GetAllRawShardsInfoParams) (r *GetAllRawShardsInfoOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainAccountTransactions implements getBlockchainAccountTransactions operation.
//
// Get account transactions.
//
// GET /v2/blockchain/accounts/{account_id}/transactions
func (UnimplementedHandler) GetBlockchainAccountTransactions(ctx context.Context, params GetBlockchainAccountTransactionsParams) (r *Transactions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainBlock implements getBlockchainBlock operation.
//
// Get blockchain block data.
//
// GET /v2/blockchain/blocks/{block_id}
func (UnimplementedHandler) GetBlockchainBlock(ctx context.Context, params GetBlockchainBlockParams) (r *BlockchainBlock, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainBlockTransactions implements getBlockchainBlockTransactions operation.
//
// Get transactions from block.
//
// GET /v2/blockchain/blocks/{block_id}/transactions
func (UnimplementedHandler) GetBlockchainBlockTransactions(ctx context.Context, params GetBlockchainBlockTransactionsParams) (r *Transactions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainConfig implements getBlockchainConfig operation.
//
// Get blockchain config.
//
// GET /v2/blockchain/config
func (UnimplementedHandler) GetBlockchainConfig(ctx context.Context) (r *BlockchainConfig, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainConfigFromBlock implements getBlockchainConfigFromBlock operation.
//
// Get blockchain config from a specific block, if present.
//
// GET /v2/blockchain/masterchain/{masterchain_seqno}/config
func (UnimplementedHandler) GetBlockchainConfigFromBlock(ctx context.Context, params GetBlockchainConfigFromBlockParams) (r *BlockchainConfig, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainMasterchainBlocks implements getBlockchainMasterchainBlocks operation.
//
// Get all blocks in all shards and workchains between target and previous masterchain block
// according to shards last blocks snapshot in masterchain.  We don't recommend to build your app
// around this method because it has problem with scalability and will work very slow in the future.
//
// GET /v2/blockchain/masterchain/{masterchain_seqno}/blocks
func (UnimplementedHandler) GetBlockchainMasterchainBlocks(ctx context.Context, params GetBlockchainMasterchainBlocksParams) (r *BlockchainBlocks, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainMasterchainHead implements getBlockchainMasterchainHead operation.
//
// Get last known masterchain block.
//
// GET /v2/blockchain/masterchain-head
func (UnimplementedHandler) GetBlockchainMasterchainHead(ctx context.Context) (r *BlockchainBlock, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainMasterchainShards implements getBlockchainMasterchainShards operation.
//
// Get blockchain block shards.
//
// GET /v2/blockchain/masterchain/{masterchain_seqno}/shards
func (UnimplementedHandler) GetBlockchainMasterchainShards(ctx context.Context, params GetBlockchainMasterchainShardsParams) (r *BlockchainBlockShards, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainMasterchainTransactions implements getBlockchainMasterchainTransactions operation.
//
// Get all transactions in all shards and workchains between target and previous masterchain block
// according to shards last blocks snapshot in masterchain. We don't recommend to build your app
// around this method because it has problem with scalability and will work very slow in the future.
//
// GET /v2/blockchain/masterchain/{masterchain_seqno}/transactions
func (UnimplementedHandler) GetBlockchainMasterchainTransactions(ctx context.Context, params GetBlockchainMasterchainTransactionsParams) (r *Transactions, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainRawAccount implements getBlockchainRawAccount operation.
//
// Get low-level information about an account taken directly from the blockchain.
//
// GET /v2/blockchain/accounts/{account_id}
func (UnimplementedHandler) GetBlockchainRawAccount(ctx context.Context, params GetBlockchainRawAccountParams) (r *BlockchainRawAccount, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainTransaction implements getBlockchainTransaction operation.
//
// Get transaction data.
//
// GET /v2/blockchain/transactions/{transaction_id}
func (UnimplementedHandler) GetBlockchainTransaction(ctx context.Context, params GetBlockchainTransactionParams) (r *Transaction, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockchainTransactionByMessageHash implements getBlockchainTransactionByMessageHash operation.
//
// Get transaction data by message hash.
//
// GET /v2/blockchain/messages/{msg_id}/transaction
func (UnimplementedHandler) GetBlockchainTransactionByMessageHash(ctx context.Context, params GetBlockchainTransactionByMessageHashParams) (r *Transaction, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChartRates implements getChartRates operation.
//
// Get chart by token.
//
// GET /v2/rates/chart
func (UnimplementedHandler) GetChartRates(ctx context.Context, params GetChartRatesParams) (r *GetChartRatesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetDnsInfo implements getDnsInfo operation.
//
// Get full information about domain name.
//
// GET /v2/dns/{domain_name}
func (UnimplementedHandler) GetDnsInfo(ctx context.Context, params GetDnsInfoParams) (r *DomainInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// GetDomainBids implements getDomainBids operation.
//
// Get domain bids.
//
// GET /v2/dns/{domain_name}/bids
func (UnimplementedHandler) GetDomainBids(ctx context.Context, params GetDomainBidsParams) (r *DomainBids, _ error) {
	return r, ht.ErrNotImplemented
}

// GetEvent implements getEvent operation.
//
// Get an event either by event ID or a hash of any transaction in a trace. An event is built on top
// of a trace which is a series of transactions caused by one inbound message. TonAPI looks for known
// patterns inside the trace and splits the trace into actions, where a single action represents a
// meaningful high-level operation like a Jetton Transfer or an NFT Purchase. Actions are expected to
// be shown to users. It is advised not to build any logic on top of actions because actions can be
// changed at any time.
//
// GET /v2/events/{event_id}
func (UnimplementedHandler) GetEvent(ctx context.Context, params GetEventParams) (r *Event, _ error) {
	return r, ht.ErrNotImplemented
}

// GetInscriptionOpTemplate implements getInscriptionOpTemplate operation.
//
// Return comment for making operation with inscription. please don't use it if you don't know what
// you are doing.
//
// GET /v2/experimental/inscriptions/op-template
func (UnimplementedHandler) GetInscriptionOpTemplate(ctx context.Context, params GetInscriptionOpTemplateParams) (r *GetInscriptionOpTemplateOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetItemsFromCollection implements getItemsFromCollection operation.
//
// Get NFT items from collection by collection address.
//
// GET /v2/nfts/collections/{account_id}/items
func (UnimplementedHandler) GetItemsFromCollection(ctx context.Context, params GetItemsFromCollectionParams) (r *NftItems, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonHolders implements getJettonHolders operation.
//
// Get jetton's holders.
//
// GET /v2/jettons/{account_id}/holders
func (UnimplementedHandler) GetJettonHolders(ctx context.Context, params GetJettonHoldersParams) (r *JettonHolders, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonInfo implements getJettonInfo operation.
//
// Get jetton metadata by jetton master address.
//
// GET /v2/jettons/{account_id}
func (UnimplementedHandler) GetJettonInfo(ctx context.Context, params GetJettonInfoParams) (r *JettonInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonInfosByAddresses implements getJettonInfosByAddresses operation.
//
// Get jetton metadata items by jetton master addresses.
//
// POST /v2/jettons/_bulk
func (UnimplementedHandler) GetJettonInfosByAddresses(ctx context.Context, req OptGetJettonInfosByAddressesReq) (r *Jettons, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonTransferPayload implements getJettonTransferPayload operation.
//
// Get jetton's custom payload and state init required for transfer.
//
// GET /v2/jettons/{jetton_id}/transfer/{account_id}/payload
func (UnimplementedHandler) GetJettonTransferPayload(ctx context.Context, params GetJettonTransferPayloadParams) (r *JettonTransferPayload, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettons implements getJettons operation.
//
// Get a list of all indexed jetton masters in the blockchain.
//
// GET /v2/jettons
func (UnimplementedHandler) GetJettons(ctx context.Context, params GetJettonsParams) (r *Jettons, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJettonsEvents implements getJettonsEvents operation.
//
// Get only jetton transfers in the event.
//
// GET /v2/events/{event_id}/jettons
func (UnimplementedHandler) GetJettonsEvents(ctx context.Context, params GetJettonsEventsParams) (r *Event, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMarketsRates implements getMarketsRates operation.
//
// Get the TON price from markets.
//
// GET /v2/rates/markets
func (UnimplementedHandler) GetMarketsRates(ctx context.Context) (r *GetMarketsRatesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMultisigAccount implements getMultisigAccount operation.
//
// Get multisig account info.
//
// GET /v2/multisig/{account_id}
func (UnimplementedHandler) GetMultisigAccount(ctx context.Context, params GetMultisigAccountParams) (r *Multisig, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftCollection implements getNftCollection operation.
//
// Get NFT collection by collection address.
//
// GET /v2/nfts/collections/{account_id}
func (UnimplementedHandler) GetNftCollection(ctx context.Context, params GetNftCollectionParams) (r *NftCollection, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftCollectionItemsByAddresses implements getNftCollectionItemsByAddresses operation.
//
// Get NFT collection items by their addresses.
//
// POST /v2/nfts/collections/_bulk
func (UnimplementedHandler) GetNftCollectionItemsByAddresses(ctx context.Context, req OptGetNftCollectionItemsByAddressesReq) (r *NftCollections, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftCollections implements getNftCollections operation.
//
// Get NFT collections.
//
// GET /v2/nfts/collections
func (UnimplementedHandler) GetNftCollections(ctx context.Context, params GetNftCollectionsParams) (r *NftCollections, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftHistoryByID implements getNftHistoryByID operation.
//
// Get the transfer nfts history for account.
//
// GET /v2/nfts/{account_id}/history
func (UnimplementedHandler) GetNftHistoryByID(ctx context.Context, params GetNftHistoryByIDParams) (r *AccountEvents, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftItemByAddress implements getNftItemByAddress operation.
//
// Get NFT item by its address.
//
// GET /v2/nfts/{account_id}
func (UnimplementedHandler) GetNftItemByAddress(ctx context.Context, params GetNftItemByAddressParams) (r *NftItem, _ error) {
	return r, ht.ErrNotImplemented
}

// GetNftItemsByAddresses implements getNftItemsByAddresses operation.
//
// Get NFT items by their addresses.
//
// POST /v2/nfts/_bulk
func (UnimplementedHandler) GetNftItemsByAddresses(ctx context.Context, req OptGetNftItemsByAddressesReq) (r *NftItems, _ error) {
	return r, ht.ErrNotImplemented
}

// GetOutMsgQueueSizes implements getOutMsgQueueSizes operation.
//
// Get out msg queue sizes.
//
// GET /v2/liteserver/get_out_msg_queue_sizes
func (UnimplementedHandler) GetOutMsgQueueSizes(ctx context.Context) (r *GetOutMsgQueueSizesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRates implements getRates operation.
//
// Get the token price in the chosen currency for display only. Don’t use this for financial
// transactions.
//
// GET /v2/rates
func (UnimplementedHandler) GetRates(ctx context.Context, params GetRatesParams) (r *GetRatesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawAccountState implements getRawAccountState operation.
//
// Get raw account state.
//
// GET /v2/liteserver/get_account_state/{account_id}
func (UnimplementedHandler) GetRawAccountState(ctx context.Context, params GetRawAccountStateParams) (r *GetRawAccountStateOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawBlockProof implements getRawBlockProof operation.
//
// Get raw block proof.
//
// GET /v2/liteserver/get_block_proof
func (UnimplementedHandler) GetRawBlockProof(ctx context.Context, params GetRawBlockProofParams) (r *GetRawBlockProofOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawBlockchainBlock implements getRawBlockchainBlock operation.
//
// Get raw blockchain block.
//
// GET /v2/liteserver/get_block/{block_id}
func (UnimplementedHandler) GetRawBlockchainBlock(ctx context.Context, params GetRawBlockchainBlockParams) (r *GetRawBlockchainBlockOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawBlockchainBlockHeader implements getRawBlockchainBlockHeader operation.
//
// Get raw blockchain block header.
//
// GET /v2/liteserver/get_block_header/{block_id}
func (UnimplementedHandler) GetRawBlockchainBlockHeader(ctx context.Context, params GetRawBlockchainBlockHeaderParams) (r *GetRawBlockchainBlockHeaderOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawBlockchainBlockState implements getRawBlockchainBlockState operation.
//
// Get raw blockchain block state.
//
// GET /v2/liteserver/get_state/{block_id}
func (UnimplementedHandler) GetRawBlockchainBlockState(ctx context.Context, params GetRawBlockchainBlockStateParams) (r *GetRawBlockchainBlockStateOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawBlockchainConfig implements getRawBlockchainConfig operation.
//
// Get raw blockchain config.
//
// GET /v2/blockchain/config/raw
func (UnimplementedHandler) GetRawBlockchainConfig(ctx context.Context) (r *RawBlockchainConfig, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawBlockchainConfigFromBlock implements getRawBlockchainConfigFromBlock operation.
//
// Get raw blockchain config from a specific block, if present.
//
// GET /v2/blockchain/masterchain/{masterchain_seqno}/config/raw
func (UnimplementedHandler) GetRawBlockchainConfigFromBlock(ctx context.Context, params GetRawBlockchainConfigFromBlockParams) (r *RawBlockchainConfig, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawConfig implements getRawConfig operation.
//
// Get raw config.
//
// GET /v2/liteserver/get_config_all/{block_id}
func (UnimplementedHandler) GetRawConfig(ctx context.Context, params GetRawConfigParams) (r *GetRawConfigOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawListBlockTransactions implements getRawListBlockTransactions operation.
//
// Get raw list block transactions.
//
// GET /v2/liteserver/list_block_transactions/{block_id}
func (UnimplementedHandler) GetRawListBlockTransactions(ctx context.Context, params GetRawListBlockTransactionsParams) (r *GetRawListBlockTransactionsOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawMasterchainInfo implements getRawMasterchainInfo operation.
//
// Get raw masterchain info.
//
// GET /v2/liteserver/get_masterchain_info
func (UnimplementedHandler) GetRawMasterchainInfo(ctx context.Context) (r *GetRawMasterchainInfoOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawMasterchainInfoExt implements getRawMasterchainInfoExt operation.
//
// Get raw masterchain info ext.
//
// GET /v2/liteserver/get_masterchain_info_ext
func (UnimplementedHandler) GetRawMasterchainInfoExt(ctx context.Context, params GetRawMasterchainInfoExtParams) (r *GetRawMasterchainInfoExtOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawShardBlockProof implements getRawShardBlockProof operation.
//
// Get raw shard block proof.
//
// GET /v2/liteserver/get_shard_block_proof/{block_id}
func (UnimplementedHandler) GetRawShardBlockProof(ctx context.Context, params GetRawShardBlockProofParams) (r *GetRawShardBlockProofOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawShardInfo implements getRawShardInfo operation.
//
// Get raw shard info.
//
// GET /v2/liteserver/get_shard_info/{block_id}
func (UnimplementedHandler) GetRawShardInfo(ctx context.Context, params GetRawShardInfoParams) (r *GetRawShardInfoOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawTime implements getRawTime operation.
//
// Get raw time.
//
// GET /v2/liteserver/get_time
func (UnimplementedHandler) GetRawTime(ctx context.Context) (r *GetRawTimeOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRawTransactions implements getRawTransactions operation.
//
// Get raw transactions.
//
// GET /v2/liteserver/get_transactions/{account_id}
func (UnimplementedHandler) GetRawTransactions(ctx context.Context, params GetRawTransactionsParams) (r *GetRawTransactionsOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetReducedBlockchainBlocks implements getReducedBlockchainBlocks operation.
//
// Get reduced blockchain blocks data.
//
// GET /v2/blockchain/reduced/blocks
func (UnimplementedHandler) GetReducedBlockchainBlocks(ctx context.Context, params GetReducedBlockchainBlocksParams) (r *ReducedBlocks, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStakingPoolHistory implements getStakingPoolHistory operation.
//
// Pool history.
//
// GET /v2/staking/pool/{account_id}/history
func (UnimplementedHandler) GetStakingPoolHistory(ctx context.Context, params GetStakingPoolHistoryParams) (r *GetStakingPoolHistoryOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStakingPoolInfo implements getStakingPoolInfo operation.
//
// Stacking pool info.
//
// GET /v2/staking/pool/{account_id}
func (UnimplementedHandler) GetStakingPoolInfo(ctx context.Context, params GetStakingPoolInfoParams) (r *GetStakingPoolInfoOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStakingPools implements getStakingPools operation.
//
// All pools available in network.
//
// GET /v2/staking/pools
func (UnimplementedHandler) GetStakingPools(ctx context.Context, params GetStakingPoolsParams) (r *GetStakingPoolsOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStorageProviders implements getStorageProviders operation.
//
// Get TON storage providers deployed to the blockchain.
//
// GET /v2/storage/providers
func (UnimplementedHandler) GetStorageProviders(ctx context.Context) (r *GetStorageProvidersOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTonConnectPayload implements getTonConnectPayload operation.
//
// Get a payload for further token receipt.
//
// GET /v2/tonconnect/payload
func (UnimplementedHandler) GetTonConnectPayload(ctx context.Context) (r *GetTonConnectPayloadOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTrace implements getTrace operation.
//
// Get the trace by trace ID or hash of any transaction in trace.
//
// GET /v2/traces/{trace_id}
func (UnimplementedHandler) GetTrace(ctx context.Context, params GetTraceParams) (r *Trace, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWalletBackup implements getWalletBackup operation.
//
// Get backup info.
//
// GET /v2/wallet/backup
func (UnimplementedHandler) GetWalletBackup(ctx context.Context, params GetWalletBackupParams) (r *GetWalletBackupOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWalletsByPublicKey implements getWalletsByPublicKey operation.
//
// Get wallets by public key.
//
// GET /v2/pubkeys/{public_key}/wallets
func (UnimplementedHandler) GetWalletsByPublicKey(ctx context.Context, params GetWalletsByPublicKeyParams) (r *Accounts, _ error) {
	return r, ht.ErrNotImplemented
}

// ReindexAccount implements reindexAccount operation.
//
// Update internal cache for a particular account.
//
// POST /v2/accounts/{account_id}/reindex
func (UnimplementedHandler) ReindexAccount(ctx context.Context, params ReindexAccountParams) error {
	return ht.ErrNotImplemented
}

// SearchAccounts implements searchAccounts operation.
//
// Search by account domain name.
//
// GET /v2/accounts/search
func (UnimplementedHandler) SearchAccounts(ctx context.Context, params SearchAccountsParams) (r *FoundAccounts, _ error) {
	return r, ht.ErrNotImplemented
}

// SendBlockchainMessage implements sendBlockchainMessage operation.
//
// Send message to blockchain.
//
// POST /v2/blockchain/message
func (UnimplementedHandler) SendBlockchainMessage(ctx context.Context, req *SendBlockchainMessageReq) error {
	return ht.ErrNotImplemented
}

// SendRawMessage implements sendRawMessage operation.
//
// Send raw message to blockchain.
//
// POST /v2/liteserver/send_message
func (UnimplementedHandler) SendRawMessage(ctx context.Context, req *SendRawMessageReq) (r *SendRawMessageOK, _ error) {
	return r, ht.ErrNotImplemented
}

// SetWalletBackup implements setWalletBackup operation.
//
// Set backup info.
//
// PUT /v2/wallet/backup
func (UnimplementedHandler) SetWalletBackup(ctx context.Context, req SetWalletBackupReq, params SetWalletBackupParams) error {
	return ht.ErrNotImplemented
}

// Status implements status operation.
//
// Status.
//
// GET /v2/status
func (UnimplementedHandler) Status(ctx context.Context) (r *ServiceStatus, _ error) {
	return r, ht.ErrNotImplemented
}

// TonConnectProof implements tonConnectProof operation.
//
// Account verification and token issuance.
//
// POST /v2/wallet/auth/proof
func (UnimplementedHandler) TonConnectProof(ctx context.Context, req *TonConnectProofReq) (r *TonConnectProofOK, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
