package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/schollz/progressbar/v3"
)

var version = "0.1.0"

func main() {
	guiMode := flag.Bool("gui", false, "Abrir interface gráfica")
	outputName := flag.String("name", "output.pdf", "Nome do PDF final")
	flag.StringVar(outputName, "n", "output.pdf", "Nome do PDF final (atalho)")
	showVersion := flag.Bool("version", false, "Exibir versão")

	flag.Parse()

	if *showVersion {
		fmt.Println("join-pdf", version)
		return
	}

	if *guiMode || len(os.Args) == 1 {
		runGUI()
		return
	}

	if flag.NArg() == 0 {
		fmt.Println("Erro: forneça arquivos PDF para mesclar (ou use --gui)")
		os.Exit(1)
	}

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
		fmt.Println("Erro: arquivos não encontrados:")
		for _, f := range notFound {
			fmt.Printf(" - %s\n", f)
		}
		os.Exit(1)
	}

	if len(pdfPaths) == 0 {
		fmt.Println("Nenhum arquivo PDF válido encontrado.")
		os.Exit(1)
	}

	bar := progressbar.NewOptions(len(pdfPaths),
		progressbar.OptionSetDescription("Mesclando PDFs"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "=", BarStart: "[", BarEnd: "]"}),
	)

	for _, path := range pdfPaths {
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			log.Fatalf("Arquivo inválido: %s", path)
		}
		bar.Add(1)
	}

	conf := model.NewDefaultConfiguration()
	err := api.MergeCreateFile(pdfPaths, *outputName, false, conf)
	if err != nil {
		log.Fatalf("Erro ao mesclar PDFs: %v", err)
	}

	fmt.Printf("\nPDF final criado com sucesso: %s\n", *outputName)
}

func containsGlob(s string) bool {
	return strings.ContainsAny(s, "*?[")
}

// ---------------- GUI MODE ----------------

func runGUI() {
	a := app.New()
	w := a.NewWindow("PDF Joiner")

	var pdfFiles []string
	list := widget.NewMultiLineEntry()
	list.SetPlaceHolder("Nenhum arquivo selecionado.")

	btnSelect := widget.NewButton("Selecionar PDFs", func() {
		fd := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil || r == nil {
				return
			}
			path := r.URI().Path()
			pdfFiles = append(pdfFiles, path)
			list.SetText(strings.Join(pdfFiles, "\n"))
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fd.Resize(fyne.NewSize(800, 600))
		fd.Show()
	})

	btnMerge := widget.NewButton("Salvar PDF Final", func() {
		if len(pdfFiles) == 0 {
			dialog.ShowInformation("Aviso", "Nenhum PDF selecionado.", w)
			return
		}

		dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil || writer == nil {
				return
			}
			output := writer.URI().Path()
			conf := model.NewDefaultConfiguration()
			err = api.MergeCreateFile(pdfFiles, output, false, conf)
			if err != nil {
				dialog.ShowError(err, w)
			} else {
				dialog.ShowInformation("Sucesso", "PDF criado com sucesso!", w)
			}
		}, w).Show()
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Selecionador de PDFs"),
		btnSelect,
		list,
		btnMerge,
	))
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
