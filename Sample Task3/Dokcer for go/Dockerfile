FROM alpine:latest

# Copy only the required directories and files
COPY main ./main

RUN chmod 777 /main

EXPOSE 8000

CMD ["/main", "start"]
