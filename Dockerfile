FROM scratch 
ADD application /
WORKDIR /
CMD ["/application"]
EXPOSE 5000
