# Build frontend
FROM node:22-alpine3.18 as frontend
WORKDIR /app/frontend
COPY . .
RUN npm ci
RUN npm run build

# Serve frontend
FROM nginx:alpine3.18
COPY --from=frontend /app/frontend/dist /usr/share/nginx/html
COPY --from=frontend /app/frontend/nginx /etc/nginx
CMD ["nginx", "-g", "daemon off;"]