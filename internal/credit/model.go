package credit

import "time"

type CreditApplication struct {
    ID               int       `json:"id"`
    CompanyName      string    `json:"company_name"`
    CNPJ             string    `json:"cnpj"`
    RequestedAmount  float64   `json:"requested_amount"`
    ApprovedAmount   *float64  `json:"approved_amount,omitempty"`
    Status           string    `json:"status"`
    ApplicationDate  time.Time `json:"application_date"`
    ApprovalDate     *time.Time `json:"approval_date,omitempty"`
}
