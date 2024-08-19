# Build frontend
FROM node:22-alpine3.18 as frontend
WORKDIR /app/frontend
COPY . .
RUN npm ci
RUN npm run build

# Serve frontend
EXPOSE 3000
CMD ["npm", "start"]