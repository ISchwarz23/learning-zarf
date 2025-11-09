VERSION_FRONTEND = 1.0.0
VERSION_BACKEND = 1.1.0

frontend-container:
	@sudo docker build -t hello-world-frontend:$(VERSION_FRONTEND) hello-world-frontend

frontend-run: frontend-container
	@sudo docker run -p 8081:80 hello-world-frontend:$(VERSION_FRONTEND)

backend-container:
	@sudo docker build -t hello-world-backend:$(VERSION_BACKEND) hello-world-backend

backend-run: backend-container
	@sudo docker run -p 8080:8080 -e GREETING="Hey there from a container!" hello-world-backend:$(VERSION_BACKEND)
	
containers: frontend-container backend-container

package: containers
	zarf package create . --confirm