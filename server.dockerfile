FROM golang:1.14 as build-env
ARG app_env
ENV APP_ENV $app_env
WORKDIR /server
COPY ./server/src/go.mod ./
COPY ./server/src/go.sum ./
# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download
COPY ./server/src ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server/start .


FROM node:12.18.3-alpine3.12 AS node_builder
WORKDIR /app
COPY ./web-client/package*.json ./
RUN npm install
COPY ./web-client ./
ENV NODE_ENV="production"
RUN npm run build


FROM alpine
COPY --from=build-env /server/start /server/
COPY --from=node_builder /app/dist /web-client/dist

WORKDIR /server
CMD ["/server/start"]
EXPOSE 3000
