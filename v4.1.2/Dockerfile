FROM golang:1.13.1

RUN apt-get update --allow-releaseinfo-change
RUN apt-get install -y build-essential

RUN apt-get install -y cmake
RUN apt-get install -y pkg-config libgtk2.0-dev libavcodec-dev libavformat-dev libswscale-dev 
RUN apt-get install -y vim git unzip

ARG version=4.1.2

ARG opencvDL=https://github.com/opencv/opencv/archive/${version}.zip
RUN cd ${HOME} && wget ${opencvDL} -O opencv.zip && unzip opencv.zip 

ARG opencv_contrib_DL=https://github.com/opencv/opencv_contrib/archive/${version}.zip
RUN cd ${HOME} && wget ${opencv_contrib_DL} -O opencv-contrib.zip && \
    unzip opencv-contrib.zip
RUN cd ${HOME}/opencv_contrib-*/modules && \
    OPENCV_EXTRA_MODULES_PATH=$(pwd) && \
    cd -

RUN rm ${HOME}/*.zip && cd ${HOME}/opencv-* && \
    mkdir build && \
    cd build && \
    cmake -D CMAKE_BUILD_TYPE=Release \
        -D CMAKE_INSTALL_PREFIX=/usr/local \
        -D OPENCV_EXTRA_MODULES_PATH=${OPENCV_EXTRA_MODULES_PATH} \
        -D WITH_JASPER=OFF \
        -D BUILD_DOCS=OFF \
        -D BUILD_EXAMPLES=OFF \
        -D BUILD_TESTS=OFF \
        -D BUILD_PERF_TESTS=OFF \
        -D BUILD_opencv_java=NO \
        -D BUILD_opencv_python=NO \
        -D BUILD_opencv_python2=NO \
        -D BUILD_opencv_python3=NO \
        -D OPENCV_GENERATE_PKGCONFIG=ON .. && \
    make -j7 && \
    make install && \
    ldconfig

RUN printf "This OpenCV VERSION Is => " && opencv_version 