FROM golangci/build-runner
LABEL authors="gianstefanomonni"
ENV DATAPATH="/ldapsvc"

CMD mkdir /ldapsvc

COPY data /ldapsvc
COPY release/ldapsvc /bin/ldapsvc
ENTRYPOINT ["/bin/ldapsvc"]
