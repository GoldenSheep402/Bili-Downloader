# Bili-Downloader

write by go (For fun)

This project needs ffmpeg dependence

It needs SESSDATA to get  High Resolution Video(without it, it could only get 480p)

write bid and your SESSDATA in config.json, it will load it when exec

- [x] Episode video download
- [ ] Download by url
- [ ] Get High Resolution without SESSDATA

You could find your SESSDATA after your log in bilibili and find it in request header in cookie 
![Snipaste_2023-03-12_15-14-29](https://user-images.githubusercontent.com/67376942/224530080-f1d37a74-3cba-433a-8ef4-259755f1b550.png)

![Snipaste_2023-03-12_15-16-17](https://user-images.githubusercontent.com/67376942/224530152-766a2f6d-6d0a-48c8-8d00-d66037385df1.png)

```
make
 .\BiliDownloader.exe
```

- the file will be saved at ./download_path named as `{bid}{name}.mp4`

![Snipaste_2023-03-14_23-11-22](https://user-images.githubusercontent.com/67376942/225046480-1b7552f8-acef-47a7-ac94-17a0cedaefb5.png)

You could also paste url here if there is nothing in config(But SESSDATA is NESSARY)
![Snipaste_2023-03-17_22-14-30](https://user-images.githubusercontent.com/67376942/225930086-94321702-c8a0-4bc7-8f70-775eb3619be6.png)
