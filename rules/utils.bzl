"""
bazel utility rules
"""

def symlink_impl(ctx):
    link = ctx.actions.declare_file(ctx.attr.name)
    ctx.actions.symlink(
        output = link,
        target_file = ctx.file.src,
    )
    return [DefaultInfo(files = depset([link]))]

symlink = rule(
    implementation = symlink_impl,
    attrs = {
        "src": attr.label(allow_single_file = True, mandatory = True),
    },
)

def _write_impl(ctx):
    ctx.actions.write(
        ctx.outputs.out_file,
        ctx.attr.content,
    )
    return [DefaultInfo(files = depset([ctx.outputs.out_file]))]

write = rule(
    implementation = _write_impl,
    attrs = {
        "content": attr.string(mandatory = True),
        "out_file": attr.output(doc = None, mandatory = True),
    },
)
