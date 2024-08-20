package credit

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateApplicationService(t *testing.T) {
    application := CreditApplication{
        CompanyName:     "Empresa ABC",
        CNPJ:            "12.345.678/0001-99",
        RequestedAmount: 100000.00,
    }
    createdApplication := CreateApplicationService(application)

    assert.NotNil(t, createdApplication)
    assert.Equal(t, createdApplication.CompanyName, application.CompanyName)
    assert.Equal(t, createdApplication.CNPJ, application.CNPJ)
    assert.Equal(t, createdApplication.RequestedAmount, application.RequestedAmount)
    assert.Equal(t, "Pending", createdApplication.Status)
}

func TestApproveApplicationService(t *testing.T) {
    application := CreateApplicationService(CreditApplication{
        CompanyName:     "Empresa XYZ",
        CNPJ:            "98.765.432/0001-10",
        RequestedAmount: 50000.00,
    })

    approvedApplication := ApproveApplicationService(application.ID, 45000.00)
    assert.NotNil(t, approvedApplication)
    assert.Equal(t, "Approved", approvedApplication.Status)
    assert.Equal(t, 45000.00, *approvedApplication.ApprovedAmount)
}

func TestRejectApplicationService(t *testing.T) {
    application := CreateApplicationService(CreditApplication{
        CompanyName:     "Empresa DEF",
        CNPJ:            "23.456.789/0001-88",
        RequestedAmount: 75000.00,
    })

    rejectedApplication := RejectApplicationService(application.ID)
    assert.NotNil(t, rejectedApplication)
    assert.Equal(t, "Rejected", rejectedApplication.Status)
}
