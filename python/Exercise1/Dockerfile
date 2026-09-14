FROM python:3.11-slim 
WORKDIR /tmp/app
COPY app.py .
COPY requirement.txt .
RUN pip3 install -r requirement.txt
RUN [ "python" , "app.py" ]