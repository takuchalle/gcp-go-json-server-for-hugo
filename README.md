# Usage

```
$ go mod vendor
$ gcloud alpha functions deploy ogp --entry-point OgpHandler --runtime go111 --trigger-http
```
And access `https://[region]-[project name].cloudfunctions.net/ogp?url=https://hogehoge.com` with a browser.

For detail, please check [my blog](https://blog.takuchalle.me/post/2018/11/01/trial_cloud_function_for_go_alpha/)(in Japanese).
