package domain

import "time"

type User struct {
    UserID      int    `json:"user_id" validate:"omitempty,gt=0"`     // Unikt ID
    Name        string `json:"name" validate:"required,min=2,max=100"`        // Användarens fullständiga namn
    Email       string `json:"email" validate:"required,email"`       // E-post (unikt)
    PasswordHash string `json:"password_hash" validate:"required,min=8"` // Krypterat lösenord
    Role        string `json:"role" validate:"omitempty,oneof=Admin Bookkeeper Manager"`        // T.ex. "Admin", "Bookkeeper"
}

type Account struct {
    AccountNo   int    `json:"account_no"`    // Kontonummer (Primary Key)
    AccountName string `json:"account_name"`  // Kontots namn
    AccountGroup int   `json:"account_group"` // Kontogrupp (1-8 i BAS-planen)
    TaxStandard string `json:"tax_standard"`  // Standardmomskod (t.ex. "25%")
    Type        string `json:"type"`          // "P&L" (Resultat) eller "BS" (Balans)
    StandardSide string `json:"standard_side"` // "Debit" eller "Credit"
}

type LineItem struct {
    LineID       int     `json:"line_id"`         // Unikt ID för raden
    VoucherID    int     `json:"voucher_id"`      // Foreign Key till VoucherID
    AccountNo    int     `json:"account_no"`      // Foreign Key till Account
    DebitAmount  float64 `json:"debit_amount"`    // Belopp i Debet
    CreditAmount float64 `json:"credit_amount"`   // Belopp i Kredit
    TaxCode      int     `json:"tax_code"`        // Momskod (t.ex. 25, 12, 6, 0)
    ProjectID    int     `json:"project_id"`      // Valfri: För projektspårning
    CostCenterID int     `json:"cost_center_id"`  // Valfri: För kostnadsställe
}

type Voucher struct {
    VoucherID   int       `json:"voucher_id"`   // Unikt ID
    Date        time.Time `json:"date"`         // Datum då händelsen inträffade
    Description string    `json:"description"`  // Beskrivning av transaktionen
    Reference   string    `json:"reference"`    // Fakturanummer, kvitto-ID, etc.
    TotalAmount float64   `json:"total_amount"` // Totalbelopp
    Period      string    `json:"period"`       // Period (t.ex. "2025-01")
    CreatedBy   int       `json:"created_by"`   // Foreign Key till UserID
    Lines       []LineItem `json:"lines"`        // Lista över Verifikatraderna
}