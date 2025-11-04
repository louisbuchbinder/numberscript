"""
"""
gazelle_prefix = "github.com/louisbuchbinder/core"

def wasm_js(name, wasm_exports, module):
    native.genrule(
        name = name,
        srcs = [":pkg"],
        outs = ["wasm.js", "sha256.wasm.js"],
        tools = ["//wasm/js/script"],
        cmd = """
            $(location //wasm/js/script) \
                {wasm_exports} \
                -module {module} \
                -wasm-go $(location :pkg) \
                -output $(location :wasm.js) \
                -output-sha256 $(location :sha256.wasm.js)
        """.format(
            wasm_exports = " ".join(["-wasm-export {}".format(e) for e in wasm_exports]),
            module = module,
        ),
    )

go_export_template = '\tjs.Global(){go_export_gets}.Set("{func_name}", js.FuncOf(func(_ js.Value, args []js.Value) any {{\n\t\treturn jsutil.OrError(pkg.{func_name}(jsutil.ToValues(args)))\n\t}}))'

wasm_go_main_template = """
package main

import (
	"syscall/js"
    jsutil "github.com/louisbuchbinder/core/wasm/js"
	pkg "{gazelle_prefix}/{module}"
)

func main() {{
{go_exports}
	select{{}}
}}
"""

def _wasm_go_impl(ctx):
    go_out = ctx.actions.declare_file(ctx.attr.name)
    module = "/".join(ctx.build_file_path.split("/")[:-2])

    go_exports = "\n".join([
        go_export_template.format(
            func_name = wasm_export.split(".")[-1],
            go_export_gets = "".join([
                '.Get("{}")'.format(e)
                for e in wasm_export.split(".")[:-1]
            ]),
        )
        for wasm_export in ctx.attr.wasm_exports
    ])

    ctx.actions.write(
        go_out,
        wasm_go_main_template.format(
            gazelle_prefix = gazelle_prefix,
            module = module,
            go_exports = go_exports,
        ),
    )
    return [DefaultInfo(files = depset([go_out]))]

wasm_go = rule(
    implementation = _wasm_go_impl,
    attrs = {
        "wasm_exports": attr.string_list(),
    },
)

def wasm_html(name, module):
    output = name + ".html"
    output_sha256 = "sha256." + output
    native.genrule(
        name = name,
        srcs = [":sha256.wasm.js"],
        outs = [output, output_sha256],
        tools = ["//wasm/html/script"],
        cmd = """
            $(location //wasm/html/script) \
                -file $(location :sha256.wasm.js) \
                -module {module} \
                -output $(location {output}) \
                -output-sha256 $(location {output_sha256})
        """.format(module = module, output = output, output_sha256 = output_sha256),
    )

def wasm_dist(name, module):
    native.genrule(
        name = name,
        srcs = [
            ":pkg",
            ":wasm.js",
            ":sha256.wasm.js",
            ":script.partial.html",
            ":sha256.script.partial.html",
        ],
        outs = ["pkg.tar.gz"],
        cmd = """
        D={module}/pkg
        function copy() {{
            cp $$1 $$D/$$2
        }}
        function sha256_dist() {{
            cp $$1 $$D/$$(sha256sum $$1 | cut -f 1 -d " ").$$2
        }}
        mkdir -p $$D
        copy $(location :pkg) main.wasm
        sha256_dist $(location :pkg) main.wasm
        copy $(location :script.partial.html) script.partial.html
        sha256_dist $(location :sha256.script.partial.html) script.partial.html
        copy $(location wasm.js) wasm.js
        sha256_dist $(location sha256.wasm.js) wasm.js
        tar -czf $(location :pkg.tar.gz) $$D/*
        """.format(module = module),
        visibility = ["//visibility:public"],
    )
