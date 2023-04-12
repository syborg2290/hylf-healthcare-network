package healthcare

type Appointment struct {
	AppointmentID           string  `json:"AppointmentID"`
	AppointmentBookingTime  string  `json:"AppointmentBookingTime"`
	AppointmentTime         string  `json:"AppointmentTime"`
	AppointmentDoctorID     string  `json:"AppointmentDoctorID"`
	AppointmentPatientID    string  `json:"AppointmentPatientID"`
	AppointmentAmount       float64 `json:"AppointmentAmount"`
	AppointmentPaid         bool    `json:"AppointmentPaid"`
	AppointmentStatus       bool    `json:"AppointmentStatus"`
	AppointmentShareEHR     bool    `json:"AppointmentShareEHR"`
	AppointmentUseInsurance bool    `json:"AppointmentUseInsurance"`
}

func (appointment *Appointment) getAppointmentID() string {
	return appointment.AppointmentID
}
func (appointment *Appointment) getAppointmentShareEHR() bool {
	return appointment.AppointmentShareEHR
}
func (appointment *Appointment) getAppointmentUseInsurance() bool {
	return appointment.AppointmentUseInsurance
}
func (appointment *Appointment) getAppointmentPaid() bool {
	return appointment.AppointmentPaid
}
func (appointment *Appointment) getAppointmentBookingTime() string {
	return appointment.AppointmentBookingTime
}
func (appointment *Appointment) getAppointmentTime() string {
	return appointment.AppointmentTime
}
func (appointment *Appointment) getAppointmentDoctorID() string {
	return appointment.AppointmentDoctorID
}
func (appointment *Appointment) getAppointmentPatientID() string {
	return appointment.AppointmentPatientID
}
func (appointment *Appointment) getAppointmentAmount() float64 {
	return appointment.AppointmentAmount
}
func (appointment *Appointment) getAppointmentStatus() bool {
	return appointment.AppointmentStatus
}

func (appointment *Appointment) setAppointmentID(aid string) {
	appointment.AppointmentID = aid
}

func (appointment *Appointment) setAppointmentBookingTime(btime string) {
	appointment.AppointmentBookingTime = btime
}
func (appointment *Appointment) setAppointmentTime(at string) {
	appointment.AppointmentTime = at
}
func (appointment *Appointment) setAppointmentDoctorID(doctorID string) {
	appointment.AppointmentDoctorID = doctorID
}
func (appointment *Appointment) setAppointmentPatientID(patientID string) {
	appointment.AppointmentPatientID = patientID
}
func (appointment *Appointment) setAppointmentAmount(amount float64) {
	appointment.AppointmentAmount = amount
}
func (appointment *Appointment) setAppointmentStatus(status bool) {
	appointment.AppointmentStatus = status
}
func (appointment *Appointment) setAppointmentShareEHR(share bool) {
	appointment.AppointmentShareEHR = share
}
func (appointment *Appointment) setAppointmentUseInsurance(use bool) {
	appointment.AppointmentUseInsurance = use
}
func (appointment *Appointment) setAppointmentPaid(paid bool) {
	appointment.AppointmentPaid = paid
}
