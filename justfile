
# Justfile — tarefas automatizadas para build e release

# Define o nome do projeto
PROJECT_NAME := "join-pdf"

# Versão para release (pode vir de tag Git)
VERSION := "0.1.0"

# Ambiente padrão
set shell := ["bash", "-cu"]

# Garante que o diretório 'dist' está limpo
clean:
	rm -rf dist/

# Build local multiplataforma (sem publicar)
build: clean
	echo "🔧 Building snapshot for {{PROJECT_NAME}}..."
	goreleaser build --clean --snapshot --rm-dist

# Release real (gera changelog e pacotes, mas não publica)
release-local: clean
	echo "📦 Releasing local version {{VERSION}}..."
	goreleaser release --clean --skip-publish

# Release real + publicação no GitHub
release: clean
	if [ -z "$GITHUB_TOKEN" ]; then echo "❌ GITHUB_TOKEN não definido"; exit 1; fi
	echo "🚀 Publishing {{PROJECT_NAME}} {{VERSION}} to GitHub..."
	goreleaser release --clean

# Executa e imprime versão do binário local (debug)
version:
	./dist/{{PROJECT_NAME}}_linux_amd64/{{PROJECT_NAME}} --version

icon:
  rsrc -ico PDF.ico -o icon.syso
