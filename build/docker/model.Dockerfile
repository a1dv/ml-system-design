FROM python:3.11.5-slim

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1

WORKDIR /app

COPY /api/requirements.txt ./

RUN pip install -r requirements.txt

COPY python/api /app

EXPOSE 8088
EXPOSE 8080
ENTRYPOINT ["python3", "cmd/api.py", "--configPath", "configs/config.yaml"]
