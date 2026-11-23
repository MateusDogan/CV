# Currículo: Mateus Oliveira Dogan

Currículos em LaTeX, de uma página cada, pensados para leitura humana e para filtros de ATS (Gupy, LinkedIn). São três versões, uma por área:

| Vaga | PDF |
|---|---|
| Análise, Ciência e Engenharia de Dados | [`CV_Mateus_Dogan_Dados.pdf`](CV_Mateus_Dogan_Dados.pdf) |
| FP&A | [`CV_Mateus_Dogan_FPA.pdf`](CV_Mateus_Dogan_FPA.pdf) |
| Product Manager / Product Owner | [`CV_Mateus_Dogan_Produto.pdf`](CV_Mateus_Dogan_Produto.pdf) |

## Objetivo

Vagas de **júnior e trainee** nas três áreas, para depois da formatura (dez. 2026) e do fim do estágio na Les. Enquanto o Mateus estiver cursando, o texto "Graduando" e "Dez. 2026 (previsão)" está correto.

### Ao se formar ou terminar o estágio, atualizar

- [ ] Resumo: "Graduando em..." vira "Bacharel em Matemática Aplicada e Computacional".
- [ ] Formação: tirar "(previsão)" da data.
- [ ] Estágio (termina em dez. 2026): se for efetivado, manter o estágio com data final e acrescentar o novo cargo na Les com a data de início e "Presente", no formato dos dois cargos da SEMAP, com as tarefas atuais no presente. Só escrever o novo cargo depois de confirmado, com o nome oficial. Sem efetivação, usar a data final e os verbos no passado.
- [ ] Monitoria e Iniciação Científica: o mesmo, se tiverem terminado.
- [ ] Conferir o título do cabeçalho de cada versão com o nome das vagas que estiver mirando (por exemplo, "Analista de Dados Júnior").

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

Ordem das seções: resumo → formação → experiência → projetos → habilidades → certificações. A experiência segue a ordem cronológica reversa nas três versões (estágio → iniciação científica → monitoria → SEMAP coordenador → SEMAP assessor).

## O que muda em cada versão

| | Dados | FP&A | Produto |
|---|---|---|---|
| Título | Análise de Dados \| Ciência de Dados \| Engenharia de Dados | FP&A \| Planejamento Financeiro \| Forecasting | Product Manager \| Product Owner \| Dados |
| Cargo no estágio | Estagiário de Análise de Dados (FP&A) | Estagiário de Planejamento Financeiro (FP&A) | Estagiário de FP&A e Produtos de Dados |
| Foco do estágio | pipelines de ETL, qualidade de dados, forecasting e KPIs | forecast e orçamento, fechamento (DRE, orçado x realizado, conciliação), relatórios automáticos | ferramenta interna, redesenho com a diretoria, definição de KPIs e ciclos de feedback |
| Iniciação científica | 2 bullets | 2 bullets | 1 bullet |
| Projetos | Previsão de Demanda e PINNs (2 bullets cada) | Previsão de Demanda (2) e PINNs (1) | só Previsão de Demanda, voltada à decisão de negócio |
| Habilidades | Dados, Machine Learning e Estatística, Ferramentas | Finanças, Ferramentas, Análise | Produto, Dados, Idiomas |

Iguais nas três: formação, monitoria, SEMAP e certificações (Google Data Analytics e Google Project Management, numa linha).

Espaço livre no fim da página: Dados e FP&A estão quase no limite; Produto tem cerca de 7 linhas, reservadas para algo de produto (projeto com métricas de uso, curso de Product Management, Jira, Figma, Amplitude, GA4).

## Compilar

Requer LaTeX (MiKTeX no Windows, TeX Live no Linux) e Go.

```
go run build.go            # as três vagas
go run build.go fpa        # só uma
```

Os PDFs vão para a raiz e são versionados: sempre compile e faça commit dos PDFs junto com o `.tex`.

## Regras

- **Uma página por currículo.** O `build.go` avisa quando passa.
- **Impressão:** A4, fonte 10pt e margens de 15 mm (definidas em `sty/loader.sty`), para não cortar na impressora.
- Se o `build.go` disser que o PDF está em uso, feche-o no leitor e rode de novo.
- **ATS:** uma coluna, texto selecionável, fonte vetorial (`lmodern`), seções com nomes padrão, datas em mês e ano, siglas por extenso na primeira vez, palavras-chave exatas das vagas.
- Mudanças no conteúdo passam pela aprovação do Mateus. O texto deve soar natural, sem frases de efeito.
- A pasta da vaga muda o que é específico; o que for igual em todas fica em `comum/` ou `sty/`.

## Pendências

- [ ] Links dos certificados do Coursera em `sty/variables.sty` (`\certAnalyticsURL`, `\certProjectURL`, hoje `URL_HERE`).
- [ ] Manter o perfil do LinkedIn igual ao currículo: o LinkedIn usa os dados do perfil, não só o PDF anexado.
