package midtrans

// ItemDetail : Represent the transaction details
type ItemDetail struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Qty          int32  `json:"quantity"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	MerchantName string `json:"merchant_name,omitempty"`
}

// CustAddress : Represent the customer address
type CustAddress struct {
	FName       string `json:"first_name"`
	LName       string `json:"last_name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

// CustDetail : Represent the customer detail
type CustDetail struct {
	// first name
	FName string `json:"first_name,omitempty"`

	// last name
	LName string `json:"last_name,omitempty"`

	Email    string       `json:"email,omitempty"`
	Phone    string       `json:"phone,omitempty"`
	BillAddr *CustAddress `json:"billing_address,omitempty"`
	ShipAddr *CustAddress `json:"customer_address,omitempty"`
}

// TransactionDetails : Represent transaction details
type TransactionDetails struct {
	OrderID  string `json:"order_id"`
	GrossAmt int64  `json:"gross_amount"`
}

// CreditCardDetail : Represent credit card detail
type CreditCardDetail struct {
	SaveCard        bool     `json:"save_card,omitempty"`
	Secure          bool     `json:"secure,omitempty"`
	Channel         string   `json:"channel,omitempty"`
	Bank            string   `json:"bank,omitempty"`
	TokenID         string   `json:"token_id,omitempty"`
	Bins            []string `json:"bins,omitempty"`
	InstallmentTerm int8     `json:"installment_term,omitempty"`
	Type            string   `json:"type,omitempty"`
	// indicate if generated token should be saved for next charge
	SaveTokenID          bool   `json:"save_token_id,omitempty"`
	SavedTokenIDExpireAt string `json:"saved_token_id_expired_at,omitempty"`
}

// PermataBankTransferDetail : Represent Permata bank_transfer detail
type PermataBankTransferDetail struct {
	Bank          Bank   `json:"bank,omitempty"`
	VaNumber      string `json:"va_number,omitempty"`
	RecipientName string `json:"recipient_name ,omitempty"`
}

// BNIBankTransferDetail : Represent BNI bank_transfer detail
type BNIBankTransferDetail struct {
	Bank     Bank   `json:"bank,omitemtpy"`
	VaNumber string `json:"va_number,omitempty"`
}

// BCABankTransferLangDetail : Represent BCA bank_transfer lang detail
type BCABankTransferLangDetail struct {
	LangID string `json:"id,omitempty"`
	LangEN string `json:"en,omitempty"`
}

/*
   Example of usage syntax:
   midtrans.BCABankTransferDetail{
       FreeText: {
           Inquiry: []midtrans.BCABankTransferLangDetail{
               {
                   LangEN: "Test",
                   LangID: "Coba",
               },
           },
       },
   }
*/

// BCABankTransferDetailFreeText : Represent BCA bank_transfer detail free_text
type BCABankTransferDetailFreeText struct {
	Inquiry []BCABankTransferLangDetail `json:"inquiry,omitempty"`
	Payment []BCABankTransferLangDetail `json:"payment,omitempty"`
}

// BCABankTransferDetail : Represent BCA bank_transfer detail
type BCABankTransferDetail struct {
	Bank           Bank                          `json:"bank"`
	VaNumber       string                        `json:"va_number,omitempty"`
	SubCompanyCode string                        `json:"sub_company_code,omitempty"`
	FreeText       BCABankTransferDetailFreeText `json:"free_text"`
}

// MandiriBillBankTransferDetail : Represent Mandiri Bill bank_transfer detail
type MandiriBillBankTransferDetail struct {
	BillInfo1 string `json:"bill_info1,omitempty"`
	BillInfo2 string `json:"bill_info2,omitempty"`
}

// BankTransferDetail : Represent bank_transfer detail
type BankTransferDetail struct {
	Bank     Bank                           `json:"bank,omitempty"`
	VaNumber string                         `json:"va_number,omitempty"`
	FreeText *BCABankTransferDetailFreeText `json:"free_text,omitempty"`
	*MandiriBillBankTransferDetail
}

// BCAKlikPayDetail : Represent Internet Banking for BCA KlikPay
type BCAKlikPayDetail struct {
	// 1 = normal, 2 = installment, 3 = normal + installment
	Type    string `json:"type"`
	Desc    string `json:"description"`
	MiscFee int64  `json:"misc_fee,omitempty"`
}

// BCAKlikBCADetail : Represent BCA KlikBCA detail
type BCAKlikBCADetail struct {
	Desc   string `json:"description"`
	UserID string `json:"user_id"`
}

// MandiriClickPayDetail : Represent Mandiri ClickPay detail
type MandiriClickPayDetail struct {
	CardNumber string `json:"card_number"`
	Input1     string `json:"input1"`
	Input2     string `json:"input2"`
	Input3     string `json:"input3"`
	Token      string `json:"token"`
}

// CIMBClicksDetail : Represent CIMB Clicks detail
type CIMBClicksDetail struct {
	Desc string `json:"description"`
}

// TelkomselCashDetail : Represent Telkomsel Cash detail
type TelkomselCashDetail struct {
	Promo      bool   `json:"promo"`
	IsReversal int8   `json:"is_reversal"`
	Customer   string `json:"customer"`
}

