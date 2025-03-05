all: bin/xor-brute-force

test:
	@docker build . --target test

bin/xor-brute-force:
	@docker build . \
		--output bin/ \
		--platform local
