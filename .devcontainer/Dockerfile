# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.205.2/containers/go/.devcontainer/base.Dockerfile
# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ENV SDK_URL2="https://dl.google.com/android/repository/commandlinetools-linux-7583922_latest.zip"

ENV SDK_URL="https://r2---sn-vgqsknlz.gvt1.com/edgedl/android/studio/ide-zips/2020.3.1.25/android-studio-2020.3.1.25-linux.tar.gz" \
    ANDROID_HOME="/usr/local/android-sdk" \
    ANDROID_SDK=$ANDROID_HOME \
    ANDROID_VERSION=29 \
    ANDROID_BUILD_TOOLS_VERSION=30.0.2


# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# RUN echo 'deb http://ftp.de.debian.org/debian/ stretch main contrib non-free' > /etc/apt/sources.list.d/stretch.list && \
#     echo "deb http://ppa.launchpad.net/openjdk-r/ppa/ubuntu jammy main" > /etc/apt/sources.list.d/openjdk.list && \
#     apt-key adv --keyserver keyserver.ubuntu.com --recv 86F44E2A && \
#     apt-get update
# RUN export DEBIAN_FRONTEND=noninteractive \
#     && apt-get install -y --no-install-recommends software-properties-common
# RUN add-apt-repository -y ppa:openjdk-r/ppa
# RUN apt-get update
# RUN apt-get install -y --no-install-recommends libc6 openjdk-8-jdk openjdk-8-jre openjdk-8-jdk-headless \
#     openjdk-8-jre-headless libjpeg62-turbo-dev
# RUN update-alternatives --config java
# RUN update-alternatives --config javac

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends openjdk-11-jdk

## Download Android SDK
RUN mkdir "$ANDROID_HOME" .android \
    && cd "$ANDROID_HOME" \
    && curl -o sdk.zip $SDK_URL2 \
    && unzip sdk.zip \
    && rm sdk.zip \
    && mv $ANDROID_HOME/cmdline-tools $ANDROID_HOME/latest \
    && mkdir $ANDROID_HOME/cmdline-tools \
    && cp -R $ANDROID_HOME/latest $ANDROID_HOME/cmdline-tools/latest
#RUN curl -o sdk.tar.gz $SDK_URL
#RUN tar -xf sdk.tar.gz 
# RUN rm sdk.tar.gz
RUN yes | $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager --licenses

## Install Android Build Tool and Libraries
RUN $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager --update
RUN $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager "build-tools;${ANDROID_BUILD_TOOLS_VERSION}" \
    "platforms;android-${ANDROID_VERSION}" \
    "platform-tools"

# Install NDK
ENV NDK_VER="19.2.5345600"
RUN $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager "ndk;$NDK_VER"
RUN ln -sf $ANDROID_HOME/ndk/$NDK_VER $ANDROID_HOME/ndk-bundle

# [Optional] Uncomment the next lines to use go get to install anything else you need
WORKDIR /workspace/go_mobile_hello_world
# USER vscode
RUN go get -x golang.org/x/mobile/cmd/gomobile

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
