FROM python
WORKDIR /tmp/app

# 1. Copy dependencies and install first
COPY requirement.txt .
RUN pip3 install -r requirement.txt

# 2. Copy app code last
COPY app.py .

CMD [ "python" , "app.py" ]