#  Minimal Image
FROM scratch

# Set working directory
WORKDIR /app

# Copy the statically built and compressed binary
COPY main .

# Expose the port your application will use
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./main"]
