# JUST BOT ON GOLANG
FROM alpine

# OR WITH WKHTMLTOPDF SUPPORT
# but this extends image build time
# FROM ubuntu:18.04
# RUN apt-get update
# RUN apt-get upgrade -y
# RUN apt-get install -y xvfb libfontconfig wkhtmltopdf \
#     && apt-get install xvfb \
#     && printf '#!/bin/bash\nxvfb-run -a --server-args="-screen 0, 1024x768x24" /usr/bin/wkhtmltopdf -q $*' > /usr/bin/wkhtmltopdf.sh \
#     && chmod a+x /usr/bin/wkhtmltopdf.sh \
#     && ln -s /usr/bin/wkhtmltopdf.sh /usr/local/bin/wkhtmltopdf


# 

WORKDIR /bot
# main is compiled executable
ADD main /bot

CMD [ "/bot/main" ]
