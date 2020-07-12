package main

import (
	"encoding/json"
	"io/ioutil"
)

const direccion = "simpleJSON.json"

func main() {
	var bank Banco
	err := json.Unmarshal(cargarJSON(direccion), &bank)
	if err != nil {
		println("Error ", err)
	}
	println("Banco :", bank.Banco)
	println("Los billetes de 20  son de color", bank.Billete[20])
}

//cargarJSON carga el fichero especificado devolviendo un array de bytes
func cargarJSON(file string) []byte {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return dat
}
