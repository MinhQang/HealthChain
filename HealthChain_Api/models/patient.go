package models

type EmergencyContact struct {
	ID           uint64 `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	PhoneNumber  string `json:"phoneNumber"`
}

type Patient struct {
	ID                 uint64           `json:"id" gorm:"primary_key"`
	Name               string           `json:"name"`
	Age                uint64           `json:"age"`
	MedicalHistory     string           `json:"medicalHistory"`
	Address            string           `json:"address"`
	PhoneNumber        string           `json:"phoneNumber"`
	Email              string           `json:"email"`
	Allergies          string           `json:"allergies"`
	CurrentMedication  string           `json:"currentMedication"`
	EmergencyContactID uint64           `json:"emergencyContactId"`
	EmergencyContact   EmergencyContact `json:"emergencyContact"`
}
