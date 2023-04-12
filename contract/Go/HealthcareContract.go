package healthcare

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	patient := Patient{
		PatientID:          "P1",
		PatientName:        "Kshitij",
		PatientAge:         20,
		PatientAddress:     "Greater Noida",
		PatientAadhar:      "63759927XXXX",
		PatientAppoitments: []string{},
		PatientEHRs:        []string{},
		PatientOrders:      []string{},
	}

	patientJSON, err := json.Marshal(patient)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(patient.PatientID, patientJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	doctors := []Doctor{
		{DoctorID: "D1", DoctorName: "Vijay", DoctorContact: "985473XXXX", DoctorSpeciality: "Respiratory Physician"},
		{DoctorID: "D2", DoctorName: "Rajeev", DoctorContact: "674533XXXX", DoctorSpeciality: "Orthopedic"},
		{DoctorID: "D3", DoctorName: "Nidhi", DoctorContact: "780473XXXX", DoctorSpeciality: "Gyncecoligist"},
		{DoctorID: "D4", DoctorName: "Nikhil", DoctorContact: "9325767XXXX", DoctorSpeciality: "Cardiologist"},
	}
	ic := InsuranceCompany{
		InsuranceCompanyName:    "Good Life",
		InsuranceCompanyContact: "980890XXXX",
		InsuranceCompanyAddress: "Mumbai",
	}
	eph := Epharma{
		EPharmaName:    "Life Medicine",
		EPharmaContact: "756980XXXX",
		EPharmaAddress: "New Delhi",
	}

	for _, doctor := range doctors {
		doctorJSON, err := json.Marshal(doctor)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(doctor.DoctorID, doctorJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	icJSON, err := json.Marshal(ic)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(ic.InsuranceCompanyName, icJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	ephJSON, err := json.Marshal(eph)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(eph.EPharmaName, ephJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	var appointmentCount int64 = 0
	acJSON, err := json.Marshal(appointmentCount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState("AppointmentCount", acJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	ehrCount := 0
	ehrcJSON, err := json.Marshal(ehrCount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState("EHRCount", ehrcJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	orderCount := 0
	orderCountJSON, err := json.Marshal(orderCount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState("OrderCount", orderCountJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	fmt.Println("Patient: ", patient)
	fmt.Println("Doctors:", doctors)
	fmt.Println("InsuranceCompany: ", ic)
	fmt.Println("Epharma: ", eph)

	return nil
}

func (s *SmartContract) BookAppointment(ctx contractapi.TransactionContextInterface, patientID string, shareEHR bool, useInsurance bool, t string) error {

	patientJSON, err := ctx.GetStub().GetState(patientID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if patientJSON == nil {
		return fmt.Errorf("the asset %s does not exist", patientID)
	}
	appointmentCountJSON, err := ctx.GetStub().GetState("AppointmentCount")
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if appointmentCountJSON == nil {
		return fmt.Errorf("the asset %s does not exist", "AppointmentCount")
	}

	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return err
	}
	var appointmentCount int64
	err = json.Unmarshal(appointmentCountJSON, &appointmentCount)
	if err != nil {
		return err
	}
	appointmentCount = appointmentCount + 1
	newAppointmentID := "A" + strconv.FormatInt(appointmentCount, 10)
	newAppointment := Appointment{
		AppointmentID:           newAppointmentID,
		AppointmentPatientID:    patient.PatientID,
		AppointmentBookingTime:  t,
		AppointmentTime:         "",
		AppointmentStatus:       false,
		AppointmentShareEHR:     shareEHR,
		AppointmentUseInsurance: useInsurance,
		AppointmentAmount:       0,
		AppointmentDoctorID:     "",
		AppointmentPaid:         false,
	}
	newAppointmentJSON, err := json.Marshal(newAppointment)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(newAppointment.AppointmentID, newAppointmentJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	appointmentCountJSON, err = json.Marshal(appointmentCount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState("AppointmentCount", appointmentCountJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	patient.addPatientAppointment(newAppointmentID)
	patientJSON, err = json.Marshal(patient)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(patient.PatientID, patientJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

func (s *SmartContract) ApproveAppointment(ctx contractapi.TransactionContextInterface, appointmentID string, doctorID string, t string) error {
	appointmentJSON, err := ctx.GetStub().GetState(appointmentID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if appointmentJSON == nil {
		return fmt.Errorf("the asset %s does not exist", appointmentID)
	}
	var appointment Appointment
	err = json.Unmarshal(appointmentJSON, &appointment)
	if err != nil {
		return err
	}

	appointment.setAppointmentDoctorID(doctorID)
	appointment.setAppointmentStatus(true)
	appointment.setAppointmentTime(t)
	appointmentJSON, err = json.Marshal(appointment)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(appointment.AppointmentID, appointmentJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

func (s *SmartContract) IssueEHR(ctx contractapi.TransactionContextInterface, appointmentID string, comments string, medicines []Medicine, t string) error {
	appointmentJSON, err := ctx.GetStub().GetState(appointmentID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if appointmentJSON == nil {
		return fmt.Errorf("the asset %s does not exist", appointmentID)
	}
	var appointment Appointment
	err = json.Unmarshal(appointmentJSON, &appointment)
	if err != nil {
		return err
	}
	patientJSON, err := ctx.GetStub().GetState(appointment.AppointmentPatientID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if patientJSON == nil {
		return fmt.Errorf("the asset %s does not exist", appointment.AppointmentPatientID)
	}
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return err
	}
	ehrCountJSON, err := ctx.GetStub().GetState("EHRCount")
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if ehrCountJSON == nil {
		return fmt.Errorf("the asset %s does not exist", "EHRCount")
	}
	var ehrCount int64
	err = json.Unmarshal(ehrCountJSON, &ehrCount)
	if err != nil {
		return err
	}
	ehrCount = ehrCount + 1
	ehrid := "EHR" + strconv.FormatInt(ehrCount, 10)
	newEHR := EHR{
		EHRID:        ehrid,
		EHRTime:      t,
		EHRPatientID: appointment.AppointmentPatientID,
		EHRDoctorID:  appointment.AppointmentDoctorID,
		EHRComments:  comments,
		EHRMedicines: medicines,
	}
	patient.addPatientEHR(ehrid)
	newEHRJSON, err := json.Marshal(newEHR)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(newEHR.EHRID, newEHRJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	ehrCountJSON, err = json.Marshal(ehrCount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState("EHRCount", ehrCountJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	patientJSON, err = json.Marshal(patient)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(patient.PatientID, patientJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

func (s *SmartContract) PlaceOrder(ctx contractapi.TransactionContextInterface, ehrID string, t string) error {
	ehrJSON, err := ctx.GetStub().GetState(ehrID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if ehrJSON == nil {
		return fmt.Errorf("the asset %s does not exist", ehrID)
	}
	var ehr EHR
	err = json.Unmarshal(ehrJSON, &ehr)
	if err != nil {
		return err
	}
	patientJSON, err := ctx.GetStub().GetState(ehr.EHRPatientID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if patientJSON == nil {
		return fmt.Errorf("the asset %s does not exist", ehr.EHRPatientID)
	}
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return err
	}
	ordercJSON, err := ctx.GetStub().GetState("OrderCount")
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if ordercJSON == nil {
		return fmt.Errorf("the asset %s does not exist", "OrderCount")
	}
	var orderCount int64
	err = json.Unmarshal(ordercJSON, &orderCount)
	if err != nil {
		return err
	}
	orderCount = orderCount + 1
	orderID := "O" + strconv.FormatInt(orderCount, 10)
	newOrder := Order{
		OrderID:        orderID,
		OrderTime:      t,
		OrderMedicines: ehr.EHRMedicines,
		OrderCompleted: false,
	}
	for _, medicine := range newOrder.OrderMedicines {
		newOrder.OrderAmount = newOrder.OrderAmount + medicine.MedicineCost
	}

	newOrderJSON, err := json.Marshal(newOrder)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(newOrder.OrderID, newOrderJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	ordercJSON, err = json.Marshal(orderCount)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState("OrderCount", ordercJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	patient.addPatientOrder(orderID)
	patientJSON, err = json.Marshal(patient)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(patient.PatientID, patientJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

func (s *SmartContract) CompleteOrder(ctx contractapi.TransactionContextInterface, orderID string) error {
	orderJSON, err := ctx.GetStub().GetState(orderID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if orderJSON == nil {
		return fmt.Errorf("the asset %s does not exist", orderID)
	}
	var order Order
	err = json.Unmarshal(orderJSON, &order)
	if err != nil {
		return err
	}
	order.setOrderCompleted(true)
	orderJSON, err = json.Marshal(order)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(order.OrderID, orderJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

func (s *SmartContract) RaiseBill(ctx contractapi.TransactionContextInterface, appointmentID string, amount float64) error {
	appointmentJSON, err := ctx.GetStub().GetState(appointmentID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if appointmentJSON == nil {
		return fmt.Errorf("the asset %s does not exist", appointmentID)
	}
	var appointment Appointment
	err = json.Unmarshal(appointmentJSON, &appointment)
	if err != nil {
		return err
	}
	appointment.setAppointmentAmount(amount)
	appointment.setAppointmentPaid(false)
	appointmentJSON, err = json.Marshal(appointment)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(appointment.AppointmentID, appointmentJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}
func (s *SmartContract) ClearBill(ctx contractapi.TransactionContextInterface, appointmentID string) error {
	appointmentJSON, err := ctx.GetStub().GetState(appointmentID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if appointmentJSON == nil {
		return fmt.Errorf("the asset %s does not exist", appointmentID)
	}
	var appointment Appointment
	err = json.Unmarshal(appointmentJSON, &appointment)
	if err != nil {
		return err
	}
	appointment.setAppointmentPaid(true)
	appointmentJSON, err = json.Marshal(appointment)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(appointment.AppointmentID, appointmentJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

func (s *SmartContract) ReadPatient(ctx contractapi.TransactionContextInterface, patientID string) (*Patient, error) {
	patientJSON, err := ctx.GetStub().GetState(patientID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if patientJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", patientID)
	}
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *SmartContract) ReadAppointment(ctx contractapi.TransactionContextInterface, appointmentID string) (*Appointment, error) {
	appointmentJSON, err := ctx.GetStub().GetState(appointmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if appointmentJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", appointmentID)
	}
	var appointment Appointment
	err = json.Unmarshal(appointmentJSON, &appointment)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (s *SmartContract) ReadOrder(ctx contractapi.TransactionContextInterface, orderID string) (*Order, error) {
	orderJSON, err := ctx.GetStub().GetState(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if orderJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", orderID)
	}
	var order Order
	err = json.Unmarshal(orderJSON, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *SmartContract) ReadEHR(ctx contractapi.TransactionContextInterface, ehrID string) (*EHR, error) {
	ehrJSON, err := ctx.GetStub().GetState(ehrID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if ehrJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", ehrID)
	}
	var ehr EHR
	err = json.Unmarshal(ehrJSON, &ehr)
	if err != nil {
		return nil, err
	}

	return &ehr, nil
}
