# gatsby-blog-vercel
Reponsitory for deploying to Vercel

## golang

利用 embed 嵌入数据文件到 serverless 函数中，被嵌入文件为 public/site.json

在 golang 中利用 go generate （因为 go embed 只支持同级目录或子级目录）和 go embed 在编译器保持和 public/site.json 同步。

保证每次 public/site.json 更新时，go gennerate 被执行一次即可（交给ci）。