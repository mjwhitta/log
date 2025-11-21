-include gomk/main.mk
-include local/Makefile

ifneq ($(unameS),windows)
spellcheck:
	@codespell -f -L erro,hilighter -S "*.pem,.git,go.*,gomk"
endif
