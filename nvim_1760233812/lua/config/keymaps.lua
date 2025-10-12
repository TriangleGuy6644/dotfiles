-- Keymaps are automatically loaded on the VeryLazy event
-- Default keymaps that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/keymaps.lua
-- Add any additional keymaps here

vim.keymap.set("n", "<leader>r", ":GRun<CR>", { noremap = true, silent = true, desc = "Run main.go in terminal" })
vim.keymap.set("t", "<Esc>", [[<C-\><C-n>]], { noremap = true })
