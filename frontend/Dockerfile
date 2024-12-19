# Use an official Node.js runtime as the base image
FROM node:18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json into the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Install DOM
RUN npm install react-router-dom @types/react-router-dom

# Secure Babel
RUN npm install --save-dev @babel/plugin-proposal-private-property-in-object

# Copy the rest of the application code
COPY . .

# Build the React app for production
RUN npm run build

# Use a lightweight web server to serve the React app
# In this case, we'll use nginx
FROM nginx:1.21-alpine

# Copy the build output from the previous stage to the nginx html directory
COPY --from=0 /app/build /usr/share/nginx/html

# Expose the default nginx port
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
