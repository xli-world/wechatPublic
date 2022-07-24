// 场景描述：
// 医院中会有多个部门，如：
// 1. 前台
// 2. 医生
// 3. 药房
// 4. 收银
// 病人来访时，他们首先都会去前台，然后是看医生、取药，最后结账。也就是说，病人需要通过一条部门链，每个部门都在完成其职能后将病人进一步沿着链条输送。

package responsibilitychain

func main() {
	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
