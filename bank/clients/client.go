package clients

import "strings"

type Holder struct {
	fullName, cpf, profession string
}

// GETTERS
func (h Holder) GetFullName() string {
	return h.fullName
}

func (h Holder) GetFirstName() string {
	return strings.Split(h.fullName, " ")[0]
}

func (h Holder) GetLastName() string {
	return strings.Split(h.fullName, " ")[1]
}

func (h Holder) GetCPF() string {
	return h.cpf
}

func (h Holder) GetProfession() string {
	return h.profession
}

func (h Holder) GetHolder() string {
	return "Nome: " + h.GetFullName() + " CPF: " + h.GetCPF() + " Profissão: " + h.GetProfession()
}

// SETTERS
func (h *Holder) SetFirstName(name string) {
	splittedName := strings.Split(h.fullName, " ")
	splittedName[0] = name

	h.fullName = strings.Join(splittedName, " ")
}

func (h *Holder) SetLastName(lastName string) {
	splittedName := strings.Split(h.fullName, " ")
	splittedName[1] = lastName

	h.fullName = strings.Join(splittedName, " ")
}

func (h *Holder) SetCPF(cpf string) {
	h.cpf = cpf
}

func (h *Holder) SetProfession(profession string) {
	h.profession = profession
}

func (h *Holder) SetHolder(name, cpf, profession string) {
	h.fullName = name
	h.cpf = cpf
	h.profession = profession
}
