distroless\_sh\_bug
===================

Test this with `bazel run //hello:prod` and then `docker run -it
distroless_sh_bug/hello:prod` and see the error that pops out.

Compare that to `docker run -it distroless_sh_bug/hello:prod ./hello`.
