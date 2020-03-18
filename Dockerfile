FROM scratch

COPY dist/ .

CMD ["/service"]