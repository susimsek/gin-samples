#  Minimal Image
FROM scratch

# Set working directory
WORKDIR /app

# Copy the statically built and compressed binary
COPY main .

# Command to run the application
ENTRYPOINT ["./main"]
a
