package model

type Health_insurance struct {
	InsuranceID                string `json:"web-scraper-start-url" gorm:"primaryKey" gorm:"column:web-scraper-start-url" validate:"required"`
	Name                       string `json:"name" gorm:"column:name"`
	Provider                   string `json:"provider" gorm:"column:provider"`
	NetworkHospitals           string `json:"network_hospitals" gorm:"column:network_hospitals"`
	ClaimSettlementRatio       string `json:"claim_settlement_ratio" gorm:"column:claim_settlement_ratio"`
	CoPayment                  string `json:"co_payment" gorm:"column:co_payment"`
	RoomRent                   string `json:"room_rent" gorm:"column:room_rent"`
	DiseaseSubLimit            string `json:"disease_sub_limit" gorm:"column:disease_sub_limit"`
	PreExistingDiseasesWaiting string `json:"pre_existing_diseases_waiting" gorm:"column:pre_existing_diseases_waiting"`
	NoClaimBonus               string `json:"no_claim_bonus" gorm:"column:no_claim_bonus"`
	HealthCheckUp              string `json:"health_check_up" gorm:"column:health_check_up"`
}
