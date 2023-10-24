//chame o doc de main.go
package main

import (
	"fmt"
)

type Node struct {
	valor   int
	proximo *Node
}

type ListaCircular struct {
	cabeca *Node
}

func (l *ListaCircular) ExibirNos() {
	if l.cabeca == nil {
		fmt.Println("Lista vazia")
		return
	}

	noInicial := l.cabeca
	for {
		fmt.Println(noInicial.valor)
		noInicial = noInicial.proximo
		if noInicial == l.cabeca {
			break
		}
	}
}

func (l *ListaCircular) InserirNoInicio(valor int) {
	novoNo := &Node{
		valor: valor,
	}

	if l.cabeca == nil {
		novoNo.proximo = novoNo
		l.cabeca = novoNo
		return
	}

	noAtual := l.cabeca
	for noAtual.proximo != l.cabeca {
		noAtual = noAtual.proximo
	}

	noAtual.proximo = novoNo
	novoNo.proximo = l.cabeca
	l.cabeca = novoNo
}

func (l *ListaCircular) ExcluirPrimeiroNo() {
	if l.cabeca == nil {
		fmt.Println("A lista está vazia.")
		return
	}

	if l.cabeca.proximo == l.cabeca {
		l.cabeca = nil
		return
	}

	noAtual := l.cabeca
	for noAtual.proximo != l.cabeca {
		noAtual = noAtual.proximo
	}

	noAtual.proximo = l.cabeca.proximo
	l.cabeca = noAtual.proximo
}

func main() {
	lista := &ListaCircular{}
	lista.InserirNoInicio(1)
	lista.InserirNoInicio(2)
	lista.InserirNoInicio(3)

	lista.ExibirNos() 

	lista.ExcluirPrimeiroNo()
	lista.ExibirNos() 
}
