## Dummy LSP Server

This project contains a **dummy implementation of a custom Language Server Protocol (LSP) server**, exploring different types of I/O happening between the client and the server.  
You can find the examples in the `lsp/` folder.

Some of the interactions covered:

- **Initialization & Synchronization**  
  When a client connects, it sends an `initialize` request, and the server attaches to the workspace/buffer. The server and client then synchronize document state (`textDocument/didOpen`, `didChange`, `didSave`).

- **Document Events**  
  Every time we edit or save a file, the client notifies the server (e.g., `textDocument/didChange`, `textDocument/didSave`).

- **Hover Information**  
  When hovering over a symbol, the client sends a `textDocument/hover` request, and the server responds with documentation or type info.

- **Go to Definition / References**  
  Clicking on a function, variable, or method triggers requests like `textDocument/definition` or `textDocument/references`, allowing navigation across files.

- **Completions**  
  As we type, the client sends a `textDocument/completion` request, and the server suggests completions (functions, variables, snippets, etc.).

- **Diagnostics**  
  On each edit or save, the server can publish diagnostics (`textDocument/publishDiagnostics`), which the client shows as errors, warnings, or hints.

This dummy server demonstrates how these requests and notifications are exchanged between an LSP client (e.g., VSCode, Neovim) and the server.

## Architecture

```text
┌─────────────┐    JSON-RPC     ┌─────────────┐    JSON-RPC     ┌─────────────┐
│             │ ──────────────→ │             │ ←────────────── │             │
│   VSCode    │                 │ GLSP Server │                 │   Neovim    │
│ Extension   │ ←────────────── │   Binary    │ ──────────────→ │ Lua Client  │
└─────────────┘                 └─────────────┘                 └─────────────┘
```

### Lua-Client setup

To attach the custom `glsp` server to Markdown files in Neovim, add the following configuration in:

and add the following code:

```lua
-- lua/after/plugin/glsp-client.lua
-- Example: attach custom GLSP server to Markdown files

-- Find project root (search upward for .git or go.mod, fallback to cwd)
-- it will prvent starting new lsp server for diff file in same project
local root = vim.fs.dirname(vim.fs.find({ ".git", "go.mod" }, { upward = true })[1]) or vim.loop.cwd()

-- Attach custom keymaps when LSP starts
local on_attach = require("shekhar.lsp_keymaps")

-- Start GLSP client
local client = vim.lsp.start_client({
  name = "glsp",
  cmd = { "/home/shekhar/Personal/glsp/glsp" }, -- path to your server binary
  root_dir = root,
  on_attach = on_attach,
})

if not client then
  vim.notify("Error while starting GLSP client!", vim.log.levels.ERROR)
end

-- Attach GLSP only for Markdown files
vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function(args)
    vim.lsp.buf_attach_client(args.buf, client)
  end,
})

```

## Example logs

```log
[ glsp ]2025/09/08 23:55:38 main.go:15: glsp is started
[ glsp ]2025/09/08 23:55:38 handler.go:11: A request of initialize type came...
[ glsp ]2025/09/08 23:55:38 handler.go:19: Connected to: Neovim~0.11.3+v0.11.3
[ glsp ]2025/09/08 23:55:38 handler.go:11: A request of initialized type came...
[ glsp ]2025/09/09 00:07:03 handler.go:11: A request of textDocument/didOpen type came...
[ glsp ]2025/09/09 00:07:03 handler.go:29: opend file :file://
[ glsp ]2025/09/09 00:07:07 handler.go:11: A request of textDocument/didClose type came...
[ glsp ]2025/09/09 00:17:21 handler.go:11: A request of textDocument/didOpen type came...
[ glsp ]2025/09/09 00:17:21 handler.go:29: opend file :file://
[ glsp ]2025/09/09 00:17:25 handler.go:11: A request of textDocument/didClose type came...
[ glsp ]2025/09/09 00:48:17 handler.go:11: A request of shutdown type came...
st.md
[ glsp ]2025/09/08 23:55:40 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:40 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:40 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:40 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:41 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:41 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:41 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:41 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:42 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:42 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:43 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:43 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:44 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:44 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:44 handler.go:11: A request of textDocument/didChange type came...
[ glsp ]2025/09/08 23:55:44 handler.go:37: changed file :file:///home/shekhar/Personal/glsp/test.md
[ glsp ]2025/09/08 23:55:46 handler.go:11: A request of textDocument/definition type came...
[ glsp ]2025/09/08 23:55:47 handler.go:11: A request of textDocument/definition type came...
[ glsp ]2025/09/08 23:55:52 handler.go:11: A request of textDocument/definition type came...
[ glsp ]2025/09/08 23:56:03 handler.go:11: A request of textDocument/definition type came...
[ glsp ]2025/09/08 23:56:05 handler.go:11: A request of textDocument/hover type came...
[ glsp ]2025/09/08 23:56:05 handler.go:11: A request of textDocument/didOpen type came...
[ glsp ]2025/09/08 23:56:05 handler.go:29: opend file :file://

```
