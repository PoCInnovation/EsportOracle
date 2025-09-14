# Multi-stage build for Vue 3 + Vite app
FROM node:22-alpine AS builder

WORKDIR /app

# Build-time env (VITE_*) picked up by Vite
ARG VITE_SUPABASE_URL
ARG VITE_SUPABASE_KEY
ARG VITE_API_URL=http://backend:8080

ENV VITE_SUPABASE_URL=${VITE_SUPABASE_URL}
ENV VITE_SUPABASE_KEY=${VITE_SUPABASE_KEY}
ENV VITE_API_URL=${VITE_API_URL}

COPY package.json package-lock.json* .npmrc* ./
RUN npm ci || npm install

COPY . .
# Skip TS type-check in container to avoid blocking image build
RUN npm run build-only

# Runtime: serve static files with Nginx
FROM nginx:alpine AS runtime

COPY --from=builder /app/dist /usr/share/nginx/html

# Basic SPA config (history fallback)
RUN printf 'server {\n\
	listen 80;\n\
	server_name _;\n\
	root /usr/share/nginx/html;\n\
	index index.html;\n\
\n\
	location / {\n\
		try_files $uri $uri/ /index.html;\n\
	}\n\
}\n' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
