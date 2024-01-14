workspace(name = "temporalite")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

gazelle_version = "v0.35.0"

gazelle_hash = "32938bda16e6700063035479063d9d24c60eda8d79fd4739563f50d331cb3209"

rules_go_version = "v0.45.0"

rules_go_hash = "de7974538c31f76658e0d333086c69efdf6679dbc6a466ac29e65434bf47076d"

golang_version = "1.21.6"

bazel_skylib_version = "1.5.0"

bazel_skylib_hash = "cd55a062e763b9349921f0f5db8c3933288dc8ba4f76dd9416aac68acee3cb94"

protobuf_all_version = "21.7"

protobuf_all_hash = "e07046fbac432b05adc1fd1318c6f19ab1b0ec0655f7f4e74627d9713959a135"

rules_protobuf_version = "5.3.0-21.7"

rules_protobuf_hash = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd"

rules_grpc_version = "4.6.0"

rules_grpc_hash = "c0d718f4d892c524025504e67a5bfe83360b3a982e654bc71fed7514eb8ac8ad"

########################################
# Definition for Skylib
########################################

http_archive(
    name = "bazel_skylib",
    sha256 = bazel_skylib_hash,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/%s/bazel-skylib-%s.tar.gz" % (bazel_skylib_version, bazel_skylib_version),
        "https://github.com/bazelbuild/bazel-skylib/releases/download/%s/bazel-skylib-%s.tar.gz" % (bazel_skylib_version, bazel_skylib_version),
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

########################################
# Definition for Golang
########################################
http_archive(
    name = "io_bazel_rules_go",
    sha256 = rules_go_hash,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/%s/rules_go-%s.zip" % (rules_go_version, rules_go_version),
        "https://github.com/bazelbuild/rules_go/releases/download/%s/rules_go-%s.zip" % (rules_go_version, rules_go_version),
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = gazelle_hash,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/%s/bazel-gazelle-%s.tar.gz" % (gazelle_version, gazelle_version),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/%s/bazel-gazelle-%s.tar.gz" % (gazelle_version, gazelle_version),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

########################################
# Definition for Protobuf
########################################
http_archive(
    name = "com_google_protobuf",
    sha256 = protobuf_all_hash,
    strip_prefix = "protobuf-%s" % protobuf_all_version,
    urls = [
        "https://github.com/protocolbuffers/protobuf/releases/download/v%s/protobuf-all-%s.tar.gz" % (protobuf_all_version, protobuf_all_version),
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "rules_proto",
    sha256 = rules_protobuf_hash,
    strip_prefix = "rules_proto-%s" % rules_protobuf_version,
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/%s.tar.gz" % rules_protobuf_version,
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

########################################
# Definition for GRPC
########################################

http_archive(
    name = "rules_proto_grpc",
    sha256 = rules_grpc_hash,
    strip_prefix = "rules_proto_grpc-%s" % rules_grpc_version,
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/%s.tar.gz" % rules_grpc_version],
)

load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

http_archive(
    name = "googleapis",
    sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",
    strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
    urls = [
        "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
    ],
)

load("@googleapis//:repository_rules.bzl", "switched_rules_by_language")

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    go = True,
    grpc = True,
    java = True,
    python = True,
)

go_repository(
    name = "com_github_google_s2a_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/s2a-go",
    sum = "h1:1kZ/sQM3srePvKs3tXAvQzo66XfcReoqFpIpIccE7Oc=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_alecthomas_kingpin_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/alecthomas/kingpin/v2",
    sum = "h1:H0aULhgmSzN8xQ3nX1uxtdlTHYoPLu5AhHxWrKI6ocU=",
    version = "v2.3.2",
)

go_repository(
    name = "com_github_cactus_go_statsd_client_v5",
    build_file_proto_mode = "disable",
    importpath = "github.com/cactus/go-statsd-client/v5",
    sum = "h1:KqvIQtc9qt34uq+nu4nd1PwingWfBt/IISgtUQ2nSJk=",
    version = "v5.0.0",
)

go_repository(
    name = "com_github_google_go_pkcs11",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/go-pkcs11",
    sum = "h1:5meDPB26aJ98f+K9G21f0AqZwo/S5BJMJh8nuhMbdsI=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_xhit_go_str2duration_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/xhit/go-str2duration/v2",
    sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
    version = "v2.1.0",
)

go_repository(
    name = "com_google_cloud_go_dataproc_v2",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dataproc/v2",
    sum = "h1:BPjIIkTCAOHUkMtWKqae55qEku5K09LVbQ46LYt7r1s=",
    version = "v2.2.1",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_api",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/genproto/googleapis/api",
    sum = "h1:CIC2YMXmIhYw6evmhPxBKJ4fmLbOFtXQN/GV3XOZR8k=",
    version = "v0.0.0-20231016165738-49dd2c1f3d0b",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_bytestream",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/genproto/googleapis/bytestream",
    sum = "h1:g3hIDl0jRNd9PPTs2uBzYuaD5mQuwOkZY0vSc0LR32o=",
    version = "v0.0.0-20230530153820-e85fd2cbaebc",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_rpc",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/genproto/googleapis/rpc",
    sum = "h1:ZlWIi1wSK56/8hn4QcBp/j9M7Gt3U/3hZw3mC7vDICo=",
    version = "v0.0.0-20231016165738-49dd2c1f3d0b",
)

go_repository(
    name = "org_uber_go_automaxprocs",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/automaxprocs",
    sum = "h1:2LxUOGiR3O6tw8ui5sZa2LAaHnsviZdVOUZw4fvbnME=",
    version = "v1.5.2",
)

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

go_repository(
    name = "co_honnef_go_tools",
    build_file_proto_mode = "disable",
    importpath = "honnef.co/go/tools",
    sum = "h1:3JgtbtFHMiCmsznwGVTUWbgGov+pVqnlf1dEJTNAXeM=",
    version = "v0.0.1-2019.2.3",
)

go_repository(
    name = "com_github_ajstarks_svgo",
    build_file_proto_mode = "disable",
    importpath = "github.com/ajstarks/svgo",
    sum = "h1:wVe6/Ea46ZMeNkQjjBW6xcqyQA/j5e0D6GytH95g0gQ=",
    version = "v0.0.0-20180226025133-644b8db467af",
)

go_repository(
    name = "com_github_alecthomas_template",
    build_file_proto_mode = "disable",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    build_file_proto_mode = "disable",
    importpath = "github.com/alecthomas/units",
    sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
    version = "v0.0.0-20211218093645-b94a6e3cc137",
)

go_repository(
    name = "com_github_antihax_optional",
    build_file_proto_mode = "disable",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_apache_thrift",
    build_file_proto_mode = "disable",
    importpath = "github.com/apache/thrift",
    sum = "h1:lNhK/1nqjbwbiOPDBPFJVKxgDEGSepKuTh6OLiXW8kg=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:5CVEjiHFvdiVlKPBzv0rjG4zH/21W/onT18R5AH/qx0=",
    version = "v1.44.289",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    build_file_proto_mode = "disable",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:VvXlSJBzZpA/zum6Sj74hxwYI2DIxRWuNIoXAzHZz5o=",
    version = "v1.3.5",
)

go_repository(
    name = "com_github_beorn7_perks",
    build_file_proto_mode = "disable",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bitly_go_hostpool",
    build_file_proto_mode = "disable",
    importpath = "github.com/bitly/go-hostpool",
    sum = "h1:XKmsF6k5el6xHG3WPJ8U0Ku/ye7njX7W81Ng7O2ioR0=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_blang_semver_v4",
    build_file_proto_mode = "disable",
    importpath = "github.com/blang/semver/v4",
    sum = "h1:1PFHFE6yCCTv8C1TeyNNarDzntLi7wMI5i/pzqYIsAM=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_bmizerany_assert",
    build_file_proto_mode = "disable",
    importpath = "github.com/bmizerany/assert",
    sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
    version = "v0.0.0-20160611221934-b7ed37b82869",
)

go_repository(
    name = "com_github_bmizerany_perks",
    build_file_proto_mode = "disable",
    importpath = "github.com/bmizerany/perks",
    sum = "h1:AP/Y7sqYicnjGDfD5VcY4CIfh1hRXBUavxrvELjTiOE=",
    version = "v0.0.0-20141205001514-d9a9656a3a4b",
)

go_repository(
    name = "com_github_brianvoe_gofakeit_v6",
    build_file_proto_mode = "disable",
    importpath = "github.com/brianvoe/gofakeit/v6",
    sum = "h1:BzOsDot1o3cufTfOk+fWKE9nFYojyDV+XHdCWL2+uyE=",
    version = "v6.22.0",
)

go_repository(
    name = "com_github_burntsushi_toml",
    build_file_proto_mode = "disable",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:o7IhLm0Msx3BaB+n3Ag7L8EVlByGnpq14C4YWiu/gL8=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    build_file_proto_mode = "disable",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_cactus_go_statsd_client_statsd",
    build_file_proto_mode = "disable",
    importpath = "github.com/cactus/go-statsd-client/statsd",
    sum = "h1:HIGF0r/56+7fuIZw2V4isE22MK6xpxWx7BbV8dJ290w=",
    version = "v0.0.0-20200423205355-cb0885a1018c",
)

go_repository(
    name = "com_github_cactus_go_statsd_client_v4",
    build_file_proto_mode = "disable",
    importpath = "github.com/cactus/go-statsd-client/v4",
    sum = "h1:cjO9CI3GAHtj/Vmmt1/BRq8+hf5MXy/Pr/CM6Dy4dp0=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    build_file_proto_mode = "disable",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
    version = "v4.2.1",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    build_file_proto_mode = "disable",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_client9_misspell",
    build_file_proto_mode = "disable",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
    version = "v0.0.0-20220112060539-c52dc94e7fbe",
)

go_repository(
    name = "com_github_cncf_xds_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:/inchEIKaYC1Akx+H+gqO04wryn5h75LSazbRlnya1k=",
    version = "v0.0.0-20230607035331-e9ce68804cb4",
)

go_repository(
    name = "com_github_coreos_go_oidc_v3",
    build_file_proto_mode = "disable",
    importpath = "github.com/coreos/go-oidc/v3",
    sum = "h1:6avEvcdvTa1qYsOZ6I5PRkSYHzpTNWgKYmaJfaYbrRw=",
    version = "v3.1.0",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_creack_pty",
    build_file_proto_mode = "disable",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_crossdock_crossdock_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/crossdock/crossdock-go",
    sum = "h1:WR1qVJzbvrVywhAk4kMQKRPx09AZVI0NdEdYs59iHcA=",
    version = "v0.0.0-20160816171116-049aabb0122b",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    build_file_proto_mode = "disable",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_dgryski_go_farm",
    build_file_proto_mode = "disable",
    importpath = "github.com/dgryski/go-farm",
    sum = "h1:fAjc9m62+UWV/WAFKLNi6ZS0675eEUC9y3AlwSbQu1Y=",
    version = "v0.0.0-20200201041132-a6ae2369ad13",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    build_file_proto_mode = "disable",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_emirpasic_gods",
    build_file_proto_mode = "disable",
    importpath = "github.com/emirpasic/gods",
    sum = "h1:FXtiHYKDGKCW2KzwZKx0iC0PQmdlorYgdFG9jPXJ1Bc=",
    version = "v1.18.1",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    build_file_proto_mode = "disable",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:wSUXTlLfiAQRWs2F+p+EKOY9rUyis1MyGqJ2DIk5HpM=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    build_file_proto_mode = "disable",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:QkIBuU5k+x7/QXPvPPnWXWlCdaBFApVqftFV6k087DA=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_facebookgo_clock",
    build_file_proto_mode = "disable",
    importpath = "github.com/facebookgo/clock",
    sum = "h1:yDWHCSQ40h88yih2JAcL6Ls/kVkSE8GFACTGVnMPruw=",
    version = "v0.0.0-20150410010913-600d898af40a",
)

go_repository(
    name = "com_github_fatih_color",
    build_file_proto_mode = "disable",
    importpath = "github.com/fatih/color",
    sum = "h1:kOqh6YHBtK8aywxGerMG2Eq3H6Qgoqeo13Bk2Mv/nBs=",
    version = "v1.15.0",
)

go_repository(
    name = "com_github_fogleman_gg",
    build_file_proto_mode = "disable",
    importpath = "github.com/fogleman/gg",
    sum = "h1:WXb3TSNmHp2vHoCroCIB1foO/yQ36swABL8aOVeDpgg=",
    version = "v1.2.1-0.20190220221249-0403632d5b90",
)

go_repository(
    name = "com_github_fortytw2_leaktest",
    build_file_proto_mode = "disable",
    importpath = "github.com/fortytw2/leaktest",
    sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_ghodss_yaml",
    build_file_proto_mode = "disable",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_gl_glfw",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_kit_kit",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-kit/kit",
    sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_go_kit_log",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-kit/log",
    sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_go_logr_logr",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-logr/logr",
    sum = "h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_go_logr_stdr",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-logr/stdr",
    sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_go_stack_stack",
    build_file_proto_mode = "disable",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gocql_gocql",
    build_file_proto_mode = "disable",
    importpath = "github.com/gocql/gocql",
    sum = "h1:WnKf8xRQImcT/KLaEWG2pjEeryDB7K0qQN9mPs1C58Q=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_gogo_gateway",
    build_file_proto_mode = "disable",
    importpath = "github.com/gogo/gateway",
    sum = "h1:u0SuhL9+Il+UbjM9VIE3ntfRujKbvVpFvNB4HbjeVQ0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_gogo_googleapis",
    build_file_proto_mode = "disable",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:1Yx4Myt7BxzvUr5ldGSbwYiZG6t9wGBZ+8/fX3Wvtq0=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_gogo_protobuf",
    build_file_proto_mode = "disable",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_gogo_status",
    build_file_proto_mode = "disable",
    importpath = "github.com/gogo/status",
    sum = "h1:DuHXlSFHNKqTQ+/ACf5Vs6r4X/dH2EgIzR9Vr+H65kg=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_golang_freetype",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_golang_glog",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/glog",
    sum = "h1:DVjP2PbBOzHyzA+dn3WhHIq4NdVu3Q+pvivFICf/7fo=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_golang_groupcache",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/groupcache",
    sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
)

go_repository(
    name = "com_github_golang_jwt_jwt",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang-jwt/jwt",
    sum = "h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=",
    version = "v3.2.2+incompatible",
)

go_repository(
    name = "com_github_golang_jwt_jwt_v4",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang-jwt/jwt/v4",
    sum = "h1:7cYmW1XlMY7h7ii7UhUyChSgS5wUJEnm9uZVTGqOWzg=",
    version = "v4.5.0",
)

go_repository(
    name = "com_github_golang_mock",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/mock",
    sum = "h1:YojYx61/OLFsiv6Rw1Z96LpldJIy31o+UHmwAUMJ6/U=",
    version = "v1.7.0-rc.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/protobuf",
    sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
    version = "v1.5.3",
)

go_repository(
    name = "com_github_golang_snappy",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/snappy",
    sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_google_go_cmp",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/go-cmp",
    sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
    version = "v0.5.9",
)

go_repository(
    name = "com_github_google_gofuzz",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/gofuzz",
    sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian_v3",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/martian/v3",
    sum = "h1:IqNFLAmvJOgVlpdEBiQbDc2EwKW77amAycfTuWKdfvw=",
    version = "v3.3.2",
)

go_repository(
    name = "com_github_google_pprof",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/pprof",
    sum = "h1:Xim43kblpZXfIBQsbuBVKCudVG457BR2GZFIz3uw3hQ=",
    version = "v0.0.0-20221118152302-e6195bd50e26",
)

go_repository(
    name = "com_github_google_renameio",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    build_file_proto_mode = "disable",
    importpath = "github.com/google/uuid",
    sum = "h1:KjJaJ9iWZ3jOFZIf1Lqf4laDRCasjl0BCmnEGxkdLb4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_googleapis_enterprise_certificate_proxy",
    build_file_proto_mode = "disable",
    importpath = "github.com/googleapis/enterprise-certificate-proxy",
    sum = "h1:UR4rDjcgpgEnqpIEvkiqTYKBCKLNmlge2eVjoZfySzM=",
    version = "v0.2.5",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:A+gCJKdRfqXkr+BIRGtZLibNXf0m1f9E4HG56etFpas=",
    version = "v2.12.0",
)

go_repository(
    name = "com_github_gorilla_securecookie",
    build_file_proto_mode = "disable",
    importpath = "github.com/gorilla/securecookie",
    sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    build_file_proto_mode = "disable",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:UH//fgunKIs4JdUbpDl1VZCDaL56wXCB/5+wF6uHfaI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    build_file_generation = "on",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    build_file_generation = "on",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
    sum = "h1:YBftPWNWd4WwGqtY2yeZL2ef8rHAxPBD8KFhJpmcqms=",
    version = "v2.16.0",
)

go_repository(
    name = "com_github_hailocab_go_hostpool",
    build_file_proto_mode = "disable",
    importpath = "github.com/hailocab/go-hostpool",
    sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
    version = "v0.0.0-20160125115350-e80d13ce29ed",
)

go_repository(
    name = "com_github_hdrhistogram_hdrhistogram_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/HdrHistogram/hdrhistogram-go",
    sum = "h1:5IcZpTvzydCQeHzK4Ef/D5rrSqwxob0t8PQPMybUNFM=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_iancoleman_strcase",
    build_file_proto_mode = "disable",
    importpath = "github.com/iancoleman/strcase",
    sum = "h1:05I4QRnGpI0m37iZQRuskXh+w77mr6Z41lwQzuHLwW0=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    build_file_proto_mode = "disable",
    importpath = "github.com/jessevdk/go-flags",
    sum = "h1:1jKYvbxEjfUl0fmqTCOfonvskHHXMjBySTLW4y9LFvc=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    build_file_proto_mode = "disable",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    build_file_proto_mode = "disable",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    build_file_proto_mode = "disable",
    importpath = "github.com/jmoiron/sqlx",
    sum = "h1:wv+0IJZfL5z0uZoUjlpKgHkgaFSYD+r9CfrXjEXsO7w=",
    version = "v1.3.4",
)

go_repository(
    name = "com_github_josharian_intern",
    build_file_proto_mode = "disable",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jpillora_backoff",
    build_file_proto_mode = "disable",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    build_file_proto_mode = "disable",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_jung_kurt_gofpdf",
    build_file_proto_mode = "disable",
    importpath = "github.com/jung-kurt/gofpdf",
    sum = "h1:PJr+ZMXIecYc1Ey2zucXdR73SMBtgjPgwa31099IMv0=",
    version = "v1.0.3-0.20190309125859-24315acbbda5",
)

go_repository(
    name = "com_github_kballard_go_shellquote",
    build_file_proto_mode = "disable",
    importpath = "github.com/kballard/go-shellquote",
    sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
    version = "v0.0.0-20180428030007-95032a82bc51",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    build_file_proto_mode = "disable",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    build_file_proto_mode = "disable",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:sxCkb+qR91z4vsqw4vGGZlDgPz3G7gjaLyK3V8y70BU=",
    version = "v2.2.3",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    build_file_proto_mode = "disable",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_kr_logfmt",
    build_file_proto_mode = "disable",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    build_file_proto_mode = "disable",
    importpath = "github.com/kr/pretty",
    sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_kr_pty",
    build_file_proto_mode = "disable",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    build_file_proto_mode = "disable",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_labstack_echo_v4",
    build_file_proto_mode = "disable",
    importpath = "github.com/labstack/echo/v4",
    sum = "h1:GliPYSpzGKlyOhqIbG8nmHBo3i1saKWFOgh41AN3b+Y=",
    version = "v4.9.1",
)

go_repository(
    name = "com_github_labstack_gommon",
    build_file_proto_mode = "disable",
    importpath = "github.com/labstack/gommon",
    sum = "h1:y7cvthEAEbU0yHOf4axH8ZG2NH8knB9iNSoTO8dyIk8=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_lib_pq",
    build_file_proto_mode = "disable",
    importpath = "github.com/lib/pq",
    sum = "h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=",
    version = "v1.10.9",
)

go_repository(
    name = "com_github_mailru_easyjson",
    build_file_proto_mode = "disable",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    build_file_proto_mode = "disable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
    version = "v0.1.13",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    build_file_proto_mode = "disable",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:JITubQf0MOLdlGRuRq+jtsDlekdYPia9ZFsB8h/APPA=",
    version = "v0.0.19",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    build_file_proto_mode = "disable",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:+xnbZSEeDbOIg5/mE6JF0w6n9duR1l3/WmbinWVwUuU=",
    version = "v0.0.14",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    build_file_proto_mode = "disable",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:yOQRA0RpS5PFz/oikGwBEqvAWhWg5ufRz4ETLjwpU1Y=",
    version = "v1.14.16",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    build_file_proto_mode = "disable",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    build_file_proto_mode = "disable",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    build_file_proto_mode = "disable",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    build_file_proto_mode = "disable",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
    version = "v0.0.0-20190716064945-2f068394615f",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    build_file_proto_mode = "disable",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    build_file_proto_mode = "disable",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_olivere_elastic_v7",
    build_file_proto_mode = "disable",
    importpath = "github.com/olivere/elastic/v7",
    sum = "h1:R7CXvbu8Eq+WlsLgxmKVKPox0oOwAE/2T9Si5BnvK6E=",
    version = "v7.0.32",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pborman_uuid",
    build_file_proto_mode = "disable",
    importpath = "github.com/pborman/uuid",
    sum = "h1:+ZZIw58t/ozdjRaXh/3awHfmWRbzYxJoAdNJxe/3pvw=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_pkg_errors",
    build_file_proto_mode = "disable",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    build_file_proto_mode = "disable",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prashantv_protectmem",
    build_file_proto_mode = "disable",
    importpath = "github.com/prashantv/protectmem",
    sum = "h1:AA9vgIBDjMHPC2McaGPojgV2dcI78ZC0TLNhYCXEKH8=",
    version = "v0.0.0-20171002184600-e20412882b3a",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    build_file_proto_mode = "disable",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:yk/hx9hDbrGHovbci4BY+pRMfSuuat626eFsHb7tmT8=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    build_file_proto_mode = "disable",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:5lQXD3cAg1OXBf4Wq03gTrXHeaV0TQvGfUooCfx1yqY=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_prometheus_common",
    build_file_proto_mode = "disable",
    importpath = "github.com/prometheus/common",
    sum = "h1:+5BrQJwiBB9xsMygAB3TNvpQKOwlkc25LbISbrdOOfY=",
    version = "v0.44.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    build_file_proto_mode = "disable",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:5EAgkfkMl659uZPbe9AS2N68a7Cc1TJbPEuGzFuRbyk=",
    version = "v0.11.0",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    build_file_proto_mode = "disable",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
    version = "v0.0.0-20201227073835-cf1acfcdf475",
)

go_repository(
    name = "com_github_remyoudompheng_bigfft",
    build_file_proto_mode = "disable",
    importpath = "github.com/remyoudompheng/bigfft",
    sum = "h1:W09IVJc94icq4NjY3clb7Lk8O1qJ8BdBEF8z0ibU0rE=",
    version = "v0.0.0-20230129092748-24d4a6f8daec",
)

go_repository(
    name = "com_github_rivo_uniseg",
    build_file_proto_mode = "disable",
    importpath = "github.com/rivo/uniseg",
    sum = "h1:8TfxU8dW6PdqD27gjM8MVNuicgxIjxpm4K7x4jp8sis=",
    version = "v0.4.4",
)

go_repository(
    name = "com_github_robfig_cron",
    build_file_proto_mode = "disable",
    importpath = "github.com/robfig/cron",
    sum = "h1:ZjScXvvxeQ63Dbyxy76Fj3AT3Ut0aKsyd2/tl3DTMuQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_robfig_cron_v3",
    build_file_proto_mode = "disable",
    importpath = "github.com/robfig/cron/v3",
    sum = "h1:WdRxkvbJztn8LMz/QEvLN5sBU+xKpSqwwUO1Pjr4qDs=",
    version = "v3.0.1",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    build_file_proto_mode = "disable",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    build_file_proto_mode = "disable",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_samuel_go_thrift",
    build_file_proto_mode = "disable",
    importpath = "github.com/samuel/go-thrift",
    sum = "h1:gD4vkYmuoWVgdV6UwI3tPo9MtMfVoIRY+Xn9919SJBg=",
    version = "v0.0.0-20190219015601-e8b6b52668fe",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    build_file_proto_mode = "disable",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
    version = "v1.9.3",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    build_file_proto_mode = "disable",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:T/YLemO5Yp7KPzS+lVtu+WsHn8yoSwTfItdAd1r3cck=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_smartystreets_go_aws_auth",
    build_file_proto_mode = "disable",
    importpath = "github.com/smartystreets/go-aws-auth",
    sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
    version = "v0.0.0-20180515143844-0c1422d1fdb9",
)

go_repository(
    name = "com_github_smartystreets_gunit",
    build_file_proto_mode = "disable",
    importpath = "github.com/smartystreets/gunit",
    sum = "h1:tyWYZffdPhQPfK5VsMQXfauwnJkqg7Tv5DLuQVYxq3Q=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_stretchr_objx",
    build_file_proto_mode = "disable",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    build_file_proto_mode = "disable",
    importpath = "github.com/stretchr/testify",
    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
    version = "v1.8.4",
)

go_repository(
    name = "com_github_temporalio_ringpop_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/temporalio/ringpop-go",
    sum = "h1:V1U9fvhusDJ1pyAvQWg0+u6mQ+o5WtRfMbnnTIZe0Fo=",
    version = "v0.0.0-20230606200434-b5c079f412d3",
)

go_repository(
    name = "com_github_temporalio_tchannel_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/temporalio/tchannel-go",
    sum = "h1:Fs3LdlF7xbnOWHymbFmvIEuxIEt1dNRCfaDkoajSaZk=",
    version = "v1.22.1-0.20220818200552-1be8d8cffa5b",
)

go_repository(
    name = "com_github_temporalio_tctl_kit",
    build_file_proto_mode = "disable",
    importpath = "github.com/temporalio/tctl-kit",
    sum = "h1:E1iAre7/4VvSJri8uOnItKVsMKnP+WEQourm+zVO0cc=",
    version = "v0.0.0-20230328153839-577f95d16fa0",
)

go_repository(
    name = "com_github_temporalio_ui_server_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/temporalio/ui-server/v2",
    sum = "h1:KS5UIZ/2WElWgyw7ClBmW7/sxBEo4qAlVcMSV/elrrw=",
    version = "v2.21.4",
)

go_repository(
    name = "com_github_twmb_murmur3",
    build_file_proto_mode = "disable",
    importpath = "github.com/twmb/murmur3",
    sum = "h1:8Yt9taO/WN3l08xErzjeschgZU2QSrwm1kclYq+0aRg=",
    version = "v1.1.8",
)

go_repository(
    name = "com_github_uber_common_bark",
    build_file_proto_mode = "disable",
    importpath = "github.com/uber-common/bark",
    sum = "h1:DkuZCBaQS9LWuNAPrCO6yQVANckIX3QI0QwLemUnzCo=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_uber_go_tally",
    build_file_proto_mode = "disable",
    importpath = "github.com/uber-go/tally",
    sum = "h1:9hLSgNBP28CjIaDmAuRTq9qV+UZY+9PcvAkXO4nNMwg=",
    version = "v3.3.15+incompatible",
)

go_repository(
    name = "com_github_uber_go_tally_v4",
    build_file_proto_mode = "disable",
    importpath = "github.com/uber-go/tally/v4",
    sum = "h1:YiKvvMKCCXlCKXI0i1hVk+xda8YxdIpjeFXohpvn8Zo=",
    version = "v4.1.7",
)

go_repository(
    name = "com_github_uber_jaeger_client_go",
    build_file_proto_mode = "disable",
    importpath = "github.com/uber/jaeger-client-go",
    sum = "h1:D6wyKGCecFaSRUpo8lCVbaOOb6ThwMmTEbhRwtKR97o=",
    version = "v2.30.0+incompatible",
)

go_repository(
    name = "com_github_uber_jaeger_lib",
    build_file_proto_mode = "disable",
    importpath = "github.com/uber/jaeger-lib",
    sum = "h1:td4jdvLcExb4cBISKIpHuGoVXh+dVKhn2Um6rjCsSsg=",
    version = "v2.4.1+incompatible",
)

go_repository(
    name = "com_github_urfave_cli",
    build_file_proto_mode = "disable",
    importpath = "github.com/urfave/cli",
    sum = "h1:ebbhrRiGK2i4naQJr+1Xj92HXZCrK7MsyTS/ob3HnAk=",
    version = "v1.22.14",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    build_file_proto_mode = "disable",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:8xSQ6szndafKVRmfyeUMxkNUJQMjL1F2zmsZ+qHpfho=",
    version = "v2.27.1",
)

go_repository(
    name = "com_github_valyala_bytebufferpool",
    build_file_proto_mode = "disable",
    importpath = "github.com/valyala/bytebufferpool",
    sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_valyala_fasttemplate",
    build_file_proto_mode = "disable",
    importpath = "github.com/valyala/fasttemplate",
    sum = "h1:TVEnxayobAdVkhQfrfes2IzOB6o+z4roRkPF52WA1u4=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_xrash_smetrics",
    build_file_proto_mode = "disable",
    importpath = "github.com/xrash/smetrics",
    sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
    version = "v0.0.0-20201216005158-039620a65673",
)

go_repository(
    name = "com_github_xwb1989_sqlparser",
    build_file_proto_mode = "disable",
    importpath = "github.com/xwb1989/sqlparser",
    sum = "h1:zzrxE1FKn5ryBNl9eKOeqQ58Y/Qpo3Q9QNxKHX5uzzQ=",
    version = "v0.0.0-20180606152119-120387863bf2",
)

go_repository(
    name = "com_github_yuin_goldmark",
    build_file_proto_mode = "disable",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
    version = "v1.4.13",
)

go_repository(
    name = "com_google_cloud_go",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go",
    sum = "h1:tyNdfIxjzaWctIiLYOTalaLKZ17SI44SKFW26QbOhME=",
    version = "v0.110.8",
)

go_repository(
    name = "com_google_cloud_go_accessapproval",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/accessapproval",
    sum = "h1:W55SFrY6EVlcmmRGUk0rGhuy3j4fn7UtEocib/zADVE=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_accesscontextmanager",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/accesscontextmanager",
    sum = "h1:jcOXen2u13aHgOHibUjxyPI+fZzVhElxy2gzJJlOOHg=",
    version = "v1.8.2",
)

go_repository(
    name = "com_google_cloud_go_aiplatform",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/aiplatform",
    sum = "h1:g+y03dll9HnX9U0oBKIqUOI+8VQWT1QJF12VGxkal0Q=",
    version = "v1.51.1",
)

go_repository(
    name = "com_google_cloud_go_analytics",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/analytics",
    sum = "h1:SScWR8i/M8h7h3lFKtOYcj0r4272aL+KvRRrsu39Vec=",
    version = "v0.21.4",
)

go_repository(
    name = "com_google_cloud_go_apigateway",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/apigateway",
    sum = "h1:I46jVrhr2M1JJ1lK7JGn2BvybN44muEh+LSjBQ1l9hw=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_apigeeconnect",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/apigeeconnect",
    sum = "h1:7LzOTW34EH2julg0MQVt+U9ZdmiCKcg6fef/ugKL2Xo=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_apigeeregistry",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/apigeeregistry",
    sum = "h1:MESEjKSfz4TvLAzT2KPimDDvhOyQlcq7aFFREG2PRt4=",
    version = "v0.7.2",
)

go_repository(
    name = "com_google_cloud_go_appengine",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/appengine",
    sum = "h1:0/OFV0FQKgi0AB4E8NuYN0JY3hJzND4ftRpK7P26uaw=",
    version = "v1.8.2",
)

go_repository(
    name = "com_google_cloud_go_area120",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/area120",
    sum = "h1:h/wMtPPsgFJfMce1b9M24Od8RuKt8CWENwr+X24tBhE=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_artifactregistry",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/artifactregistry",
    sum = "h1:Ssv6f+jgfhDdhu43AaHUaSosIYpQ+TPCJNwqYSJT1AE=",
    version = "v1.14.3",
)

go_repository(
    name = "com_google_cloud_go_asset",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/asset",
    sum = "h1:+9f5/s/U0AGZSPLTOMcXSZ5NDB5jQ2Szr+WQPgPA8bk=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_assuredworkloads",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/assuredworkloads",
    sum = "h1:EbPyk3fC8sTxSIPoFrCR9P1wRTVdXcRxvPqFK8/wdso=",
    version = "v1.11.2",
)

go_repository(
    name = "com_google_cloud_go_automl",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/automl",
    sum = "h1:kUN4Y6N61AsNdXsdZIug1c+2pTJ5tg9xUA6+yn0Wf8Y=",
    version = "v1.13.2",
)

go_repository(
    name = "com_google_cloud_go_baremetalsolution",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/baremetalsolution",
    sum = "h1:uRpZsKiWFDyT1sARZVRKqnOmf2mpRfVas7KMC3/MA4I=",
    version = "v1.2.1",
)

go_repository(
    name = "com_google_cloud_go_batch",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/batch",
    sum = "h1:+8ZogCLFauglOE5ybTCWscoexD7Z8k4XW27RVTKNEoo=",
    version = "v1.5.1",
)

go_repository(
    name = "com_google_cloud_go_beyondcorp",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/beyondcorp",
    sum = "h1:uQpsXwttlV0+AXHdB5qaZl1mz2SsyYV1PKgTR74noaQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:LHIc9E7Kw+ftFpQFKzZYBB88IAFz7qONawXXx0F3QBo=",
    version = "v1.56.0",
)

go_repository(
    name = "com_google_cloud_go_billing",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/billing",
    sum = "h1:ozS/MNj6KKz8Reuw7tIG8Ycucq/YpSf3u3XCqrupbcg=",
    version = "v1.17.2",
)

go_repository(
    name = "com_google_cloud_go_binaryauthorization",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/binaryauthorization",
    sum = "h1:i2S+/G36VA1UG8gdcQLpq5I58/w/RzAnjQ65scKozFg=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_certificatemanager",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/certificatemanager",
    sum = "h1:Xytp8O0/EDh2nVscHhFQpicY9YAT3f3R7D7pv/z29uE=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_channel",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/channel",
    sum = "h1:+1B+Gj/3SJSLGJZXCp3dWiseMVHoSZ7Xo6Klg1fqM64=",
    version = "v1.17.1",
)

go_repository(
    name = "com_google_cloud_go_cloudbuild",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/cloudbuild",
    sum = "h1:Tp0ITIlFam7T8K/TyeceITtpw1f8+KxVKwYyiyWDPK8=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_clouddms",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/clouddms",
    sum = "h1:LrtqeR2xKV3juG5N7eeUgW+PqdMClOWH2U9PN3EpfFw=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_cloudtasks",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/cloudtasks",
    sum = "h1:IoJI49JClvv2+NYvcABRgTO9y4veAUFlaOTigm+xXqE=",
    version = "v1.12.2",
)

go_repository(
    name = "com_google_cloud_go_compute",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/compute",
    sum = "h1:V97tBoDaZHb6leicZ1G6DLK2BAaZLJ/7+9BB/En3hR0=",
    version = "v1.23.1",
)

go_repository(
    name = "com_google_cloud_go_compute_metadata",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/compute/metadata",
    sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
    version = "v0.2.3",
)

go_repository(
    name = "com_google_cloud_go_contactcenterinsights",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/contactcenterinsights",
    sum = "h1:dEfCjtdYjS3n8/1HEKbJaOL31l3dEs3q9aeaNsyrJBc=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_container",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/container",
    sum = "h1:1CXjOL/dZZ2jXX1CYWqlxmXqJbZo8HwQX4DJxLzgQWo=",
    version = "v1.26.1",
)

go_repository(
    name = "com_google_cloud_go_containeranalysis",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/containeranalysis",
    sum = "h1:PHh4KTcMpCjYgxfV+TzvP24wolTGP9lGbqh9sBNHxjs=",
    version = "v0.11.1",
)

go_repository(
    name = "com_google_cloud_go_datacatalog",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/datacatalog",
    sum = "h1:xJp9mZrc2HPaoxIz3sP9pCmf/impifweQ/yGG9VBfio=",
    version = "v1.18.1",
)

go_repository(
    name = "com_google_cloud_go_dataflow",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dataflow",
    sum = "h1:cpu2OeNxnYVadAIXETLRS5riz3KUR8ErbTojAQTFJVg=",
    version = "v0.9.2",
)

go_repository(
    name = "com_google_cloud_go_dataform",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dataform",
    sum = "h1:l155O3DS7pfyR91maS4l92bEjKbkbWie3dpgltZ1Q68=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_datafusion",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/datafusion",
    sum = "h1:CIIXp4bbwck49ZTV/URabJaV48jVB86THyVBWGgeDjw=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_datalabeling",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/datalabeling",
    sum = "h1:4N5mbjauemzaatxGOFVpV2i8HiXSUUhyNRBU+dCBHl0=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_dataplex",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dataplex",
    sum = "h1:8Irss8sIalm/X8r0Masv5KJRkddcxov3TiW8W96FmC4=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_dataqna",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dataqna",
    sum = "h1:vJ9JVKDgDG7AQMbTD8pdWaogJ4c/yHn0qer+q0nFIaw=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:0P9WcsQeTWjuD1H14JIY7XQscIPQ4Laje8ti96IC5vg=",
    version = "v1.15.0",
)

go_repository(
    name = "com_google_cloud_go_datastream",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/datastream",
    sum = "h1:XWiXV1hzs8oAd54//wcb1L15Jl7MnZ/cY2B8XCmu0xE=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_deploy",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/deploy",
    sum = "h1:eV5MdoQJGdac/k7D97SDjD8iLE4jCzL42UCAgG6j0iE=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_dialogflow",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dialogflow",
    sum = "h1:Ml/hgEzU3AN0tjNSSv4/QmG1nqwYEsiCySKMkWMqUmI=",
    version = "v1.44.1",
)

go_repository(
    name = "com_google_cloud_go_dlp",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/dlp",
    sum = "h1:sWOATigjZOKmA2rVOSjIcKLCtL2ifdawaukx+H9iffk=",
    version = "v1.10.2",
)

go_repository(
    name = "com_google_cloud_go_documentai",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/documentai",
    sum = "h1:IAKWBngDFTxABdAH52uAn0osPDemyegyRmf5IQKznHw=",
    version = "v1.23.2",
)

go_repository(
    name = "com_google_cloud_go_domains",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/domains",
    sum = "h1:SjpTtaTNRPPajrGiZEtxz9dpElO4PxuDWFvU4JpV1gk=",
    version = "v0.9.2",
)

go_repository(
    name = "com_google_cloud_go_edgecontainer",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/edgecontainer",
    sum = "h1:B+Acb/0frXUxc60i6lC0JtXrBFAKoS7ZELmet9+ySo8=",
    version = "v1.1.2",
)

go_repository(
    name = "com_google_cloud_go_errorreporting",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/errorreporting",
    sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_essentialcontacts",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/essentialcontacts",
    sum = "h1:xrGTLRTzunQk5XhBIkdftuC00B9MUoEXi7Pjgeu1kMM=",
    version = "v1.6.3",
)

go_repository(
    name = "com_google_cloud_go_eventarc",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/eventarc",
    sum = "h1:FmEcxG5rX3LaUB2nRjf2Pas5J5TtVrVznaHN5rxYxnQ=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_filestore",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/filestore",
    sum = "h1:/Nnk5pOoY1Lx6A42hJ2eBYcBfqKvLcnh8fV4egopvY4=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:/3S4RssUV4GO/kvgJZB+tayjhOfyAHs+KcpJgRVu/Qk=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_functions",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/functions",
    sum = "h1:DpT51zU3UMTt64efB4a9hE9B98Kb0fZC3IfaVp7GnkE=",
    version = "v1.15.2",
)

go_repository(
    name = "com_google_cloud_go_gkebackup",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/gkebackup",
    sum = "h1:1fnA934a/0oz7nU22gTzmGYFVi6V13Q/hCkdC99K178=",
    version = "v1.3.2",
)

go_repository(
    name = "com_google_cloud_go_gkeconnect",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/gkeconnect",
    sum = "h1:AuR3YNK0DgLVrmcc8o4sBrU0dVs/SULSuLh4Gmn1e10=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_gkehub",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/gkehub",
    sum = "h1:7rddjV52z0RbToFYj1B39R9dsn+6IXgx4DduEH7N25Q=",
    version = "v0.14.2",
)

go_repository(
    name = "com_google_cloud_go_gkemulticloud",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/gkemulticloud",
    sum = "h1:V82LxEvFIGJnebn7BBdOUKcVlNQqBaubbKtLgRicHow=",
    version = "v1.0.1",
)

go_repository(
    name = "com_google_cloud_go_gsuiteaddons",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/gsuiteaddons",
    sum = "h1:vR7E1gR85x0wlbUek3cZYJ67U67GpNrboNCRiF/VSSc=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_iam",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:18tKG7DzydKWUnLjonWcJO6wjSCAtzh4GcRKlH/Hrzc=",
    version = "v1.1.3",
)

go_repository(
    name = "com_google_cloud_go_iap",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/iap",
    sum = "h1:J5r6CL6EakRmsMRIm2yV0PF5zfIm4sMQbQfPhSTnRzA=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_ids",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/ids",
    sum = "h1:KqvR28pAnIss6d2pmGOQ+Fcsi3FOWDVhqdr6QaVvqsI=",
    version = "v1.4.2",
)

go_repository(
    name = "com_google_cloud_go_iot",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/iot",
    sum = "h1:qFNv3teWkONIPmuY2mzodEnHb6E67ch2OZ6216ycUiU=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_kms",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:RYsbxTRmk91ydKCzekI2YjryO4c5Y2M80Zwcs9/D/cI=",
    version = "v1.15.3",
)

go_repository(
    name = "com_google_cloud_go_language",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/language",
    sum = "h1:BjU7Ljhh0ZYnZC8jZwiezf1FH75yijJ4raAScseqCns=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_lifesciences",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/lifesciences",
    sum = "h1:0naTq5qUWoRt/b5P+SZ/0mun7ZTlhpJZJsUxhCmLv1c=",
    version = "v0.9.2",
)

go_repository(
    name = "com_google_cloud_go_logging",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/logging",
    sum = "h1:26skQWPeYhvIasWKm48+Eq7oUqdcdbwsCVwz5Ys0FvU=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_longrunning",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/longrunning",
    sum = "h1:u+oFqfEwwU7F9dIELigxbe0XVnBAo9wqMuQLA50CZ5k=",
    version = "v0.5.2",
)

go_repository(
    name = "com_google_cloud_go_managedidentities",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/managedidentities",
    sum = "h1:QijSmmWHb3EzYQr8SrjWe941ba9G5sTCF5PvhhMM8CM=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_maps",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/maps",
    sum = "h1:/wp8wImC3tHIHOoaQGRA+KyH3as/Dvp+3J/NqJQBiPQ=",
    version = "v1.4.1",
)

go_repository(
    name = "com_google_cloud_go_mediatranslation",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/mediatranslation",
    sum = "h1:nyBZbNX1j34H00n+irnQraCogrkRWntQsDoA6s8OfKo=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_memcache",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/memcache",
    sum = "h1:WLJALO3FxuStMiYdSQwiQBDBcs4G8DDwZQmXK+YzAWk=",
    version = "v1.10.2",
)

go_repository(
    name = "com_google_cloud_go_metastore",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/metastore",
    sum = "h1:tLemzNMjKY+xdJUDQt9v5+fQqSufTNgKHHQmihG5ay8=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_monitoring",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/monitoring",
    sum = "h1:CTklIuUkS5nCricGojPwdkSgPsCTX2HmYTxFDg+UvpU=",
    version = "v1.16.1",
)

go_repository(
    name = "com_google_cloud_go_networkconnectivity",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/networkconnectivity",
    sum = "h1:uR+ASueYNodsPCd9wcYEedqjH4+LaCkKqltRBF6CmB4=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_networkmanagement",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/networkmanagement",
    sum = "h1:ZK6i6FVQNc1t3fecM3hf9Nu6Kr9C95xr+zMVORYd8ak=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_networksecurity",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/networksecurity",
    sum = "h1:fA73AX//KWaqNKOvuQ00WUD3Z/XMhiMhHSFTEl2Wxec=",
    version = "v0.9.2",
)

go_repository(
    name = "com_google_cloud_go_notebooks",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/notebooks",
    sum = "h1:j/G3r6SPoWzD6CZZrDffZGwgGALvxWwtKJHJ4GF17WA=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_optimization",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/optimization",
    sum = "h1:71wTxJz8gRrVEHF4fw18sGynAyNQwatxCJBI3m3Rd4c=",
    version = "v1.5.1",
)

go_repository(
    name = "com_google_cloud_go_orchestration",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/orchestration",
    sum = "h1:lb+Vphr+x2V9ukHwLjyaXJpbPuPhaKdobQx3UAOeSsQ=",
    version = "v1.8.2",
)

go_repository(
    name = "com_google_cloud_go_orgpolicy",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/orgpolicy",
    sum = "h1:Dnfh5sj3aIAuJzH4Q4rBp6lCJ/IdXRBbwQ0/nQsUySE=",
    version = "v1.11.2",
)

go_repository(
    name = "com_google_cloud_go_osconfig",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/osconfig",
    sum = "h1:AjHbw8MgKKaTFAEJWGdOYtMED3wUXKLtvdfP8Uzbuy0=",
    version = "v1.12.2",
)

go_repository(
    name = "com_google_cloud_go_oslogin",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/oslogin",
    sum = "h1:r3JYeLf004krfXhRMDfYKlBdMgDDc2q2PM1bomb5Luw=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_phishingprotection",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/phishingprotection",
    sum = "h1:BIv/42ooQXh/jW8BW2cgO0E6yRPbEdvqH3JzKV7BlmI=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_policytroubleshooter",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/policytroubleshooter",
    sum = "h1:92YSoPZE62QkNM0G6Nl6PICKUyv4aNgsdtWWceJR6ys=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_privatecatalog",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/privatecatalog",
    sum = "h1:gxL4Kn9IXt3tdIOpDPEDPI/kBBLVzaAX5wq6IbOYi8A=",
    version = "v0.9.2",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:6SPCPvWav64tj0sVX/+npCBKhUi/UjJehy9op/V3p2g=",
    version = "v1.33.0",
)

go_repository(
    name = "com_google_cloud_go_pubsublite",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/pubsublite",
    sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise_v2",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/recaptchaenterprise/v2",
    sum = "h1:06V6+edT20PcrFJfH0TVWMZpZCUpSCADgwGwhkMsGmY=",
    version = "v2.8.1",
)

go_repository(
    name = "com_google_cloud_go_recommendationengine",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/recommendationengine",
    sum = "h1:odf0TZXtwoZ5kJaWBlaE9D0AV+WJLLs+/SRSuE4T/ds=",
    version = "v0.8.2",
)

go_repository(
    name = "com_google_cloud_go_recommender",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/recommender",
    sum = "h1:GI4EBCMTLfC8I8R+e13ZaTAa8ZZ0KRPdS99hGtJYyaU=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_redis",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/redis",
    sum = "h1:2ZtIGspMT65wern2rjX35XPCCJxVKF4J0P1S99bac3k=",
    version = "v1.13.2",
)

go_repository(
    name = "com_google_cloud_go_resourcemanager",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/resourcemanager",
    sum = "h1:lC3PjJMHLPlZKqLfan6FkEb3X1F8oCRc1ylY7vRHvDQ=",
    version = "v1.9.2",
)

go_repository(
    name = "com_google_cloud_go_resourcesettings",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/resourcesettings",
    sum = "h1:feqx2EcLRgtmwNHzeLw5Og4Wcy4vcZxw62b0x/QNu60=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_retail",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/retail",
    sum = "h1:ed5hWjpOwfsi6E9kj2AFzkz5ScT3aZs7o3MUM0YITUM=",
    version = "v1.14.2",
)

go_repository(
    name = "com_google_cloud_go_run",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/run",
    sum = "h1:xc46W9kxJI2De9hmpqHEBSSLJhP3bSZl86LdlJa5zm8=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_scheduler",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/scheduler",
    sum = "h1:lgUd1D84JEgNzzHRlcZEIoQ6Ny10YWe8RNH1knhouNk=",
    version = "v1.10.2",
)

go_repository(
    name = "com_google_cloud_go_secretmanager",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/secretmanager",
    sum = "h1:52Z78hH8NBWIqbvIG0wi0EoTaAmSx99KIOAmDXIlX0M=",
    version = "v1.11.2",
)

go_repository(
    name = "com_google_cloud_go_security",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/security",
    sum = "h1:VNpdJNfMeHSJZ+647QtzPrvZ6rWChBklLm/NY64RVW8=",
    version = "v1.15.2",
)

go_repository(
    name = "com_google_cloud_go_securitycenter",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/securitycenter",
    sum = "h1:Epx7Gm9ZRPRiFfwDFplka2zKCS0J3cpm0Et1KwI2tvY=",
    version = "v1.23.1",
)

go_repository(
    name = "com_google_cloud_go_servicedirectory",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/servicedirectory",
    sum = "h1:SXhbxsfQJBsUDeo743x5AnVe8ifC7qjXU3bSTT6t/+Q=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_shell",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/shell",
    sum = "h1:zk0Cf2smbFlAdhBQ5tXESZzzmsTfGc31fJfI6a0SVD8=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:QrJFOpaxCXdXF+GkiruLz642PHxkdj68PbbnLw3O2Zw=",
    version = "v1.50.0",
)

go_repository(
    name = "com_google_cloud_go_speech",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/speech",
    sum = "h1:z035FMLs98jpnqcP5xZZ6Es+g6utbeVoUH64BaTzTSU=",
    version = "v1.19.1",
)

go_repository(
    name = "com_google_cloud_go_storage",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:uOdMxAs8HExqBlnLtnQyP0YkvbiDpdGShGKtx6U/oNM=",
    version = "v1.30.1",
)

go_repository(
    name = "com_google_cloud_go_storagetransfer",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/storagetransfer",
    sum = "h1:CU03oYLauu7xRV25fFmozHZHA/SokLQlC20Ip/UvFro=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_talent",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/talent",
    sum = "h1:TyJqwhmncdW5CL4rzYSYKJrR9YAe0iNqHtJTnnOaEyM=",
    version = "v1.6.3",
)

go_repository(
    name = "com_google_cloud_go_texttospeech",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/texttospeech",
    sum = "h1:Ac53sRkUo8UMSuhyyWRFJvWEaX8vm0EFwwiTAxeVYuU=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_tpu",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/tpu",
    sum = "h1:SAFzyGp6mU37lfLTV0cNQwu7tqH4X8b4RCpQZ1s+mYM=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_trace",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/trace",
    sum = "h1:80Rh4JSqJLfe/xGNrpyO4MQxiFDXcHG1XrsevfmrIRQ=",
    version = "v1.10.2",
)

go_repository(
    name = "com_google_cloud_go_translate",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/translate",
    sum = "h1:gNPBVMINs+aZMB8BW+IfrHLLTfdq0t0GMwa31NmOXY4=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_video",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/video",
    sum = "h1:yMfxQ4N/fXNDsCKNKw9W+FpdrJPj5CDu+FuAJBmGuoo=",
    version = "v1.20.1",
)

go_repository(
    name = "com_google_cloud_go_videointelligence",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/videointelligence",
    sum = "h1:vAKuM4YHwZy1W5P7hGJdfXriovqHHUZKhDBq8o4nqfg=",
    version = "v1.11.2",
)

go_repository(
    name = "com_google_cloud_go_vision_v2",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/vision/v2",
    sum = "h1:o8iiH4UsI6O8wO2Ax2r88fLG1RzYQIFevUQY7hXPZeM=",
    version = "v2.7.3",
)

go_repository(
    name = "com_google_cloud_go_vmmigration",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/vmmigration",
    sum = "h1:ObE8VWzL+xkU22IsPEMvPCWArnSQ85dEwR5fzgaOvA4=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_vmwareengine",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/vmwareengine",
    sum = "h1:Bj9WECvQk1fkx8IG7gqII3+g1CzhqkPOV84WXvifpFg=",
    version = "v1.0.1",
)

go_repository(
    name = "com_google_cloud_go_vpcaccess",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/vpcaccess",
    sum = "h1:3qKiWvzK07eIa943mCvkcZB4gimxaQKKGdNoX01ps7A=",
    version = "v1.7.2",
)

go_repository(
    name = "com_google_cloud_go_webrisk",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/webrisk",
    sum = "h1:1NZppagzdGO0hVMJsUhZQ5a3Iu2cNyNObu85VFcvIVA=",
    version = "v1.9.2",
)

go_repository(
    name = "com_google_cloud_go_websecurityscanner",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/websecurityscanner",
    sum = "h1:V7PhbJ2OvpGHINL67RBhpwU3+g4MOoqOeL/sFYrogeE=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_workflows",
    build_file_proto_mode = "disable",
    importpath = "cloud.google.com/go/workflows",
    sum = "h1:jvhSfcfAoOt0nILm7aZPJAHdpoe571qrJyc2ZlngaJk=",
    version = "v1.12.1",
)

go_repository(
    name = "com_lukechampine_uint128",
    build_file_proto_mode = "disable",
    importpath = "lukechampine.com/uint128",
    sum = "h1:cDdUVfRwDUDovz610ABgFD17nXD4/uDgVHl2sC3+sbo=",
    version = "v1.3.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    build_file_proto_mode = "disable",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/check.v1",
    sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
    version = "v1.0.0-20201130134442-10cb98267c6c",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_inf_v0",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_resty_v1",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
    version = "v2.6.0",
)

go_repository(
    name = "in_gopkg_validator_v2",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/validator.v2",
    sum = "h1:xF0KWyGWXm/LM2G1TrEjqOu4pa6coO9AlWSf3msVfDY=",
    version = "v2.0.1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    build_file_proto_mode = "disable",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "io_opencensus_go",
    build_file_proto_mode = "disable",
    importpath = "go.opencensus.io",
    sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
    version = "v0.24.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    sum = "h1:ZOLJc06r4CB42laIXg/7udr0pbZyuAihN10A/XuiQRY=",
    version = "v0.42.0",
)

go_repository(
    name = "io_opentelemetry_go_otel",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:Z7GVAX/UkAXPKsy94IU+i6thsQS4nb7LviLpnaNeW8s=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
    sum = "h1:t4ZwRPU+emrcvM2e9DHd0Fsf0JTPVcbfa/BhTDF03d0=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlpmetric",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric",
    sum = "h1:f6BwB2OACc3FCbYVznctQ9V6KK7Vq6CjmYXJ7DeSs4E=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlpmetric_otlpmetricgrpc",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc",
    sum = "h1:rm+Fizi7lTM2UefJ1TO347fSRcwmIsUAaZmYmIGBRAo=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
    sum = "h1:cbsD4cUcviQGXdw8+bo5x2wazq10SKz8hEbtCRPcU78=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
    sum = "h1:TVQp/bboR4mhZSav+MdgXB8FaRho1RC8UwVn3T0vjVc=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_prometheus",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/prometheus",
    sum = "h1:whAaiHxOatgtKd+w0dOi//1KUxj3KoPINZdtDaDj3IA=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_metric",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/metric",
    sum = "h1:RbrpwVG1Hfv85LgnZ7+txXioPDoh6EdbZHo26Q3hqOo=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/sdk",
    sum = "h1:Z1Ok1YsijYL0CSJpHt4cS3wDDh7p572grzNrBMiMWgE=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk_metric",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/sdk/metric",
    sum = "h1:Kun8i1eYf48kHH83RucG93ffz0zGV1sh46FAScOTuDI=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_trace",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/trace",
    sum = "h1:8JRpaObFoW0pxuVPapkgH8UhHQj+bJW8jJsCZEu5MQs=",
    version = "v1.16.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:BLOA1cZBAGSbRiNuGCCKiFrCdYB7deeHDeD1SueyOfA=",
    version = "v0.20.0",
)

go_repository(
    name = "io_rsc_pdf",
    build_file_proto_mode = "disable",
    importpath = "rsc.io/pdf",
    sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
    version = "v0.1.1",
)

go_repository(
    name = "io_temporal_go_api",
    build_file_proto_mode = "disable",
    importpath = "go.temporal.io/api",
    sum = "h1:cTq1Fqv4SKCvYuvkF3ZdkLxYgECwGEeCzRqcr4qjRd4=",
    version = "v1.24.1-0.20231020155655-224b196919ea",
)

go_repository(
    name = "io_temporal_go_sdk",
    build_file_proto_mode = "disable",
    importpath = "go.temporal.io/sdk",
    sum = "h1:jC9l9vHHz5OJ7PR6OjrpYSN4+uEG0bLe5rdF9nlMSGk=",
    version = "v1.25.1",
)

go_repository(
    name = "io_temporal_go_server",
    build_file_proto_mode = "disable",
    importpath = "go.temporal.io/server",
    sum = "h1:OrZcCxFjHBBGAdp9ySgvURMdNx5v4+pYCTIhwRkhIXk=",
    version = "v1.23.0-rc0",
)

go_repository(
    name = "io_temporal_go_version",
    build_file_proto_mode = "disable",
    importpath = "go.temporal.io/version",
    sum = "h1:dMrei9l9NyHt8nG6EB8vAwDLLTwx2SvRyucCSumAiig=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_google_api",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/api",
    sum = "h1:RjPESny5CnQRn9V6siglged+DZCgfu9l6mO9dkX9VOg=",
    version = "v0.128.0",
)

go_repository(
    name = "org_golang_google_appengine",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/genproto",
    sum = "h1:+YaDE2r2OG8t/z5qmsh7Y+XXwCbvadxxZ0YY6mTdrVA=",
    version = "v0.0.0-20231016165738-49dd2c1f3d0b",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/grpc",
    sum = "h1:Z5Iec2pjwb+LEOqzpB2MR12/eKFhDPhuqW91O+4bwUk=",
    version = "v1.59.0",
)

go_repository(
    name = "org_golang_google_grpc_examples",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/grpc/examples",
    sum = "h1:soFSbKG2H9FiKip4adXyJL8FEsA+wcFvu1YYMDcKKWY=",
    version = "v0.0.0-20230623203957-0b3a81eabc28",
)

go_repository(
    name = "org_golang_google_protobuf",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/protobuf",
    sum = "h1:g0LDEJHgrBl9N9r17Ru3sqWhkIx2NB67okBHPwC7hs8=",
    version = "v1.31.0",
)

go_repository(
    name = "org_golang_x_crypto",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/crypto",
    sum = "h1:wBqGXzWJW6m1XrIKlAH0Hs1JJ7+9KBwnIO8v66Q9cHc=",
    version = "v0.14.0",
)

go_repository(
    name = "org_golang_x_exp",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/exp",
    sum = "h1:/yRP+0AN7mf5DkD3BAI6TOFnd51gEoDEb8o35jIFtgw=",
    version = "v0.0.0-20230728194245-b0cb94b80691",
)

go_repository(
    name = "org_golang_x_image",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/lint",
    sum = "h1:5hukYrvBGR8/eNkX5mdUezrA6JiaEZDtJb9Ei+1LlBs=",
    version = "v0.0.0-20190930215403-16217165b5de",
)

go_repository(
    name = "org_golang_x_mobile",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/mod",
    sum = "h1:bUO06HqtnRcc/7l71XBe4WcqTZ+3AH1J59zWDDwLKgU=",
    version = "v0.11.0",
)

go_repository(
    name = "org_golang_x_net",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/net",
    sum = "h1:pVaXccu2ozPjCXewfr1S7xza/zcXTity9cCdXQYSjIM=",
    version = "v0.17.0",
)

go_repository(
    name = "org_golang_x_oauth2",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/oauth2",
    sum = "h1:vPL4xzxBM4niKCW6g9whtaWVXTJf1U5e4aZxxFx/gbU=",
    version = "v0.11.0",
)

go_repository(
    name = "org_golang_x_sync",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/sync",
    sum = "h1:ftCYgMx6zT/asHUrPw8BLLscYtGznsLAnjq5RH9P66E=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_sys",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/sys",
    sum = "h1:Af8nKPmuFypiUBjVoU9V20FiaFXOcuZI21p0ycVYYGE=",
    version = "v0.13.0",
)

go_repository(
    name = "org_golang_x_term",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/term",
    sum = "h1:bb+I9cTfFazGW51MZqBVmZy7+JEJMouUHTUSKVQLBek=",
    version = "v0.13.0",
)

go_repository(
    name = "org_golang_x_text",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/text",
    sum = "h1:ablQoSUd0tRdKxZewP80B+BaqeKJuVhuRxj/dkrun3k=",
    version = "v0.13.0",
)

go_repository(
    name = "org_golang_x_time",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/time",
    sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/tools",
    sum = "h1:tvDr/iQoUqNdohiYm0LmmKcBk+q86lb9EprIUFhHHGg=",
    version = "v0.10.0",
)

go_repository(
    name = "org_golang_x_xerrors",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/xerrors",
    sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
    version = "v0.0.0-20220907171357-04be3eba64a2",
)

go_repository(
    name = "org_gonum_v1_gonum",
    build_file_proto_mode = "disable",
    importpath = "gonum.org/v1/gonum",
    sum = "h1:CCXrcPKiGGotvnN6jfUsKk4rRqm7q09/YbKb5xCEvtM=",
    version = "v0.8.2",
)

go_repository(
    name = "org_gonum_v1_netlib",
    build_file_proto_mode = "disable",
    importpath = "gonum.org/v1/netlib",
    sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
    version = "v0.0.0-20190313105609-8cb42192e0e0",
)

go_repository(
    name = "org_gonum_v1_plot",
    build_file_proto_mode = "disable",
    importpath = "gonum.org/v1/plot",
    sum = "h1:Qh4dB5D/WpoUUp3lSod7qgoyEHbDGPUWjIbnqdqqe1k=",
    version = "v0.0.0-20190515093506-e2840ee46a6b",
)

go_repository(
    name = "org_modernc_cc_v3",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/cc/v3",
    sum = "h1:QoR1Sn3YWlmA1T4vLaKZfawdVtSiGx8H+cEojbC7v1Q=",
    version = "v3.41.0",
)

go_repository(
    name = "org_modernc_ccgo_v3",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/ccgo/v3",
    sum = "h1:af6KNtFgsVmnDYrWk3PQCS9XT6BXe7o3ZFJKkIKvXNQ=",
    version = "v3.16.14",
)

go_repository(
    name = "org_modernc_ccorpus",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/ccorpus",
    sum = "h1:J16RXiiqiCgua6+ZvQot4yUuUy8zxgqbqEEUuGPlISk=",
    version = "v1.11.6",
)

go_repository(
    name = "org_modernc_httpfs",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/httpfs",
    sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
    version = "v1.0.6",
)

go_repository(
    name = "org_modernc_libc",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/libc",
    sum = "h1:uvJSeCKL/AgzBo2yYIPPTy82v21KgGnizcGYfBHaNuM=",
    version = "v1.24.1",
)

go_repository(
    name = "org_modernc_mathutil",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/mathutil",
    sum = "h1:rV0Ko/6SfM+8G+yKiyI830l3Wuz1zRutdslNoQ0kfiQ=",
    version = "v1.5.0",
)

go_repository(
    name = "org_modernc_memory",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/memory",
    sum = "h1:i6mzavxrE9a30whzMfwf7XWVODx2r5OYXvU46cirX7o=",
    version = "v1.6.0",
)

go_repository(
    name = "org_modernc_opt",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/opt",
    sum = "h1:3XOZf2yznlhC+ibLltsDGzABUGVx8J6pnFMS3E4dcq4=",
    version = "v0.1.3",
)

go_repository(
    name = "org_modernc_sqlite",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/sqlite",
    sum = "h1:nrSBg4aRQQwq59JpvGEQ15tNxoO5pX/kUjcRNwSAGQM=",
    version = "v1.23.1",
)

go_repository(
    name = "org_modernc_strutil",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/strutil",
    sum = "h1:fNMm+oJklMGYfU9Ylcywl0CO5O6nTfaowNsh2wpPjzY=",
    version = "v1.1.3",
)

go_repository(
    name = "org_modernc_tcl",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/tcl",
    sum = "h1:C4ybAYCGJw968e+Me18oW55kD/FexcHbqH2xak1ROSY=",
    version = "v1.15.2",
)

go_repository(
    name = "org_modernc_token",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/token",
    sum = "h1:Xl7Ap9dKaEs5kLoOQeQmPWevfnk/DM5qcLcYlA8ys6Y=",
    version = "v1.1.0",
)

go_repository(
    name = "org_modernc_z",
    build_file_proto_mode = "disable",
    importpath = "modernc.org/z",
    sum = "h1:zDJf6iHjrnB+WRD88stbXokugjyc0/pB91ri1gO6LZY=",
    version = "v1.7.3",
)

go_repository(
    name = "org_uber_go_atomic",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/atomic",
    sum = "h1:ZvwS0R+56ePWxUNi+Atn9dWONBPp/AUETXlHW0DxSjE=",
    version = "v1.11.0",
)

go_repository(
    name = "org_uber_go_dig",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/dig",
    sum = "h1:5Chju+tUvcC+N7N6EV08BJz41UZuO3BmHcN4A287ZLI=",
    version = "v1.17.0",
)

go_repository(
    name = "org_uber_go_fx",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/fx",
    sum = "h1:ZMC/pnRvhsthOZh9MZjMq5U8Or3mA9zBSPaLnzs3ihQ=",
    version = "v1.20.0",
)

go_repository(
    name = "org_uber_go_goleak",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/goleak",
    sum = "h1:NBol2c7O1ZokfZ0LEU9K6Whx/KnwvepVetCUhtKja4A=",
    version = "v1.2.1",
)

go_repository(
    name = "org_uber_go_multierr",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/multierr",
    sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
    version = "v1.11.0",
)

go_repository(
    name = "org_uber_go_tools",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/tools",
    sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
    version = "v0.0.0-20190618225709-2cfd321de3ee",
)

go_repository(
    name = "org_uber_go_zap",
    build_file_proto_mode = "disable",
    importpath = "go.uber.org/zap",
    sum = "h1:sI7k6L95XOKS281NhVKOFCUNIvv9e0w4BF8N3u+tCRo=",
    version = "v1.26.0",
)

go_rules_dependencies()

go_register_toolchains(version = golang_version)

gazelle_dependencies()
