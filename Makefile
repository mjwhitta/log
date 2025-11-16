-include gomk/main.mk
-include local/Makefile

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L ERRO -S ".git,gomk,Makefile,*.pem,README.md"
endif
