# Synopsis

- Official : https://docs.opencv.org/4.2.0/d7/d9f/tutorial_linux_install.html
- Build opencv in Debian & use GoCV:v0.21.0 to operate it.
	- example at : `/go/gocv/cmd/`
- Please make sure the OpenCV version before you build your own app to work with OpenCV.

# Notice
## v4.2 
- OpenCV v4.2 have this warning, use 4.1.2 is fine
    - https://github.com/hybridgroup/gocv/issues/582
    - https://github.com/hybridgroup/gocv/pull/577#issuecomment-568183200
- Choose GoCV:v0.21.0

# How to work with X11 : show result on the windows
## pre
> 順序重要

1. socat : 用於建立通道
    1. `brew install socat`
    2. `socat TCP-LISTEN:6000,reuseaddr,fork UNIX-CLIENT:\"$DISPLAY\"`
2. xquartz
    1. `brew install Caskroom/cask/xquartz`
    2. `open -a xquartz` 
3. Check tunnel `lsof -i TCP:6000`
    ![img](asserts/img.png)

4. RUN, choose opencv version
    - `docker-compose up -d`
    - `docker exec -it opencv /bin/bash`
    - `go run showimg/main.go /home/imgs/yujin.jpg`

![final](./asserts/final.png)

# How to generate new opencv version?
> Need to check the version release from `https://github.com/opencv/opencv/releases`
> for example v4.2->v4.3, 

1. copy the `v4.2` folder, and named it `v4.3`
2. update the image version:`opencv:v4.3` of the docker-compose.yml
3. update the download version:`ARG version=4.3` of the Dockerfile
4. Run `docker-compose up -d` and check it.

# install issue
## Can't generate Makefile
- Please make sure the `OPENCV_EXTRA_MODULES_PATH` is correct.