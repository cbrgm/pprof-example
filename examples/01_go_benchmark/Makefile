.PHONY: profile
profile:
	go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof

.PHONY: pprof
pprof:
	go tool pprof cpu.prof

.PHONY: cleanup
cleanup:
	rm *.prof
	rm go_benchmark.test

