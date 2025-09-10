local cmd = { "/home/you/code/glsp/glsp" }

local root = vim.fs.dirname(vim.fs.find({ ".git", "go.mod" }, { upward = true })[1]) or vim.loop.cwd()

vim.lsp.start({
	name = "glsp-dev",
	cmd = cmd,
	root_dir = root,
	on_attach = function(client, bufnr)
		vim.keymap.set("n", "K", vim.lsp.buf.hover, { buffer = bufnr })
		vim.keymap.set("n", "gd", vim.lsp.buf.definition, { buffer = bufnr })
		vim.keymap.set("n", "<leader>rn", vim.lsp.buf.rename, { buffer = bufnr })
		print("glsp-dev attached to buffer " .. bufnr)
	end,
})
