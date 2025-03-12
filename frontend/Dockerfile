FROM node:20-alpine AS build

WORKDIR /app

RUN npm install -g pnpm
COPY package.json pnpm-lock.yaml* ./
RUN pnpm install --frozen-lockfile
COPY . .

RUN pnpm build

# prune dev dependencies
RUN pnpm prune --prod

# Stage 2: Setup the runtime environment
FROM node:20-alpine AS runtime

WORKDIR /app

# Create a non-root user to run the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy only the build output and necessary files from the build stage
COPY --from=build --chown=appuser:appgroup /app/build ./build
COPY --from=build --chown=appuser:appgroup /app/node_modules ./node_modules
COPY --from=build --chown=appuser:appgroup /app/package.json ./

USER appuser

# Expose the port your app runs on
EXPOSE 3030:3030

# Set environment variables
ENV PORT=3030
ENV NODE_ENV=production

# Start the application
CMD ["node", "build"]