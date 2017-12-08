FROM scratch 
ADD application /
ADD static/index.html /static/index.html
WORKDIR /
CMD ["/application"]
EXPOSE 5000
