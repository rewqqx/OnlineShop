FROM ubuntu:20.04



ENV PATH=$PATH:/opt/gradle/gradle-7.4.2/bin


RUN apt-get update && apt-get install -y apt-transport-https
RUN apt-get install -y wget
RUN apt-get install zip unzip

RUN apt install default-jdk -y

RUN wget -c https://services.gradle.org/distributions/gradle-7.4.2-bin.zip -P /tmp
RUN unzip -d /opt/gradle /tmp/gradle-7.4.2-bin.zip
COPY gradle.sh /etc/profile.d/gradle.sh
RUN chmod +x /etc/profile.d/gradle.sh
RUN /etc/profile.d/gradle.sh


RUN apt-get install -y wget
RUN wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

RUN wget https://chromedriver.storage.googleapis.com/112.0.5615.49/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip
RUN mv chromedriver /usr/bin/chromedriver
RUN chown root:root /usr/bin/chromedriver
RUN chmod +x /usr/bin/chromedriver

RUN google-chrome --product-version
RUN chromedriver --version

WORKDIR /usr/app/