package healthcare

type Epharma struct {
	EPharmaName    string `json:"EPharmaName"`
	EPharmaContact string `json:"EPharmaContact"`
	EPharmaAddress string `json:"EPharmaAddress"`
}

func (ep *Epharma) getEPharmaName() string {
	return ep.EPharmaName
}
func (ep *Epharma) getEPharmaAddress() string {
	return ep.EPharmaAddress
}
func (ep *Epharma) getEPharmaContact() string {
	return ep.EPharmaContact
}

func (ep *Epharma) setEPharmaName(name string) {
	ep.EPharmaName = name
}
func (ep *Epharma) setEPharmaAddress(address string) {
	ep.EPharmaAddress = address
}
func (ep *Epharma) setEPharmaContact(contact string) {
	ep.EPharmaContact = contact
}

func newEPharma(name string, contact string, address string) Epharma {
	var newep Epharma = Epharma{
		EPharmaName:    name,
		EPharmaContact: contact,
		EPharmaAddress: address,
	}
	return newep
}

/*func main() {
	var ep Epharma = newEPharma("Life Medicine", "756388XXXX", "Mumbai")
	fmt.Println(ep.getEPharmaName())
	fmt.Println(ep.getEPharmaAddress())
	fmt.Println(ep.getEPharmaContact())
}
*/
