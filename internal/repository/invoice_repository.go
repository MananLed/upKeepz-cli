package repository

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/MananLed/upKeepz-cli/constants"
	"github.com/MananLed/upKeepz-cli/internal/model"
	"github.com/MananLed/upKeepz-cli/internal/utils"
)

type InvoiceRepositoryInterface interface {
	SaveInvoice(model.Invoice) error
	GetInvoiceByMonthAndYear(string, string) (*model.Invoice, error)
	GetInvoicesByYear(string) ([]model.Invoice, error)
}

type InvoiceRepository struct {
	mu sync.Mutex
}

func (r *InvoiceRepository) loadInvoices() ([]model.Invoice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	fileData, err := os.ReadFile(string(constants.InvoiceDataPath))

	if err != nil {
		if os.IsNotExist(err) {
			return []model.Invoice{}, nil
		}
		return nil, err
	}

	var invoices []model.Invoice

	if err := json.Unmarshal(fileData, &invoices); err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *InvoiceRepository) SaveInvoice(invoice model.Invoice) error {
	invoices, err := r.loadInvoices()

	if err != nil {
		return err
	}

	invoice.ID = utils.GenerateUUID()
	invoices = append(invoices, invoice)

	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := json.MarshalIndent(invoices, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(string(constants.InvoiceDataPath), data, 0644)
}

func (r *InvoiceRepository) GetInvoiceByMonthAndYear(month string, year string) (*model.Invoice, error) {
	invoices, err := r.loadInvoices()

	if err != nil {
		return nil, err
	}

	for _, i := range invoices {
		if i.Month == month && i.Year == year {
			return &i, nil
		}
	}

	return nil, errors.New("invoice not found")
}

func (r *InvoiceRepository) GetInvoicesByYear(year string) ([]model.Invoice, error) {
	invoices, err := r.loadInvoices()

	if err != nil {
		return nil, err
	}

	var invoicesOfYear []model.Invoice

	for _, i := range invoices {
		if i.Year == year {
			invoicesOfYear = append(invoicesOfYear, i)
		}
	}
	return invoicesOfYear, nil
}
