# Build phase
FROM node:20-alpine as build

WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# Production phase
FROM nginx:1.25-alpine
COPY --from=build /app/build /usr/share/nginx/html

# Replace default nginx config if needed
# COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
