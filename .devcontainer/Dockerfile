# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.205.2/containers/go/.devcontainer/base.Dockerfile
# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ARG ANDROID_VERSION=29
ARG ANDROID_BUILD_TOOLS_VERSION=30.0.2

ARG SDK_URL="https://dl.google.com/android/repository/commandlinetools-linux-7583922_latest.zip"
ENV ANDROID_HOME="/usr/local/android-sdk" \
    ANDROID_SDK=$ANDROID_HOME


# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends openjdk-11-jdk libegl1-mesa-dev libgles2-mesa-dev libx11-dev

## Download Android SDK
RUN mkdir "$ANDROID_HOME" .android \
    && cd "$ANDROID_HOME" \
    && curl -o sdk.zip $SDK_URL \
    && unzip sdk.zip \
    && rm sdk.zip \
    && mv $ANDROID_HOME/cmdline-tools $ANDROID_HOME/latest \
    && mkdir $ANDROID_HOME/cmdline-tools \
    && cp -R $ANDROID_HOME/latest $ANDROID_HOME/cmdline-tools/latest \
    && yes | $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager --licenses

## Install Android Build Tool and Libraries
RUN $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager --update
RUN $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager "build-tools;${ANDROID_BUILD_TOOLS_VERSION}" \
    "platforms;android-${ANDROID_VERSION}" \
    "platform-tools"

# Install NDK
ARG NDK_VER="19.2.5345600"
RUN $ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager "ndk;$NDK_VER"
RUN ln -sf $ANDROID_HOME/ndk/$NDK_VER $ANDROID_HOME/ndk-bundle

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x golang.org/x/mobile/cmd/gomobile

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
