package healthcare

type Patient struct {
	PatientName        string   `json:"PatientName"`
	PatientAge         int      `json:"PatientAge"`
	PatientAadhar      string   `json:"PatientAadhar"`
	PatientAddress     string   `json:"PatientAddress"`
	PatientID          string   `json:"PatientID"`
	PatientOrders      []string `json:"PatientMedicines"`
	PatientEHRs        []string `json:"PatientEHRs"`
	PatientAppoitments []string `json:"PatientAppoitments"`
}

//type Contract struct {
//	contractapi.Contract
//}

func (p *Patient) getPatientName() string {
	return p.PatientName
}
func (p *Patient) getPatientAge() int {
	return p.PatientAge
}
func (p *Patient) getPatientAadhar() string {
	return p.PatientAadhar
}
func (p *Patient) getPatientAddress() string {
	return p.PatientAddress
}
func (p *Patient) getPatientID() string {
	return p.PatientID
}
func (p *Patient) getPatientEHRs() []string {
	return p.PatientEHRs
}
func (p *Patient) getPatientAppoitments() []string {
	return p.PatientAppoitments
}
func (p *Patient) getPatientOrders() []string {
	return p.PatientOrders
}
func (p *Patient) setPatientName(name string) {
	p.PatientName = name
}
func (p *Patient) setPatientAge(age int) {
	p.PatientAge = age
}
func (p *Patient) setPatientAadhar(aadhar string) {
	p.PatientAadhar = aadhar
}
func (p *Patient) setPatientAddress(address string) {
	p.PatientAddress = address
}
func (p *Patient) setPatientID(patientID string) {
	p.PatientID = patientID
}
func newPatient(patientID string, name string, age int, aadhar string, address string) Patient {
	var newp Patient = Patient{
		PatientID:      patientID,
		PatientName:    name,
		PatientAge:     age,
		PatientAddress: address,
		PatientAadhar:  aadhar,
	}
	return newp
}
func (p *Patient) addPatientEHR(ehrid string) {
	p.PatientEHRs = append(p.PatientEHRs, ehrid)
}
func (p *Patient) addPatientAppointment(appointmentid string) {
	p.PatientAppoitments = append(p.PatientAppoitments, appointmentid)
}
func (p *Patient) addPatientOrder(orderid string) {
	p.PatientOrders = append(p.PatientOrders, orderid)
}

/*func main() {

	var p Patient = newPatient("P1", "danith", 20, "Matara", "63759927XXXX")
	fmt.Println(p.getPatientName())
	fmt.Println(p.getPatientAge())
	fmt.Println(p.getPatientAadhar())
	fmt.Println(p.getPatientAddress())
	fmt.Println(p.getPatientID())
	p.setPatientAddress("Mumbai")
	fmt.Println(p.getPatientAddress())
	//fmt.Println(ep._name)
	s := time.Now()
	fmt.Println(s)
}*/
