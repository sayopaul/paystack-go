package paystack

import (
	"fmt"
	"time"
)

// BankAccountService handles operations related to dedicated bank accounts
// For more details see https://paystack.com/docs/api/dedicated-virtual-account#create
type BankAccountService service

// BankAccountRequest represents a request to create a bank account.
type BankAccountRequest struct {
	Email         string  `json:"email,omitempty"`
	FirstName     float32 `json:"first_name,omitempty"`
	LastName      string  `json:"last_name,omitempty"`
	Phone         string  `json:"phone,omitempty"`
	PreferredBank string  `json:"preferred_bank,omitempty"`
	Country       string  `json:"country,omitempty"`
}

// BankAccount is the resource representing your Paystack transfer.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account
type BankAccountResponse struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type DedicatedBankAccount struct {
	Customer struct {
		ID                       int         `json:"id"`
		FirstName                string      `json:"first_name"`
		LastName                 string      `json:"last_name"`
		Email                    string      `json:"email"`
		CustomerCode             string      `json:"customer_code"`
		Phone                    string      `json:"phone"`
		RiskAction               string      `json:"risk_action"`
		InternationalFormatPhone interface{} `json:"international_format_phone"`
	} `json:"customer"`
	Bank struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"bank"`
	ID            int       `json:"id"`
	AccountName   string    `json:"account_name"`
	AccountNumber string    `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Currency      string    `json:"currency"`
	SplitConfig   struct {
		Subaccount string `json:"subaccount"`
	} `json:"split_config"`
	Active   bool `json:"active"`
	Assigned bool `json:"assigned"`
}

type SingleDedicatedBankAccount struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Transactions          []interface{} `json:"transactions"`
		Subscriptions         []interface{} `json:"subscriptions"`
		Authorizations        []interface{} `json:"authorizations"`
		FirstName             interface{}   `json:"first_name"`
		LastName              interface{}   `json:"last_name"`
		Email                 string        `json:"email"`
		Phone                 interface{}   `json:"phone"`
		Metadata              interface{}   `json:"metadata"`
		Domain                string        `json:"domain"`
		CustomerCode          string        `json:"customer_code"`
		RiskAction            string        `json:"risk_action"`
		ID                    int           `json:"id"`
		Integration           int           `json:"integration"`
		CreatedAt             time.Time     `json:"createdAt"`
		UpdatedAt             time.Time     `json:"updatedAt"`
		CreatedAtSnake        time.Time     `json:"created_at"`
		UpdatedAtSnake        time.Time     `json:"updated_at"`
		TotalTransactions     int           `json:"total_transactions"`
		TotalTransactionValue []interface{} `json:"total_transaction_value"`
		DedicatedAccount      struct {
			ID            int       `json:"id"`
			AccountName   string    `json:"account_name"`
			AccountNumber string    `json:"account_number"`
			CreatedAt     time.Time `json:"created_at"`
			UpdatedAt     time.Time `json:"updated_at"`
			Currency      string    `json:"currency"`
			Active        bool      `json:"active"`
			Assigned      bool      `json:"assigned"`
			Provider      struct {
				ID           int    `json:"id"`
				ProviderSlug string `json:"provider_slug"`
				BankID       int    `json:"bank_id"`
				BankName     string `json:"bank_name"`
			} `json:"provider"`
			Assignment struct {
				AssigneeID   int    `json:"assignee_id"`
				AssigneeType string `json:"assignee_type"`
				AccountType  string `json:"account_type"`
				Integration  int    `json:"integration"`
			} `json:"assignment"`
		} `json:"dedicated_account"`
	} `json:"data"`
}

type BankAccountList struct {
	Meta   ListMeta
	Values []DedicatedBankAccount `json:"data,omitempty"`
}

// create creates a new banka ccount
// For more details see https://paystack.com/docs/api/dedicated-virtual-account#create
func (s *BankAccountService) Create(req *BankAccountRequest) (*BankAccountResponse, error) {
	bankAccount := &BankAccountResponse{}
	err := s.client.Call("POST", "/dedicated_account/assign", req, bankAccount)
	return bankAccount, err
}

// Get returns the details of a bank account.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account#fetch
func (s *BankAccountService) Get(idCode string) (*SingleDedicatedBankAccount, error) {
	u := fmt.Sprintf("/dedicated_account/%s", idCode)
	dedicatedBankAccount := &SingleDedicatedBankAccount{}
	err := s.client.Call("GET", u, nil, dedicatedBankAccount)
	return dedicatedBankAccount, err
}

// List returns a list of bank accounts.
// For more details see https://paystack.com/docs/api/dedicated-virtual-account#list
func (s *BankAccountService) List() (*BankAccountList, error) {
	return s.ListN(10, 0)
}

// ListN returns a list of bank accounts
// For more details see https://paystack.com/docs/api/dedicated-virtual-account#list
func (s *BankAccountService) ListN(count, offset int) (*BankAccountList, error) {
	u := paginateURL("/dedicated_account?active=true&currency=NGN", count, offset)
	bankAccounts := &BankAccountList{}
	err := s.client.Call("GET", u, nil, bankAccounts)
	return bankAccounts, err
}
