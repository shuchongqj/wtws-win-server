FROM dark86/alpine:https

WORKDIR /usr/src/app

COPY ["./" , "./"]

RUN chmod 777 /usr/src/app
RUN chmod +x ./tos

CMD ["./tos"]
