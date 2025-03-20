ARG NODE_VERSION=22.14

# Build stage 
FROM node:${NODE_VERSION}-slim AS base

# Setup PNPM 
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable 

# Copy package.json and pnpm-lock.yaml to working dir 
WORKDIR /app 
COPY ./package.json /app/
COPY ./pnpm-lock.yaml /app/

# Install dependencies 
RUN pnpm install --shamefully-hoist 

# Copy the rest of the application to the working dir 
COPY . ./ 

# Build the application 
RUN pnpm run build 

# New stage for production image 
FROM node:${NODE_VERSION}-slim

# Working dir inside the container 
WORKDIR /app 

# Copy the output from the build stage to the working dir
COPY --from=base /app/.output ./ 

# Define envrionment variables and port 
ENV HOST=0.0.0.0
ENV NODE_ENV=production
EXPOSE 3000
# Start the application
CMD ["node", "/app/server/index.mjs"]
