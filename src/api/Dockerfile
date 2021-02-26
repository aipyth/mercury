# FROM python:3.9
FROM python:3.9-slim-buster

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

##################################
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y xvfb libfontconfig wkhtmltopdf \
    && apt-get install xvfb \
    && printf '#!/bin/bash\nxvfb-run -a --server-args="-screen 0, 1024x768x24" /usr/bin/wkhtmltopdf -q $*' > /usr/bin/wkhtmltopdf.sh \
    && chmod a+x /usr/bin/wkhtmltopdf.sh \
    && ln -s /usr/bin/wkhtmltopdf.sh /usr/local/bin/wkhtmltopdf
#####################################

WORKDIR /code

COPY ./Pipfile ./Pipfile.lock /code/
RUN pip install pipenv && pipenv install --system

COPY . /code/