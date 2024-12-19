FROM python:alpine3.21 AS base

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY server.py server.py

ARG FLASK_PORT=5000
EXPOSE $FLASK_PORT

CMD ["python", "server.py", "--port", "$FLASK_PORT"]