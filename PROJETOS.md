# Projetos para o portfólio

Plano de projetos para mostrar no GitHub (github.com/MateusDogan) e reforçar as três versões do currículo: **Dados**, **FP&A** e **Produto**. Revisado em 27/11/2026.

Regra geral: nada do estágio na Les vai para o GitHub com dados de clientes. Quando um projeto repetir o trabalho do estágio, ele usa dados públicos ou dados sintéticos.

---

## 1. Deixar visível o que já existe

Estado do GitHub em 27/11/2026: 5 repositórios públicos, sem README de perfil, sem bio e sem repositórios fixados.

| Repositório | Situação | O que fazer |
|---|---|---|
| `MS571` | Só tem o README. É o projeto de previsão de demanda do RU, o mais citado nos três currículos | Subir o código e os notebooks, renomear para `previsao-demanda-ru`, README com problema, dados, modelos e o resultado (RMSE 12% menor) |
| `Oscilatory-Analysis` | Pronto (Python e C, FFT, método de Prony) e fora do currículo | Descrição e tópicos no repositório; avaliar entrada no CV de Dados |
| `IGA` | Site real de cliente, com backlog e decisões registradas no README | Descrição e tópicos; usar como caso de produto (requisitos, priorização, entrega) |
| `MS629` | Métodos de otimização (gradiente com Armijo, BFGS, espectral), sem README útil | README com a tabela de convergência dos métodos |
| `CV` | Fonte LaTeX dos currículos | Manter |

Ainda não publicados:

- [ ] **Iniciação científica (GMST)**: formulações de PL/PI, heurística em C e benchmarks em instâncias de 500+ nós. Se o grupo não liberar o código, uma versão didática com instâncias públicas.
- [ ] **PINNs (WeakPINN para Burgers)**: notebook com a comparação contra o método de malha e a animação da onda de choque.
- [ ] **bank-parser**: parser de extratos bancários em PDF (validado com 264 de 264 transações). Hoje privado ou fora do perfil. Publicar só com extratos de exemplo, sem dados reais.

Perfil:

- [ ] README de perfil (repositório `MateusDogan/MateusDogan`): uma linha por área, stack e links para os projetos.
- [ ] Bio e 4 a 6 repositórios fixados.
- [ ] No currículo, link direto para o repositório de cada projeto.

---

## 2. Projeto principal: drivers macroeconômicos para FP&A

Pipeline que coleta indicadores econômicos públicos, projeta esses indicadores, monta cenários e aplica tudo no orçamento de uma empresa fictícia. Cobre as três frentes: engenharia de dados (pipeline, validação, SQL), ciência de dados (séries temporais, backtest) e FP&A (premissas, cenários, DRE).

### Fontes públicas

| Fonte | Dados |
|---|---|
| Banco Central, SGS | IPCA mensal (433), IGP-M (189), Selic meta (432), CDI (12), dólar médio mensal (3698), IBC-Br (24363), taxa de desocupação (24369) |
| Banco Central, Boletim Focus (API Olinda) | Expectativas de mercado de IPCA, Selic, câmbio, PIB e IGP-M, anuais e mensais, com mediana e desvio padrão |
| IBGE, SIDRA | IPCA por grupo, PIB trimestral (opcional) |

### Etapas

1. **Extração**: baixa as séries pelas APIs e guarda os dados brutos.
2. **Validação**: interrompe o pipeline quando falta mês, quando um valor sai da faixa esperada ou quando a série está desatualizada.
3. **Armazenamento**: DuckDB ou SQLite, com as transformações em SQL.
4. **Previsão**: modelos estatísticos (SARIMA, ETS) comparados com o Focus em backtest nos últimos 24 meses.
5. **Cenários**: base (mediana do Focus), otimista e pessimista (mediana mais ou menos um desvio padrão).
6. **Aplicação em FP&A**, empresa fictícia com orçamento de 2027:
   - receita reajustada pelo IPCA e volume ligado ao PIB ou ao IBC-Br;
   - salários com dissídio pelo IPCA;
   - aluguel reajustado pelo IGP-M;
   - custo importado em dólar;
   - dívida a CDI mais spread e caixa aplicado a CDI.
7. **Sensibilidade**: efeito no EBITDA de +1 p.p. de IPCA, +10% no dólar e +1 p.p. na Selic (gráfico de tornado).
8. **Saída**: relatório HTML, exportação para Excel e atualização mensal pelo GitHub Actions, com testes.

Entra nos três currículos: Dados (pipeline e modelos), FP&A (cenários e DRE) e Produto (relatório pensado para quem decide).

---

## 3. Outros projetos que misturam as áreas

### Dados + FP&A

