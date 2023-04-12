package healthcare

type Doctor struct {
	DoctorName         string `json:"DoctorName"`
	DoctorContact      string `json:"DoctorContact"`
	DoctorSpeciality   string `json:"DoctorSpeciality"`
	DoctorID           string `json:"DoctorID"`
	DoctorAvailability bool   `json:"DoctorAvailability"`
}

func (doctor *Doctor) getDoctorName() string {
	return doctor.DoctorName
}
func (doctor *Doctor) getDoctorContact() string {
	return doctor.DoctorContact
}
func (doctor *Doctor) getDoctorSpeciality() string {
	return doctor.DoctorSpeciality
}
func (doctor *Doctor) getDoctorID() string {
	return doctor.DoctorID
}
func (doctor *Doctor) getDoctorAvailability() bool {
	return doctor.DoctorAvailability
}

func (doctor *Doctor) setDoctorName(name string) {
	doctor.DoctorName = name
}
func (doctor *Doctor) setDoctorContact(contact string) {
	doctor.DoctorContact = contact
}
func (doctor *Doctor) setDoctorSpeciality(speciality string) {
	doctor.DoctorSpeciality = speciality
}
func (doctor *Doctor) setDoctorID(doctorID string) {
	doctor.DoctorID = doctorID
}
func (doctor *Doctor) setDoctorAvailability(availability bool) {
	doctor.DoctorAvailability = availability
}

func newDoctor(doctorID string, name string, contact string, speciality string) Doctor {
	var newDoc Doctor = Doctor{
		DoctorID:         doctorID,
		DoctorName:       name,
		DoctorContact:    contact,
		DoctorSpeciality: speciality,
	}
	return newDoc
}

/*func main() {
	doctors := make([]Doctor, 0)
	id := "D" + strconv.FormatInt(int64(len(doctors)+1), 10)
	var doctor = newDoctor(id, "kasun", "985473XXXX", "Orthopadeic")
	doctors = append(doctors, doctor)
	doctors = append(doctors, newDoctor("D2", "tharindu", "985473XXXX", "Respiratory Physician"))
	fmt.Println(doctors[1].DoctorID)
}*/