// IndosatDompetkuDetail : Represent Indosat Dompetku detail
type IndosatDompetkuDetail struct {
	MSISDN string `json:"msisdn"`
}

// MandiriEcashDetail : Represent Mandiri e-Cash detail
type MandiriEcashDetail struct {
	Desc string `json:"description"`
}

// ConvStoreDetail : Represent cstore detail
type ConvStoreDetail struct {
	Store   string `json:"store"`
	Message string `json:"message"`
}

// Callbacks : Represent callbacks url after transaction successfully paid
type Callbacks struct {
	Finish string `json:"finish,omitempty"`
}

// Expiry : Represent transaction expiry time
type Expiry struct {
	StartTime string `json:"start_time,omitempty"`
	Unit      string `json:"unit,omitempty"`
	Duration  int64  `json:"duration,omitempty"`
}

// ChargeReq : Represent Charge request payload
type ChargeReq struct {
	PaymentType                   PaymentType                    `json:"payment_type"`
	TransactionDetails            TransactionDetails             `json:"transaction_details"`
	CreditCard                    *CreditCardDetail              `json:"credit_card,omitempty"`
	BankTransfer                  *BankTransferDetail            `json:"bank_transfer,omitempty"`
	MandiriBillBankTransferDetail *MandiriBillBankTransferDetail `json:"echannel,omitempty"`
	BCAKlikPay                    *BCAKlikPayDetail              `json:"bca_klikpay,omitempty"`
	BCAKlikBCA                    *BCAKlikBCADetail              `json:"bca_klikbca,omitempty"`
	MandiriClickPay               *MandiriClickPayDetail         `json:"mandiri_clickpay,omitempty"`
	MandiriEcash                  *MandiriEcashDetail            `json:"mandiri_ecash,omitempty"`
	CIMBClicks                    *CIMBClicksDetail              `json:"cimb_clicks,omitempty"`
	TelkomselCash                 *TelkomselCashDetail           `json:"telkomsel_cash,omitempty"`
	IndosatDompetku               *IndosatDompetkuDetail         `json:"indosat_dompetku,omitempty"`
	CustomerDetail                *CustDetail                    `json:"customer_details,omitempty"`
	ConvStore                     *ConvStoreDetail               `json:"cstore,omitempty"`
	Items                         *[]ItemDetail                  `json:"item_details,omitempty"`
	CustField1                    string                         `json:"custom_field1,omitempty"`
	CustField2                    string                         `json:"custom_field2,omitempty"`
	CustField3                    string                         `json:"custom_field3,omitempty"`
}

// SnapReq : Represent SNAP API request payload
type SnapReq struct {
	TransactionDetails  TransactionDetails         `json:"transaction_details"`
	EnabledPayments     []PaymentType              `json:"enabled_payments"`
	Items               *[]ItemDetail              `json:"item_details,omitempty"`
	CustomerDetail      *CustDetail                `json:"customer_details,omitempty"`
	CreditCard          *CreditCardDetail          `json:"credit_card,omitempty"`
	BCABankTransfer     *BCABankTransferDetail     `json:"bca_va,omitempty"`
	BNIBankTransfer     *BNIBankTransferDetail     `json:"bni_va,omitempty"`
	PermataBankTransfer *PermataBankTransferDetail `json:"permata_va,omitempty"`
	Expiry              *Expiry                    `json:"expiry,omitempty"`
	Callbacks           *Callbacks                 `json:"callbacks,omitempty"`
	CustomField1        string                     `json:"custom_field1,omitempty"`
	CustomField2        string                     `json:"custom_field2,omitempty"`
	CustomField3        string                     `json:"custom_field3,omitempty"`
}

// CaptureReq : Represent Capture request payload
type CaptureReq struct {
	TransactionID string  `json:"transaction_id"`
	GrossAmt      float64 `json:"gross_amount"`
}

// BeneficiariesReq : Represent create or Beneficiaries request payload
type BeneficiariesReq struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Bank      string `json:"bank"`
	AliasName string `json:"alias_name"`
	Email     string `json:"email,omitempty"`
}

// PayoutReq : Represent create payout
type PayoutReq struct {
	Payouts []PayoutDetails `json:"payouts"`
}

// PayoutDetails : Represent Payout Details
type PayoutDetails struct {
	BeneficiaryName    string `json:"beneficiary_name"`
	BeneficiaryAccount string `json:"beneficiary_account"`
	BneficiaryBank     string `json:"beneficiary_bank"`
	BeneficiaryEmail   string `json:"beneficiary_email,omitempty"`
	Amount             string `json:"amount"`
	Notes              string `json:"notes"`
	BankAccountID      string `json:"bank_account_id,omitempty"`
}

// ApprovePayoutReq : Represent Approve Payout Requests
type ApprovePayoutReq struct {
	ReferenceNos []string `json:"reference_nos"`
	OTP          string   `json:"otp"`
}

// RejectPayoutReq : Represent Reject Payout Requests
type RejectPayoutReq struct {
	ReferenceNos []string `json:"reference_nos"`
	RejectReason string   `json:"reject_reason"`
}
