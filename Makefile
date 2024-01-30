.PHONY: Short_url

include ./backend/.env

install_app:
	@echo Install frontend and backend app
	cd frontend && bun install; \
	cd ../backend && go mod tidy

start_f:
	@echo Running frontend app
	cd frontend && bun run dev

start_b:
	@echo Running backend app
	cd backend && go run .