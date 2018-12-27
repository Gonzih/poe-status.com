figwheel: clean
	clojure -m figwheel.main -b dev -r

release: clean
	clojure -m figwheel.main -O advanced -bo dev

generate: release
	go generate

clean:
	rm assets/cljs/* -rf
