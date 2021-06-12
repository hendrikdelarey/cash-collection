FROM gcr.io/distroless/base
COPY ./service /service
CMD ["/service", "serve"]