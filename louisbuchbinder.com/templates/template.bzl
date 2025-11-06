"""
"""

template_struct_field = provider(
    doc = "definition for a single struct field",
    fields = ["name", "type"],
)

template = provider(
    doc = "definition for the template",
    fields = ["name", "const_name", "filename", "struct_fields_list"],
)

go_lib_template = """
//go:embed {filename}
var {const_name}_TEMPLATE string

type {name}TemplateInput struct{{{struct_fields}}}

var {name}Template *template.Template

func init() {{
	{name}Template = template.Must(template.New("{const_name}_TEMPLATE").Parse({const_name}_TEMPLATE))
	var _ = MustRender{name}Template({name}TemplateInput{{}})
}}

func Render{name}Template(in {name}TemplateInput) ([]byte, error) {{
	return util.ExecuteTemplate({name}Template, in)
}}

func MustRender{name}Template(in {name}TemplateInput) []byte {{
	return util.Must(Render{name}Template(in))
}}
"""

def render_struct_fields(struct_fields_list):
    if len(struct_fields_list) == 0:
        return ""
    content = "\n".join([
        "\t{name} {type}".format(name = f.name, type = f.type)
        for f in struct_fields_list
    ])
    return "\n" + content + "\n"

def render_templates(templates):
    return "".join([
        go_lib_template.format(
            name = t.name,
            const_name = t.const_name,
            filename = t.filename,
            struct_fields = render_struct_fields(t.struct_fields_list),
        )
        for t in templates
    ])

embed_go_template = """package templates

import (
	_ "embed"
	"html/template"

	"github.com/louisbuchbinder/core/lib/util"
)
{content}
"""

def render_templates_embed_go(templates):
    return embed_go_template.format(content = render_templates(templates))
