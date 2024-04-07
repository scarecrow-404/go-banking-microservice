package service

import (
	"time"

	"github.com/scarecrow-404/banking/domain"
	"github.com/scarecrow-404/banking/dto"
	"github.com/scarecrow-404/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest)(*dto.NewAccountResponse,*errs.AppError){
	err :=req.Validate()
	if err != nil {
		return nil,err
	}
	
	a:=domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}

	newAccount,err := s.repo.Save(a)
	if err != nil {
		return nil,err
	}
	response := newAccount.ToNewAccountResponseDto()
	return &response,nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService{
	return DefaultAccountService{repo}
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()

	if err!= nil{
		return nil, err
	}

	if req.IsWithdraw(){
		account, err := s.repo.ById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount){
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}

	}

	t := domain.Transaction{
		TransactionId: "",
		AccountId:     req.AccountId,
		Amount:        req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := transaction.ToDto()
	return &response, nil
}