FROM eraac/golang

ADD bot /bot

CMD ["/bot", "-config", "/config.json"]
