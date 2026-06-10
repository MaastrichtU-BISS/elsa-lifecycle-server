# Makefile for elsa-lifecycle-server

.PHONY: test

test:
	gotestsum --format=standard-verbose -- ./controllers/tests/...