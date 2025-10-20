#!/bin/bash

set -euo pipefail

ASCII=$(cat << "EOF"
 ███▄ ▄███▓ ▄▄▄        ▄████  ███▄    █  ██▓▄▄▄█████▓ █    ██ ▓█████▄ ▓█████ 
▓██▒▀█▀ ██▒▒████▄     ██▒ ▀█▒ ██ ▀█   █ ▓██▒▓  ██▒ ▓▒ ██  ▓██▒▒██▀ ██▌▓█   ▀ 
▓██    ▓██░▒██  ▀█▄  ▒██░▄▄▄░▓██  ▀█ ██▒▒██▒▒ ▓██░ ▒░▓██  ▒██░░██   █▌▒███   
▒██    ▒██ ░██▄▄▄▄██ ░▓█  ██▓▓██▒  ▐▌██▒░██░░ ▓██▓ ░ ▓▓█  ░██░░▓█▄   ▌▒▓█  ▄ 
▒██▒   ░██▒ ▓█   ▓██▒░▒▓███▀▒▒██░   ▓██░░██░  ▒██▒ ░ ▒▒█████▓ ░▒████▓ ░▒████▒
░ ▒░   ░  ░ ▒▒   ▓▒█░ ░▒   ▒ ░ ▒░   ▒ ▒ ░▓    ▒ ░░   ░▒▓▒ ▒ ▒  ▒▒▓  ▒ ░░ ▒░ ░
░  ░      ░  ▒   ▒▒ ░  ░   ░ ░ ░░   ░ ▒░ ▒ ░    ░    ░░▒░ ░ ░  ░ ▒  ▒  ░ ░  ░
░      ░     ░   ▒   ░ ░   ░    ░   ░ ░  ▒ ░  ░       ░░░ ░ ░  ░ ░  ░    ░   
       ░         ░  ░      ░          ░  ░              ░        ░       ░  ░
EOF
)

DESKTOPS=("KDE Plasma" "GNOME" "Budgie" "i3" "XFCE" "Cinnamon" "MATE" "LXQt" "Deepin" "Hyprland" "Wayfire" "Sway")

echo "$ASCII"
echo
echo "Select a Desktop Environment by number:"

for i in "${!DESKTOPS[@]}"; do
    echo "$((i+1))) ${DESKTOPS[i]}"
done

read -rp "Enter your choice [1-${#DESKTOPS[@]}]: " choice

if ! [[ "$choice" =~ ^[0-9]+$ ]] || (( choice < 1 || choice > ${#DESKTOPS[@]} )); then
    echo "Invalid selection."
    exit 1
fi

CHOICE="${DESKTOPS[choice-1]}"
echo "Installing $CHOICE..."

case "$CHOICE" in
    "KDE Plasma")
        sudo pacman -S --needed plasma kde-applications sddm --noconfirm
        sudo systemctl enable sddm --now
        ;;
    "GNOME")
        sudo pacman -S --needed gnome gdm --noconfirm
        sudo systemctl enable gdm --now
        ;;
    "Budgie")
        sudo pacman -S --needed budgie-desktop gdm --noconfirm
        sudo systemctl enable gdm --now
        ;;
    "i3")
        sudo pacman -S --needed i3-wm i3status dmenu --noconfirm
        ;;
    "XFCE")
        sudo pacman -S --needed xfce4 xfce4-goodies lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "Cinnamon")
        sudo pacman -S --needed cinnamon lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "MATE")
        sudo pacman -S --needed mate mate-extra lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "LXQt")
        sudo pacman -S --needed lxqt sddm --noconfirm
        sudo systemctl enable sddm --now
        ;;
    "Deepin")
        sudo pacman -S --needed deepin deepin-extra lightdm --noconfirm
        sudo systemctl enable lightdm --now
        ;;
    "Hyprland")
        sudo pacman -S --needed hyprland waybar --noconfirm
        ;;
    "Wayfire")
        sudo pacman -S --needed wayfire waybar --noconfirm
        ;;
    "Sway")
        sudo pacman -S --needed sway waybar --noconfirm
        ;;
esac

echo "$CHOICE installation complete!"
