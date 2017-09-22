# Use 'make image' to build
FROM scratch
COPY releases/pimba-static /pimba

ENTRYPOINT ["/pimba"]
EXPOSE 8080
CMD ["api", "--storage", "/var/www", "--secret", "$JWT_SECRET"]
