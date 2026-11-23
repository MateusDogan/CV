# Currículo — Mateus Oliveira Dogan

Currículo em LaTeX, de uma página, pensado para leitura humana e para filtros de ATS.

- **PDF:** [`CV.pdf`](CV.pdf), recompilado pelo GitHub Actions a cada push.
- **Compilar localmente:** `pdflatex -output-directory=build main.tex` (MiKTeX), ou a skill `compilar-curriculo` no Claude Code.
- **Estrutura:** `main.tex` junta as seções de `src/`; layout e comandos em `sty/`; dados pessoais em `sty/variables.sty`.

## Pendências

- [ ] Adicionar os links dos certificados do Coursera em `sty/variables.sty` (`\certAnalyticsURL`, `\certProjectURL`, `\certITURL`, hoje `URL_HERE`).
- [ ] Reformular Habilidades e Certificações (`src/5-skill.tex`) em uma coluna, depois de medir o espaço que sobra na página.
- [ ] No futuro: três versões do currículo (Dados, FP&A e Produto).
