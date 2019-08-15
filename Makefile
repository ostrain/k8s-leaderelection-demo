IMAGE=owenstrain/k8s-leaderelection-demoapp:latest

demoapp: main.go
	GOOS=linux go build -o $@ .

clean:
	rm demoapp

.PHONY: image
image: Dockerfile demoapp
	docker build -t $(IMAGE) .

.PHONY: push
push: image
	docker push $(IMAGE)
