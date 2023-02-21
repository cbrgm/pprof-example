.PHONY: profile
profile:
	go build main.go
	sleep 5
	./main -cpuprofile=cpu.prof

.PHONY: pprof
pprof:
	go tool pprof cpu.prof

.PHONY: pprof/web
pprof/web:
	go tool pprof -http=localhost:8080 cpu.prof

.PHONY: cleanup
cleanup:
	rm main
	rm *.prof
