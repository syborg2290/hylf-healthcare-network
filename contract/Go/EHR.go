package healthcare

type EHR struct {
	EHRID        string     `json:"EHRID"`
	EHRTime      string     `json:"EHRTime"`
	EHRPatientID string     `json:"EHRPatientID"`
	EHRDoctorID  string     `json:"EHRDoctorID"`
	EHRComments  string     `json:"EHRComments"`
	EHRMedicines []Medicine `json:"EHRMedicines"`
}

func (ehr *EHR) getEHRTime() string {
	return ehr.EHRTime
}

/*func (ehr *EHR) getEHRDate() string {
	return ehr.EHRDate
}*/
func (ehr *EHR) getEHRPatientID() string {
	return ehr.EHRPatientID
}
func (ehr *EHR) getEHRDoctorID() string {
	return ehr.EHRDoctorID
}
func (ehr *EHR) getEHRComments() string {
	return ehr.EHRComments
}
func (ehr *EHR) getEHRID() string {
	return ehr.EHRID
}
func (ehr *EHR) getEHRMedicines() []Medicine {
	return ehr.EHRMedicines
}

func (ehr *EHR) setEHRTime(t string) {
	ehr.EHRTime = t
}

/*func (ehr *EHR) setEHRDate(date string) {
	ehr.EHRTime = time
}*/
func (ehr *EHR) setEHRPatientID(patientID string) {
	ehr.EHRPatientID = patientID
}
func (ehr *EHR) setEHRDoctorID(doctorID string) {
	ehr.EHRDoctorID = doctorID
}
func (ehr *EHR) setEHRComments(comment string) {
	ehr.EHRComments = comment
}
func (ehr *EHR) setEHRMedicines(medicines []Medicine) {
	ehr.EHRMedicines = medicines
}
