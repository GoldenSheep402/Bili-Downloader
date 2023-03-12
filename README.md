# Bili-Downloader

write by go (For fun)

I will achieve High Resolution Video later(it seems that it didn't responce all data(it needs cookie))

~~This project needs ffmpeg dependence~~

put ffmpeg in /lib

write bid and your SESSDATA in config.json, it will load it when exec

You could find your SESSDATA after your log in bilibili and find it in request header in cookie 
![Snipaste_2023-03-12_15-14-29](https://user-images.githubusercontent.com/67376942/224530080-f1d37a74-3cba-433a-8ef4-259755f1b550.png)

![Snipaste_2023-03-12_15-16-17](https://user-images.githubusercontent.com/67376942/224530152-766a2f6d-6d0a-48c8-8d00-d66037385df1.png)

```
make
 .\BiliDownloader.exe
```

- the file will be saved at ./download_path named as `{bid}{name}.mp4`

![Snipaste_2023-03-12_22-06-51](https://user-images.githubusercontent.com/67376942/224549903-155a47a7-d9d3-44cd-955d-50616fa64df9.png)


