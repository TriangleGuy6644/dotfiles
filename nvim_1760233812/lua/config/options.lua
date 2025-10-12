-- Options are automatically loaded before lazy.nvim startup
-- Default options that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/options.lua
-- Add any additional options here

vim.api.nvim_create_user_command("GRun", function()
  -- If we already have a stored terminal buffer and it's still valid
  if vim.g.go_run_term_buf and vim.api.nvim_buf_is_valid(vim.g.go_run_term_buf) then
    -- Switch to it and rerun
    vim.api.nvim_set_current_buf(vim.g.go_run_term_buf)
    vim.cmd("startinsert")
    vim.fn.chansend(vim.b.terminal_job_id, "clear && go run main.go\n")
  else
    -- Otherwise open a new split terminal
    vim.cmd("botright split | resize 10")
    vim.cmd("terminal")
    -- Remember this terminal buffer ID
    vim.g.go_run_term_buf = vim.api.nvim_get_current_buf()
    -- Run your Go file
    vim.fn.chansend(vim.b.terminal_job_id, "go run main.go\n")
  end
end, {})
