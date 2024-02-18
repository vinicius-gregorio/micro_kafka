package create_transaction

import (
	"context"

	"github.com/vinicius-gregorio/fc_walletcore/internal/entity"
	"github.com/vinicius-gregorio/fc_walletcore/internal/gateway"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/events"
	"github.com/vinicius-gregorio/fc_walletcore/pkg/uow"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	TransactionID string  `json:"transaction_id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type CreateTransactionUsecase struct {
	Uow                uow.UowInterface
	eventDispatcher    events.EventDispatcherInterface
	transactionCreated events.EventInterface
	balanceUpdated     events.EventInterface
}

func NewCreateTransactionUsecase(Uow uow.UowInterface, eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface, balanceUpdated events.EventInterface) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		Uow:                Uow,
		eventDispatcher:    eventDispatcher,
		transactionCreated: transactionCreated,
		balanceUpdated:     balanceUpdated,
	}
}

func (uc *CreateTransactionUsecase) Execute(ctx context.Context, input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}
	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)

		accFrom, err := accountRepository.FindByID(input.AccountIDFrom)
		if err != nil {
			return err

		}
		accTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}

		tr, err := entity.NewTransaction(accFrom, accTo, input.Amount)
		if err != nil {
			return err
		}
		err = accountRepository.UpdateBalance(accFrom)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accTo)
		if err != nil {
			return err
		}
		err = transactionRepository.Create(tr)
		if err != nil {
			return err
		}
		output.TransactionID = tr.ID
		output.AccountIDFrom = tr.AccountFrom.ID
		output.AccountIDTo = tr.AccountTo.ID
		output.Amount = tr.Amount

		balanceUpdatedOutput.AccountIDFrom = tr.AccountFrom.ID
		balanceUpdatedOutput.AccountIDTo = tr.AccountTo.ID
		balanceUpdatedOutput.BalanceAccountIDFrom = tr.AccountFrom.Balance
		balanceUpdatedOutput.BalanceAccountIDTo = tr.AccountTo.Balance
		return nil
	})
	if err != nil {
		return nil, err
	}
	uc.transactionCreated.SetPayload(output)
	uc.eventDispatcher.Dispatch(uc.transactionCreated)

	uc.balanceUpdated.SetPayload(balanceUpdatedOutput)
	uc.eventDispatcher.Dispatch(uc.balanceUpdated)

	return output, nil
}

func (uc *CreateTransactionUsecase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *CreateTransactionUsecase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
