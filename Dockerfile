FROM node:16-alpine AS build

WORKDIR /app

# Copy npm files
COPY package.json package-lock.json ./

# Install dependencies using npm
RUN npm ci

# Copy all project files
COPY . .

# Build the React app
RUN npm run build

# Production stage with nginx
FROM nginx:alpine

# Copy React build to Nginx
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

