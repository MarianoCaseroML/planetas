FROM scratch 
ADD application /
ADD static/index.html /static
WORKDIR /
CMD ["/application"]
EXPOSE 5000
