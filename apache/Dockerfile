FROM ubuntu
WORKDIR /app
RUN apt update -y
RUN apt install apache2 -y
EXPOSE 80
CMD [ "apachectl" , "-D" , "FOREGROUND" ]