- **Análise de empresas da B3 com dados da CVM**: pipeline com os demonstrativos públicos (DFP e ITR), margens, alavancagem e capital de giro, comparação entre empresas do mesmo setor e forecast de receita. Reaproveita a estrutura do projeto principal.
- **Risco de crédito e inadimplência**: modelo de classificação com dados públicos e o efeito na provisão e no fluxo de caixa. Liga Machine Learning ao indicador de inadimplência do currículo.
- **Fluxo de caixa com Monte Carlo**: milhares de cenários de recebimentos e pagamentos para estimar a chance de faltar caixa em cada mês.

### Dados + Produto

- **Métricas de produto com o dataset público da Olist**: funil, coortes de retenção, LTV, recompra e um painel de métricas.
- **Teste A/B do começo ao fim**: tamanho de amostra, teste de hipótese e uma recomendação escrita, no formato de documento de decisão de produto.

### Produto + FP&A

- **Business case de um produto**: preço, unit economics, payback e cenários. Pode usar os cenários macro do projeto principal para os reajustes.
- **Estudo de caso do redesenho dos relatórios do estágio**, anonimizado: problema, usuários, versões testadas, métricas escolhidas e resultado. Ocupa o espaço livre reservado no CV de Produto.

### Otimização + negócio

Usa o que é mais raro no currículo, a experiência da iniciação científica.

- **Alocação de orçamento ou mix de produção**: programação linear (OR-Tools) para maximizar a margem com restrições de capacidade e verba.
- **Roteirização ou estoque**: heurísticas para reduzir custo logístico, com a economia mostrada em reais.

---

## 4. Projetos escolhidos

Escolhidos em 27/11/2026. É a lista que vale para começar; a seção 3 fica como banco de ideias.

### Churn

- [ ] **Churn de clientes B2B de serviços**: contratos mensais de uma empresa de serviços fictícia. Prever quem vai cancelar, explicar o motivo de cada previsão (SHAP) e medir o efeito na receita recorrente (MRR) e no forecast. Liga churn a FP&A.

### Produto pequeno e real

Ferramentas que outras pessoas usem de fato, com documento de produto (PRD), métricas de uso e histórico de versões.

- [ ] **Calculadora de reajuste de aluguel e contratos**: corrige valores pelo IGP-M, pelo IPCA, pelo INCC ou pelo CDI e mostra o histórico, com os índices puxados das APIs do Banco Central.
- [ ] **Organizador de finanças compartilhadas para repúblicas**: divide as contas, mostra quem deve para quem e guarda o histórico mensal.
- [ ] **Calculadora de orçamento de projeto para o escritório IGA**: área, tipo de projeto e etapas, usada no atendimento. O cliente já existe, então há feedback real.

### Fechamento mensal automatizado

Versão pública do trabalho do estágio, com dados sintéticos.

- [ ] **Fechamento de um e-commerce**: vendas por canal, taxas de marketplace e de cartão, antecipação de recebíveis e devoluções, com conciliação entre pedido, pagamento e extrato.
- [ ] **Fechamento de uma loja de roupas física**: vendas no caixa e na maquininha, taxas de cartão, conciliação entre maquininha e extrato, CMV, inventário e coleções sazonais.
- [ ] **Conciliação bancária a partir de PDFs de extrato**: usa o bank-parser como entrada, cruza com os lançamentos e aponta as diferenças.

### Pipeline de dados abertos

- [ ] **Clima e agro**: estações meteorológicas do INMET com a produção agrícola do IBGE (PAM), cruzando chuva e temperatura com a safra de cada município. Coleta recorrente, modelagem em camadas, validação e agendamento.

### Da seção 3

- [ ] **Simulação de teste A/B**: tamanho de amostra, teste de hipótese e recomendação escrita no formato de documento de decisão de produto.
- [ ] **Otimização de orçamento**: programação linear (OR-Tools) para alocar verba e maximizar a margem dentro das restrições.
- [ ] **Roteirização ou estoque**: heurísticas para reduzir o custo logístico ou de estoque, com a economia em reais. O estoque pode usar os dados da loja de roupas.
- [ ] **Business case de um produto**: preço, unit economics, payback e cenários, com os reajustes vindos dos cenários macro do projeto principal.

### Combinações

- **Conciliação por PDF + bank-parser**: um projeto maior em vez de dois menores.
- **Churn B2B + fechamento**: o churn alimenta o forecast de receita do fechamento.
- **Loja de roupas + estoque**: o fechamento gera os dados de vendas e a otimização decide a reposição.
- **Calculadora de reajuste + projeto principal**: as duas usam as mesmas séries do Banco Central e podem dividir o código de extração.
- **Planejador de grade para alunos da UNICAMP**: produto com usuários reais que usa a otimização da iniciação científica.

---

## 5. Ordem sugerida

1. [ ] README de perfil, bio e `MS571` completo.
2. [ ] GMST e PINNs publicados.
3. [ ] Drivers macroeconômicos para FP&A.
4. [ ] Análise de empresas da B3 com dados da CVM.
5. [ ] Métricas de produto com a Olist.
6. [ ] Alocação de orçamento com programação linear.

Com esses itens, cada versão do currículo fica com pelo menos dois projetos próprios no GitHub.
