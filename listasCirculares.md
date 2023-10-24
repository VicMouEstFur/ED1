Uma lista circular é uma variação da lista encadeada onde o último elemento aponta para o primeiro, formando um loop contínuo. Esse design permite navegação cíclica pelos elementos, sem um início ou fim distintos. Ela é útil em aplicações que requerem rotação ou processamento cíclico de elementos.

Func exibirNos(no) {
    Se no == nil :
        print(“Lista vazia”)
        Return
    noInicial = no
    Faça: 
        print(no.valor)
        no = no.próximo
    Enquanto no != noInicial
}

Func inserirNoInicio(valor, noCabeça) {
    novoNo = CriaNó()
    novoNo.valor = valor

    Se noCabeça == nil :
        novoNo.próximo = novoNo
    Senão:
        noAtual = noCabeça
        Enquanto noAtual.próximo != noCabeça:
            noAtual = noAtual.próximo
        
        noAtual.próximo = novoNo
        novoNo.próximo = noCabeça

    Return novoNo
}

Func excluirPrimeiroNo(noCabeça) {
    Se noCabeça == nil :
        print("A lista está vazia.")
        Return nil

    Se noCabeça.próximo == noCabeça:
        Libere noCabeça
        Return nil

    noAtual = noCabeça
    Enquanto noAtual.próximo != noCabeça:
        noAtual = noAtual.próximo

    noAtual.próximo = noCabeça.próximo
    Libere noCabeça
    Return noAtual.próximo
}
