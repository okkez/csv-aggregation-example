FROM ruby:2.7.1-buster

RUN apt update && apt install -y ca-certificates lsb-release wget curl \
        && wget https://apache.bintray.com/arrow/$(lsb_release --id --short | tr 'A-Z' 'a-z')/apache-arrow-archive-keyring-latest-$(lsb_release --codename --short).deb \
        && apt install -y ./apache-arrow-archive-keyring-latest-$(lsb_release --codename --short).deb \
        && rm -f ./apache-arrow-archive-keyring-latest-$(lsb_release --codename --short).deb
RUN apt update \
        && apt install -y libarrow-dev libarrow-glib-dev libarrow-dataset-dev libparquet-dev libparquet-glib-dev libgirepository1.0-dev \
        && apt install -y libgandiva-dev libgandiva-glib-dev

RUN groupadd app && useradd app -g app -m
USER app
WORKDIR /app
COPY Gemfile /app/
