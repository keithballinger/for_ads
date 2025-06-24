# Build the frontend
FROM node:18-alpine AS frontend
WORKDIR /app
COPY frontend /app
RUN npm install
RUN npm run build

# Build the backend
FROM golang:1.22-alpine AS backend
WORKDIR /app
COPY backend /app
RUN go build -o /app/bin/backend /app/main.go

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=frontend /app/dist ./frontend/dist
COPY --from=backend /app/bin/backend .
EXPOSE 8080
CMD ["/app/backend"]
