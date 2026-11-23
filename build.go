// Compila os currículos com pdflatex e confere se cada um tem uma página.
//
//	go run build.go            # todas as vagas
//	go run build.go dados fpa  # só as escolhidas
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var vagas = map[string]string{
	"dados":   "CV_Mateus_Dogan_Dados.pdf",
	"fpa":     "CV_Mateus_Dogan_FPA.pdf",
	"produto": "CV_Mateus_Dogan_Produto.pdf",
}

var ordem = []string{"dados", "fpa", "produto"}

var paginas = regexp.MustCompile(`Output written on .*?\((\d+) page`)

func main() {
	escolhidas := os.Args[1:]
	if len(escolhidas) == 0 {
		escolhidas = ordem
	}

	pdflatex, err := acharPdflatex()
	if err != nil {
		falhar(err.Error())
	}

	tmp, err := os.MkdirTemp("", "cv-build-")
	if err != nil {
		falhar(err.Error())
	}
	defer os.RemoveAll(tmp)

	erros := 0
	for _, vaga := range escolhidas {
		saida, ok := vagas[vaga]
		if !ok {
			fmt.Printf("✗ %s: vaga desconhecida (use %s)\n", vaga, strings.Join(ordem, ", "))
			erros++
			continue
		}
		if err := compilar(pdflatex, tmp, vaga, saida); err != nil {
			fmt.Printf("✗ %s: %v\n", vaga, err)
			erros++
		}
	}
	if erros > 0 {
		os.Exit(1)
	}
}

func compilar(pdflatex, tmp, vaga, saida string) error {
	job := "cv_" + vaga
	cmd := exec.Command(pdflatex,
		"-interaction=nonstopmode", "-halt-on-error",
		"-output-directory="+tmp, "-jobname="+job,
		`\def\vaga{`+vaga+`}\input{main.tex}`)
	cmd.Run() // o resultado é lido do log, que diz mais que o código de saída

	log, err := os.ReadFile(filepath.Join(tmp, job+".log"))
	if err != nil {
		return fmt.Errorf("pdflatex não gerou log: %w", err)
	}
	if erro := primeiroErro(string(log)); erro != "" {
		return fmt.Errorf("erro do LaTeX:\n%s", erro)
	}
	// o LaTeX quebra o log em 79 colunas, inclusive no meio de "(1 page"
	m := paginas.FindStringSubmatch(strings.NewReplacer("\r", "", "\n", "").Replace(string(log)))
	if m == nil {
		return fmt.Errorf("nenhum PDF gerado")
	}

	pdf, err := os.ReadFile(filepath.Join(tmp, job+".pdf"))
	if err != nil {
		return err
	}
	if err := os.WriteFile(saida, pdf, 0o644); err != nil {
		return err
	}

	if m[1] != "1" {
		fmt.Printf("⚠ %s → %s: %s páginas (precisa caber em 1)\n", vaga, saida, m[1])
	} else {
		fmt.Printf("✓ %s → %s: 1 página\n", vaga, saida)
	}
	return nil
}

func primeiroErro(log string) string {
	linhas := strings.Split(log, "\n")
	for i, l := range linhas {
		if strings.HasPrefix(l, "!") {
			fim := min(i+4, len(linhas))
			return strings.Join(linhas[i:fim], "\n")
		}
	}
	return ""
}

func acharPdflatex() (string, error) {
	if p, err := exec.LookPath("pdflatex"); err == nil {
		return p, nil
	}
	if runtime.GOOS == "windows" {
		p := filepath.Join(os.Getenv("LOCALAPPDATA"), "Programs", "MiKTeX", "miktex", "bin", "x64", "pdflatex.exe")
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}
	return "", fmt.Errorf("pdflatex não encontrado: instale o MiKTeX (Windows) ou o TeX Live (Linux)")
}

func falhar(msg string) {
	fmt.Println("✗", msg)
	os.Exit(1)
}
