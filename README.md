# PDF Joiner

PDF Joiner é uma ferramenta de linha de comando e GUI para unir múltiplos arquivos PDF em um único documento, preservando totalmente o conteúdo — incluindo texto Unicode, imagens, fontes, metadados e estrutura interna.

Criado em Go, multiplataforma e leve, com interface gráfica integrada via Fyne e motor PDF confiável via pdfcpu.

---

## Funcionalidades

- Junta múltiplos PDFs mantendo a ordem informada
- Suporte a globs (ex: "docs/*.pdf")
- Valida existência dos arquivos informados
- Barra de progresso no terminal (modo CLI)
- Interface gráfica simples e intuitiva (modo GUI)
- Suporte total a Unicode e fontes incorporadas
- Binários multiplataforma

---

## Instalação

Baixe a versão mais recente na página de releases: https://github.com/suderio/join-pdf/releases

### Windows

- Baixe e execute `join-pdf-gui.exe` → abre a interface gráfica diretamente
- Ou use `join-pdf.exe` com linha de comando

### Linux / macOS

# CLI
chmod +x join-pdf
./join-pdf -n final.pdf input1.pdf "outros/*.pdf"

# GUI (se compilado com Fyne)
./join-pdf --gui

---

## Uso via Linha de Comando (CLI)

join-pdf -n final.pdf arquivo1.pdf arquivo2.pdf "docs/*.pdf"

### Opções:

  -n, --name     Nome do PDF final (ex: saida.pdf)
  --gui          Abre a interface gráfica
  --version      Exibe a versão do programa

---

## Uso via Interface Gráfica (GUI)

1. Execute `join-pdf-gui.exe` (Windows) ou `./join-pdf --gui` (Linux/macOS)
2. Clique em “Selecionar PDFs” e escolha os arquivos
3. Clique em “Salvar PDF Final”
4. Pronto! 🎉

---

## Requisitos para compilar (desenvolvedores)

- Go 1.18 ou superior
- Fyne (go get fyne.io/fyne/v2)
- pdfcpu
- goreleaser (para empacotar binários)

---

## Compilando localmente

go build -ldflags="-s -w -X main.version=dev" -o join-pdf

### Para GUI no Windows (sem console):

go build -ldflags="-H=windowsgui -s -w -X main.version=dev" -o join-pdf-gui.exe

---

## Licença

MIT License. Veja o arquivo LICENSE.

---

## Autor

Desenvolvido por Seu Nome (https://github.com/suderio) com 💙 e PDFs rebeldes.

