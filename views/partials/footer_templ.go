// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Footer() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<footer class=\"footer -mt-2 p-10 bg-base-300 text-accent-content\" data-theme=\"\"><div><img src=\"https://img.icons8.com/?size=512&amp;id=X3XGYoBQVt1Q&amp;format=png\" width=\"70\" height=\"50\"><p class=\"font-extrabold text-[18px]\">Fitness app</p></div><div><span class=\"footer-title\">FrontEnd</span> <a class=\"link link-hover flex\">Htmx</a> <a class=\"link link-hover flex\">Go Templ</a> <a class=\"link link-hover flex\">Tailwind</a> <a class=\"link link-hover flex\">Daisy UI</a></div><div><span class=\"footer-title\">BackEnd</span> <a class=\"link link-hover flex\">GoLang</a> <a class=\"link link-hover flex\">Chi-Router</a> <a class=\"link link-hover flex\">JWT</a> <a class=\"link link-hover flex\">SQL</a></div><div><span class=\"footer-title\">Tools</span> <a class=\"link link-hover flex\">Docker</a> <a class=\"link link-hover flex\">sqlc</a> <a class=\"link link-hover flex\">goose</a> <a class=\"link link-hover flex\">neovim</a></div></footer><script type=\"text/javascript\" src=\"https://cdn.jsdelivr.net/npm/toastify-js\"></script><script>\n  document.body.addEventListener(\"warnToast\", function (event) {\n    Warning(event.detail.value)\n  })\n\n  document.body.addEventListener(\"errorToast\", function (event) {\n    Error(event.detail.value)\n  })\n\n  document.body.addEventListener(\"successToast\", function (event) {\n    Success(event.detail.value)\n  })\n\n  document.body.addEventListener(\"infoToast\", function (event) {\n    Info(event.detail.value)\n  })\n\n  document.getElementById(\"calories_calc_form\").addEventListener(\"htmx:afterSwap\", function (event) {\n    let response = JSON.parse(event.detail.target.innerText)\n    event.detail.target.innerHTML = `\n    <iframe name=\"dummyframe\" id=\"dummyframe\" style=\"display: none;\"></iframe>\n    <form target=\"dummyframe\" id=\"kcalForm\" action=\"/server/daily-input\" method=\"post\">\n      <p>Calories:</p> <p class=\"text-accent\">${response.calories}</p>\n      <input name=\"calories\" type=\"hidden\" value=${response.calories} />\n    <br>\n      <p>CarboHydrates:</p> <p class=\"text-accent\">${response.totalNutrients.CHOCDF.quantity} ${response.totalNutrients.CHOCDF.unit}</p>\n      <input name=\"carbohydrates\" type=\"hidden\" value=${response.totalNutrients.CHOCDF.quantity} />\n    <br>\n      <p>Protein:</p> <p class=\"text-accent\">${response.totalNutrients.PROCNT.quantity} ${response.totalNutrients.PROCNT.unit}</p>\n      <input name=\"protein\" type=\"hidden\" value=${response.totalNutrients.PROCNT.quantity} />\n    <br>\n      <p>Fats:</p> <p class=\"text-accent\">${response.totalNutrients.FAT.quantity} ${response.totalNutrients.FAT.unit}</p>\n      <input name=\"fat\" type=\"hidden\" value=${response.totalNutrients.FAT.quantity} />\n    <br>\n      <p>Fiber:</p> <p class=\"text-accent\">${response.totalNutrients.FIBTG.quantity} ${response.totalNutrients.FIBTG.unit}</p>\n      <input name=\"fiber\" type=\"hidden\" value=${response.totalNutrients.FIBTG.quantity} />\n    <br>\n      <p>Sugar:</p> <p class=\"text-accent\">${response.totalNutrients.SUGAR.quantity} ${response.totalNutrients.SUGAR.unit}</p>\n    <br>\n      <p>Saturated Fats:</p> <p class=\"text-accent\">${response.totalNutrients.FASAT.quantity} ${response.totalNutrients.FASAT.unit}</p>\n    <br>\n      <button class=\"btn btn-lg btn-primary btn-outline\" type=\"button\" onclick=\"submit_form('kcalForm')\">Enter into logs</button>\n    </form>\n    `\n  })\n\n  function submit_form(id) {\n    const form = document.forms[id];\n    const formData = new URLSearchParams(new FormData(form));\n    fetch(form.action, {\n      method: 'post',\n      headers: {\n        'Content-Type': 'application/x-www-form-urlencoded'\n      },\n      body: formData\n    })\n      .then(response => {\n        if (response.status === 200) {\n          Success(\"Entered in logs\")\n        } else {\n          Error(\"Couldn't enter in logs\")\n        }\n      })\n      .catch(error => {\n        console.error('Error:', error);\n      });\n  }\n\n  function DrawerClose() {\n    document.getElementById(\"my-drawer\").checked = false\n  }\n\n  function SetTheme() {\n    const elements = document.querySelectorAll(\"[data-theme]\")\n    elements.forEach(element => {\n      let theme = window.localStorage.getItem(\"theme\")\n      if (theme === null) {\n        theme = \"retro\"\n      }\n      element.setAttribute(\"data-theme\", theme)\n    })\n  }\n\n  function ChangeTheme() {\n    if (window.localStorage.getItem(\"theme\") == \"retro\") {\n      current_theme = \"coffee\"\n    } else {\n      current_theme = \"retro\"\n    }\n    window.localStorage.setItem(\"theme\", current_theme)\n    const elements = document.querySelectorAll(\"[data-theme]\")\n    elements.forEach(element => {\n      const theme = window.localStorage.getItem(\"theme\") || \"luxury\"\n      element.setAttribute(\"data-theme\", theme)\n    })\n  }\n\n  function Warning(message) {\n    Toastify({\n      text: message,\n      duration: 1500,\n      close: false,\n      gravity: \"bottom\", // `top` or `bottom`\n      position: \"center\", // `left`, `center` or `right`\n      stopOnFocus: true, // Prevents dismissing of toast on hover\n      style: {\n        background: \"oklch(var(--wa))\",\n      },\n    }).showToast();\n  }\n\n  function Success(message) {\n    Toastify({\n      text: message,\n      duration: 1500,\n      close: false,\n      gravity: \"bottom\", // `top` or `bottom`\n      position: \"center\", // `left`, `center` or `right`\n      stopOnFocus: true, // Prevents dismissing of toast on hover\n      style: {\n        background: \"oklch(var(--su))\",\n      },\n    }).showToast();\n  }\n\n  function Error(message) {\n    Toastify({\n      text: message,\n      duration: 1500,\n      close: false,\n      gravity: \"bottom\", // `top` or `bottom`\n      position: \"center\", // `left`, `center` or `right`\n      stopOnFocus: true, // Prevents dismissing of toast on hover\n      style: {\n        background: \"oklch(var(--er))\",\n      },\n    }).showToast();\n  }\n\n  function Info(message) {\n    Toastify({\n      text: message,\n      duration: 1500,\n      close: false,\n      gravity: \"bottom\", // `top` or `bottom`\n      position: \"center\", // `left`, `center` or `right`\n      stopOnFocus: false, // Prevents dismissing of toast on hover\n      style: {\n        background: \"oklch(var(--in))\",\n      },\n    }).showToast();\n  }\n\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
