
# 📎 PDF Joiner

**PDF Joiner** é uma ferramenta de linha de comando e GUI para unir múltiplos
arquivos PDF em um único documento, preservando totalmente o conteúdo —
incluindo texto Unicode, imagens, fontes, metadados e estrutura interna.

Criado em **Go**, multiplataforma e leve, com interface gráfica integrada
via [Fyne](https://fyne.io) e motor PDF confiável via [pdfcpu](https://github.com/pdfcpu/pdfcpu).

---

## ✨ Funcionalidades

- 🔗 Junta múltiplos PDFs mantendo a ordem informada
- 📂 Suporte a *globs* (ex: `"docs/*.pdf"`)
- ✅ Valida existência dos arquivos informados
- 📊 Barra de progresso no terminal (modo CLI)
- 🪟 Interface gráfica simples e intuitiva (modo GUI)
- 🧾 Suporte total a Unicode e fontes incorporadas
- 📦 Binários multiplataforma

---

## 🚀 Instalação

Baixe a versão mais recente na [página de releases](https://github.com/seu-usuario/pdf-joiner/releases).

### Windows

- Baixe e execute `pdf-joiner-gui.exe` → abre a interface gráfica diretamente  
- Ou use `pdf-joiner.exe` com linha de comando

### Linux / macOS

```bash
# CLI
chmod +x pdf-joiner
./pdf-joiner -n final.pdf input1.pdf "outros/*.pdf"

# GUI (se compilado com Fyne)
./pdf-joiner --gui
