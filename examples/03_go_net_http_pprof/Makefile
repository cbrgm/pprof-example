.PHONY: profile
profile:
	go run *.go

.PHONY: pprof/cpu
pprof/cpu:
	go tool pprof http://localhost:6060/debug/pprof/profile

.PHONY: pprof/cpu
pprof/cpu/web:
	go tool pprof -http=localhost:8080  http://localhost:6060/debug/pprof/profile

.PHONY: pprof/memory
pprof/memory:
	go tool pprof http://localhost:6060/debug/pprof/heap

.PHONY: pprof/memory
pprof/memory/web:
	go tool pprof -http=localhost:8080 http://localhost:6060/debug/pprof/heap

.PHONY: pprof/goroutines
pprof/goroutines:
	go tool pprof http://localhost:6060/debug/pprof/goroutine

.PHONY: pprof/goroutines
pprof/goroutines/web:
	go tool pprof -http=localhost:8080  -seconds=10 http://localhost:6060/debug/pprof/goroutine

.PHONY: cleanup
cleanup:
	rm main
	rm *.prof
