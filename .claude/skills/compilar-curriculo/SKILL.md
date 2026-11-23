---
name: compilar-curriculo
description: Compila o currículo LaTeX (main.tex) em CV.pdf localmente com o MiKTeX, confere se cabe em uma página e mostra o resultado. Use depois de qualquer mudança em src/ ou sty/, ou quando o usuário pedir para gerar, compilar ou ver o PDF do currículo.
---

# Compilar o currículo

1. Rode da raiz do repositório (pasta `CV`), em PowerShell:

   ```powershell
   pdflatex -interaction=nonstopmode -halt-on-error -output-directory=build main.tex
   ```

   Se `pdflatex` não for encontrado, o MiKTeX fica em `$env:LOCALAPPDATA\Programs\MiKTeX\miktex\bin\x64\pdflatex.exe`. Crie a pasta `build` antes, se não existir. O MiKTeX instala sozinho os pacotes que faltarem na primeira compilação.

2. Confira o número de páginas na última linha do log `build/main.log`: `Output written on build/main.pdf (N page...)`.
   - O currículo tem que ter **1 página**. Se passar, diga quanto sobrou ou faltou e proponha o que cortar; não corte nada sem aprovação do usuário.

3. Copie `build/main.pdf` para `CV.pdf` na raiz (é o arquivo versionado).

4. Abra `CV.pdf` com a ferramenta Read para ver o layout e confira:
   - nada passou para a segunda página;
   - nenhuma palavra cortada com hífen em lugar estranho nem linha com uma palavra só;
   - o texto extraído está legível (é o que o ATS lê).

5. Para medir espaço livre, veja onde o texto termina na página renderizada e informe em linhas aproximadas.

Regras: mudanças no conteúdo do currículo só com aprovação do usuário. O GitHub Actions também compila a cada push, mas a conferência de uma página é feita aqui, antes do commit.
