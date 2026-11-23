# Currículo: Mateus Oliveira Dogan

Currículos em LaTeX, de uma página cada, pensados para leitura humana e para filtros de ATS (Gupy, LinkedIn). São três versões, uma por área:

| Vaga | PDF |
|---|---|
| Análise, Ciência e Engenharia de Dados | [`CV_Mateus_Dogan_Dados.pdf`](CV_Mateus_Dogan_Dados.pdf) |
| FP&A | [`CV_Mateus_Dogan_FPA.pdf`](CV_Mateus_Dogan_FPA.pdf) |
| Product Manager / Product Owner | [`CV_Mateus_Dogan_Produto.pdf`](CV_Mateus_Dogan_Produto.pdf) |

## Estrutura

```
main.tex           monta o currículo da vaga escolhida (\vaga)
build.go           compilador: gera os PDFs e confere se cada um tem 1 página
sty/               layout, comandos e dados pessoais (variables.sty)
comum/             igual em todas as vagas: formação e certificações
dados/ fpa/ produto/
  resumo.tex       título do cabeçalho e resumo
  experiencia.tex  bullets da experiência, com as palavras-chave da área
  projetos.tex
  habilidades.tex
```

Ordem das seções: resumo → formação → experiência → projetos → habilidades → certificações.

## Compilar

Requer LaTeX (MiKTeX no Windows, TeX Live no Linux) e Go.

```
go run build.go            # as três vagas
go run build.go fpa        # só uma
```

Os PDFs vão para a raiz e são versionados: sempre compile e faça commit dos PDFs junto com o `.tex`.

## Regras

- **Uma página por currículo.** O `build.go` avisa quando passa.
- **ATS:** uma coluna, texto selecionável, fonte vetorial (`lmodern`), seções com nomes padrão, datas em mês e ano, siglas por extenso na primeira vez, palavras-chave exatas das vagas.
- Mudanças no conteúdo passam pela aprovação do Mateus. O texto deve soar natural, sem frases de efeito.
- A pasta da vaga muda o que é específico; o que for igual em todas fica em `comum/` ou `sty/`.

## Pendências

- [ ] Links dos certificados do Coursera em `sty/variables.sty` (`\certAnalyticsURL`, `\certProjectURL`, hoje `URL_HERE`).
- [ ] Adaptar resumo, experiência, projetos e habilidades de `fpa/` e `produto/` (hoje são cópias de `dados/`, só o título muda).
