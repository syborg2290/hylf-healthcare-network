package healthcare

type Medicine struct {
	MedicineName      string  `json:"MedicineName"`
	MedicineAmount    float32 `json:"MedicineAmount"`
	MedicineNooOfDays int     `json:"MedicineNooOfDays"`
	MedicineFrequency int     `json:"MedicineFrequency"`
	MedicineActive    bool    `json:"MedicineActive"`
	MedicineCost      float64 `json:"MedicineCost"`
}

func (medicine *Medicine) getMedicineName() string {
	return medicine.MedicineName
}
func (medicine *Medicine) getMedicineAmount() float32 {
	return medicine.MedicineAmount
}
func (medicine *Medicine) getMedicineNoOfDays() int {
	return medicine.MedicineNooOfDays
}
func (medicine *Medicine) getMedicineFrequency() int {
	return medicine.MedicineFrequency
}
func (medicine *Medicine) getMedicineActive() bool {
	return medicine.MedicineActive
}
func (medicine *Medicine) getMedicineCost() float64 {
	return medicine.MedicineCost
}

func (medicine *Medicine) setMedicineName(name string) {
	medicine.MedicineName = name
}
func (medicine *Medicine) setMedicineAmount(amount float32) {
	medicine.MedicineAmount = amount
}
func (medicine *Medicine) setMedicineNoOfDays(nod int) {
	medicine.MedicineNooOfDays = nod
}
func (medicine *Medicine) setMedicineFrequency(frequency int) {
	medicine.MedicineFrequency = frequency
}
func (medicine *Medicine) setMedicineActive(active bool) {
	medicine.MedicineActive = active
}
func (medicine *Medicine) setMedicineCost(cost float64) {
	medicine.MedicineCost = cost
}
