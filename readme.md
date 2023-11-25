```
cpu: AMD Ryzen 9 5950X 16-Core Processor

std:

BenchmarkGzipCompressUnsafe-32              3225            313061 ns/op          55.44 MB/s
BenchmarkGzipCompress-32                    3172            320519 ns/op          54.15 MB/s
BenchmarkGzipCompressStream-32              3518            340196 ns/op          51.02 MB/s
BenchmarkGzipCompressStdZlib-32             2523            508630 ns/op          34.13 MB/s
BenchmarkGzipDecompressUnsafe-32          246291              4979 ns/op        3485.94 MB/s
BenchmarkGzipDecompress-32                142844              7848 ns/op        2211.71 MB/s
BenchmarkGzipDecompressStream-32           47031             23777 ns/op         729.98 MB/s
BenchmarkGzipDecompressStdZlib-32          76522             17548 ns/op         989.11 MB/s
BenchmarkCompressUnsafe-32                  3745            311609 ns/op          55.70 MB/s
BenchmarkCompress-32                        3709            325418 ns/op          53.34 MB/s
BenchmarkCompressStream-32                  3374            331544 ns/op          52.35 MB/s
BenchmarkCompressStdZlib-32                 2678            488616 ns/op          35.52 MB/s
BenchmarkDecompressUnsafe-32              215644              5870 ns/op        2957.11 MB/s
BenchmarkDecompress-32                    153452              7419 ns/op        2339.60 MB/s
BenchmarkDecompressStream-32               54178             24188 ns/op         717.58 MB/s
BenchmarkDecompressStdZlib-32              64855             18439 ns/op         941.30 MB/s

cloudflare:

BenchmarkGzipCompressUnsafe-32              3289            374924 ns/op          46.29 MB/s
BenchmarkGzipCompress-32                    3212            370161 ns/op          46.89 MB/s
BenchmarkGzipCompressStream-32              3046            387642 ns/op          44.78 MB/s
BenchmarkGzipCompressStdZlib-32             2538            458199 ns/op          37.88 MB/s
BenchmarkGzipDecompressUnsafe-32          538898              3693 ns/op        4700.35 MB/s
BenchmarkGzipDecompress-32                252193              4728 ns/op        3670.96 MB/s
BenchmarkGzipDecompressStream-32           62512             19553 ns/op         887.71 MB/s
BenchmarkGzipDecompressStdZlib-32          66260             17852 ns/op         972.25 MB/s
BenchmarkCompressUnsafe-32                  3296            371052 ns/op          46.78 MB/s
BenchmarkCompress-32                        3192            363557 ns/op          47.74 MB/s
BenchmarkCompressStream-32                  3063            378185 ns/op          45.90 MB/s
BenchmarkCompressStdZlib-32                 2701            467410 ns/op          37.13 MB/s
BenchmarkDecompressUnsafe-32              633384              1791 ns/op        9688.94 MB/s
BenchmarkDecompress-32                    320124              3463 ns/op        5012.67 MB/s
BenchmarkDecompressStream-32               67143             18934 ns/op         916.72 MB/s
BenchmarkDecompressStdZlib-32              63448             18821 ns/op         922.21 MB/s

chromium:

BenchmarkGzipCompressUnsafe-32              3661            336720 ns/op          51.55 MB/s
BenchmarkGzipCompress-32                    3543            344490 ns/op          50.38 MB/s
BenchmarkGzipCompressStream-32              3054            359438 ns/op          48.29 MB/s
BenchmarkGzipCompressStdZlib-32             2260            472063 ns/op          36.77 MB/s
BenchmarkGzipDecompressUnsafe-32          245685              4902 ns/op        3541.00 MB/s
BenchmarkGzipDecompress-32                170240              7471 ns/op        2323.37 MB/s
BenchmarkGzipDecompressStream-32           51409             24155 ns/op         718.56 MB/s
BenchmarkGzipDecompressStdZlib-32          69213             17872 ns/op         971.19 MB/s
BenchmarkCompressUnsafe-32                  3553            332198 ns/op          52.25 MB/s
BenchmarkCompress-32                        3424            332883 ns/op          52.14 MB/s
BenchmarkCompressStream-32                  3391            355412 ns/op          48.84 MB/s
BenchmarkCompressStdZlib-32                 2734            458269 ns/op          37.88 MB/s
BenchmarkDecompressUnsafe-32              207081              5765 ns/op        3010.76 MB/s
BenchmarkDecompress-32                    161613              7337 ns/op        2365.83 MB/s
BenchmarkDecompressStream-32               50998             22794 ns/op         761.47 MB/s
BenchmarkDecompressStdZlib-32              62338             19650 ns/op         883.30 MB/s
```