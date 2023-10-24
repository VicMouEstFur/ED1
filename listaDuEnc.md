Uma lista duplamente encadeada é uma estrutura de dados linear composta por uma sequência de elementos chamados nós. Cada nó contém três informações: um valor, uma referência ao próximo nó e uma referência ao nó anterior. Em uma lista duplamente encadeada ordenada, os nós são organizados com base nos seus valores, de forma crescente ou decrescente.

Func exibirNos(cabeca) {
    atual = cabeca
    Enquanto atual != nil:
        print(atual.valor)
        atual = atual.proximo
}

Func buscarNo(cabeca, valorBuscado) {
    atual = cabeca
    Enquanto atual != nil:
        Se atual.valor == valorBuscado:
            Return atual
        atual = atual.proximo
    Return nil  
}

Func inserirNo(cabeca, valor) {
    novoNo = CriaNo(valor)
    Se cabeca == nil:
        Return novoNo  

    atual = cabeca
    anterior = nil
    Enquanto atual != nil E atual.valor < valor:
        anterior = atual
        atual = atual.proximo

    Se anterior == nil:  
        novoNo.proximo = cabeca
        cabeca.anterior = novoNo
        Return novoNo
    Else:
        anterior.proximo = novoNo
        novoNo.anterior = anterior
        novoNo.proximo = atual
        Se atual != nil:
            atual.anterior = novoNo
}
Func removerNo(cabeca, valor) {
    atual = buscarNo(cabeca, valor)
    Se atual == nil:
        print("Valor não encontrado")
        Return cabeca

    Se atual.anterior != nil:
        atual.anterior.proximo = atual.proximo
    Else:
        cabeca = atual.proximo

    Se atual.proximo != nil:
        atual.proximo.anterior = atual.anterior

    Libere atual
    Return cabeca
}
