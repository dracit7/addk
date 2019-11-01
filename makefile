
img-%:
	go run main.go
	dot -Tpng -o instances/img/$*.png instances/dot/$*.dot

clean:
	rm instances/img/*.png
	rm instances/dot/*.dot