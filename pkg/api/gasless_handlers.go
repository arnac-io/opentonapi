package api

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/arnac-io/opentonapi/pkg/oas"
	"github.com/tonkeeper/tongo/ton"
)

func (h *Handler) GaslessConfig(ctx context.Context) (*oas.GaslessConfig, error) {
	if h.gasless == nil {
		return nil, toError(http.StatusNotImplemented, fmt.Errorf("not implemented"))
	}
	config, err := h.gasless.Config(ctx)
	if err != nil {
		return nil, toError(http.StatusInternalServerError, fmt.Errorf("failed to get gasless config"))
	}
	o := &oas.GaslessConfig{
		GasJettons:   make([]oas.GaslessConfigGasJettonsItem, 0, len(config.SupportedJettons)),
		RelayAddress: config.RelayAddress,
	}
	for _, jetton := range config.SupportedJettons {
		o.GasJettons = append(o.GasJettons, oas.GaslessConfigGasJettonsItem{MasterID: jetton})
	}
	return o, nil
}

func (h *Handler) GaslessEstimate(ctx context.Context, req *oas.GaslessEstimateReq, params oas.GaslessEstimateParams) (*oas.SignRawParams, error) {
	if h.gasless == nil {
		return nil, toError(http.StatusNotImplemented, fmt.Errorf("not implemented"))
	}
	masterID, err := ton.ParseAccountID(params.MasterID)
	if err != nil {
		return nil, toError(http.StatusBadRequest, fmt.Errorf("invalid master_id"))
	}
	walletAddress, err := ton.ParseAccountID(req.WalletAddress)
	if err != nil {
		return nil, toError(http.StatusBadRequest, fmt.Errorf("invalid wallet address"))
	}
	publicKey, err := hex.DecodeString(req.WalletPublicKey)
	if err != nil {
		return nil, toError(http.StatusBadRequest, fmt.Errorf("invalid public key"))
	}
	messages := make([]string, 0, len(req.Messages))
	for _, msg := range req.Messages {
		messages = append(messages, msg.Boc)
	}
	signParams, err := h.gasless.Estimate(ctx, masterID, walletAddress, publicKey, messages)
	if err != nil {
		return nil, toError(http.StatusBadRequest, err)
	}
	o := &oas.SignRawParams{
		RelayAddress: signParams.RelayAddress,
		Commission:   signParams.Commission,
		From:         walletAddress.ToRaw(),
		ValidUntil:   time.Now().UTC().Add(4 * time.Minute).Unix(),
	}
	o.Messages = make([]oas.SignRawMessage, 0, len(signParams.Messages))
	for _, msg := range signParams.Messages {
		message := oas.SignRawMessage{
			Address: msg.Address,
			Amount:  msg.Amount,
		}
		if len(msg.Payload) > 0 {
			message.Payload = oas.NewOptString(msg.Payload)
		}
		if len(msg.StateInit) > 0 {
			message.StateInit = oas.NewOptString(msg.StateInit)
		}
		o.Messages = append(o.Messages, message)
	}
	return o, nil
}

func (h *Handler) GaslessSend(ctx context.Context, req *oas.GaslessSendReq) error {
	if h.gasless == nil {
		return toError(http.StatusNotImplemented, fmt.Errorf("not implemented"))
	}
	msg, err := decodeMessage(req.Boc)
	if err != nil {
		return toError(http.StatusBadRequest, err)
	}
	pubkey, err := hex.DecodeString(req.WalletPublicKey)
	if err != nil {
		return toError(http.StatusBadRequest, err)
	}
	if len(pubkey) != ed25519.PublicKeySize {
		return toError(http.StatusBadRequest, fmt.Errorf("invalid public key"))
	}
	if err := h.gasless.Send(ctx, pubkey, msg.payload); err != nil {
		return toError(http.StatusInternalServerError, err)
	}
	return nil
}
