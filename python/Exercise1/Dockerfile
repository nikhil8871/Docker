FROM ubuntu:latest
RUN apt update -y
RUN apt install python3 -y
RUN apt install python3-pip -y
WORKDIR /app
COPY requirements.txt .
COPY app.py .
RUN pip3 install --break-system-packages -r requirements.txt
CMD ["python3" , "app.py"]