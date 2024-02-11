package main

import lib "github.com/Diegiwg/winapi_lib"

func main() {

	w := lib.NewWindow()
	println(w.Instance)

	lib.MessageBox("Minha Caixa de Mensagem", "Minha Super E Legal Mensagem")
}
