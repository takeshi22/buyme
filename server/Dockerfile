FROM ruby:3.0.2
RUN apt-get update -qq && apt-get install -y nodejs postgresql-client
RUN mkdir -p /buyme/server
WORKDIR /buyme/server
COPY Gemfile /buyme/server/Gemfile
COPY Gemfile.lock /buyme/server/Gemfile.lock
RUN bundle install
COPY . /buyme/server

COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 2222

CMD ["rails", "server", "-b", "0.0.0.0"]