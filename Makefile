help: ## Show this help
	@ echo 'Usage: make <target>'
	@ echo
	@ echo 'Available targets:'
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

YEAR ?= $(shell date +%Y)
DAY ?= $(shell expr `ls -v $(YEAR)/go/ | grep '^day' | tail -n 1 | sed 's/day//g'` + 1)

skeleton:  ## make skeleton main.go and data.txt files, optional: $DAY and $YEAR
	@ mkdir -p $(YEAR)/go/day$(DAY)
	@ cp $(YEAR)/go/template/template.go $(YEAR)/go/day$(DAY)/main.go
	@ touch $(YEAR)/go/day$(DAY)/data.txt
	@ echo "Created directory and files for year $(YEAR) and day $(DAY)"

run: ## run the go code for the current day
	@ if [ -z "$(PART)" ]; then echo "PART is not set"; exit 1; fi
	@ go run $(YEAR)/go/day$(shell expr $(DAY) - 1)/main.go -part $(PART)
