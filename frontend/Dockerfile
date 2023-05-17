FROM node:lts-alpine AS builder
WORKDIR /app


RUN apk add curl && \
    curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
# Files required by pnpm install
COPY .npmrc package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile --prod
COPY . .
RUN pnpm run build && pnpm prune --production
RUN ls -al

FROM node:lts-alpine
USER node:node
WORKDIR /app
# COPY --from=builder --chown=node:node /app/build ./build
COPY --from=builder --chown=node:node /app/.svelte-kit/output/ .
COPY --from=builder --chown=node:node /app/node_modules ./node_modules
COPY --chown=node:node package.json .
EXPOSE 3000
# still needs fixing
CMD ["node","./server/index.js"]
