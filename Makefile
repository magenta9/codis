.DEFAULT_GOAL := build-all

export GO15VENDOREXPERIMENT=1

build-all: dbserver dashboard proxy codis-admin codis-ha fe clean-gotest

codis-deps:
	@mkdir -p bin config && bash version
	@make --no-print-directory -C vendor/github.com/spinlock/jemalloc-go/

dashboard: codis-deps
	go build -i -o bin/dashboard ./cmd/dashboard
	@./bin/dashboard --default-config > config/dashboard.toml

proxy: codis-deps
	go build -i -tags "cgo_jemalloc" -o bin/proxy ./cmd/proxy
	@./bin/proxy --default-config > config/proxy.toml

codis-admin: codis-deps
	go build -i -o bin/config ./cmd/admin

codis-ha: codis-deps
	go build -i -o bin/haservice ./cmd/ha

fe: codis-deps
	go build -i -o bin/fe ./cmd/fe
	@rm -rf bin/assets; cp -rf cmd/fe/assets bin/

dbserver:
	@mkdir -p bin
	@rm -f bin/dbserver*
	make -j4 -C extern/redis-3.2.11/
	@cp -f extern/redis-3.2.11/src/redis-server  bin/dbserver
	@cp -f extern/redis-3.2.11/src/redis-benchmark bin/
	@cp -f extern/redis-3.2.11/src/redis-cli bin/
	@cp -f extern/redis-3.2.11/src/redis-sentinel bin/
	@cp -f extern/redis-3.2.11/redis.conf config/
	@sed -e "s/^sentinel/# sentinel/g" extern/redis-3.2.11/sentinel.conf > config/sentinel.conf

clean-gotest:
	@rm -rf ./pkg/topom/gotest.tmp

clean: clean-gotest
	@rm -rf bin
	@rm -rf scripts/tmp

distclean: clean
	@make --no-print-directory --quiet -C extern/redis-3.2.11 distclean
	@make --no-print-directory --quiet -C vendor/github.com/spinlock/jemalloc-go/ distclean

gotest: codis-deps
	go test ./cmd/... ./pkg/...

gobench: codis-deps
	go test -gcflags -l -bench=. -v ./pkg/...

docker:
	docker build --force-rm -t codis-image .

demo:
	pushd example && make
