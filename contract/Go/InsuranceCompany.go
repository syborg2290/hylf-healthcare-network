package healthcare

type InsuranceCompany struct {
	InsuranceCompanyName    string `json:"InsuranceCompanyName"`
	InsuranceCompanyContact string `json:"InsuranceCompanyContact"`
	InsuranceCompanyAddress string `json:"InsuranceCompanyAddress"`
}

func (ic *InsuranceCompany) getInsuranceCompanyName() string {
	return ic.InsuranceCompanyName
}
func (ic *InsuranceCompany) getInsuranceCompanyAddress() string {
	return ic.InsuranceCompanyAddress
}
func (ic *InsuranceCompany) getInsuranceCompanyContact() string {
	return ic.InsuranceCompanyContact
}

func (ic *InsuranceCompany) setInsuranceCompanyName(name string) {
	ic.InsuranceCompanyName = name
}
func (ic *InsuranceCompany) setInsuranceCompanyAddress(address string) {
	ic.InsuranceCompanyAddress = address
}
func (ic *InsuranceCompany) setInsuranceCompanyContact(contact string) {
	ic.InsuranceCompanyContact = contact
}

func newIC(name string, contact string, address string) InsuranceCompany {
	var newic InsuranceCompany = InsuranceCompany{
		InsuranceCompanyName:    name,
		InsuranceCompanyContact: contact,
		InsuranceCompanyAddress: address,
	}
	return newic
}

/*func main() {
	var ic InsuranceCompany = newIC("Good Life", "754848XXXX", "Srilanka")
	fmt.Println(ic.getInsuranceCompanyName())
	fmt.Println(ic.getInsuranceCompanyAddress())
	fmt.Println(ic.getInsuranceCompanyContact())
}
*/
