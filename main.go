package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/schollz/progressbar/v3"
)

var version = "dev" // será sobrescrito pelo goreleaser

func main() {
	// Flags
	outputName := flag.String("name", "output.pdf", "Nome do PDF final")
	flag.StringVar(outputName, "n", "output.pdf", "Nome do PDF final (atalho)")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Erro: forneça pelo menos um arquivo PDF de entrada (pode usar globs)")
		os.Exit(1)
	}

	// Resolver globs
	files := flag.Args()
	var pdfPaths []string
	var notFound []string

	for _, pattern := range files {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatalf("Erro no glob '%s': %v", pattern, err)
		}
		if len(matches) == 0 && !containsGlob(pattern) {
			if _, err := os.Stat(pattern); os.IsNotExist(err) {
				notFound = append(notFound, pattern)
				continue
			}
		}
		pdfPaths = append(pdfPaths, matches...)
	}

	if len(notFound) > 0 {
		fmt.Println("Erro: os seguintes arquivos não foram encontrados:")
		for _, f := range notFound {
			fmt.Printf(" - %s\n", f)
		}
		os.Exit(1)
	}

	if len(pdfPaths) == 0 {
		fmt.Println("Nenhum arquivo PDF válido encontrado.")
		os.Exit(1)
	}

	// Barra de progresso
	bar := progressbar.NewOptions(len(pdfPaths),
		progressbar.OptionSetDescription("Mesclando PDFs"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerPadding: "-",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// Verifica se os arquivos existem e são PDFs
	for _, path := range pdfPaths {
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			log.Fatalf("Arquivo inválido: %s", path)
		}
		bar.Add(1)
	}

	// Merge com pdfcpu
	conf := model.NewDefaultConfiguration()
	err := api.MergeCreateFile(pdfPaths, *outputName, false, conf)
	if err != nil {
		log.Fatalf("Erro ao mesclar PDFs: %v", err)
	}

	fmt.Printf("\nPDF final criado com sucesso: %s\n", *outputName)
}

func containsGlob(s string) bool {
	return filepath.Base(s) != filepath.Clean(s)
}
