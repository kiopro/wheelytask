FROM ruby:2.5.1

WORKDIR /app
ADD . /app

ENV RACK_ENV production

RUN bundle install --without development test

EXPOSE 4567

#CMD ["ruby", "app.rb"]
CMD ["thin", "--threaded", "start"]

# docker build -t kiopro/price-service:1.0.0 .
# docker run --rm -it -p 4444:4567 kiopro/price-service:latest
# docker rmi kiopro/price-service:1.0.0
