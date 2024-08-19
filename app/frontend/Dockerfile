# Build frontend
FROM node:22-alpine3.18 as frontend
WORKDIR /app/frontend
COPY . .
RUN npm install
RUN npm run build

# Serve frontend
FROM nginx:alpine3.18
COPY --from=frontend /app/frontend/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]