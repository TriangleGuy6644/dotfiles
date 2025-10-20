#!/bin/bash

set -euo pipefail

if ! command -v dialog &> /dev/null; then
    sudo pacman -S --noconfirm dialog
fi

cat << "EOF"
 ███▄ ▄███▓ ▄▄▄        ▄████  ███▄    █  ██▓▄▄▄█████▓ █    ██ ▓█████▄ ▓█████ 
▓██▒▀█▀ ██▒▒████▄     ██▒ ▀█▒ ██ ▀█   █ ▓██▒▓  ██▒ ▓▒ ██  ▓██▒▒██▀ ██▌▓█   ▀ 
▓██    ▓██░▒██  ▀█▄  ▒██░▄▄▄░▓██  ▀█ ██▒▒██▒▒ ▓██░ ▒░▓██  ▒██░░██   █▌▒███   
▒██    ▒██ ░██▄▄▄▄██ ░▓█  ██▓▓██▒  ▐▌██▒░██░░ ▓██▓ ░ ▓▓█  ░██░░▓█▄   ▌▒▓█  ▄ 
▒██▒   ░██▒ ▓█   ▓██▒░▒▓███▀▒▒██░   ▓██░░██░  ▒██▒ ░ ▒▒█████▓ ░▒████▓ ░▒████▒
░ ▒░   ░  ░ ▒▒   ▓▒█░ ░▒   ▒ ░ ▒░   ▒ ▒ ░▓    ▒ ░░   ░▒▓▒ ▒ ▒  ▒▒▓  ▒ ░░ ▒░ ░
░  ░      ░  ▒   ▒▒ ░  ░   ░ ░ ░░   ░ ▒░ ▒ ░    ░    ░░▒░ ░ ░  ░ ▒  ▒  ░ ░  ░
░      ░     ░   ▒   ░ ░   ░    ░   ░ ░  ▒ ░  ░       ░░░ ░ ░  ░ ░  ░    ░   
       ░         ░  ░      ░          ░  ░              ░        ░       ░  ░
EOF

DE_LIST=(
"KDE Plasma" "gnome" "sudo pacman -S --needed plasma kde-applications sddm"
"GNOME" "gnome" "sudo pacman -S --needed gnome gdm"
"Budgie" "budgie-desktop" "sudo pacman -S --needed budgie-desktop gdm"
"i3" "i3-wm" "sudo pacman -S --needed i3-wm i3status dmenu"
"XFCE" "xfce4" "sudo pacman -S --needed xfce4 xfce4-goodies lightdm"
"Cinnamon" "cinnamon" "sudo pacman -S --needed cinnamon lightdm"
"MATE" "mate" "sudo pacman -S --needed mate mate-extra lightdm"
"LXQt" "lxqt" "sudo pacman -S --needed lxqt sddm"
"Deepin" "deepin" "sudo pacman -S --needed deepin deepin-extra lightdm"
"Hyprland" "hyprland" "sudo pacman -S --needed hyprland waybar"
"Wayfire" "wayfire" "sudo pacman -S --needed wayfire waybar"
"Sway" "sway" "sudo pacman -S --needed sway waybar"
)

MENU_ITEMS=()
for ((i=0;i<${#DE_LIST[@]};i+=3)); do
    MENU_ITEMS+=("${DE_LIST[i]}" "")
done

CHOICE=$(dialog --clear --title "Select Desktop Environment" \
--menu "Use arrow keys to select, Enter to confirm" 20 60 12 \
"${MENU_ITEMS[@]}" 3>&1 1>&2 2>&3)

clear

for ((i=0;i<${#DE_LIST[@]};i+=3)); do
    if [[ "${DE_LIST[i]}" == "$CHOICE" ]]; then
        INSTALL_CMD="${DE_LIST[i+2]}"
        break
    fi
done

echo "Installing $CHOICE..."
eval "$INSTALL_CMD"

case "$CHOICE" in
    "KDE Plasma"|"LXQt")
        sudo systemctl enable sddm --now
        ;;
    "GNOME"|"Budgie")
        sudo systemctl enable gdm --now
        ;;
    "XFCE"|"Cinnamon"|"MATE"|"Deepin")
        sudo systemctl enable lightdm --now
        ;;
esac

echo "$CHOICE installation complete!"
