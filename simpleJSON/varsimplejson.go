package main

//Banco estructura para datos del banco
type Banco struct {
	Billete map[int]string `json:"billetes"`
	Banco   string         `json:"banco"`
	Emp     Empresa        `json:"Empresa"`
	Cli     []Cliente      `json:"Cliente"`
}

//Empresa estructura
type Empresa struct {
	Cif       string
	Nombre    string
	Empleados []int
}

//Cliente estructura
type Cliente struct {
	ID       int
	Nombre   string
	Apellido string
	Edad     int
	Ciudad   string
	Viajes   map[string]int
}
