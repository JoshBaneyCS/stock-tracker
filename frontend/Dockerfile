# Use Node LTS image
FROM node:18

# Set working directory
WORKDIR /app

# Copy package files
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the app
COPY frontend/ .

# Build the app for production
RUN npm run build

# Use Nginx to serve the built frontend
FROM nginx:alpine
COPY --from=0 /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
