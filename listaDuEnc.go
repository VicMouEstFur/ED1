package main

import (
	"fmt"
)

type Node struct {
	valor     int
	proximo   *Node
	anterior  *Node
}

type ListaDuplamenteEncadeada struct {
	cabeca *Node
}

func (l *ListaDuplamenteEncadeada) ExibirNos() {
	atual := l.cabeca
	for atual != nil {
		fmt.Println(atual.valor)
		atual = atual.proximo
	}
}

func (l *ListaDuplamenteEncadeada) BuscarNo(valorBuscado int) *Node {
	atual := l.cabeca
	for atual != nil {
		if atual.valor == valorBuscado {
			return atual
		}
		atual = atual.proximo
	}
	return nil
}

func (l *ListaDuplamenteEncadeada) InserirNo(valor int) {
	novoNo := &Node{
		valor: valor,
	}

	if l.cabeca == nil {
		l.cabeca = novoNo
		return
	}

	atual := l.cabeca
	anterior := (*Node)(nil)

	for atual != nil && atual.valor < valor {
		anterior = atual
		atual = atual.proximo
	}

	if anterior == nil {
		novoNo.proximo = l.cabeca
		l.cabeca.anterior = novoNo
		l.cabeca = novoNo
	} else {
		anterior.proximo = novoNo
		novoNo.anterior = anterior
		novoNo.proximo = atual
		if atual != nil {
			atual.anterior = novoNo
		}
	}
}

func (l *ListaDuplamenteEncadeada) RemoverNo(valor int) {
	atual := l.BuscarNo(valor)

	if atual == nil {
		fmt.Println("Valor não encontrado")
		return
	}

	if atual.anterior != nil {
		atual.anterior.proximo = atual.proximo
	} else {
		l.cabeca = atual.proximo
	}

	if atual.proximo != nil {
		atual.proximo.anterior = atual.anterior
	}
}

func main() {
	lista := &ListaDuplamenteEncadeada{}
	lista.InserirNo(3)
	lista.InserirNo(1)
	lista.InserirNo(2)
	lista.ExibirNos() 

	no := lista.BuscarNo(2)
	if no != nil {
		fmt.Printf("Nó com valor %d encontrado!\n", no.valor)
	} else {
		fmt.Println("Nó não encontrado.")
	}

	lista.RemoverNo(2)
	lista.ExibirNos() 
}
