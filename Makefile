GOCMD ?= go

.PHONY: run-pub run-sub

run-pub:
	$(GOCMD) run publisher/main.go

run-sub:
	$(GOCMD) run subscriber/main.go
